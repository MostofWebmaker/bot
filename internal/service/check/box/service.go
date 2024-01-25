package box

import (
	"context"
	"encoding/json"
	"errors"
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
	boxberryHostInfo        = "BOXBERRY_HOST_INFO"
	boxberryHostCheckStatus = "BOXBERRY_HOST_CHECK_STATUS"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []BoxInfo {
	return allEntities
}

func (s *Service) Get(idx int) (*BoxInfo, error) {
	return &allEntities[idx], nil
}

func (s *Service) Check(trackNo string) (*BoxInfo, error) {
	httpClient := &http.Client{Timeout: 500 * time.Second}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*100)
	defer cancel()

	boxHostInfoUrl, found := os.LookupEnv(boxberryHostInfo)
	if !found {
		log.Panic("environment variable BOXBERRY_HOST_INFO not found in .env")
	}

	uri := strings.TrimRight(boxHostInfoUrl, "/") + "?"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		uri, nil)
	if err != nil {
		log.Fatal(err) // nil context or invalid method
	}

	q := url.Values{}
	q.Add("searchId", trackNo)
	req.URL.RawQuery = q.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err) // actual request error
	}
	defer resp.Body.Close()

	fmt.Println("StatusCode", resp.StatusCode)

	if resp.StatusCode == 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
		}

		infoItems := make([]BoxInfo, 1)
		err = json.Unmarshal(respBody, &infoItems)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}
		len := len(infoItems)
		if len == 0 {
			return nil, errProviderDataNotFound
		}

		info := infoItems[0]

		value, err := json.Marshal(infoItems)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil, nil
		}
		log.Printf("Got data:\n - %v", string(value))

		boxHostCheckStatusUrl, found := os.LookupEnv(boxberryHostCheckStatus)
		if !found {
			log.Panic("environment variable BOXBERRY_HOST_CHECK_STATUS not found in .env")
		}
		uri := strings.TrimRight(boxHostCheckStatusUrl, "/") + "?"

		req, err := http.NewRequestWithContext(ctx, http.MethodGet,
			uri, nil)
		if err != nil {
			log.Fatal(err) // nil context or invalid method
		}

		q := url.Values{}
		q.Add("trackId", info.Track_id)
		req.URL.RawQuery = q.Encode()

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err) // actual request error
		}
		defer resp.Body.Close()

		respBodyStatus, err := io.ReadAll(resp.Body) // НЕЛЬЗЯ СЧИТЫВАТЬ БОЛЕЕ 1 РАЗА!
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
		}

		infoStatus := new(BoxInfo)
		err = json.Unmarshal(respBodyStatus, infoStatus)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}

		info.Statuses = infoStatus.Statuses

		return &info, nil
	}

	return nil, errors.New("not found")
}
