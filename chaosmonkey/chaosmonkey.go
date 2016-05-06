package chaosmonkey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type EventRequest struct {
	EventType string `json:"eventType"`
	GroupType string `json:"groupType"`
	GroupName string `json:"groupName"`
	ChaosType string `json:"chaosType,omitempty"`
}

type EventResponse struct {
	EventRequest

	MonkeyType string `json:"monkeyType"`
	EventID    string `json:"eventId"`
	EventTime  int64  `json:"eventTime"`
	Region     string `json:"region"`
}

type Config struct {
	Endpoint string
	Username string
	Password string

	HTTPClient *http.Client
}

type Client struct {
	config *Config
}

func (c *Config) ReadEnvironment() error {
	return nil
}

func NewClient(c *Config) (*Client, error) {
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}
	return &Client{config: c}, nil
}

func (c *Client) TriggerEvent(groupName, chaosType string) error {
	e := EventRequest{
		EventType: "CHAOS_TERMINATION",
		GroupType: "ASG",
		GroupName: groupName,
		ChaosType: chaosType,
	}
	payload, err := json.Marshal(&e)
	if err != nil {
		return err
	}

	url := c.config.Endpoint + "/simianarmy/api/v1/chaos"
	resp, err := c.sendRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return decodeError(resp)
	}

	var res EventResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}

	return nil
}

func (c *Client) sendRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", "havoc")

	if c.config.Username != "" && c.config.Password != "" {
		req.SetBasicAuth(c.config.Username, c.config.Password)
	}

	return c.config.HTTPClient.Do(req)
}

func decodeError(resp *http.Response) error {
	var m struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&m); err == nil && m.Message != "" {
		return fmt.Errorf("%s", m.Message)
	}
	return fmt.Errorf("%s", resp.Status)
}