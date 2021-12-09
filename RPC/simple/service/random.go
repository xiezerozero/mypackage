package service

import (
	"log"
	"math/rand"
)

type RandomService struct {
}

func (rs *RandomService) Rand(seed int64, result *int64) error {
	rand.Seed(seed)
	*result = rand.Int63()
	log.Println("new random result: ", *result)
	return nil
}
