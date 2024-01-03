package sipuni

import "fmt"

// SipuniError
type SipuniError struct {
	HTTPCode     int
	Endpoint     string
	ErrorMessage string
}

// Error returns string representation of the YandexWebmasterError
func (e *SipuniError) Error() string {
	return fmt.Sprintf("Http code: %d, endoint: %s, message: %s", e.HTTPCode, e.Endpoint, e.ErrorMessage)
}
