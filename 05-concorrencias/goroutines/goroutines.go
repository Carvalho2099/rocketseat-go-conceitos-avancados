package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)
	wg.Wait()

	for range n {
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://www.google.com")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			fmt.Println("Ok")
		}()
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for range n {
		go func(ctx context.Context) {
			defer wg.Done()
			req, err := http.NewRequestWithContext(
				ctx,
				"GET",
				"http://www.google.com",
				nil,
			)
			if err != nil {
				panic(err)
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("Timeout")
					return
				}
			}
			defer resp.Body.Close()
		}(ctx)
	}

	fmt.Println(time.Since(start))
}
