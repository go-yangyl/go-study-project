package singleton

import "sync"

type Singleton struct {
}

var once sync.Once

func GetSingleton() (singleton *Singleton) {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
