package cloudconvert

type ConversionService struct {
	client *Client
}

type Conversion struct {
	ID        string                 `json:"id"`
	URL       string                 `json:"url"`
	Percent   string                 `json:"percent"`
	Message   string                 `json:"message"`
	Step      string                 `json:"step"`
	StartTime int                    `json:"starttime"`
	Expire    int                    `json:"expire"`
	Input     map[string]interface{} `json:"input"`
	Converter map[string]interface{} `json:"converter"`
}

type ConversionInput struct {
	Input        string `json:"input"`
	File         string `json:"file"`
	Filename     string `json:"filename"`
	OutputFormat string `json:"outputformat"`
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

func (s *ConversionService) Status(url string) (*Conversion, error) {
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	uResp := new(Conversion)
	_, err = s.client.Do(req, uResp)
	return uResp, err
}

func (s *ConversionService) Remove(url string) error {
	req, err := s.client.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}
