package utils_test

import (
	"PSU/cmd/utils"
	"testing"
)

func TestValidateProjectName(t *testing.T) {
	name := "Test Project"

	_, err := utils.ValidateProjectName(name)
	if err != nil {
		t.Errorf("%s: %s", name, err)
	}
}