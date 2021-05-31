package networking

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetCurrentNodeIP() (string, error) {
	resp, err := http.Get("https://checkip.amazonaws.com")
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}

	return strings.Trim(string(body), "\n"), nil
}

func IpApiGetIPLocation(ip string) (*IpApiIPLocationResponse, error) {
	link := url.URL{
		Scheme: "https",
		Host:   "ipapi.co",
		Path:   ip + "/json/",
	}

	resp, err := http.Get(link.String())
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(IpApiIPLocationResponse)
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	if response.Error {
		return nil, errors.New(response.Reason)
	}

	return response, nil
}

func IpStackGetIPLocation(ip string) (*IpStackIPLocationResponse, error) {
	q := url.Values{}
	q.Add("access_key", os.Getenv("IPSTACK_API_KEY"))
	q.Add("format", "1")
	link := url.URL{
		Scheme:   "http",
		Host:     "api.ipstack.com",
		Path:     ip,
		RawQuery: q.Encode(),
	}

	resp, err := http.Get(link.String())
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := new(IpStackIPLocationResponse)
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
