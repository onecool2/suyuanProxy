rm -rf *.abi
rm -rf suyuanContract.go
./solc-static-linux suyuan.solc --abi -o ./
abigen --abi=suyuanContract.abi --pkg=suyuanContract --out=suyuanContract.go
