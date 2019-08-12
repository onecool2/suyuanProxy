// FILE: compileDeploy.js
console.log('Setting up...');
const fs = require ('fs');
const solc = require ('solc');
const Web3 = require ('web3');
const web3 = new Web3(new Web3.providers.HttpProvider("http://39.100.66.220:32000"));
console.log('Reading Contract...');
const input = fs.readFileSync('../suyuan.solc');
console.log('Compiling Contract...');
const output = solc.compile(input.toString(), 1);
const bytecode = output.contracts[':suyuanContract'].bytecode;
const abi = output.contracts[':suyuanContract'].interface;
//Compile contract
console.log('abi=' +  abi)
account=web3.eth.accounts[0];
var web3_version = "0.20"
console.log(web3_version != '1.0')
if (web3_version != 1.0) {
	//Contract Object for web3.js <1.0
	const helloWorldContract =  web3.eth.contract(JSON.parse(abi));
	console.log('unlocking Coinbase account');
	const password = "abc";
	try { 
	    web3.personal.unlockAccount(account, password, 00);
} catch(e) {
	console.log(e);
	return;
}

console.log("Deploying the contract");
//var contractInstance = MyContract.new([constructorParam1] [, constructorParam2], {data: '0x12345...', from: myAccount, gas: 1000000});
    var helloWorldContractInstance = helloWorldContract.new({
        data: '0x' + bytecode,
        from: account,
        gas: 6702292, 
        //gasPrice: "111111111111111111" //For quorum chain the gasPrice must be set 0
    })
//.then(function(newContractInstance){
    console.log(helloWorldContractInstance.transactionHash) // instance with the new contract address
//		});

}else{
        console.log('web version = 1.0')
	//Contract Object for web3js > 1.0

	const helloWorldContract = new web3.eth.Contract(JSON.parse(abi)) //web3.eth.Contract(JSON.parse(abi));
	console.log('unlocking Coinbase account');
	const password = "abc";
	try {
		web3.eth.personal.unlockAccount("0x27a29f6cc38e2f1dac2e7fe8ce59bab2fc47148d", password, 600);
	} catch(e) {
		console.log(e);
		return;
	}

	console.log("Deploying the contract");
	const helloWorldContractInstance = helloWorldContract.deploy({
data: '0x' + bytecode,
})
.send({
from: "0x27a29f6cc38e2f1dac2e7fe8ce59bab2fc47148d",
gas: 8000000, 
//gasPrice: "111111111111111111" //For quorum chain the gasPrice must be set 0
})
.then(function(newContractInstance){
		console.log(newContractInstance.options.address) // instance with the new contract address
		});
}
