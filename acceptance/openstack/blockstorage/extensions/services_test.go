//go:build acceptance || blockstorage
// +build acceptance blockstorage

package extensions

import (
	"testing"

	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/clients"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/acceptance/tools"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/openstack/blockstorage/extensions/services"
	th "github.com/cloud-barista/nhncloud-sdk-for-drv/testhelper"
)

func TestServicesList(t *testing.T) {
	clients.SkipRelease(t, "stable/mitaka")
	clients.RequireAdmin(t)

	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	allPages, err := services.List(blockClient, services.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}
}
