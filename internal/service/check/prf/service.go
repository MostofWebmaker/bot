package prf

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	prfHost = "PRF_HOST"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []PRFInfo {
	return allEntities
}

func (s *Service) Get(idx int) (*PRFInfo, error) {
	return &allEntities[idx], nil
}

func (s *Service) Check(trackNo string) (*PRFInfo, error) {
	httpClient := &http.Client{Timeout: 500 * time.Second}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*100)
	defer cancel()

	prfHostUrl, found := os.LookupEnv(prfHost)
	if !found {
		log.Panic("environment variable KCE_HOST not found in .env")
	}

	uri := strings.TrimRight(prfHostUrl, "/")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		uri, nil)
	if err != nil {
		log.Fatal(err) // nil context or invalid method
	}

	q := url.Values{}
	q.Add("language", "ru")
	q.Add("track-numbers", trackNo)
	req.URL.RawQuery = q.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err) // actual request error
	}
	defer resp.Body.Close()

	fmt.Println("StatusCode", resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body) // НЕЛЬЗЯ СЧИТЫВАТЬ БОЛЕЕ 1 РАЗА!
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	info := new(PRFInfo)
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
