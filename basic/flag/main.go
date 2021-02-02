package main

import (
	"flag"
	"fmt"
	"github.com/mreiferson/go-options"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Option struct {
	Count         int64 `flag:"count"`
	InitialFields map[string]interface{}
}

var (
	_initialFields = multiFields{}
)

type multiFields map[string]interface{}

// String return fields describe string
func (mf *multiFields) String() string {
	fields := make([]string, len(*mf))
	i := 0
	for k, v := range *mf {
		fields[i] = k + fmt.Sprintf(": %v", v)
		i++
	}
	return strings.Join(fields, ",")
}

func (mf *multiFields) Set(fields string) error {
	temp := strings.Split(fields, ",")
	for _, s := range temp {
		field := strings.Split(s, ":")
		if len(field) == 2 {
			(*mf)[field[0]] = field[1]
		}
	}
	return nil
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
	opts.InitialFields = _initialFields
	options.Resolve(opts, flagSet, nil)
	count := opts.Count
	fmt.Println(count, *opts)
}

func flagSet(opts *Option) *flag.FlagSet {
	flagSet := flag.NewFlagSet("rpc test", flag.ExitOnError)
	flagSet.Var(&_initialFields, "initial-fields", "")
	fmt.Println(*flagSet.Int64("count", opts.Count, "run count"))
	return flagSet
}
