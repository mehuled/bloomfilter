package main

import (
	"fmt"
	"log"
)

type BloomFilter struct {
	bitmask       uint64
	hashFunctions []func(key string) uint64
}

func New() *BloomFilter {
	return &BloomFilter{
		bitmask:       0,
		hashFunctions: []func(key string) uint64{SimpleHashFunc},
	}
}

func (bf *BloomFilter) Add(key string) error {
	for _, fn := range bf.hashFunctions {
		bit := fn(key)
		if bit < 0 || bit > 63 {
			return fmt.Errorf("invalid value for bitmask")
		}
		bf.bitmask = bf.bitmask | (1 << bit)
	}
	return nil
}

func (bf *BloomFilter) Has(key string) bool {
	for _, fn := range bf.hashFunctions {
		bit := fn(key)
		if bf.bitmask&(1<<bit) == 0 {
			return false
		}
	}

	return true
}

func (bf *BloomFilter) Print() {
	log.Println(fmt.Sprintf("bitmask is : %b", bf.bitmask))
}

func SimpleHashFunc(key string) uint64 {
	r := []rune(key)
	res := r[0] % 64
	return uint64(res)
}
