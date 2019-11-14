package terraform_kintone

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("KINTONE_HOST", ""),
				Description:  "kintone host. ex sushi.kintone-dev.ninja",
				ValidateFunc: validation.NoZeroValues,
			},
			"user": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("KINTONE_USER", ""),
				Description:  "kintone user",
				ValidateFunc: validation.NoZeroValues,
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("KINTONE_PASSWORD", ""),
				Description:  "kintone password",
				ValidateFunc: validation.NoZeroValues,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"kintone_application": resourceKintoneApplication(),
			"kintone_record":      resourceKintoneRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &Config{
		Host:     d.Get("host").(string),
		User:     d.Get("user").(string),
		Password: d.Get("password").(string),
	}, nil
}
