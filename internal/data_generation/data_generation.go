package main

import (
	"encoding/json"
	"main/internal/types"
	"math/rand/v2"
	"os"
	"strconv"

	"github.com/brianvoe/gofakeit/v7"
)

func generateOrder() types.Order {
	var order types.Order
	err := gofakeit.Struct(&order)
	if err != nil {
		panic(err)
	}
	return order
}

func main() {
	o := generateOrder()
	data, err := json.Marshal(o)

	if err != nil {
		panic(err)
	}
	f, err := os.Create("test/testdata/" + "model" + strconv.Itoa(int(rand.Int32()%100)) + string(".json"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(data)
}
