package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Store struct {
	stock      int
	mutex      sync.Mutex
	useLocking bool
	useDelay   bool
}

func NewStore(initialStock int, useLocking, useDelay bool) Store {
	return Store{stock: initialStock, useLocking: useLocking, useDelay: useDelay}
}

func (s *Store) GetStock() int {
	return s.stock
}

func (s *Store) Buy(qty int) {
	if s.useLocking {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}

	if s.useDelay {
		time.Sleep(time.Duration(getRandomValue()) * time.Millisecond)
	}

	currentStock := s.stock
	s.stock = currentStock - qty
}

func main() {

	s := NewStore(1000, true, false)

	wg := sync.WaitGroup{}

	length := s.GetStock()

	wg.Add(length)

	for i := 0; i < length; i++ {

		go func(i int, s *Store) {

			s.Buy(1)

			wg.Done()

		}(i, &s)

	}

	wg.Wait()

	fmt.Printf(">>>> %v\n", s.GetStock())

}

const (
	min = 1
	max = 1
)

func getRandomValue() int {
	return rand.Intn(max-min+1) + min
}
