package role

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aerospike_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "aerospike"
		r.ShortGroup = "aerospike"
		r.Kind = "Role"
		if s, ok := r.TerraformResource.Schema["privileges"]; ok {
			s.Type = schema.TypeSet
			s.Elem = &schema.Resource{
				Schema: map[string]*schema.Schema{
					"privilege": {
						Type:     schema.TypeString,
						Required: true,
					},
					"namespace": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"set": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			}
		}
		// r.TerraformResource.Schema["privileges"].Elem.(*tfjson.Schema).Attributes["namespace"].Type = "string"
		// r.TerraformResource.Schema["privileges"].Elem.(*tfjson.Schema).Attributes["set"].Type = "string"
		// r.TerraformResource.Schema["privileges"].Elem.(*tfjson.Schema).Attributes["privilege"].Type = "string"
	})
}
