package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Birthday already occurred this year",
			dob:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 2000,
		},
		{
			name:     "Birthday not occurred yet this year",
			dob:      time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: func() int {
				now := time.Now()
				age := now.Year() - 2000
				if now.Month() < 12 || (now.Month() == 12 && now.Day() < 31) {
					age--
				}
				return age
			}(),
		},
		{
			name:     "Born today years ago",
			dob:      time.Date(2000, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 2000,
		},
		{
			name:     "Born yesterday years ago",
			dob:      time.Date(2000, time.Now().Month(), time.Now().Day()-1, 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 2000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAge(tt.dob)
			if got != tt.expected {
				t.Errorf("CalculateAge() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestCalculateAge_EdgeCases(t *testing.T) {
	// Test leap year birthday (Feb 29)
	leapYearDOB := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
	age := CalculateAge(leapYearDOB)
	
	now := time.Now()
	expectedAge := now.Year() - 2000
	if now.Month() < 2 || (now.Month() == 2 && now.Day() < 29) {
		expectedAge--
	}
	
	if age != expectedAge {
		t.Errorf("Leap year CalculateAge() = %v, expected %v", age, expectedAge)
	}
}
