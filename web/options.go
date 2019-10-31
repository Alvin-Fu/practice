package web

type Option struct {
	ID         int64
	ConfigFile string
}

func NewOption() *Option {
	return &Option{}
}
