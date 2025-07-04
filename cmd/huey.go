package cmd

import (
	"Huey/models"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	HueAppHeader    = "hue-application-key"
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

func makeRequest(method string, url string, body []byte) (response *http.Response, e error) {
	client := &http.Client{
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

func Color(lightId string, color string) (e error) {
	var x float64
	var y float64

	switch strings.ToLower(color) {

	case "red":
		x = 0.6718
		y = 0.3195
	case "green":
		x = 0.2425
		y = 0.6561
	case "blue":
		x = 0.0971
		y = 0.1225
	case "magenta":
		x = 0.4143
		y = 0.2038
	case "purple":
		x = 0.1315
		y = 0.0813
	case "orange":
		x = 0.5654
		y = 0.3962
	case "yellow":
		x = 0.4394
		y = 0.5359
	default:
		fmt.Printf("No color coordinates found for %s\n", color)
		return nil
	}

	url := fmt.Sprintf("https://%s/clip/v2/resource/light/%s", os.Getenv("HUE_IP_ADDRESS"), lightId)

	body := models.HueDevice{On: &models.OnState{On: true}, Color: &models.ColorState{XY: models.XYState{X: x, Y: y}}}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return fmt.Errorf("error marshalling json for request: %s\n", err)
	}

	res, err := makeRequest("PUT", url, jsonBody)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("error: %s\n%s\n", res.Status, body)
	}
	return nil
}

func Devices(roomsFlag bool) (result []models.Device, e error) {

	var url string

	if roomsFlag {
		url = fmt.Sprintf("https://%s/clip/v2/resource/room", os.Getenv("HUE_IP_ADDRESS"))
	} else {
		url = fmt.Sprintf("https://%s/clip/v2/resource/light", os.Getenv("HUE_IP_ADDRESS"))
	}

	res, err := makeRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error: %s\n%s\n", res.Status, body)
	}

	var response models.HueData
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %s\n", err)
	}

	var devices []models.Device

	for _, data := range response.Data {
		if roomsFlag {
			devices = append(devices, models.Device{Name: data.Metadata.Name, Type: data.Metadata.Archetype, Id: (*data.Services)[0].RID})
		} else {
			devices = append(devices, models.Device{Name: data.Metadata.Name, Type: data.Metadata.Archetype, Id: data.Id})
		}
	}

	return devices, nil
}

func Dim(lightId string, brightness float64) (e error) {

	url := fmt.Sprintf("https://%s/clip/v2/resource/light/%s", os.Getenv("HUE_IP_ADDRESS"), lightId)

	body := models.HueDevice{On: &models.OnState{On: true}, Dimming: &models.DimmingState{Brightness: brightness}}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return fmt.Errorf("error marshalling json for request to hue-api-v2 %s", err)
	}

	res, err := makeRequest("PUT", url, jsonBody)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("❗ Error: %s\n%s\n", res.Status, body)
	}
	return nil
}

func Off(deviceId string, isRoom bool) (e error) {

	var url string
	if isRoom {
		url = fmt.Sprintf("https://%s/clip/v2/resource/grouped_light/%s", os.Getenv("HUE_IP_ADDRESS"), deviceId)
	} else {
		url = fmt.Sprintf("https://%s/clip/v2/resource/light/%s", os.Getenv("HUE_IP_ADDRESS"), deviceId)
	}

	body := []byte(`{"on":{"on":false}}`)

	res, err := makeRequest("PUT", url, body)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("❗ Error: %s\n%s\n", res.Status, body)
	}

	return nil
}

func On(deviceId string, isRoom bool) (e error) {

	var url string
	if isRoom {
		url = fmt.Sprintf("https://%s/clip/v2/resource/grouped_light/%s", os.Getenv("HUE_IP_ADDRESS"), deviceId)
	} else {
		url = fmt.Sprintf("https://%s/clip/v2/resource/light/%s", os.Getenv("HUE_IP_ADDRESS"), deviceId)
	}

	body := []byte(`{"on":{"on":true}}`)

	res, err := makeRequest("PUT", url, body)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("error: %s\n%s\n", res.Status, body)
	}

	return nil
}

func Register() (response []byte, e error) {

	url := fmt.Sprintf("https://%s/api", os.Getenv("HUE_IP_ADDRESS"))

	body := []byte(`{"devicetype": "huey#go_cli","generateclientkey": true}`)

	res, err := makeRequest("POST", url, body)

	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}
