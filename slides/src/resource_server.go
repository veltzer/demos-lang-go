// vet: expect "terraform-plugin-sdk" - needs the Terraform plugin SDK, which this repository does not vendor OMIT
package main

// START OMIT
import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//END OMIT
