package main

import (
        "fmt"
	"log"
	"context"
	"strings"
	//"strconv"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
	"encoding/json"
	"net/http"
	"math/big"
	"io/ioutil"
	"crypto/ecdsa"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	suyuanContract "github.com/onecool2/suyuanProxy/contract"
)

type Person struct {
        Name string
        Phone string
}

type SyItem struct {
	Qrcode     *big.Int
	Fn_name    *big.Int
	Infos []map[string]interface{}
	Images []map[int]interface{}
	Blocknumber *big.Int;
}
type Sy struct {
	SySlice []SyItem
	Index	*big.Int;
	BlockNum *big.Int;
        CurrentHash    string;
	PreviousHash   string;

}
type Block struct {
	Id_  bson.ObjectId `bson:"_id"`
        Number int64
	Hash string
	ParentHash string `bson:"parentHash"`
}

type Transaction struct {
	_id string
	Hash string
	BlockHash string
	BlockNumber int `bson:"blockNumber"`
	CumulativeGasUsed string
	Fn_name int64
	From string
	Qr_code int64
	Gas string
	GasPrice string
	GasUsed string
	Input string
	Log string
	LogBloom string
	Nonce string
	R string
	S string
	Params string
	To string
	TransactionIndex string
	V string
	Value string
	Images string
}
type RetMsg struct {
	code    int    //400
	message string //"具体错误信息",
	data    string //"data":null
}
const (
	RPC_HOST           string = "http://quorumrpc.suyuan.svc.cluster.local:22000"
	WS_HOST            string = "ws://127.0.0.1:32002"
	CONTRACT_ADDRESS   string = "0x1b8d742a7a45364ba6b9132d460b814d0fc43722"
	OWNER_PUBLIC_KEY   string = "0x25cde39d96684e2a681ae0289b37af8e9859ed99"
	OWNER_PRIVATE_KEY  string = "38cd3eef7f9040c6c1b3d1dc203c7070196787ade42877d7cd48df05c5158809"
	EMPTY              int    = 0
	READY              int    = 1
	SENT               int    = 2
	MAX_ON_CHAIN_DELAY        = "-30s"
)

var (
	EthClient  *ethclient.Client
	Contract   *suyuanContract.SuyuanContract
	privateKey *ecdsa.PrivateKey
	publicKey  common.Address
	owner      common.Address
)

