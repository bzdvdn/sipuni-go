package sipuni

import (
	"strconv"
	"time"
)

// StatisticService - service for statistic management
type StatisticService struct {
	client *Client
}

// init StatisticService
func newStatisticService(cl *Client) *StatisticService {
	return &StatisticService{client: cl}
}

// ExportRequest - struct for statistic export query
type ExportRequest struct {
	DtmfUserAnswer  int
	From            time.Time
	To              time.Time
	Type            int
	State           int
	Tree            string
	ShowTreeID      int
	FromNumber      string
	ToNumber        string
	NumbersRinged   int
	NumbersInvolved int
	Names           int
	OutgoingLine    int
	ToAnswer        string
	Anonymous       int
	FirstTime       int
}

func (e *ExportRequest) tosipuniRequest() sipuniRequest {
	sr := sipuniRequest{
		endpoint: "statistic/export",
		keyPath:  []string{"anonymous", "firstTime", "from", "fromNumber", "state", "to", "toAnswer", "toNumber", "tree", "type"},
		query: sipuniRequestQuery{
			"anonymous":  strconv.Itoa(e.Anonymous),
			"firstTime":  strconv.Itoa(e.FirstTime),
			"from":       e.From.Format(DDMMYYYY),
			"fromNumber": e.FromNumber,
			"state":      strconv.Itoa(e.State),
			"to":         e.To.Format(DDMMYYYY),
			"toAnswer":   e.ToAnswer,
			"toNumber":   e.ToNumber,
			"tree":       e.Tree,
			"type":       strconv.Itoa(e.Type),
		},
	}
	return sr
}

// Export - export statistic, doc: https://help.sipuni.com/articles/134-182-112--poluchenie-statistiki-po-zvonkam-zapisej-razgovorov-i-statusov-sotrudnikov/
func (s *StatisticService) Export(exportRequest ExportRequest) (string, error) {
	req := exportRequest.tosipuniRequest()
	resp, err := s.client.sendRequest(req)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// ExportAll - export all stats, doc: https://help.sipuni.com/articles/134-182-112--poluchenie-statistiki-po-zvonkam-zapisej-razgovorov-i-statusov-sotrudnikov/
func (s *StatisticService) ExportAll(limit int, order string, page int) (string, error) {
	req := sipuniRequest{
		endpoint: "statistic/export/all",
		keyPath:  []string{"limit", "order", "page"},
		query: sipuniRequestQuery{
			"limit": strconv.Itoa(limit),
			"order": order,
			"page":  strconv.Itoa(page),
		},
	}
	resp, err := s.client.sendRequest(req)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// GetRecord - get mp3 file record from stats, doc: https://help.sipuni.com/articles/134-182-112--poluchenie-statistiki-po-zvonkam-zapisej-razgovorov-i-statusov-sotrudnikov/
func (s *StatisticService) GetRecord(ID string) ([]byte, error) {
	req := sipuniRequest{
		endpoint: "statistic/record",
		keyPath:  []string{"id"},
		query: sipuniRequestQuery{
			"id": ID,
		},
	}
	resp, err := s.client.sendRequest(req)
	if err != nil {
		return []byte(""), err
	}
	return resp, nil
}

// GetOperatorsStats - get statistic by operators, doc: https://help.sipuni.com/articles/134-182-112--poluchenie-statistiki-po-zvonkam-zapisej-razgovorov-i-statusov-sotrudnikov/
func (s *StatisticService) GetOperatorsStats() (string, error) {
	req := sipuniRequest{
		endpoint: "statistic/operators",
	}
	resp, err := s.client.sendRequest(req)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
