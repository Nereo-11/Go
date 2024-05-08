package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	/*
		canal := make(chan int)
		canal <- 15  //Para enviar datos es con esa flechita
		valor := <- canal // para recibir datos se declara asi
	*/
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)

	for _, api := range apis {
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch) // cantidad de apis que se reciben
	}

	time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Tomo %v segundos!\n", elapsed.Seconds())
}

func checkApi(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error: %s está caído \n", api)
		return
	}

	ch <- fmt.Sprintf("Correcto: %s está funcionando \n", api)
}
