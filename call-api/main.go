package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	/* */ log.Println("call-api: begin")
	defer log.Println("call-api: end")

	ctx, cancelContextFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelContextFunc()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://api.vercel.com/v9/projects"),
		nil,
	)
	if err != nil {
		log.Panicf("error forming request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("VERCEL_TOKEN")))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("error reading body: %v", err)
	}

	fmt.Println(string(body))
}
