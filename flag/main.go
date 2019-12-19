package main

import (
	"flag"
	"fmt"
	"github.com/mreiferson/go-options"
	"math/rand"
	"os"
	"time"
)

type Option struct {
	Count int64 `flag:"count"`
}

func NewOption() *Option {
	return &Option{
		//Count: 10000,
	}
}

func main() {
	opts := NewOption()
	flagSet := flagSet(opts)
	flagSet.Parse(os.Args[1:])
	// NOTE: clear all args after we have parsed, avoid RPCServer parse these args again
	os.Args = os.Args[:1]
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(*flagSet)
	options.Resolve(opts, flagSet, nil)
	count := opts.Count
	fmt.Println(count, opts)
}

func flagSet(opts *Option) *flag.FlagSet {
	flagSet := flag.NewFlagSet("rpc test", flag.ExitOnError)
	fmt.Println(*flagSet.Int64("count", opts.Count, "run count"))
	return flagSet
}
