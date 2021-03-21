package nullable

import (
	//"errors"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Because sql.NullString is not serializable
// We create new similar type
type OurNewNullString sql.NullString


// For Selects
func (s *OurNewNullString) Scan(dbValue interface{}) error {
	// Use built-in for parsing
	existing := sql.NullString{String: s.String}
	err := existing.Scan(dbValue)

	s.Valid, s.String = existing.Valid, existing.String

	return err
}

// For inserting NULL values
func (s *OurNewNullString) Value() (driver.Value, error) {
	existing := sql.NullString{String: s.String}
	return existing.Value()
}

// for JSON serialization
func (ourNS OurNewNullString) MarshalJSON() ([]byte, error) {
	if ourNS.Valid {
		return json.Marshal(ourNS.String)
	}
	return json.Marshal(nil)
}

// for JSON deserialization
func (ourNS *OurNewNullString) UnmarshalJSON(text []byte) error {
	ourNS.Valid = false
	if string(text) == "null" {
		return nil
	}
	s := ""
	err := json.Unmarshal(text, &s)
	if err != nil {
		return err
	}
	// Yahooo
	ourNS.Valid = true
	ourNS.String = s
	return nil
}