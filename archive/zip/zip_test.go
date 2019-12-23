package zip

import "testing"

func TestUnCompressDir(t *testing.T) {
	err := UnCompress("../../../rpcx/", "rpcx.zip", "rpcx/")
	if err != nil {
		t.Fatalf("uncompress dir err: %v", err)
	}
}

func TestUnCompressFile(t *testing.T) {
	err := UnCompress("../readme.md", "read.zip", "read/")
	if err != nil {
		t.Fatalf("uncompress file err: %v", err)
	}
}

func TestCompressDir(t *testing.T) {
	err := Compress("rpcx.zip", "rpcx/")
	if err != nil {
		t.Fatalf("compress err: %v", err)
	}
}

func TestCompressFile(t *testing.T) {
	err := Compress("read.zip", "read/")
	if err != nil {
		t.Fatalf("compress err: %v", err)
	}
}
