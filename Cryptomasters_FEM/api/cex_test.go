package api_test

import (
	"testing"

	"learninggo.com/go/Cryptomasters/api"
)

func TestAPICALL(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not found")
	}

}
