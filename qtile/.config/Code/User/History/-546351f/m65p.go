package concurrent

import (
	"fmt"
	"net/http"
	"sync"
)

func Check(wg *sync.WaitGroup, endpoints ...string) {
	for _, endpoint := range endpoints {
		wg.Add(1)
		go func(endpoint string, wg *sync.WaitGroup) {
			rep, err := http.Get(endpoint)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(rep.Status)
			wg.Done()
		}(endpoint, wg)

	}
	wg.Wait()

}
