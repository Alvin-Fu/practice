package singleton

import "sync"

type Singleton struct {
	Str string
}

func newSingleton()*Singleton{
	return &Singleton{}
}

var singleton *Singleton
var once sync.Once
func GetSingleton()*Singleton{
	once.Do(func() {
		singleton = newSingleton()
	})
	return singleton
}