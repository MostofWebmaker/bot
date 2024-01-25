package kce

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	kceHost = "KCE_HOST"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []KCEInfo {
	return allEntities
}

func (s *Service) Get(idx int) (*KCEInfo, error) {
	return &allEntities[idx], nil
}

func (s *Service) Check(trackNo string) (*KCEInfo, error) {
	httpClient := &http.Client{Timeout: 500 * time.Second}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*100)
	defer cancel()

	kceHostUrl, found := os.LookupEnv(kceHost)
	if !found {
		log.Panic("environment variable KCE_HOST not found in .env")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		strings.TrimRight(kceHostUrl, "/")+"/"+trackNo, nil)
	if err != nil {
		log.Fatal(err) // nil context or invalid method
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err) // actual request error
	}
	defer resp.Body.Close()

	fmt.Println("StatusCode", resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	info := new(KCEInfo)
	err = json.Unmarshal(respBody, info)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	value, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, nil
	}
	log.Printf("Got data:\n - %v", string(value))

	return info, nil
}
