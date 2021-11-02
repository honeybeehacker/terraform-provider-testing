package testing

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func Provider() *schema.Provider {

	return &schema.Provider{

		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			"testing_record": resourceRecord(),
		},
	}
}
