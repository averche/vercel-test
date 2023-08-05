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

	// fetch user
	if err := request(
		ctx,
		token,
		http.MethodGet,
		"/v2/user",
		nil,
	); err != nil {
		log.Panicln(err)
	}
	//
	// // fetch team
	if err := request(
		ctx,
		token,
		http.MethodGet,
		"/v2/teams/team_kRChqe754SOcJudYPP6rroKH",
		nil,
	); err != nil {
		log.Panicln(err)
	}
	//
	// // fetch configuration
	// if err := request(
	// 	ctx,
	// 	token,
	// 	http.MethodGet,
	// 	"/v1/integrations/configuration/icfg_sZhrM2YWkdRdz0LvLLpupITC?teamId=team_kRChqe754SOcJudYPP6rroKH",
	// 	nil,
	// ); err != nil {
	// 	log.Panicln(err)
	// }

	// fetch all envs
	// if err := request(
	// 	ctx,
	// 	token,
	// 	http.MethodGet,
	// 	"/v9/projects",
	// 	nil,
	// ); err != nil {
	// 	log.Panicln(err)
	// }

	return
	// create a secret
	if err := request(
		ctx,
		token,
		http.MethodPost,
		"/v10/projects/prj_30mxVfZKN5oYRcrGk5rku3hUSlQY/env?upsert=true",
		map[string]interface{}{
			"key":    "TEST_ENV",
			"type":   "encrypted",
			"target": []string{"production"},
			"value":  "production",
		},
	); err != nil {
		log.Panicln(err)
	}

	// update a secret
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
