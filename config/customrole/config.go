package customrole

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_custom_role", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "CustomRole"
	})
}
