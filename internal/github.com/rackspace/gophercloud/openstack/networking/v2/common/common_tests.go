package common

import (
	"github.com/jrperritt/rack/internal/github.com/rackspace/gophercloud"
	"github.com/jrperritt/rack/internal/github.com/rackspace/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
