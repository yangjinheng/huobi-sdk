package utils

// AK/SK
const (
	Endpoint  = "https://api.huobi.pro"
	ProxyAddr = "http://127.0.0.1:9000"
	AccessKey = "b0dab8f1-c6073ad0-a7e12771-b1rkuf4drg"
	SecretKey = "9c9ed648-5a9420b1-0564382d-4a775"
)

// var
var (
	ReqBuilder Signer
)

func init() {
	ReqBuilder = Signer{
		Endpoint:  Endpoint,
		AccessKey: AccessKey,
		SecretKey: SecretKey,
	}
}
