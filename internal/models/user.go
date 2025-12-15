package models

import "time"

// CreateUserRequest is the request body for creating a user
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB  string `json:"dob" validate:"required"` // Format: YYYY-MM-DD
}

// UpdateUserRequest is the request body for updating a user
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB  string `json:"dob" validate:"required"` // Format: YYYY-MM-DD
}

// UserResponse is the response body for a user
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}

// ErrorResponse is the response body for errors
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// ParseDOB parses date string to time.Time
func ParseDOB(dob string) (time.Time, error) {
	return time.Parse("2006-01-02", dob)
}

// FormatDOB formats time.Time to date string
func FormatDOB(t time.Time) string {
	return t.Format("2006-01-02")
}
