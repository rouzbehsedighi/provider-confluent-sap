package confluent_schema

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_schema", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.References["credentials.0.key_secret_ref"] = config.Reference{
			Type: "Secret",
		}

		r.References["credentials.0.secret_secret_ref"] = config.Reference{
			Type: "Secret",
		}

		// Important: explicitly map status fields
		r.TerraformResource.Schema["credentials"].Computed = true
		r.TerraformResource.Schema["credentials"].Optional = true
		r.UseAsync = true
		r.Kind = "Schema"
	})
}
