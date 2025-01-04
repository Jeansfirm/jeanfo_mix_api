package util_test

import (
	"jeanfo_mix/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidiate(t *testing.T) {
	validPasswd := "abcEfg1011"
	valid := util.IsValidPassword(validPasswd)
	assert.True(t, valid)

}
