package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	empieza := time.Now()

	apis := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/users",
		"https://management.azure.com",
		"https://graph.microsoft.com",
		"https://dev.azure.com",
		"https://api.github.com",
	}

	for _, api := range apis {
		go revisarApis(api)
	}

	time.Sleep(1 * time.Second)

	transcurrido := time.Since(empieza)
	fmt.Printf("Tiempo transcurrido: %.2f segundos\n", transcurrido.Seconds())
}

func revisarApis(api string) {
	resp, err := http.Get(api)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error!!!! ", api, err)
	} else {
		fmt.Println("Correcto!!! ", api)

	}
}
