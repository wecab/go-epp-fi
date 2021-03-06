package registry

import (
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

const reqIDChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const reqIDLength = 5

func createRequestID(length int) string {
	reqID := make([]byte, length)
	for i := range reqID {
		reqID[i] = reqIDChars[rand.Intn(len(reqIDChars))]
	}
	return string(reqID)
}

func parseDate(rawDate string) (time.Time, error) {
	emptyDateFormat := "0001-01-01T00:00:00"
	greetingDateFormat := time.RFC3339Nano
	pollDateFormat := "2006-01-02T15:04:05"
	domainDateFormat := "2006-01-02T15:04:05.000"
	renewalDateFormat := "2006-01-03T15:04:05.0Z"

	if rawDate == "" {
		return time.Time{}, nil
	}

	if rawDate == emptyDateFormat {
		return time.Time{}, nil
	}

	if date, err := time.Parse(greetingDateFormat, rawDate); err == nil {
		return date, nil
	}
	if date, err := time.Parse(domainDateFormat, rawDate); err == nil {
		return date, nil
	}
	if date, err := time.Parse(pollDateFormat, rawDate); err == nil {
		return date, nil
	}
	if date, err := time.Parse(renewalDateFormat, rawDate); err == nil {
		return date, nil
	}

	return time.Time{}, errors.New("Unrecognised date format: " + rawDate)
}

func (s *Client) logAPIConnectionError(err error, args ...string) {
	s.log.Error("API connection failed when making a request", "error", err, args)
}

