/*
 * @title String & slice utility library for Solidity contracts.
 * @author Nick Johnson <arachnid@notdot.net>
 *
 * @dev Functionality in this library is largely implemented using an
 *      abstraction called a 'slice'. A slice represents a part of a string -
 *      anything from the entire string to a single character, or even no
 *      characters at all (a 0-length slice). Since a slice only has to specify
 *      an offset and a length, copying and manipulating slices is a lot less
 *      expensive than copying and manipulating the strings they reference.
 *
 *      To further reduce gas costs, most functions on slice that need to return
 *      a slice modify the original one instead of allocating a new one; for
 *      instance, `s.split(".")` will return the text up to the first '.',
 *      modifying s to only contain the remainder of the string after the '.'.
 *      In situations where you do not want to modify the original slice, you
 *      can make a copy first with `.copy()`, for example:
 *      `s.copy().split(".")`. Try and avoid using this idiom in loops; since
 *      Solidity has no memory management, it will result in allocating many
 *      short-lived slices that are later discarded.
 *
 *      Functions that return two slices come in two versions: a non-allocating
 *      version that takes the second slice as an argument, modifying it in
 *      place, and an allocating version that allocates and returns the second
 *      slice; see `nextRune` for example.
 *
 *      Functions that have to copy string data will return strings rather than
 *      slices; these can be cast back to slices for further processing if
 *      required.
 *
 *      For convenience, some functions are provided with non-modifying
 *      variants that create a new slice and return both; for instance,
 *      `s.splitNew('.')` leaves s unmodified, and returns two values
 *      corresponding to the left and right parts of the string.
 */

pragma solidity ^0.4.14;

