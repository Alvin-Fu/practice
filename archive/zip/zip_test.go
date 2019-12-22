package zip

import "testing"


func TestUnCompressDir(t *testing.T) {
	UnCompressDir("../../../rpcx/", "rpcx.zip")
}
func TestCompress(t *testing.T) {
	Compress("rpcx.zip", "rpcx/")
}