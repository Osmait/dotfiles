package concurrent

import (
	"fmt"
	"net/http"
)

func Check(c chan string, endpoints ...string) {
	for _, endpoint := range endpoints {
		c <- endpoint
		go func(endpoint string, c chan string) {
			rep, err := http.Get(<-c)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(rep.Status)

		}(endpoint, c)

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
