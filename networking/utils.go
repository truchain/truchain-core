package networking

import (
	"encoding/json"
	"io"
	"net/http"
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

func GetIPLocation(ip string) (*IpStackIPLocationResponse, error) {
	resp, err := http.Get("http://api.ipstack.com/" + ip + "?access_key=" + os.Getenv("IPSTACK_API_KEY") + "&format=1")
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
