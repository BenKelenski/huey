package cmd

import (
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
	HUE_APP_HEADER = "hue-application-key"
)

type HueResponse struct {
	Data   []HueDevice `json:"data"`
	Errors []any       `json:"errors"` // or define properly if you care about errors
}

type HueDevice struct {
	ID          string          `json:"offId"`
	IDV1        string          `json:"id_v1,omitempty"`
	Identify    json.RawMessage `json:"identify,omitempty"` // Empty objects → use RawMessage or `map[string]any`
	Metadata    Metadata        `json:"metadata"`
	ProductData ProductData     `json:"product_data"`
	Services    []Service       `json:"services"`
	Type        string          `json:"type"`
}

type Metadata struct {
	Archetype string `json:"archetype"`
	Name      string `json:"name"`
}

type ProductData struct {
	Certified            bool   `json:"certified"`
	HardwarePlatformType string `json:"hardware_platform_type,omitempty"`
	ManufacturerName     string `json:"manufacturer_name"`
	ModelID              string `json:"model_id"`
	ProductArchetype     string `json:"product_archetype"`
	ProductName          string `json:"product_name"`
	SoftwareVersion      string `json:"software_version"`
}

type Service struct {
	RID   string `json:"rid"`
	RType string `json:"rtype"`
}

type Light struct {
	Name string
	Type string
	Id   string
}

type HueRequest struct {
	On    OnState    `json:"on,omitempty"`
	Color ColorState `json:"color,omitempty"`
}

type OnState struct {
	On bool `json:"on"`
}

type ColorState struct {
	XY XYState `json:"xy"`
}

type XYState struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func makeRequest(method string, url string, body []byte) *http.Response {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Error while creating request to hue-api-v2", err)
		os.Exit(1)
	}

	req.Header.Add(HUE_APP_HEADER, os.Getenv("HUE_USERNAME"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	return res
}

func Color(lightId string, color string) {
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
		return
	}

	url := fmt.Sprintf("https://%s/clip/v2/resource/light/%s", os.Getenv("HUE_IP_ADDRESS"), lightId)

	body := HueRequest{On: OnState{On: true}, Color: ColorState{XY: XYState{X: x, Y: y}}}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		fmt.Println("Error marshalling json for request to hue-api-v2", err)
		os.Exit(1)
	}

	res := makeRequest("PUT", url, jsonBody)

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		fmt.Printf("❗ Error: %s\n%s\n", res.Status, body)
		os.Exit(1)
	}
}

func Devices() (result []Light) {

	url := fmt.Sprintf("https://%s/clip/v2/resource/device", os.Getenv("HUE_IP_ADDRESS"))

	res := makeRequest("GET", url, nil)

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		fmt.Printf("❗ Error: %s\n%s\n", res.Status, body)
		os.Exit(1)
	}

	var response HueResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println("❗ Error decoding JSON:", err)
		os.Exit(1)
	}

	var lights []Light

	for _, device := range response.Data {
		for _, service := range device.Services {
			if service.RType == "light" {
				lights = append(lights, Light{device.Metadata.Name, device.Metadata.Archetype, service.RID})
			}
		}
	}

	return lights
}

func Off(lightId string) {

	url := fmt.Sprintf("https://%s/clip/v2/resource/light/%s", os.Getenv("HUE_IP_ADDRESS"), lightId)

	body := []byte(`{"on":{"on":false}}`)

	res := makeRequest("PUT", url, body)

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		fmt.Printf("❗ Error: %s\n%s\n", res.Status, body)
		os.Exit(1)
	}
}

func On(lightId string) {

	url := fmt.Sprintf("https://%s/clip/v2/resource/light/%s", os.Getenv("HUE_IP_ADDRESS"), lightId)

	body := []byte(`{"on":{"on":true}}`)

	res := makeRequest("PUT", url, body)

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		fmt.Printf("❗ Error: %s\n%s\n", res.Status, body)
		os.Exit(1)
	}
}

func Register() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	url := fmt.Sprintf("https://%s/api", os.Getenv("HUE_IP_ADDRESS"))

	body := []byte(`{"devicetype": "huey#go_cli","generateclientkey": true}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("Error while creating request to hue-api-v2", err)
		os.Exit(1)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: response body - %s", resBody)
}
