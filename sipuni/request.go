package sipuni

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

type sipuniRequestQuery map[string]string

type sipuniRequest struct {
	endpoint string
	keyPath  []string
	query    sipuniRequestQuery
}

func (r *sipuniRequest) generateHash(values []string) string {
	preHash := strings.Join(values, "+")
	hash := md5.Sum([]byte(preHash))
	return hex.EncodeToString(hash[:])
}

func (r *sipuniRequest) createQueryParams(user string, secretKey string) string {
	var buf strings.Builder
	var values []string
	for _, key := range r.keyPath {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		keyEscaped := url.QueryEscape(key)
		value := r.query[key]
		values = append(values, value)
		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(value))
	}
	values = append(values, user, secretKey)
	hash := r.generateHash(values)
	queryParams := fmt.Sprintf("%s&user=%s&hash=%s", buf.String(), user, hash)
	return queryParams
}
