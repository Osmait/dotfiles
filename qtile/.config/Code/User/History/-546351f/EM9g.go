package concurrent

import (
	"fmt"
	"net/http"
)

func Check(c chan string, endpoints ...string) {

	numResponses := make(chan int, 1)

	// Enviar el número de respuestas esperadas
	numResponses <- len(endpoints)

	for _, endpoint := range endpoints {

		go func(endpoint string, c chan string) {
			rep, err := http.Get(endpoint)
			if err != nil {
				c <- fmt.Sprintf("error checking %s: %v", endpoint, err)
				return
			}

			c <- rep.Status

		}(endpoint, c)
		numResponses <- <-numResponses - 1

	}

}

func Check2(endpoints ...string) {
	for _, endpoint := range endpoints {

		func(endpoint string) {
			rep, err := http.Get(endpoint)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(rep.Status)

		}(endpoint)

	}

}
