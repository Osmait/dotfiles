package concurrent

import (
	"fmt"
	"net/http"
)

func check(endpoints ...string) {
	for _, endpoint := range endpoints {

		go func(endpoint string) {
			rep, err := http.Get(endpoint)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(rep.Status)
		}(endpoint)
	}

}
