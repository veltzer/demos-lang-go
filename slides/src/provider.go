// vet: expect "terraform-plugin-sdk" - needs the Terraform plugin SDK, which this repository does not vendor OMIT
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
	}
}
