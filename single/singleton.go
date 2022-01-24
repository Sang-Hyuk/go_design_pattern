package main

import (
	"fmt"
	"sync"
)

// package level variable
var instance *component
var once sync.Once
var mx sync.RWMutex

type component struct {
	count int
}

func GetInstance() *component {
	once.Do(func() {
		fmt.Println("인스턴스 생성.")
		instance = &component{count: 0}
	})
	return instance
}

func (s *component) add() {
	mx.Lock()
	defer mx.Unlock()
	s.count++
}

func (s *component) getCount() int {
	mx.RLock()
	defer mx.RUnlock()
	return s.count
}

func main() {
	GetInstance().add()
	fmt.Println(GetInstance().getCount())

	GetInstance().add()
	fmt.Println(GetInstance().getCount())

	GetInstance().add()
	fmt.Println(GetInstance().getCount())

	wait := new(sync.WaitGroup)
	count := 100
	wait.Add(count)

	for i := 0; i < count; i++ {
		go func(idx int) {
			GetInstance().add()
			fmt.Printf("[%d] go routine : instance count(%d)\n", idx, GetInstance().getCount())
			defer wait.Done()
		}(i)
	}

	wait.Wait()
}
