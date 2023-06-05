package main

import (
	"fmt"

	"github.com/rovezuka/go_module/storage"
)

func main() {
	ms := storage.NewMemoryStorage()
	spawnEmployees(ms)
	fmt.Println(ms.Get(3))
	fmt.Println(ms)
}

func spawnEmployees(s storage.Storage) {
	for i := 1; i <= 10; i++ {
		s.Insert(storage.Employee{Id: i})
	}
}
