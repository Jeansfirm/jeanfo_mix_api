package util_test

import (
	"jeanfo_mix/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"ValidPassword1", "Passw0rd!", true},
		{"ValidPassword2", "Pssw0rd34", true},
		{"TooShort", "P@ss1", false},
		{"NoNumber", "Password!", false},
		{"NoUppercase", "p@ssw0rd", false},
		{"NoLowercase", "P@SSW0RD", false},
		{"NoSpecialChar", "Password1", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.IsValidPassword(tt.password)
			assert.Equal(t, tt.want, got)
		})
	}
}
