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

	if err := request(ctx, http.MethodGet, "/projects"); err != nil {
		log.Panicln(err)
	}
}

func request(ctx context.Context, method, path string) error {
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		fmt.Sprintf("https://api.vercel.com/v9%s", path),
		nil,
	)
	if err != nil {
		return fmt.Errorf("error forming request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("VERCEL_TOKEN")))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	fmt.Println(string(body))

	return nil
}
