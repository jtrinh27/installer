package aws

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// IsForbidden returns true if and only if the input error is an HTTP
// 403 error from the AWS API.
func IsForbidden(err error) bool {
	requestError, ok := err.(awserr.RequestFailure)
	return ok && requestError.StatusCode() == http.StatusForbidden
}

// GetBaseDomain returns a base domain chosen from among the account's
// public routes.
func GetBaseDomain() (string, error) {
	return "sd.spawar.navy.mil", nil
}
