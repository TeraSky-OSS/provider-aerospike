package user

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aerospike_user", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "aerospike"
		r.ShortGroup = "aerospike"
		r.Kind = "User"
	})
}
