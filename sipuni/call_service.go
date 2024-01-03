package sipuni

import (
	"encoding/json"
	"strconv"
)

// CallService : service for calls management
type CallService struct {
	client *Client
}

// func for init CallService
func newCallService(cl *Client) *CallService {
	return &CallService{client: cl}
}

type CallResponse struct {
	ID     string `json:"id,omitempty"`
	CallID string `json:"callId,omitempty"`
}

func (s *CallService) sendCallRequest(req sipuniRequest) (CallResponse, error) {
	var callResp CallResponse
	resp, err := s.client.sendRequest(req)
	if err != nil {
		return callResp, err
	}
	if err = json.Unmarshal(resp, &callResp); err != nil {
		return callResp, &SipuniError{422, req.endpoint, "json Unmarshal error"}
	}
	return callResp, nil
}

// MakeCall - make basic call, doc: https://help.sipuni.com/articles/134-182-113--sozdanie-zvonka-na-nomer-s-pomoshyu-api/
func (s *CallService) MakeCall(phone string, sipnumber string, reverse int, antiaon int) (CallResponse, error) {
	req := sipuniRequest{
		endpoint: "callback/call_number",
		keyPath:  []string{"antiaon", "phone", "reverse", "sipnumber"},
		query: sipuniRequestQuery{
			"antiaon":   strconv.Itoa(antiaon),
			"phone":     phone,
			"reverse":   strconv.Itoa(reverse),
			"sipnumber": sipnumber,
		},
	}
	return s.sendCallRequest(req)
}

// MakeTreeCall - make tree call, doc: https://help.sipuni.com/articles/134-182-113--sozdanie-zvonka-na-nomer-s-pomoshyu-api/
func (s *CallService) MakeTreeCall(phone string, sipnumber string, tree string, reverse int, callAttemptTime int) (CallResponse, error) {
	req := sipuniRequest{
		endpoint: "callback/call_tree",
		keyPath:  []string{"callAttemptTime", "phone", "reverse", "sipnumber", "tree"},
		query: sipuniRequestQuery{
			"callAttemptTime": strconv.Itoa(callAttemptTime),
			"phone":           phone,
			"reverse":         strconv.Itoa(reverse),
			"sipnumber":       sipnumber,
			"tree":            tree,
		},
	}
	return s.sendCallRequest(req)
}

// MakeExternalCall - make external call, doc: https://help.sipuni.com/articles/134-182-113--sozdanie-zvonka-na-nomer-s-pomoshyu-api/
func (s *CallService) MakeExternalCall(phoneFrom string, phoneTo string, sipnumber string, sipnumber2 string) (CallResponse, error) {
	req := sipuniRequest{
		endpoint: "callback/call_external",
		keyPath:  []string{"phoneFrom", "phoneTo", "sipnumber", "sipnumber2"},
		query: sipuniRequestQuery{
			"phoneFrom":  phoneFrom,
			"phoneTo":    phoneTo,
			"sipnumber":  sipnumber,
			"sipnumber2": sipnumber2,
		},
	}
	return s.sendCallRequest(req)
}

// HangUpCall - hungup call, doc: https://help.sipuni.com/articles/134-182-108--zapros-na-zavershenie-zvonka/
func (s *CallService) HangUpCall(callID string) error {
	req := sipuniRequest{
		endpoint: "events/call/hangup",
		keyPath:  []string{"callId"},
		query: sipuniRequestQuery{
			"callId": callID,
		},
	}
	_, err := s.client.sendRequest(req)
	return err
}

// MakeVoiceCall - make voice call, doc: https://help.sipuni.com/articles/134-182-107--generaciya-golosovogo-zvonka-s-pomoshyu-api/
func (s *CallService) MakeVoiceCall(phone string, message string, voice string, sipnumber string) (CallResponse, error) {
	req := sipuniRequest{
		endpoint: "voicecall/call",
		keyPath:  []string{"message", "phone", "sipnumber", "voice"},
		query: sipuniRequestQuery{
			"message":   message,
			"phone":     phone,
			"voice":     voice,
			"sipnumber": sipnumber,
		},
	}
	return s.sendCallRequest(req)
}

// CallSpyMode doc: https://help.sipuni.com/articles/134-182-464--suflirovanie/
func (s *CallService) CallSpyMode(callID string, sipnumber string, whisper string) error {
	req := sipuniRequest{
		endpoint: "events/call/spymode",
		keyPath:  []string{"callId", "sipnumber", "whisper"},
		query: sipuniRequestQuery{
			"callId":    callID,
			"sipnumber": sipnumber,
			"whisper":   whisper,
		},
	}
	_, err := s.client.sendRequest(req)
	return err
}
