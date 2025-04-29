package auth_test

import (
	"net/http"
	"testing"

	"github.com/pjagbuya/learn-cicd-starter/internal/auth" // Ensure your import path is correct
)

func TestGetAPIKey_MalformedAuthorizationHeader_TooManyParts(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey part1 part2 part3")
	key, err := auth.GetAPIKey(headers)
	t.Logf("GetAPIKey returned: key='%s', error='%v'", key, err)

	if err != nil {
		t.Fatalf("Test failed: unexpected error: %v", err)
	}

}
