package apiclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const HOSTNAME = "reqres.in"

type ApiClient struct {
	client *http.Client
	token  string
}

func NewApiClient(timeout time.Duration, token string) (*ApiClient, error) {
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
		token: token,
	}, nil
}

func (ac *ApiClient) GetUser(id int) (User, error) {
	url := fmt.Sprintf("https://%s/api/users/%d", HOSTNAME, id)
	
	// resp, err := ac.client.Get(url)
	// if err != nil {
	// 	return User{}, err
	// }
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return User{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", ac.token)

	resp, err := ac.client.Do(req)
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

// Request
// /api/users

// {
//     "name": "morpheus",
//     "job": "leader"
// }

// Response
// 201

// {
//     "name": "morpheus",
//     "job": "leader",
//     "id": "88",
//     "createdAt": "2025-06-03T02:47:34.275Z"
// }

func (ac *ApiClient) AddUser(name, job string) (int, error) {
	data := &postData{
		Name: name,
		Job:  job,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("https://%s/api/users", HOSTNAME)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", ac.token)

	resp, err := ac.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return 0, fmt.Errorf("request status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil
	}

	var user createdUser

	err = json.Unmarshal(body, &user)
	if err != nil {
		return 0, err
	}
	fmt.Printf("created user: %#v\n", user)

	id, err := strconv.Atoi(user.ID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
