package confluent_schema

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_schema", func(r *config.Resource) {
		r.ShortGroup = "confluent"
		r.Kind = "Schema"
		r.UseAsync = true

		// Configure references for credentials
		r.References["credentials.key_secret_ref"] = config.Reference{
			Type: "Secret",
		}
		r.References["credentials.secret_secret_ref"] = config.Reference{
			Type: "Secret",
		}

		// Mark credentials as computed and optional
		if schema, ok := r.TerraformResource.Schema["credentials"]; ok {
			schema.Computed = true
			schema.Optional = true
		}
	})
}
