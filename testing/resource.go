package testing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRecord() *schema.Resource {

	return &schema.Resource{

		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,

		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var data string
	if val, ok := rd.GetOk("data"); ok {
		data = val.(string)
	}
	rd.SetId(data)

	return resourceRead(ctx, rd, meta)
}

func resourceRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRead(ctx, rd, meta)
}

func resourceDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}
