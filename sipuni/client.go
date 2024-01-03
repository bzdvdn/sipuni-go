package sipuni

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	apiBaseUrl = "https://sipuni.com/api/"
	DDMMYYYY   = "02.01.2006"
)

// Client : main client for Sipuni management
type Client struct {
	client    *http.Client
	user      string
	secretKey string
	Call      *CallService
	Auxiliary *AuxiliaryService
	Autocall  *AutocallService
	Statistic *StatisticService
}

// NewClient creates new Client to Sipuni
func NewClient(user string, secretKey string) *Client {
	cl := &Client{
		client:    http.DefaultClient,
		secretKey: secretKey,
		user:      user,
	}
	cl.Call = newCallService(cl)
	cl.Auxiliary = newAuxiliaryService(cl)
	cl.Autocall = newAutocallService(cl)
	cl.Statistic = newStatisticService(cl)
	return cl
}

func (c *Client) sendRequest(req sipuniRequest) ([]byte, error) {
	queryParams := req.createQueryParams(c.user, c.secretKey)
	queryUrl := fmt.Sprintf("%s%s?%s", apiBaseUrl, req.endpoint, queryParams)
	// fmt.Println(queryUrl)
	buf := &bytes.Buffer{}
	resp, err := http.Post(queryUrl, "application/json", buf)
	if err != nil {
		return nil, &SipuniError{http.StatusServiceUnavailable, req.endpoint, err.Error()}
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &SipuniError{resp.StatusCode, req.endpoint, string(respBody)}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &SipuniError{resp.StatusCode, req.endpoint, string(respBody)}
	}
	return respBody, nil
}
