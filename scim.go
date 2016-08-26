package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Scim struct {
	Target     string
	Auth       string
	httpClient *http.Client
}

func NewScim(target Target, context Context) Scim {
	return Scim{
		Target:     target.URL(),
		Auth:       context.TokenType() + " " + context.AccessToken(),
		httpClient: NewHTTPClient(NewTLSConfig(target.SkipSslVerify(), target.CaCertFile())),
	}
}

func (s Scim) GetClients(start int, count int) ([]Client, error) {
	var query string
	if start > 0 {
		query = fmt.Sprintf("?startIndex=%d", start)
	}
	if count > 0 {
		if query == "" {
			query = "?"
		} else {
			query = query + "&"
		}
		query = query + fmt.Sprintf("count=%d", count)
	}
	_, body, err := s.Request("GET", "/oauth/clients"+query, nil, nil)

	if err != nil {
		return nil, err
	}

	resources := ClientResources{}
	_ = json.Unmarshal(body, &resources)
	return resources.Clients, nil
}

func (s Scim) GetClient(clientID string) (*Client, error) {
	status, body, err := s.Request("GET", "/oauth/clients/"+clientID, nil, nil)
	if status == http.StatusNotFound {
		return nil, nil
	}
	if err == nil {
		client := &Client{}
		_ = json.Unmarshal(body, client)
		return client, nil
	}
	return nil, err
}

func (s Scim) CreateClient(client *Client) error {
	_, _, err := s.Request("POST", "/oauth/clients", client, nil)
	return err
}

func (s Scim) PutClient(client *Client) error {
	_, _, err := s.Request("PUT", "/oauth/clients/"+client.ID, client, nil)
	return err
}

func (s Scim) DeleteClient(clientID string) error {
	_, _, err := s.Request("DELETE", "/oauth/clients/"+clientID, nil, nil)
	return err
}

func (s Scim) GetUsers(start int, count int) ([]User, error) {
	var query string
	if start > 0 {
		query = fmt.Sprintf("?startIndex=%d", start)
	}
	if count > 0 {
		if query == "" {
			query = "?"
		} else {
			query = query + "&"
		}
		query = query + fmt.Sprintf("count=%d", count)
	}
	_, body, err := s.Request("GET", "/Users"+query, nil, nil)

	if err != nil {
		return nil, err
	}

	resources := UserResources{}
	_ = json.Unmarshal(body, &resources)
	return resources.Users, nil
}

func (s Scim) GetUser(userName string) (*User, error) {
	query := fmt.Sprintf("userName eq \"%s\"", userName)
	status, body, err := s.Request("GET", "/Users?filter="+url.QueryEscape(query), nil, nil)
	if status == http.StatusNotFound {
		return nil, nil
	}
	if err == nil {
		resources := UserResources{}
		_ = json.Unmarshal(body, &resources)
		if len(resources.Users) == 1 {
			return &resources.Users[0], nil
		}
	}
	return nil, err
}

func (s Scim) CreateUser(user *User) error {
	_, _, err := s.Request("POST", "/Users", user, nil)
	return err
}

func (s Scim) PutUser(user *User) error {
	_, _, err := s.Request("PUT", "/Users/"+user.ID, user, map[string]string{
		"If-Match": strconv.Itoa(user.Meta.Version),
	})
	return err
}

func (s Scim) DeleteUser(userID string) error {
	_, _, err := s.Request("DELETE", "/Users/"+userID, nil, nil)
	return err
}

func (s Scim) Request(method, path string, body interface{}, headers map[string]string) (status int, responseBody []byte, err error) {
	var bodyReader io.Reader
	if body != nil {
		bodyJSON, _ := json.Marshal(body)
		bodyReader = bytes.NewReader(bodyJSON)
	}
	request, err := http.NewRequest(method, s.Target+path, bodyReader)
	if err != nil {
		return
	}

	if headers != nil {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}
	request.Header.Set("Authorization", s.Auth)
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := s.httpClient.Do(request)
	if err != nil {
		return
	}

	status = response.StatusCode
	responseBody, err = ioutil.ReadAll(response.Body)

	if status >= http.StatusBadRequest && err == nil {
		responseErr := Error{}
		if json.Unmarshal(responseBody, &responseErr) == nil {
			err = fmt.Errorf("Invalid status response: %d. %s", status, responseErr.Description)
		} else {
			err = fmt.Errorf("Invalid status response: %d", status)
		}
	}
	return
}
