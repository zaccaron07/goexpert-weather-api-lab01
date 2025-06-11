package entity

import (
	"testing"
)

func TestZipcodeEntity(t *testing.T) {
	t.Run("should create a valid Zipcode entity", func(t *testing.T) {
		zipcode := "88820000"

		z, err := NewZipcode(zipcode)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if z.CEP != zipcode {
			t.Errorf("expected Zipcode %s, got %s", zipcode, z.CEP)
		}

	})

	t.Run("should return error for invalid Zipcode", func(t *testing.T) {
		invalidZipcodes := []string{
			"1234567",
			"123456789",
			"abcdefgh",
			"1234-567",
			"",
		}

		for _, zipcode := range invalidZipcodes {
			_, err := NewZipcode(zipcode)
			if err == nil {
				t.Errorf("expected error for invalid Zipcode %q, got nil", zipcode)
			}
		}
	})
}