library strings {
    struct slice {
        uint _len;
        uint _ptr;
    }

    function memcpy(uint dest, uint src, uint len) private pure {
        // Copy word-length chunks while possible
        for(; len >= 32; len -= 32) {
            assembly {
                mstore(dest, mload(src))
            }
            dest += 32;
            src += 32;
        }

        // Copy remaining bytes
        uint mask = 256 ** (32 - len) - 1;
        assembly {
            let srcpart := and(mload(src), not(mask))
            let destpart := and(mload(dest), mask)
            mstore(dest, or(destpart, srcpart))
        }
    }

    /*
     * @dev Returns a slice containing the entire string.
     * @param self The string to make a slice from.
     * @return A newly allocated slice containing the entire string.
     */
    function toSlice(string memory self) internal pure returns (slice memory) {
        uint ptr;
        assembly {
            ptr := add(self, 0x20)
        }
        return slice(bytes(self).length, ptr);
    }

    /*
     * @dev Returns the length of a null-terminated bytes32 string.
     * @param self The value to find the length of.
     * @return The length of the string, from 0 to 32.
     */
    function len(bytes32 self) internal pure returns (uint) {
        uint ret;
        if (self == 0)
            return 0;
        if (self & 0xffffffffffffffffffffffffffffffff == 0) {
            ret += 16;
            self = bytes32(uint(self) / 0x100000000000000000000000000000000);
        }
        if (self & 0xffffffffffffffff == 0) {
            ret += 8;
            self = bytes32(uint(self) / 0x10000000000000000);
        }
        if (self & 0xffffffff == 0) {
            ret += 4;
            self = bytes32(uint(self) / 0x100000000);
        }
        if (self & 0xffff == 0) {
            ret += 2;
            self = bytes32(uint(self) / 0x10000);
        }
        if (self & 0xff == 0) {
            ret += 1;
        }
        return 32 - ret;
    }

    /*
     * @dev Returns a slice containing the entire bytes32, interpreted as a
     *      null-terminated utf-8 string.
     * @param self The bytes32 value to convert to a slice.
     * @return A new slice containing the value of the input argument up to the
     *         first null.
     */
    function toSliceB32(bytes32 self) internal pure returns (slice memory ret) {
        // Allocate space for `self` in memory, copy it there, and point ret at it
        assembly {
            let ptr := mload(0x40)
            mstore(0x40, add(ptr, 0x20))
            mstore(ptr, self)
            mstore(add(ret, 0x20), ptr)
        }
        ret._len = len(self);
    }

    /*
     * @dev Returns a new slice containing the same data as the current slice.
     * @param self The slice to copy.
     * @return A new slice containing the same data as `self`.
     */
    function copy(slice memory self) internal pure returns (slice memory) {
        return slice(self._len, self._ptr);
    }

    /*
     * @dev Copies a slice to a new string.
     * @param self The slice to copy.
     * @return A newly allocated string containing the slice's text.
     */
    function toString(slice memory self) internal pure returns (string memory) {
        string memory ret = new string(self._len);
        uint retptr;
        assembly { retptr := add(ret, 32) }

        memcpy(retptr, self._ptr, self._len);
        return ret;
    }

    /*
     * @dev Returns the length in runes of the slice. Note that this operation
     *      takes time proportional to the length of the slice; avoid using it
     *      in loops, and call `slice.empty()` if you only need to know whether
     *      the slice is empty or not.
     * @param self The slice to operate on.
     * @return The length of the slice in runes.
     */
    function len(slice memory self) internal pure returns (uint l) {
        // Starting at ptr-31 means the LSB will be the byte we care about
        uint ptr = self._ptr - 31;
        uint end = ptr + self._len;
        for (l = 0; ptr < end; l++) {
            uint8 b;
            assembly { b := and(mload(ptr), 0xFF) }
            if (b < 0x80) {
                ptr += 1;
            } else if(b < 0xE0) {
                ptr += 2;
            } else if(b < 0xF0) {
                ptr += 3;
            } else if(b < 0xF8) {
                ptr += 4;
            } else if(b < 0xFC) {
                ptr += 5;
            } else {
                ptr += 6;
            }
        }
    }

    /*
     * @dev Returns true if the slice is empty (has a length of 0).
     * @param self The slice to operate on.
     * @return True if the slice is empty, False otherwise.
     */
    function empty(slice memory self) internal pure returns (bool) {
        return self._len == 0;
    }

    /*
     * @dev Returns a positive number if `other` comes lexicographically after
     *      `self`, a negative number if it comes before, or zero if the
     *      contents of the two slices are equal. Comparison is done per-rune,
     *      on unicode codepoints.
     * @param self The first slice to compare.
     * @param other The second slice to compare.
     * @return The result of the comparison.
     */
    function compare(slice memory self, slice memory other) internal pure returns (int) {
        uint shortest = self._len;
        if (other._len < self._len)
            shortest = other._len;

        uint selfptr = self._ptr;
        uint otherptr = other._ptr;
        for (uint idx = 0; idx < shortest; idx += 32) {
            uint a;
            uint b;
            assembly {
                a := mload(selfptr)
                b := mload(otherptr)
            }
            if (a != b) {
                // Mask out irrelevant bytes and check again
                uint256 mask = uint256(-1); // 0xffff...
                if(shortest < 32) {
                  mask = ~(2 ** (8 * (32 - shortest + idx)) - 1);
                }
                uint256 diff = (a & mask) - (b & mask);
                if (diff != 0)
                    return int(diff);
            }
            selfptr += 32;
            otherptr += 32;
        }
        return int(self._len) - int(other._len);
    }

    /*
     * @dev Returns true if the two slices contain the same text.
     * @param self The first slice to compare.
     * @param self The second slice to compare.
     * @return True if the slices are equal, false otherwise.
     */
    function equals(slice memory self, slice memory other) internal pure returns (bool) {
        return compare(self, other) == 0;
    }

    /*
     * @dev Extracts the first rune in the slice into `rune`, advancing the
     *      slice to point to the next rune and returning `self`.
     * @param self The slice to operate on.
     * @param rune The slice that will contain the first rune.
     * @return `rune`.
     */
    function nextRune(slice memory self, slice memory rune) internal pure returns (slice memory) {
        rune._ptr = self._ptr;

        if (self._len == 0) {
            rune._len = 0;
            return rune;
        }

        uint l;
        uint b;
        // Load the first byte of the rune into the LSBs of b
        assembly { b := and(mload(sub(mload(add(self, 32)), 31)), 0xFF) }
        if (b < 0x80) {
            l = 1;
        } else if(b < 0xE0) {
            l = 2;
        } else if(b < 0xF0) {
            l = 3;
        } else {
            l = 4;
        }

        // Check for truncated codepoints
        if (l > self._len) {
            rune._len = self._len;
            self._ptr += self._len;
            self._len = 0;
            return rune;
        }

        self._ptr += l;
        self._len -= l;
        rune._len = l;
        return rune;
    }

    /*
     * @dev Returns the first rune in the slice, advancing the slice to point
     *      to the next rune.
     * @param self The slice to operate on.
     * @return A slice containing only the first rune from `self`.
     */
    function nextRune(slice memory self) internal pure returns (slice memory ret) {
        nextRune(self, ret);
    }

    /*
     * @dev Returns the number of the first codepoint in the slice.
     * @param self The slice to operate on.
     * @return The number of the first codepoint in the slice.
     */
    function ord(slice memory self) internal pure returns (uint ret) {
        if (self._len == 0) {
            return 0;
        }

        uint word;
        uint length;
        uint divisor = 2 ** 248;

        // Load the rune into the MSBs of b
        assembly { word:= mload(mload(add(self, 32))) }
        uint b = word / divisor;
        if (b < 0x80) {
            ret = b;
            length = 1;
        } else if(b < 0xE0) {
            ret = b & 0x1F;
            length = 2;
        } else if(b < 0xF0) {
            ret = b & 0x0F;
            length = 3;
        } else {
            ret = b & 0x07;
            length = 4;
        }

        // Check for truncated codepoints
        if (length > self._len) {
            return 0;
        }

        for (uint i = 1; i < length; i++) {
            divisor = divisor / 256;
            b = (word / divisor) & 0xFF;
            if (b & 0xC0 != 0x80) {
                // Invalid UTF-8 sequence
                return 0;
            }
            ret = (ret * 64) | (b & 0x3F);
        }

        return ret;
    }

    /*
     * @dev Returns the keccak-256 hash of the slice.
     * @param self The slice to hash.
     * @return The hash of the slice.
     */
    function keccak(slice memory self) internal pure returns (bytes32 ret) {
        assembly {
            ret := keccak256(mload(add(self, 32)), mload(self))
        }
    }

    /*
     * @dev Returns true if `self` starts with `needle`.
     * @param self The slice to operate on.
     * @param needle The slice to search for.
     * @return True if the slice starts with the provided text, false otherwise.
     */
    function startsWith(slice memory self, slice memory needle) internal pure returns (bool) {
        if (self._len < needle._len) {
            return false;
        }

        if (self._ptr == needle._ptr) {
            return true;
        }

        bool equal;
        assembly {
            let length := mload(needle)
            let selfptr := mload(add(self, 0x20))
            let needleptr := mload(add(needle, 0x20))
            equal := eq(keccak256(selfptr, length), keccak256(needleptr, length))
        }
        return equal;
    }

    /*
     * @dev If `self` starts with `needle`, `needle` is removed from the
     *      beginning of `self`. Otherwise, `self` is unmodified.
     * @param self The slice to operate on.
     * @param needle The slice to search for.
     * @return `self`
     */
    function beyond(slice memory self, slice memory needle) internal pure returns (slice memory) {
        if (self._len < needle._len) {
            return self;
        }

        bool equal = true;
        if (self._ptr != needle._ptr) {
            assembly {
                let length := mload(needle)
                let selfptr := mload(add(self, 0x20))
                let needleptr := mload(add(needle, 0x20))
                equal := eq(keccak256(selfptr, length), keccak256(needleptr, length))
            }
        }

        if (equal) {
            self._len -= needle._len;
            self._ptr += needle._len;
        }

        return self;
    }

    /*
     * @dev Returns true if the slice ends with `needle`.
     * @param self The slice to operate on.
     * @param needle The slice to search for.
     * @return True if the slice starts with the provided text, false otherwise.
     */
    function endsWith(slice memory self, slice memory needle) internal pure returns (bool) {
        if (self._len < needle._len) {
            return false;
        }

        uint selfptr = self._ptr + self._len - needle._len;

        if (selfptr == needle._ptr) {
            return true;
        }

        bool equal;
        assembly {
            let length := mload(needle)
            let needleptr := mload(add(needle, 0x20))
            equal := eq(keccak256(selfptr, length), keccak256(needleptr, length))
        }

        return equal;
    }

    /*
     * @dev If `self` ends with `needle`, `needle` is removed from the
     *      end of `self`. Otherwise, `self` is unmodified.
     * @param self The slice to operate on.
     * @param needle The slice to search for.
     * @return `self`
     */
    function until(slice memory self, slice memory needle) internal pure returns (slice memory) {
        if (self._len < needle._len) {
            return self;
        }

        uint selfptr = self._ptr + self._len - needle._len;
        bool equal = true;
        if (selfptr != needle._ptr) {
            assembly {
                let length := mload(needle)
                let needleptr := mload(add(needle, 0x20))
                equal := eq(keccak256(selfptr, length), keccak256(needleptr, length))
            }
        }

        if (equal) {
            self._len -= needle._len;
        }

        return self;
    }

    // Returns the memory address of the first byte of the first occurrence of
    // `needle` in `self`, or the first byte after `self` if not found.
    function findPtr(uint selflen, uint selfptr, uint needlelen, uint needleptr) private pure returns (uint) {
        uint ptr = selfptr;
        uint idx;

        if (needlelen <= selflen) {
            if (needlelen <= 32) {
                bytes32 mask = bytes32(~(2 ** (8 * (32 - needlelen)) - 1));

                bytes32 needledata;
                assembly { needledata := and(mload(needleptr), mask) }

                uint end = selfptr + selflen - needlelen;
                bytes32 ptrdata;
                assembly { ptrdata := and(mload(ptr), mask) }

                while (ptrdata != needledata) {
                    if (ptr >= end)
                        return selfptr + selflen;
                    ptr++;
                    assembly { ptrdata := and(mload(ptr), mask) }
                }
                return ptr;
            } else {
                // For long needles, use hashing
                bytes32 hash;
                assembly { hash := keccak256(needleptr, needlelen) }

                for (idx = 0; idx <= selflen - needlelen; idx++) {
                    bytes32 testHash;
                    assembly { testHash := keccak256(ptr, needlelen) }
                    if (hash == testHash)
                        return ptr;
                    ptr += 1;
                }
            }
        }
        return selfptr + selflen;
    }

    // Returns the memory address of the first byte after the last occurrence of
    // `needle` in `self`, or the address of `self` if not found.
    function rfindPtr(uint selflen, uint selfptr, uint needlelen, uint needleptr) private pure returns (uint) {
        uint ptr;

        if (needlelen <= selflen) {
            if (needlelen <= 32) {
                bytes32 mask = bytes32(~(2 ** (8 * (32 - needlelen)) - 1));

                bytes32 needledata;
                assembly { needledata := and(mload(needleptr), mask) }

                ptr = selfptr + selflen - needlelen;
                bytes32 ptrdata;
                assembly { ptrdata := and(mload(ptr), mask) }

                while (ptrdata != needledata) {
                    if (ptr <= selfptr)
                        return selfptr;
                    ptr--;
                    assembly { ptrdata := and(mload(ptr), mask) }
                }
                return ptr + needlelen;
            } else {
                // For long needles, use hashing
                bytes32 hash;
                assembly { hash := keccak256(needleptr, needlelen) }
                ptr = selfptr + (selflen - needlelen);
                while (ptr >= selfptr) {
                    bytes32 testHash;
                    assembly { testHash := keccak256(ptr, needlelen) }
                    if (hash == testHash)
                        return ptr + needlelen;
                    ptr -= 1;
                }
            }
        }
        return selfptr;
    }

    /*
     * @dev Modifies `self` to contain everything from the first occurrence of
     *      `needle` to the end of the slice. `self` is set to the empty slice
     *      if `needle` is not found.
     * @param self The slice to search and modify.
     * @param needle The text to search for.
     * @return `self`.
     */
    function find(slice memory self, slice memory needle) internal pure returns (slice memory) {
        uint ptr = findPtr(self._len, self._ptr, needle._len, needle._ptr);
        self._len -= ptr - self._ptr;
        self._ptr = ptr;
        return self;
    }

    /*
     * @dev Modifies `self` to contain the part of the string from the start of
     *      `self` to the end of the first occurrence of `needle`. If `needle`
     *      is not found, `self` is set to the empty slice.
     * @param self The slice to search and modify.
     * @param needle The text to search for.
     * @return `self`.
     */
    function rfind(slice memory self, slice memory needle) internal pure returns (slice memory) {
        uint ptr = rfindPtr(self._len, self._ptr, needle._len, needle._ptr);
        self._len = ptr - self._ptr;
        return self;
    }

    /*
     * @dev Splits the slice, setting `self` to everything after the first
     *      occurrence of `needle`, and `token` to everything before it. If
     *      `needle` does not occur in `self`, `self` is set to the empty slice,
     *      and `token` is set to the entirety of `self`.
     * @param self The slice to split.
     * @param needle The text to search for in `self`.
     * @param token An output parameter to which the first token is written.
     * @return `token`.
     */
    function split(slice memory self, slice memory needle, slice memory token) internal pure returns (slice memory) {
        uint ptr = findPtr(self._len, self._ptr, needle._len, needle._ptr);
        token._ptr = self._ptr;
        token._len = ptr - self._ptr;
        if (ptr == self._ptr + self._len) {
            // Not found
            self._len = 0;
        } else {
            self._len -= token._len + needle._len;
            self._ptr = ptr + needle._len;
        }
        return token;
    }

    /*
     * @dev Splits the slice, setting `self` to everything after the first
     *      occurrence of `needle`, and returning everything before it. If
     *      `needle` does not occur in `self`, `self` is set to the empty slice,
     *      and the entirety of `self` is returned.
     * @param self The slice to split.
     * @param needle The text to search for in `self`.
     * @return The part of `self` up to the first occurrence of `delim`.
     */
    function split(slice memory self, slice memory needle) internal pure returns (slice memory token) {
        split(self, needle, token);
    }

    /*
     * @dev Splits the slice, setting `self` to everything before the last
     *      occurrence of `needle`, and `token` to everything after it. If
     *      `needle` does not occur in `self`, `self` is set to the empty slice,
     *      and `token` is set to the entirety of `self`.
     * @param self The slice to split.
     * @param needle The text to search for in `self`.
     * @param token An output parameter to which the first token is written.
     * @return `token`.
     */
    function rsplit(slice memory self, slice memory needle, slice memory token) internal pure returns (slice memory) {
        uint ptr = rfindPtr(self._len, self._ptr, needle._len, needle._ptr);
        token._ptr = ptr;
        token._len = self._len - (ptr - self._ptr);
        if (ptr == self._ptr) {
            // Not found
            self._len = 0;
        } else {
            self._len -= token._len + needle._len;
        }
        return token;
    }

    /*
     * @dev Splits the slice, setting `self` to everything before the last
     *      occurrence of `needle`, and returning everything after it. If
     *      `needle` does not occur in `self`, `self` is set to the empty slice,
     *      and the entirety of `self` is returned.
     * @param self The slice to split.
     * @param needle The text to search for in `self`.
     * @return The part of `self` after the last occurrence of `delim`.
     */
    function rsplit(slice memory self, slice memory needle) internal pure returns (slice memory token) {
        rsplit(self, needle, token);
    }

    /*
     * @dev Counts the number of nonoverlapping occurrences of `needle` in `self`.
     * @param self The slice to search.
     * @param needle The text to search for in `self`.
     * @return The number of occurrences of `needle` found in `self`.
     */
    function count(slice memory self, slice memory needle) internal pure returns (uint cnt) {
        uint ptr = findPtr(self._len, self._ptr, needle._len, needle._ptr) + needle._len;
        while (ptr <= self._ptr + self._len) {
            cnt++;
            ptr = findPtr(self._len - (ptr - self._ptr), ptr, needle._len, needle._ptr) + needle._len;
        }
    }

    /*
     * @dev Returns True if `self` contains `needle`.
     * @param self The slice to search.
     * @param needle The text to search for in `self`.
     * @return True if `needle` is found in `self`, false otherwise.
     */
    function contains(slice memory self, slice memory needle) internal pure returns (bool) {
        return rfindPtr(self._len, self._ptr, needle._len, needle._ptr) != self._ptr;
    }

    /*
     * @dev Returns a newly allocated string containing the concatenation of
     *      `self` and `other`.
     * @param self The first slice to concatenate.
     * @param other The second slice to concatenate.
     * @return The concatenation of the two strings.
     */
    function concat(slice memory self, slice memory other) internal pure returns (string memory) {
        string memory ret = new string(self._len + other._len);
        uint retptr;
        assembly { retptr := add(ret, 32) }
        memcpy(retptr, self._ptr, self._len);
        memcpy(retptr + self._len, other._ptr, other._len);
        return ret;
    }

    /*
     * @dev Joins an array of slices, using `self` as a delimiter, returning a
     *      newly allocated string.
     * @param self The delimiter to use.
     * @param parts A list of slices to join.
     * @return A newly allocated string containing all the slices in `parts`,
     *         joined with `self`.
     */
    function join(slice memory self, slice[] memory parts) internal pure returns (string memory) {
        if (parts.length == 0)
            return "";

        uint length = self._len * (parts.length - 1);
        for(uint i = 0; i < parts.length; i++)
            length += parts[i]._len;

        string memory ret = new string(length);
        uint retptr;
        assembly { retptr := add(ret, 32) }

        for(i = 0; i < parts.length; i++) {
            memcpy(retptr, parts[i]._ptr, parts[i]._len);
            retptr += parts[i]._len;
            if (i < parts.length - 1) {
                memcpy(retptr, self._ptr, self._len);
                retptr += self._len;
            }
        }

        return ret;
    }
}

