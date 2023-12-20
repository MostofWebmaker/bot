package cdek

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
	cdekHost    = "CDEK_HOST"
	local       = "ru"
	contentType = "application/json"
)

var reqBody = `query getTrackingInfo
{
  tracking: trackingInfo(
    trackId: "%s"
    websiteId: "ru"
    locale: "ru"
    token: null
  ) {
    success
    orderNumber
    status {
      code
      name
      note
      date
    }
    statuses {
      code
      name
      note
      date
      completed
      items {
        code
        name
        statuses {
            code
            name
            date
        }
      }
    }
    cityFrom {
      code
      name
    }
    cityTo {
      code
      name
    }
    orderDate
    tariffDateEnd
    storageDateEnd
    deliveryAgreementDate
    returnOrderNumber
    weight
    stockType
    receiver {
        initials
        address {
            title
            city {
                code
                name
            }
            office {
                systemName
                type
                worktime
                notes
            }
        }
    }
    notes {
        code
        name
    }
    nonDeliveryNote {
        code
        name
    }
    errors {
      message
      code
    }
    specialNote
  }
}
`

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []CDEKInfo {
	return allEntities
}

func (s *Service) Get(idx int) (*CDEKInfo, error) {
	return &allEntities[idx], nil
}

func (s *Service) Check(trackNo string) (*CDEKInfo, error) {
	httpClient := &http.Client{Timeout: 500 * time.Second}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*100)
	defer cancel()

	cdekHostUrl, found := os.LookupEnv(cdekHost)
	if !found {
		log.Panic("environment variable CDEK_HOST not found in .env")
	}

	preparedBody := fmt.Sprintf(reqBody, trackNo)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		cdekHostUrl, strings.NewReader(preparedBody))
	if err != nil {
		log.Fatal(err) // nil context or invalid method
	}

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

	info := new(CDEKInfo)
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
