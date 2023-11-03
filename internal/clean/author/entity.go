package author

import (
	"time"
)

type Author struct {
	Name string    `db:"name" json:"name"`
	DOB  time.Time `db:"dob" json:"dob"`
}
