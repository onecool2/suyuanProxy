curl -XPOST "http://127.0.0.1:3333/generateQRCode" -H 'Content-Type: application/json' -d'
{
	"action":"updateInfo",
	"qrcode":"123123",
	"checkpoint":"0",
	"producer":"aaaaaa",
	"time":"123123123132",
	"index":"123123123132"
}'
