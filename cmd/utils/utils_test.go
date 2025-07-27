package utils_test

import (
	"PSU/cmd/utils"
	"strconv"
	"testing"
)

func TestValidateProjectName(t *testing.T) {
	result := utils.ValidateProjectName("TEST TEST") 
	if result {
		t.Errorf("Test has spaces in name of the project got: %s want: %s", strconv.FormatBool(result), "false")
	}
}