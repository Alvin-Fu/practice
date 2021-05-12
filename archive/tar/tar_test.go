package tar

import "testing"

func TestPack(t *testing.T) {
	err := Pack("../../../rpcx/", "rpcx.tar")
	if err != nil {
		t.Fatalf("pack err: %v", err)
	}
}

func TestUnpack(t *testing.T) {
	err := Unpack("rpcx.tar")
	if err != nil {
		t.Fatalf("unpack err: %v", err)
	}
}
