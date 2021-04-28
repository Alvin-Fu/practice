package proxy

import "testing"

func TestProxy(t *testing.T) {
	p := NewProxySubject()
	p.DoSomething()
}
