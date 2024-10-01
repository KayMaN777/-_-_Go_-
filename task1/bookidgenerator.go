package main

import (
	"hash/fnv"
)

type BookIdGenerator interface {
	genID(title string) uint32
}

type FnvGenerator struct{}

func (fnvHash *FnvGenerator) genID(title string) uint32 {
	generator := fnv.New32a()
	generator.Write([]byte(title))
	return generator.Sum32()
}