//pragma solidity ^0.4.25;
//import "github.com/Arachnid/solidity-stringutils/strings.sol";

contract Ownable {


    address public owner;
    mapping (address=> bool) public mapWriter;
    
    event OwnershipRenounced(address indexed previousOwner);
    event OwnershipTransferred(
        address indexed previousOwner,
        address indexed newOwner
    );

  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
    constructor() public {
        owner = msg.sender;
    }

  /**
   * @dev Throws if called by any account other than the owner.
   */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }
    
    modifier onlyWriter() {
        require(mapWriter[msg.sender]);
        _;
    }
    
    function addWriter(address writer) public onlyOwner {
        mapWriter[writer] = true;
    }
    
  /**
   * @dev Allows the current owner to relinquish control of the contract.
   * @notice Renouncing to ownership will leave the contract without an owner.
   * It will not be possible to call the functions with the `onlyOwner`
   * modifier anymore.
   */
   
    function renounceOwnership() public onlyOwner {
        emit OwnershipRenounced(owner);
        owner = address(0);
    }

  /**
   * @dev Allows the current owner to transfer control of the contract to a newOwner.
   * @param _newOwner The address to transfer ownership to.
   */
    function transferOwnership(address _newOwner) public onlyOwner {
        _transferOwnership(_newOwner);
    }

  /**
   * @dev Transfers control of the contract to a newOwner.
   * @param _newOwner The address to transfer ownership to.
   */
    function _transferOwnership(address _newOwner) internal {
        require(_newOwner != address(0));
        emit OwnershipTransferred(owner, _newOwner);
        owner = _newOwner;
    }
}

