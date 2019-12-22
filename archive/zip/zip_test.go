package zip

import "testing"


func TestUnCompressDir(t *testing.T) {
	err := UnCompressDir("../../../rpcx/", "rpcx.zip")
	if err != nil {
		t.Fatalf("uncompress err: %v", err)
	}
}
func TestCompress(t *testing.T) {
	err := Compress("rpcx.zip", "rpcx/")
	if err != nil {
		t.Fatalf("compress err: %v", err)
	}
}