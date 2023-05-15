package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	i "github.com/jnguyen0220/here-we-go/intro"
)

func main() {
	// 2. Hello World
	fmt.Println("Hello World")

	// 3. Custome Function
	i.Welcome()

	// 4. math/rand
	fmt.Println(rand.Intn(100))

	// 6. Http Server
	port := 8080
	http.HandleFunc("/", handler)
	fmt.Printf("Http server running on :%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
