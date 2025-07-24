package confluent_schema

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_schema", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"

		// Based on your original working references, keeping the array notation
		r.References["credentials.0.key_secret_ref"] = config.Reference{
			Type: "v1:Secret",
		}

		r.References["credentials.0.secret_secret_ref"] = config.Reference{
			Type: "v1:Secret",
		}

		// Configure the credentials schema properly
		if credsSchema := r.TerraformResource.Schema["credentials"]; credsSchema != nil {
			credsSchema.Computed = true
			credsSchema.Optional = true
		}

		r.UseAsync = true
		r.Kind = "Schema"

		// Remove the problematic AdditionalConnectionDetailsFn for now
		// You can add it back once the basic functionality is working
	})
}
