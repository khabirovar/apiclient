package apiclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const HOSTNAME = "reqres.in"

type ApiClient struct {
	client *http.Client
}

func NewApiClient(timeout time.Duration) (*ApiClient, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &ApiClient{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (ac *ApiClient) GetUser(id int) (User, error) {
	url := fmt.Sprintf("https://%s/api/users/%d", HOSTNAME, id)
	resp, err := ac.client.Get(url)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var ur userResponse
	err = json.Unmarshal(body, &ur)
	if err != nil {
		return User{}, err
	}

	return ur.User, nil
}
