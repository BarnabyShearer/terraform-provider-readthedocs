package readthedocs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	rtd "github.com/BarnabyShearer/readthedocs/v3"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("READTHEDOCS_TOKEN", nil),
				Description: "API Token for authentication.",
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("READTHEDOCS_BASE_URL", "https://readthedocs.org/api/v3"),
				Description: "ReadTheDocs API base URL. Can be used to target the Read The Docs For Business API.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"readthedocs_project": resourceProject(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return rtd.NewClient(d.Get("token").(string), d.Get("base_url").(string)), nil
}
