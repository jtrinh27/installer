package aws

import (
	"strings"

	awss "github.com/aws/aws-sdk-go/aws/session"
)

//go:generate mockgen -source=./route53.go -destination=mock/awsroute53_generated.go -package=mock

// API represents the calls made to the API.
type API interface {
}

// Client makes calls to the AWS Route53 API.
type Client struct {
	ssn *awss.Session
}

// NewClient initializes a client with a session.
func NewClient(ssn *awss.Session) *Client {
	client := &Client{
		ssn: ssn,
	}
	return client
}

func skipRecord(recordName string, dottedClusterDomain string) bool {
	// skip record sets that are not sub-domains of the cluster domain. Such record sets may exist for
	// hosted zones that are used for other clusters or other purposes.
	if !strings.HasSuffix(recordName, "."+dottedClusterDomain) {
		return true
	}
	// skip record sets that are the cluster domain. Record sets for the cluster domain are fine. If the
	// hosted zone has the name of the cluster domain, then there will be NS and SOA record sets for the
	// cluster domain.
	if len(recordName) == len(dottedClusterDomain) {
		return true
	}

	return false
}
