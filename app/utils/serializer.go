package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// JSONSerializer ...
func JSONSerializer(rw http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(rw).Encode(data)
}

// Timestamp ...
type Timestamp struct {
	sql.NullTime
}

// UnmarshalJSON ...
func (t *Timestamp) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse(time.RFC1123, strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// MarshalJSON ...
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return []byte(strconv.FormatInt(t.Time.Unix(), 10)), nil
	}

	return []byte("null"), nil
}
