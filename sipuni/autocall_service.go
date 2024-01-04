package sipuni

import (
	"encoding/json"
	"strconv"
)

// AutocallService - service for autocall management
type AutocallService struct {
	client *Client
}

// init AutocallService
func newAutocallService(cl *Client) *AutocallService {
	return &AutocallService{client: cl}
}

type AutoCallAddResponse struct {
	ID string `json:"id"`
}

// AddNumbder - add number to autocall, doc: https://help.sipuni.com/articles/134-182-466--avtoobzvon-api/
func (s *AutocallService) AddNumber(again int, number string, autocallID int) (AutoCallAddResponse, error) {
	req := sipuniRequest{
		endpoint: "autocall/add_number",
		keyPath:  []string{"again", "autocallId", "number"},
		query: sipuniRequestQuery{
			"again":      strconv.Itoa(again),
			"number":     number,
			"autocallId": strconv.Itoa(autocallID),
		},
	}
	var respData AutoCallAddResponse
	resp, err := s.client.sendRequest(req)
	// fmt.Println(string(resp))
	if err != nil {
		return respData, err
	}
	if err := json.Unmarshal(resp, &respData); err != nil {
		return respData, &SipuniError{422, req.endpoint, "cant unmarshal"}
	}
	return respData, nil
}

// DeleteNumber - delete number from autocall, doc: https://help.sipuni.com/articles/134-182-466--avtoobzvon-api/
func (s *AutocallService) DeleteNumber(number string, autocallID int) error {
	req := sipuniRequest{
		endpoint: "autocall/add_number",
		keyPath:  []string{"autocallId", "number"},
		query: sipuniRequestQuery{
			"number":     number,
			"autocallId": strconv.Itoa(autocallID),
		},
	}
	_, err := s.client.sendRequest(req)
	return err
}

// AddOperator - add operator to autocall, doc: https://help.sipuni.com/articles/134-182-466--avtoobzvon-api/
func (s *AutocallService) AddOperator(operator string, autocallID int) error {
	req := sipuniRequest{
		endpoint: "autocall/add_operator",
		keyPath:  []string{"autocallId", "operator"},
		query: sipuniRequestQuery{
			"autocallId": strconv.Itoa(autocallID),
			"operator":   operator,
		},
	}
	_, err := s.client.sendRequest(req)
	return err
}

// DeleteOperator - delete operator from autocall, doc: https://help.sipuni.com/articles/134-182-466--avtoobzvon-api/
func (s *AutocallService) DeleteOperator(operator string, autocallID int) error {
	req := sipuniRequest{
		endpoint: "autocall/delete_operator",
		keyPath:  []string{"autocallId", "operator"},
		query: sipuniRequestQuery{
			"autocallId": strconv.Itoa(autocallID),
			"operator":   operator,
		},
	}
	_, err := s.client.sendRequest(req)
	return err
}

// GetStatistic - get statistic from autocall, doc: https://help.sipuni.com/articles/134-182-466--avtoobzvon-api/
func (s *AutocallService) GetStatistic(autocallID int) (string, error) {
	req := sipuniRequest{
		endpoint: "autocall/statistics",
		keyPath:  []string{"autocallId"},
		query: sipuniRequestQuery{
			"autocallId": strconv.Itoa(autocallID),
		},
	}
	resp, err := s.client.sendRequest(req)
	if err != nil {
		return "", err
	}
	return string(resp), nil // return csv data, nil
}
