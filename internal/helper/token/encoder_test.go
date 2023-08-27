package token

import (
	"testing"
)

func TestEncoder(t *testing.T) {
	_, err := EncodeToken("12345678", "john@dev.com")

	if err != nil {
		t.Error(err)
	}

}
