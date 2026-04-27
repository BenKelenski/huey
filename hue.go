package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	HueAppHeader    = "hue-application-key"
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
	GET             = "GET"
	PUT             = "PUT"
	POST            = "POST"
)

func makeRequest(method string, url string, body []byte) (response *http.Response, e error) {
	client := &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set(ContentType, ApplicationJson)
	req.Header.Set(HueAppHeader, os.Getenv("HUE_USERNAME"))

	if err != nil {
		return nil, fmt.Errorf("error while creating request to hue-api-v2: %s\n", err)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("client error sending http request: %s\n", err)
	}

	return res, nil
}

func SetRoomLights(room Room, on bool) error {
	var groupedLightID string
	for _, s := range room.Services {
		if s.RType == "grouped_light" {
			groupedLightID = s.RID
			break
		}
	}
	if groupedLightID == "" {
		return fmt.Errorf("no grouped_light service found for room %s", room.Metadata.Name)
	}

	url := fmt.Sprintf("https://%s/clip/v2/resource/grouped_light/%s", os.Getenv("HUE_IP_ADDRESS"), groupedLightID)
	body, err := json.Marshal(map[string]any{"on": map[string]any{"on": on}})
	if err != nil {
		return fmt.Errorf("error marshaling request body: %s", err)
	}

	res, err := makeRequest(PUT, url, body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	return nil
}

func GetRooms() ([]Room, error) {
	url := fmt.Sprintf("https://%s/clip/v2/resource/room", os.Getenv("HUE_IP_ADDRESS"))

	res, err := makeRequest(GET, url, nil)

	if err != nil {
		fmt.Printf("error while making request to hue-api-v2: %s\n", err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d\n", res.StatusCode)
	}

	var roomsResponse RoomsResponse
	err = json.NewDecoder(res.Body).Decode(&roomsResponse)

	if err != nil {
		return nil, fmt.Errorf("error while decoding response body: %s\n", err)
	}

	return roomsResponse.Data, nil
}
