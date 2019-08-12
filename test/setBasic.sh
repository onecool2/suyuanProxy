curl -XPOST "http://127.0.0.1:3333/huoyanjing/setInfo" -H 'Content-Type: application/json' -d'
{
	"action":"getInfo",
	"qrcode":"2",
	"basic_address":"basic_address",
	"basic_picture":"basic_picture",
	"basic_name":"打打鸭脖"
}'
