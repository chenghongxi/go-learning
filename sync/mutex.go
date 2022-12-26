package main

import (
	"fmt"
	"sync"
	"time"
)

type Book struct {
	BookName string
	L        *sync.Mutex
}

func (bk *Book) SetName(wg *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("Unlock set name:", name)
		bk.L.Unlock()
		wg.Done()
	}()
	bk.L.Lock()
	fmt.Println("Lock set name:", name)
	time.Sleep(1 * time.Second)
	bk.BookName = name
}

func main() {
	bk := Book{}
	bk.L = new(sync.Mutex)
	wg := &sync.WaitGroup{}
	books := []string{"<三国演义>", "<道德经>", "<西游记>"}
	for _, book := range books {
		wg.Add(1)
		go bk.SetName(wg, book)
	}
	wg.Wait()
}
