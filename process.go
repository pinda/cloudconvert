package cloudconvert

import (
	"errors"
	"time"
)

type ProcessService struct {
	client *Client
}

type Process struct {
	URL     string    `json:"url"`
	ID      string    `json:"id"`
	Host    string    `json:"host"`
	Expires time.Time `json:"expires"`
	MaxTime int       `json:"maxtime"`
	Minutes int       `json:"minutes"`
}

type ProcessInput struct {
	InputFormat  string `json:"inputformat"`
	OutputFormat string `json:"outputformat"`
}

func (s *ProcessService) New(p ProcessInput) (*Process, error) {
	if p.InputFormat == "" || p.OutputFormat == "" {
		return nil, errors.New("Input and outputformat are required")
	}

	req, err := s.client.NewRequest("POST", "/process", p)
	if err != nil {
		return nil, err
	}

	uResp := new(Process)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}
