package availabilityzones

import (
	"github.com/cloud-barista/nhncloud-sdk-for-drv"
	"github.com/cloud-barista/nhncloud-sdk-for-drv/pagination"
)

// List will return the existing availability zones.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.SinglePageBase(r)}
	})
}
