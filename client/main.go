package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	c := http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: 5,
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			log.Printf("making request #%d\n", i)

			resp, err := c.Get("http://localhost:8088/hello")
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("receiving response for request #%d\n", i)

			resp.Body.Close()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