func init() {
	var contractAddress common.Address
	var err error
	EthClient, err = ethclient.Dial(RPC_HOST)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress = common.HexToAddress(CONTRACT_ADDRESS)
	Contract, err = suyuanContract.NewSuyuanContract(contractAddress, EthClient)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = crypto.HexToECDSA(OWNER_PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	owner = crypto.PubkeyToAddress(*publicKeyECDSA)
}
/*
func getInfo(Data map[string]interface{}, SY *Sy) {

	fmt.Println("getInfo:", Data)
	_ = Contract
	var err error
	fmt.Println("action:", Data["qrcode"])

//	qrcode, err := strconv.ParseInt(Data["qrcode"].(string), 10, 64) 
	qrcode := big.NewInt(int64(Data["qrcode"].(float64)))

	var fn_name int64
	var Infos string
	var Images string
	for fn_name = 0; fn_name < 20; fn_name++ {
	var sy SyItem;
		sy.Qrcode, sy.Fn_name, Infos, sy.Blocknumber, Images, err = Contract.GetInfo(&bind.CallOpts{}, qrcode, big.NewInt(fn_name))
		if (sy.Blocknumber.Uint64() > 0){
		    kAndV := strings.Split(Infos, "^")
 		    for i:=0; i < len(kAndV); i++ {
		        kv := strings.Split(kAndV[i], "|")
		        m := make(map[string]interface{})
		        fmt.Println("kv:", kv[0], m[kv[0]])
	                m[kv[0]] = kv[1]
			sy.Infos = append(sy.Infos, m)
		    }
		    kAndVImages := strings.Split(Images, "^")
 		    for i:=0; i < len(kAndVImages); i++ {
	                m := make(map[int]interface{})
			m[i] = kAndVImages[i]
			sy.Images = append(sy.Images, m)
		        fmt.Println("images:", i, m[i])
		    }

		    fmt.Println("--------------", sy)
		    fmt.Println("++++++++++++++")
		    SY.SySlice = append(SY.SySlice, sy)
		}else{
		    break;
		}
	}
	SY.BlockNum, SY.CurrentHash, SY.PreviousHash, err = Contract.GetBlock(&bind.CallOpts{})
       
	if err != nil {
	    glog.Warning("getBasic error:", err)
	}
}
*/
func getInfoFromDB(Data map[string]interface{}, SY *Sy) {

	fmt.Println("getInfo:", Data)
	_ = Contract
	var err error
	fmt.Println("action:", Data["qrcode"])
	
	result := []Transaction{}
	block := []Block{}
	qrcode := int64(Data["qrcode"].(float64))
        err = c.Find(bson.M{"qr_code": qrcode}).All(&result)
        if err != nil {
                log.Fatal(err)
        }

	fmt.Println("result", len(result))
	for i:=0; i< len(result); i++ {
            var sy SyItem;
	    sy.Qrcode = big.NewInt(result[i].Qr_code)
            sy.Fn_name = big.NewInt(result[i].Fn_name)
            sy.Blocknumber = big.NewInt(int64(result[i].BlockNumber))
            fmt.Println("BlockNumber1111111", result[i].BlockNumber)
	    kAndV := strings.Split(result[i].Params, "^")
	    for i:=0; i < len(kAndV); i++ {
	        kv := strings.Split(kAndV[i], "|")
	        m := make(map[string]interface{})
	        fmt.Println("kv:", kv[0], m[kv[0]])
                m[kv[0]] = kv[1]
		sy.Infos = append(sy.Infos, m)
	    }

            kAndVImages := strings.Split(result[i].Images, "^")

 	    for i:=0; i < len(kAndVImages); i++ {
	        m := make(map[int]interface{})
		m[i] = kAndVImages[i]
		sy.Images = append(sy.Images, m)
		fmt.Println("images:", i, m[i])
	    }
	    fmt.Println("--------------", sy)
	    fmt.Println("++++++++++++++")
	    SY.SySlice = append(SY.SySlice, sy)

	}
        err = b.Find(nil).Limit(1).Sort("-_id").All(&block)
        if err != nil {
                log.Fatal("read db err:", err)
        }
	fmt.Println("blocknum", block[0].Number, "Hash", block[0].Hash, "ParentHash", block[0].ParentHash)
	fmt.Println("block", block[0])
	SY.BlockNum = big.NewInt(block[0].Number)
	SY.CurrentHash =block[0].Hash
	SY.PreviousHash = block[0].ParentHash
	//SY.CurrentHash, CurrentHash[:32]
	//SY.PreviousHash = PreviousHash[:32]

	if err != nil {
	    glog.Warning("getBasic error:", err)
	}
}


func setInfo(Data map[string]interface{}) {

	fmt.Println("setBasic:", Data)
	_ = Contract

	nonce, err := EthClient.PendingNonceAt(context.Background(), owner)
	if err != nil {
		log.Fatal(err)
	}

	var tx *types.Transaction 
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(600000) // in units
	action := Data["action"]

	if (action == "setInfo"){
		qrcode := big.NewInt(int64(Data["qrcode"].(float64)))
		fn_name := big.NewInt(int64(Data["fn_name"].(float64)))
		images := Data["images"].(string)
		infos := Data["infos"].(string)
		tx, err = Contract.SetInfo(auth, qrcode, fn_name, infos, images)


	if err != nil {
		glog.Warning("send tx:", err)
	}
	fmt.Println("###########:", tx.Hash())
	//fmt.Println("send token:%s", ss)
	//fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	//return tx.Hash()
	//result, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0xfA02a776BB22cc644AE4d78EC348702bFB5D927A"))
	//  if err != nil {
	//      log.Fatal(err)
	//  }
	//  fmt.Println("result", result) // "1.0"
	}
}

func replyBlock(writer http.ResponseWriter, code int, block *big.Int, data string) {

	if bs, err := json.Marshal(block); err == nil {
		writer.Header().Set("Content-type", "application/x-www-form-urlencoded")
		writer.WriteHeader(200)
		writer.Write([]byte(bs))

		fmt.Println(bs)
	}

}

func replyMsg(writer http.ResponseWriter, code int, sy *Sy, data string) {

	//retMsg := RetMsg{code, message, data}
	if bs, err := json.Marshal(sy); err == nil {
		//req := bytes.NewBuffer([]byte(bs))
		//body_type := "application/x-www-form-urlencoded"
		//writer.Header().Set("Content-type", "application/text")
		writer.Header().Set("Content-type", "application/x-www-form-urlencoded")
		writer.WriteHeader(200)
		writer.Write([]byte(bs))

		//fmt.Fprintln(writer, bs)
		/*resp, _ := */ //http.Post(ZUI_RI_SERVER_HOST, body_type, req)
		//fmt.Println(bs)
	}

}
/*
func generateQRCode(writer http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	body_str := string(body)
	fmt.Println(body_str)
      
	var action map[string]interface{}
	if err := json.Unmarshal(body, &action); err == nil {
		fmt.Println("action:", action["action"])
		if (action["action"] == "generateQRCode"){
		    replyMsg(writer, 200, , "ok")
		    fmt.Println("ok:", err)
		}else{
		    replyMsg(writer, 400, err.Error(), err.Error())
		}
	}else{
		fmt.Println("Unmarshal:", err)
	}
	fmt.Println("end generateQRCode")
}
*/
func handlerHttp(writer http.ResponseWriter, request *http.Request){
	body, _ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	body_str := string(body)
	fmt.Println(body_str)
      
/**************************************************/
	var reqBody map[string]interface{}
	var sy Sy
	if err := json.Unmarshal(body, &reqBody); err == nil {
		fmt.Println("action:", reqBody["action"])
		if (reqBody["action"] == "generateQRCode"){
		    fmt.Println("ok:", err)
		}else if(reqBody["action"] == "setInfo"){
		    setInfo(reqBody)
		}else if(reqBody["action"] == "getInfo"){
		    getInfoFromDB(reqBody, &sy)
	            //fmt.Println("sy:", sy)
		    replyMsg(writer, 200, &sy, "ok")
		}else{
	            fmt.Println("Unmarshal:", err)
		    replyMsg(writer, 400, &sy, err.Error())
		}
	}
	fmt.Println("end handlerHttp")


}
var c *mgo.Collection
var b *mgo.Collection
func main() {
/************************** connect mongoDB ***************************************/

        session, err := mgo.Dial("quorum-mongodb.suyuan.svc.cluster.local")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c = session.DB("quorum").C("transactions")
        b = session.DB("quorum").C("blocks")
        
/*************************************************************************/

/************************* start http server *****************************************/
	r := mux.NewRouter()
	r.HandleFunc("/generateQRCode", handlerHttp)
	http.Handle("/generateQRCode", r)
	r.HandleFunc("/huoyanjing/setInfo", handlerHttp)
	http.Handle("/huoyanjing/setInfo", r)
	r.HandleFunc("/huoyanjing/getInfo", handlerHttp)
	http.Handle("/huoyanjing/getInfo", r)

	fmt.Println("Side Car proxy start listening on 3333...")
	http.ListenAndServe(":3333", r)
/*************************************************************************/
}
