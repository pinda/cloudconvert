package cloudconvert

import (
	"encoding/json"
	"io"
)

type ConversionService struct {
	client *Client
}

type Conversion struct {
	ID        string          `json:"id"`
	URL       string          `json:"url"`
	Percent   json.RawMessage `json:"percent"`
	Message   string          `json:"message"`
	Step      string          `json:"step"`
	StartTime int64           `json:"starttime"`
	EndTime   int64           `json:"endtime"`
	Expire    int64           `json:"expire"`
	Input     StatusInput     `json:"input,omitempty"`
	Output    StatusOutput    `json:"output,omitempty"`
	Converter StatusConverter `json:"converter,omitempty"`
}

type StatusInput struct {
	Type     string `json:"type"`
	FileName string `json:"filename"`
	Size     int64  `json:"size"`
	Name     string `json:"name"`
	Ext      string `json:"ext"`
}

type StatusOutput struct {
	FileName  string   `json:"filename"`
	Ext       string   `json:"ext"`
	Files     []string `json:"files"`
	Size      int64    `json:"size"`
	URL       string   `json:"url"`
	Downloads int      `json:"downloads"`
}

type StatusConverter struct {
	Format   string            `json:"format"`
	Type     string            `json:"type"`
	Options  map[string]string `json:"options"`
	Duration float64           `json:"duration"`
}

type ConversionInput struct {
	Input        string `json:"input"`
	File         string `json:"file"`
	Filename     string `json:"filename"`
	OutputFormat string `json:"outputformat"`
}

type S3Credentials struct {
	AccessKeyID     string `json:"accesskeyid"`
	SecretAccessKey string `json:"secretaccesskey"`
	Bucket          string `json:"bucket"`
	Path            string `json:"path,omitempty"`
	ACL             string `json:"acl,omitempty"`
}

type Credentials struct {
	Credentials S3Credentials `json:"s3"`
}

type S3ConversionInput struct {
	Input        Credentials `json:"input"`
	File         string      `json:"file"`
	OutputFormat string      `json:"outputformat"`
	Output       Credentials `json:"output"`
}

func (s *ConversionService) New(url string, con ConversionInput) (*Conversion, error) {
	req, err := s.client.NewRequest("POST", url, con)
	if err != nil {
		return nil, err
	}

	uResp := new(Conversion)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

func (s *ConversionService) NewS3(url string, con S3ConversionInput) (*Conversion, error) {
	req, err := s.client.NewRequest("POST", url, con)
	if err != nil {
		return nil, err
	}

	uResp := new(Conversion)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

func (s *ConversionService) Status(url string) (*Conversion, error) {
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	uResp := new(Conversion)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

func (s *ConversionService) Download(url string) (io.ReadCloser, error) {
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp.Body, err
}

func (s *ConversionService) Remove(url string) error {
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
