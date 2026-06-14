package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {

	dob := time.Date(
		2000,
		time.January,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	age := CalculateAge(dob)

	currentYear := time.Now().Year()

	expected := currentYear - 2000

	if age != expected {
		t.Errorf(
			"expected %d, got %d",
			expected,
			age,
		)
	}
}