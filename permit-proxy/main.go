package main

import (
	"fmt"
	"log"

	"permit-proxy/internal/store"
)

func main() {
	p := store.New()
	v, err := p.Get()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(v))
	fmt.Printf("%+v\n", v[7])
}
