服务端证书生成
生成私钥
openssl genrsa -out server.key 2048
生成证书
openssl req -new -x509 -key server.key -out server.pem -days 3650

或者使用
go run $GOROOT/src/crypto/tls/generste_cert.go -host localhost


客户端证书生成
生成私钥
openssl genrsa -out client.key 2048
生成证书
openssl req -new -x509 -key client.key -out client.pem -days 3650











