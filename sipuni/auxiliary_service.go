package sipuni

import (
	"encoding/json"
	"strconv"
)

// AuxiliaryService - service for manager aux data
type AuxiliaryService struct {
	client *Client
}

// init AuxiliaryService
func newAuxiliaryService(cl *Client) *AuxiliaryService {
	return &AuxiliaryService{client: cl}
}

type Schema struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	CreatedAt int    `json:"createdAt"`
}

type SchemaResponse struct {
	Data []*Schema `json:"data"`
}

func (s *AuxiliaryService) GetUserSchemas(type_ string) (SchemaResponse, error) {
	req := sipuniRequest{
		endpoint: "schema/list",
		keyPath:  []string{"type"},
		query: sipuniRequestQuery{
			"type": type_,
		},
	}
	var respData SchemaResponse
	resp, err := s.client.sendRequest(req)
	if err != nil {
		return respData, err
	}
	if err := json.Unmarshal(resp, &respData); err != nil {
		return respData, &SipuniError{422, req.endpoint, "cant unmarshal"}
	}
	return respData, nil
}

func (s *AutocallService) SetUserStatus(sipnumber string, status int) error {
	req := sipuniRequest{
		endpoint: "workplace/setUserStatu",
		keyPath:  []string{"sipnumber", "status"},
		query: sipuniRequestQuery{
			"sipnumber": sipnumber,
			"status":    strconv.Itoa(status),
		},
	}
	_, err := s.client.sendRequest(req)
	return err
}
