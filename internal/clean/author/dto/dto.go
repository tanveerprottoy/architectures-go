package dto

import "time"

type CreateUserDTO struct {
	Name string    `json:"name" validate:"required"`
	DOB  time.Time `json:"dob" validate:"required"`
}

type UpdateUserDTO struct {
	Name string    `json:"name" validate:"required"`
	DOB  time.Time `json:"dob" validate:"required"`
}
