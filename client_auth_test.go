package huwlte

import "testing"

func TestGenerateConcentratedPasswordSHA256(t *testing.T) {
	for _, test := range []struct {
		Username          string
		Password          string
		VerificationToken string

		Result string
	}{
		{
			Username:          "admin",
			Password:          "admin1",
			VerificationToken: "3Y2mAEKBkny9Jr0Z+Oj7R/+8cq8lFzMr",
			Result:            "ZTMzYzc4MDdmM2FmYmQwNjkwYjliZWExOGRhZjI4MDAzYzM3MjVjOTYwZTI3YWYxYTBjNDU2ODYzNTI3OGQzOQ==",
		},
	} {
		result := generateConcentratedPasswordSHA256(test.Username, test.Password, test.VerificationToken)
		if result != test.Result {
			t.Errorf("Expected %s, got %s", test.Result, result)
		}

	}
}
