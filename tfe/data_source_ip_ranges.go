package tfe

import (
	"fmt"
	"log"

	tfe "github.com/hashicorp/go-tfe"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTFEIPRanges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTFEIPRangesRead,

		Schema: map[string]*schema.Schema{
			"api": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"notifications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sentinel": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vcs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceTFEIPRangesRead(d *schema.ResourceData, meta interface{}) error {
	tfeClient := meta.(*tfe.Client)

	log.Printf("[DEBUG] Reading IP Ranges")
	ipRanges, err := tfeClient.Meta.IPRanges.Read(ctx, "")
	if err != nil {
		return fmt.Errorf("Error retrieving IP ranges: %w", err)
	}

	d.SetId("ip-ranges")
	d.Set("api", ipRanges.API)
	d.Set("notifications", ipRanges.Notifications)
	d.Set("sentinel", ipRanges.Sentinel)
	d.Set("vcs", ipRanges.VCS)

	return nil
}
