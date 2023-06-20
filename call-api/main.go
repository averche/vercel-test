package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancelContextFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelContextFunc()

	token := os.Getenv("VERCEL_TOKEN")
	if token == "" {
		log.Panicln("the expected VERCEL_TOKEN environment is not set!")
	}

	// create a secret (upsert seems to be ignored)
	if err := request(
		ctx,
		token,
		http.MethodPost,
		"/v9/projects/prj_30mxVfZKN5oYRcrGk5rku3hUSlQY/env?upsert=true",
		map[string]interface{}{
			"key":    "MY_NEW_ENV2",
			"type":   "encrypted",
			"target": []string{"preview", "development", "production"},
			"value":  "some value!!!!",
		},
	); err != nil {
		log.Panicln(err)
	}

	// update a secret
	if err := request(
		ctx,
		token,
		http.MethodPatch,
		"/v9/projects/prj_30mxVfZKN5oYRcrGk5rku3hUSlQY/env/08w3OJ5CpocL9okK",
		map[string]interface{}{
			"key":   "TEST_ENV",
			"type":  "encrypted",
			"value": "encrypted new value!!!!",
		},
	); err != nil {
		log.Panicln(err)
	}
}

func request(ctx context.Context, token, method, path string, body map[string]interface{}) error {
	var reqBody bytes.Buffer

	if body != nil {
		if err := json.NewEncoder(&reqBody).Encode(body); err != nil {
			return fmt.Errorf("could not encode request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		fmt.Sprintf("https://api.vercel.com%s", path),
		&reqBody,
	)
	if err != nil {
		return fmt.Errorf("error forming request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	fmt.Println(string(respBody))

	return nil
}