contract CupboardContract is Ownable {
    using strings for *;
/*******************************
   1 终端管理：
          云柜ID(bigint/20)；名称；设备编号(varchar/100)；CPU序列号(varchar/32)；所属合伙人(varchar/100);
          Mac地址(varchar/32)；APK版本(varchar/32)；商品表类型(smallint 1-五层   2- 13层 3-九层)；在线状态；状态(smallint)。
     查询筛选项:终端管理：云柜名称、MAC地址、在线状态（全部/离线/在线）

 struct yg_terminal{
      uint256  yg_id;           //   云柜ID(bigint/20)
      string   yg_name;         //#  名称
      string   device_sn;       // * 设备编号(varchar/100)
      string   cpu_sn;          // * CPU序列号(varchar/32)
      string   partner;         // * 所属合伙人(varchar/100)
      bytes32  mac_addr;        //#* Mac地址(varchar/32)
      string   apk_version;     //   APK版本(varchar/32)
      uint8    commodity;       //   商品表类型(smallint 1-五层   2- 13层 3-九层)
      uint8    online_status;   //   在线状态
      uint8    status;          //#  状态(smallint)
  }
  *******************************/
  
  /****************************** 
  2  用户交易记录：
           用户名(varchar/50)；手机号(varchar/20)；订单编号(varchar/64)；交易号(varchar/64)；交易额(decimal/20，2)；
           创建时间(datetime)；支付类型(smallint银行类型 1-微信端 2-钱包余额3-支付宝)；支付状态(smallint 1- 待支付 ;2 - 成功3-失败 ;4-已退款;5-已取消)；
           计算异常(smallint计算异常0-未出错  大于0表示失败以及失败的代码)。
     查询筛选项:  2.智能云柜-用户交易记录：手机号、订单号、创建开始时间-创建结束时间


 struct transaction_record{
      string   user_name;         //*  用户名(varchar/50)
      uint256  cell_number;       //*  手机号(varchar/20)
      string   order_number;      //   订单编号(varchar/64)
      string   transactoin_number;//   交易号(varchar/64)
      uint256  Transaction_amount;//   交易额(decimal/20，2)
      bytes32  create_time;       //   创建时间(datetime)
      uint8    payment_type;      //   支付类型(smallint银行类型 1-微信端 2-钱包余额3-支付宝)
      uint8    payment_status;    //   支付状态(smallint 1- 待支付 ;2 - 成功3-失败 ;4-已退款;5-已取消)
      uint8    pay_exception;     //   计算异常(smallint计算异常0-未出错  大于0表示失败以及失败的代码）
  }
  ******************************/
  
  /***************************** 
  3  充值记录：用户名(varchar/50)；手机号(varchar/20)；订单编号(varchar/64)；充值额(decimal/20，2)；创建时间(datetime)；支付时间(datetime)；
     支付状态(smallint 1- 待支付 ;2 - 成功3-失败 ;4-已退款;5-已取消)；支付类型(smallint银行类型 1-微信端 2-钱包余额3-支付宝)；微信交易ID(varchar/64)；
  查询筛选项: 智能云柜-充值记录：手机号、支付状态（所有/支付中/成功/失败）、创建开始时间-创建结束时间

   struct recharge_record{
      string   user_name;         //* 用户名(varchar/50)
      uint256  cell_number;       //* 手机号(varchar/20)
      string   order_number;      //  订单编号(varchar/64)
      uint256  recharge_amount;   //  充值额(decimal/20，2)
      bytes32  create_time;       //  创建时间(datetime)
      bytes32  pay_time;          //  支付时间(datetime)
      uint8    pay_status;        //  支付状态(smallint 1- 待支付 ;2 - 成功3-失败 ;4-已退款;5-已取消)
      uint8    pay_type;          //  支付类型(smallint银行类型 1-微信端 2-钱包余额3-支付宝)
      string   wechar_id;         //* 微信交易ID(varchar/64)
  }
  ******************************/  
  /***************************** 
  4、供应链管理-大仓管理:
     仓库名称(varchar/100)；仓库地址(varchar/200)；状态；创建时间(datetime)；修改时间(datetime)
   查询筛选项: 供应链管理-大仓管理：大仓名称

   struct big_warehouse{
      string  warehouse_name;       //   仓库名称(varchar/100)
      string  warehouse_address;    //*  仓库地址(varchar/200)
      uint8   status;               //   状态
      bytes32  create_time;          //   创建时间(datetime)
      bytes32  modify_time;          //   修改时间(datetime)
  }
  ******************************/  
  /***************************** 
  5、供应链管理-云仓管理：
     仓库名称(varchar/100)；是否实仓(int仓库类型：1-真实仓；2-虚拟仓；仓库地址(varchar/200)；
     所属片区(varchar/100)；仓库管理员(varchar/100)；状态；创建时间(datetime)；修改时间(datetime)；
   查询筛选项: 供应链管理-云仓管理：仓库名称

   struct cloud_warehouse{
      string   warehouse_name;       //  仓库名称(varchar/100)
      uint8    real;                 //  是否实仓(int仓库类型：1-真实仓；2-虚拟仓）
      string   warehouse_address;    //* 仓库地址(varchar/200)
      string   warehouse_area;       //  所属片区(varchar/100)
      string   warehouse_keeper;     //  仓库管理员(varchar/100)
      uint8    status;               //  状态
      bytes32  create_time;          //  创建时间(datetime)
      bytes32  modify_time;          //  修改时间(datetime)
  }
  ******************************/  
  
  /***************************** 
  6、商圈管理-商圈信息：
     商圈名称(varchar/100)；所属城市(varchar/100)；状态(int/2状态；1-有效；2-禁用；3-无效(删除)；
     创建时间(datetime)；修改时间(datetime)；
   查询筛选项: 商圈管理-商圈信息：商圈名称

   struct trading_area{
      string   name;            //  商圈名称(varchar/100)
      string   city;            //  所属城市(varchar/100)
      uint8    status;          //  状态(int/2状态；1-有效；2-禁用；3-无效(删除)；
      bytes32  create_time;     //  创建时间(datetime)
      bytes32  modify_time;     //  修改时间(datetime)
  }
  ******************************/ 
  
  /***************************** 
  7、商圈管理-商业片区：
     商业片区名称(varchar/100)；所属商圈(varchar/100)；
     状态(int状态；1-有效；2-禁用；3-无效(删除)；创建时间(datetime)；修改时间(datetime)；
     查询筛选项:商圈管理-商业片区：商业片区名称

   struct sowntown{
      string   name;            //   商业片区名称(varchar/100)
      string   trading_area;    //   所属商圈(varchar/100)
      uint8    status;          //   状态(int状态；1-有效；2-禁用；3-无效(删除)
      bytes32  create_time;     //   创建时间(datetime)
      bytes32  modify_time;     //   修改时间(datetime)
  }
  ******************************/    
  /***************************** 
  8、商圈管理-楼宇管理：
     楼宇名称(varchar/100)；所属片区(varchar/100)；所属云仓库(varchar/100)；
     状态(int/2状态；1-有效；2-禁用；3-无效(删除)；创建时间(datetime)；修改时间(datetime)；
     查询筛选项:商圈管理-楼宇管理：楼宇名称

  struct building{
      string   name;                 //  楼宇名称(varchar/100)
      string   sowntown;             //  所属片区(varchar/100)
      string   cloud_warehouse_name; //  所属云仓库(varchar/100)
      uint8   status;                //  状态(int/2状态；1-有效；2-禁用；3-无效(删除)
      bytes32  create_time;          //  创建时间(datetime)
      bytes32  modify_time;          //  修改时间(datetime)
  }
  ******************************/   
  event searchResult (string key,string value);
  event searchFound (string key,string value);
  event insertEvent (string table, uint256 arrayIndex, uint256 mapIndex, string ss);
  event updateEvent (string table, uint256 arrayIndex, uint256 mapIndex, string ss);
  event item (string key,string value);
  event list1 (string key,string value);
  event inserted (string listname, string list, string, string, string);
  event display1(string, string, string, string, string, string);
  string DELIM = "-";
  /*****************************/  
 
  struct yg_terminal {
    mapping (string => string)  map_yg_terminal;
  }
  //mapping (uint256 => yg_terminal) yg_terminal_array;
  yg_terminal[1000][1000] yg_terminal_array;
  uint public map_yg_terminal_index;
  uint public array_yg_terminal_index;
   /*
  struct transaction_record {
    mapping (string => string)  map_transaction_record;
  }
  mapping (uint256 => transaction_record) transaction_record_array;
  uint public map_transaction_record_index;
  
  struct recharge_record {
    mapping (string => string)  map_recharge_record;
  }
  mapping (uint256 => recharge_record) recharge_record_array;
  uint public map_recharge_record_index;
  
  struct big_warehouse {
    mapping (string => string)  map_big_warehouse;
  }
  mapping (uint256 => big_warehouse) big_warehouse_array;
  uint public map_big_warehouse_index;
  
  struct cloud_warehouse {
    mapping (string => string)  map_cloud_warehouse;
  }
  mapping (uint256 => cloud_warehouse) cloud_warehouse_array;
  uint public map_cloud_warehouse_index;
  
  struct trading_area {
    mapping (string => string)  map_trading_area;
  }
  mapping (uint256 => trading_area) trading_area_array;
  uint public map_trading_area_index;
  
  struct sowntown {
    mapping (string => string)  map_sowntown;
  }
  mapping (uint256 => sowntown) sowntown_array;
  uint public map_sowntown_index;
  
  struct building {
    mapping (string => string)  map_building;
  }
  mapping (uint256 => building) building_array;
  uint public map_building_index;
  
  */
  /******************************/
  //uint public bcd;
  //string [10] public ccc ;
  
  
   function atoi (string ss) public pure returns (uint){
        uint number;
        uint num;
        strings.slice memory s = ss.toSlice();
        strings.slice memory delim = ".".toSlice();
        string[] memory parts = new string[](s.count(delim) + 1);
        for(uint i = 0; i < parts.length; i++) {
            parts[i] = s.split(delim).toString();
            
            if (parts[i].toSlice().equals("0".toSlice())) {
                     num = 0;
            }else if (parts[i].toSlice().equals("1".toSlice())) {
                     num = 1;
            }else if (parts[i].toSlice().equals("2".toSlice())) {
                     num = 2;
            }else if (parts[i].toSlice().equals("3".toSlice())){
                     num = 3;
            }else if (parts[i].toSlice().equals("4".toSlice())){
                     num = 4;
            }else if (parts[i].toSlice().equals("5".toSlice())){
                     num = 5;
            }else if (parts[i].toSlice().equals("6".toSlice())){
                     num = 6;
            }else if (parts[i].toSlice().equals("7".toSlice())){
                     num = 7;
            }else if (parts[i].toSlice().equals("8".toSlice())){
                     num = 8;
            }else if (parts[i].toSlice().equals("9".toSlice())){
                     num = 9;
            }else{
                return 0;
            }
            
            //ss = s.toSlice().concat(ss.toSlice());
            number = number * 10 + num;
            //emit display2(number, parts[i], i);
        }
        //emit display2(number, parts[i], i);
        return number;
    }
  
  
    function itoa (uint256 num) pure public returns (string )  {
        string memory ss;
        string memory s;
        uint256 remainder;
        uint256 number = num;
        if (number == 0) {
            return "0";
        }
        while (number > 0){
            remainder = number % 10; 
            if (remainder == 0) {
                s = "0";
            }else if (remainder == 1) {
                     s = "1";
            }else if (remainder == 2) {
                     s = "2";
            }else if (remainder == 3){
                     s = "3";
            }else if (remainder == 4){
                     s = "4";
            }else if (remainder == 5){
                     s = "5";
            }else if (remainder == 6){
                     s = "6";
            }else if (remainder == 7){
                     s = "7";
            }else if (remainder == 8){
                     s = "8";
            }else if (remainder == 9){
                     s = "9";
            }else{
                return "err";
            }
            ss = s.toSlice().concat(ss.toSlice());
       
            number = number / 10;
        } 
        return ss;
    }
    function getAndIncIndex () public {
        if (map_yg_terminal_index >= 999) {
            array_yg_terminal_index++;
            map_yg_terminal_index = 0;
        }else{
            map_yg_terminal_index++;
        }
        
    }
    //function itoa (uint256 num) pure public returns (string ) 
    function update(string table, uint256 mapIndex, uint256 arrayIndex, string ss)   public returns(string) {
    //var ss = "a-10-b-20-c-30"; 
        string memory list;
        strings.slice memory s = ss.toSlice();
        strings.slice memory delim = DELIM.toSlice();
        string[] memory parts = new string[](s.count(delim) + 1);
   
 
        
        for(uint i = 0; i < parts.length; i++) {
            parts[i] = s.split(delim).toString();
            if (i > 0 && i % 2 != 0) {
                if (i != 1) {
                    list = list.toSlice().concat(DELIM.toSlice());  //Add "DELIM" before columns execept the first one
                }
                yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal[parts[i-1]] = parts[i]; // create map key=a value=10
                //emit item(parts[i-1], yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal[parts[i-1]]);
                list = list.toSlice().concat(parts[i-1].toSlice()); // Add column into "list" varible
                //emit list1("list", parts[i-1]);
            }
        }
        yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal["list"] = list; // Add "list" key into map
        yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal["table"] = table; // Add "table" key into map
        //emit inserted("list", yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal["list"]);
        emit inserted("result:", yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal["table"],
                               yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal["a"],
                               yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal["b"],
                               yg_terminal_array[array_yg_terminal_index][map_yg_terminal_index].map_yg_terminal["c"]);
        
        if ((mapIndex == 0) && (arrayIndex == 0)) {
            getAndIncIndex();
            //emit insertEvent(table, arrayIndex, mapIndex, ss);
        } else {
            //emit updateEvent(table, arrayIndex, mapIndex, ss);
        }
        
        return ss;
    }
/*
    function searchRange(string table, string key, uint256 smallValue, uint256 bigValue) view  public returns (string )  {
        string memory  results; 

        for (uint256 i = 0; i < map_yg_terminal_index; i++) {
            if (table.toSlice().equals(yg_terminal_array[i].map_yg_terminal["table"].toSlice())  // find a record table = table and key= key value = value
                    && (smallValue <= atoi(yg_terminal_array[i].map_yg_terminal[key]))
                    && (bigValue >=atoi(yg_terminal_array[i].map_yg_terminal[key]))) {
                        
                strings.slice memory  s = yg_terminal_array[i].map_yg_terminal["list"].toSlice();
                strings.slice memory  delim = DELIM.toSlice();
                string[] memory parts = new  string[](s.count(delim) + 1);
                for(uint n = 0; n < parts.length; n++) {
                    parts[n] = s.split(delim).toString();
                }
                //emit display1("a", yg_terminal_array[i].map_yg_terminal["a"], "b", yg_terminal_array[i].map_yg_terminal["b"], "c", yg_terminal_array[i].map_yg_terminal["c"]);
                results = results.toSlice().concat(table.toSlice());
                results = results.toSlice().concat(DELIM.toSlice());
                results = results.toSlice().concat(itoa(i).toSlice());
                results = results.toSlice().concat(DELIM.toSlice());
                for(n = 0; n < parts.length; n++) {
                    //parts[n] = s.split(delim).toString();
                    //emit searchFound(parts[n], yg_terminal_array[i].map_yg_terminal[parts[n]]);
                    results = results.toSlice().concat(parts[n].toSlice());
                    results = results.toSlice().concat(DELIM.toSlice());
                    results = results.toSlice().concat(yg_terminal_array[i].map_yg_terminal[parts[n]].toSlice());
                    results = results.toSlice().concat(DELIM.toSlice());
                }
                results = results.toSlice().concat(DELIM.toSlice()); //Add one "DELIM" after a record
            }
        }
        //abc = results;
        //emit searchResult("results", results);
        return results;
    }
*/
    function search(string table, string key, string value) view  public returns (string )  {
        string memory  results; 

        for (uint256 i = 0; i <= array_yg_terminal_index; i++) {
            for (uint256 j = 0; j < 1000; j++) {
                 //emit display1("a", yg_terminal_array[0][1].map_yg_terminal["a"], "b", yg_terminal_array[i][j].map_yg_terminal["b"], "c", yg_terminal_array[i][j].map_yg_terminal["c"]);
                if (table.toSlice().equals(yg_terminal_array[i][j].map_yg_terminal["table"].toSlice())  // find a record table = table and key= key value = value
                    && value.toSlice().equals(yg_terminal_array[i][j].map_yg_terminal[key].toSlice())) {
                    strings.slice memory s = yg_terminal_array[i][j].map_yg_terminal["list"].toSlice();
                    strings.slice memory delim = DELIM.toSlice();
                    string[] memory parts = new string[](s.count(delim) + 1);
                    for(uint n = 0; n < parts.length; n++) {
                        parts[n] = s.split(delim).toString();
                    }
                   // emit display1("a", yg_terminal_array[i][j].map_yg_terminal["a"], "b", yg_terminal_array[i][j].map_yg_terminal["b"], "c", yg_terminal_array[i][j].map_yg_terminal["c"]);
                    results = results.toSlice().concat(table.toSlice());
                    results = results.toSlice().concat(DELIM.toSlice());
                    results = results.toSlice().concat(itoa(i).toSlice());
                    results = results.toSlice().concat(DELIM.toSlice());
 		    results = results.toSlice().concat(itoa(j).toSlice());
                    results = results.toSlice().concat(DELIM.toSlice());

                    for(n = 0; n < parts.length; n++) {
                        //parts[n] = s.split(delim).toString();
                        //emit searchFound(parts[n], yg_terminal_array[i].map_yg_terminal[parts[n]]);
    
                        results = results.toSlice().concat(parts[n].toSlice());
                        results = results.toSlice().concat(DELIM.toSlice());
                        results = results.toSlice().concat(yg_terminal_array[i][j].map_yg_terminal[parts[n]].toSlice());
                        results = results.toSlice().concat(DELIM.toSlice());
                    }
                    results = results.toSlice().concat(DELIM.toSlice()); //Add one "DELIM" after a record
                }
            }
        }
        //abc = results;
       // emit searchResult("results", results);
        return results;
    }
}




