pragma solidity ^0.4.25;
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

  event searchResult (string key,string value);
  event searchFound (string key,string value);
  event insertEvent (uint256 index, string ss);
  event updateEvent (string table, uint256 index, string ss);
  event item (string key,string value);
  event list1 (string key,string value);
  event inserted (string listname, string list);

  string DELIM = "-";
  /*****************************/

  struct yg_terminal {
    mapping (string => string)  map_yg;
  }
  string[1000000000] yg_array;
  //mapping (uint256 => yg_terminal) yg_array;

  uint public yg_index;
 

    function update(string ss) onlyWriter public returns(string) {
    //var ss = "a-10-b-20-c-30";
        //string memory list;
        //strings.slice memory s = ss.toSlice();
        //strings.slice memory delim = DELIM.toSlice();
        //string[] memory parts = new string[](s.count(delim) + 1);

        
        //uint256 index = yg_index;
    
        emit insertEvent(yg_index, ss);
     
       /* for(uint i = 0; i < parts.length; i++) {
            parts[i] = s.split(delim).toString();
            if (i > 0 && i % 2 != 0) {
                if (i != 1) {
                    list = list.toSlice().concat(DELIM.toSlice());  //Add "DELIM" before columns execept the first one
                }
                yg_array[index].map_yg[parts[i-1]] = parts[i]; // create map key=a value=10
                emit item(parts[i-1], yg_array[index].map_yg[parts[i-1]]);
                list = list.toSlice().concat(parts[i-1].toSlice()); // Add column into "list" varible
                emit list1("list", parts[i-1]);
            }
        }*/
        yg_array[yg_index] = ss; // Add "list" key into map
        yg_index++;
        //yg_array[index].map_yg["table"] = table; // Add "table" key into map
        //emit inserted("list", yg_array[index].map_yg["list"]);
        //emit inserted("table", yg_array[index].map_yg["table"]);
        return ss;
    }
}

