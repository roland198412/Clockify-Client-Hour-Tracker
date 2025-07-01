package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ClockifyReport struct {
	Totals []struct {
		TotalTime         int `json:"totalTime"`
		TotalBillableTime int `json:"totalBillableTime"`
		EntriesCount      int `json:"entriesCount"`
		Amounts           []struct {
			Type             string        `json:"type"`
			Value            float64       `json:"value"`
			AmountByCurrency []interface{} `json:"amountByCurrency"`
		} `json:"amounts"`
		NumOfCurrencies       int     `json:"numOfCurrencies"`
		Id                    string  `json:"_id"`
		TotalAmount           float64 `json:"totalAmount"`
		TotalAmountByCurrency []struct {
			Currency string `json:"currency"`
			Amount   int    `json:"amount"`
		} `json:"totalAmountByCurrency"`
	} `json:"totals"`
	GroupOne []struct {
		Currency string `json:"currency"`
		Duration int64  `json:"duration"`
		Amounts  []struct {
			Type             string  `json:"type"`
			Value            float64 `json:"value"`
			AmountByCurrency []struct {
				Currency string  `json:"currency"`
				Amount   float64 `json:"amount"`
			} `json:"amountByCurrency"`
		} `json:"amounts"`
		Amount        float64 `json:"amount"`
		Id            string  `json:"_id"`
		Name          string  `json:"name"`
		NameLowerCase string  `json:"nameLowerCase"`
		Color         string  `json:"color"`
		ClientName    string  `json:"clientName"`
	} `json:"groupOne"`
}

type HttpClient struct {
	client   http.Client
	endPoint string
	apiKey   string
	clientId string
}

func NewClockifyClient(endPoint, apiKey, clientId string) *HttpClient {
	return &HttpClient{
		client: http.Client{
			Timeout: 60 * time.Second,
		},
		endPoint: endPoint,
		apiKey:   apiKey,
		clientId: clientId,
	}
}

// GetCurrentMonthHoursWorked Fetching the amount of hours in seconds worked for the client this month
func (cr *HttpClient) GetCurrentMonthHoursWorked(startDate, endDate time.Time) (int64, error) {
	body := map[string]interface{}{
		"dateRangeStart": startDate.UTC().Format(time.RFC3339),
		"dateRangeEnd":   endDate.UTC().Format(time.RFC3339),
		"summaryFilter": map[string]interface{}{
			"groups": []string{"PROJECT"},
		},
		"clients": map[string]interface{}{
			"ids": []string{cr.clientId},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", cr.endPoint, bytes.NewBuffer(jsonBody))
	req.Header.Add("X-Api-Key", cr.apiKey)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return 0, err
	}

	resp, err := cr.client.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var reportData ClockifyReport

		err = json.Unmarshal(bodyBytes, &reportData)

		if err != nil {
			log.Fatal(err)
		}

		var duration int64

		if len(reportData.GroupOne) > 0 {
			duration = reportData.GroupOne[0].Duration
		}

		return duration, nil
	} else {
		log.Fatalf("Request failed with status code %d. Error: %v", resp.StatusCode, err)
	}
	return 0, nil
}
