package main

import (
	"log"
)

func main() {

	bf := New()

	key1 := "hello"
	key2 := "world"
	key3 := "non_existent"
	key4 := "h_non-existent"

	bf.Print()

	err := bf.Add(key1)
	if err != nil {
		panic(err)
	}
	bf.Print()

	err = bf.Add(key2)
	if err != nil {
		panic(err)
	}

	bf.Print()
	log.Printf(" key : %s exists : %v\n", key1, bf.Has(key1))
	log.Printf(" key : %s exists : %v\n", key2, bf.Has(key2))
	log.Printf(" key : %s exists : %v\n", key3, bf.Has(key3))
	log.Printf(" key : %s exists : %v\n", key4, bf.Has(key4))

}
