package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomUsage() {
	randomizer := rand.New(rand.NewSource(10))
	fmt.Println(randomizer.Int())
	fmt.Println(randomizer.Int())
	fmt.Println(randomizer.Int())
}

func uniqueSeed() {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println(randomizer.Int())
	fmt.Println(randomizer.Float32())
	fmt.Println(randomizer.Uint32())
}

func randomIndex() {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println(randomizer.Intn(10))
}

func randomString(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}
