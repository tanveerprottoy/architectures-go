package data

import (
	"database/sql"
	"encoding/json"
)

// compose sql.NullString in NullString
type NullString struct {
	sql.NullString
}

func MakeNullString(s string, v bool) NullString {
	return NullString{sql.NullString{String: s, Valid: v}}
}

// Scan implements the Scanner interface for NullString
/* func (n *NullString) Scan(val any) error {
	var str string
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &str)
		if err != nil {
			return err
		}
		n.String = str
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", val)
} */

// MarshalJSON for NullString
func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.String)
}

// UnmarshalJSON for NullString
func (n *NullString) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *string
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		n.Valid = true
		n.String = *v
	} else {
		n.Valid = false
	}
	return nil
}

/* func (n *NullString) Value() (driver.Value, error) {
    if !n.NullString.Valid {
        return nil, nil
    }
    return n.NullString.String, nil
} */

// compose sql.NullInt64 in NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Int64)
}

// UnmarshalJSON for NullInt64
func (n *NullInt64) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *int64
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		n.Valid = true
		n.Int64 = *v
	} else {
		n.Valid = false
	}
	return nil
}

// compose sql.NullFloat64 in NullInt64
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullInt64
func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Float64)
}

// UnmarshalJSON for NullInt64
func (n *NullFloat64) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *float64
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		n.Valid = true
		n.Float64 = *v
	} else {
		n.Valid = false
	}
	return nil
}

// compose sql.NullBool in NullBool
type NullBool struct {
	sql.NullBool
}

// MarshalJSON for NullBool
func (n NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Bool)
}

// UnmarshalJSON for NullBool
func (n *NullBool) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *bool
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		n.Valid = true
		n.Bool = *v
	} else {
		n.Valid = false
	}
	return nil
}
