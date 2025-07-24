package confluent_schema

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_schema", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"

		// Fix the reference paths - check the actual Terraform provider schema
		// These should match the actual field names in the Confluent provider
		r.References["credentials.key"] = config.Reference{
			Type:      "v1:Secret",
			Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("data[\"key\"]",true)`,
		}

		r.References["credentials.secret"] = config.Reference{
			Type:      "v1:Secret",
			Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("data[\"secret\"]",true)`,
		}

		// Alternative approach using secret references if the provider supports it
		// Uncomment these if the provider uses *_secret_ref pattern
		/*
			r.References["credentials.key_secret_ref"] = config.Reference{
				Type: "v1:Secret",
			}
			r.References["credentials.secret_secret_ref"] = config.Reference{
				Type: "v1:Secret",
			}
		*/

		// Configure the credentials schema properly
		if credsSchema := r.TerraformResource.Schema["credentials"]; credsSchema != nil {
			credsSchema.Computed = true
			credsSchema.Optional = true

			// If credentials is a block/object, configure its nested fields
			if credsBlock := credsSchema.Elem; credsBlock != nil {
				if resourceSchema, ok := credsBlock.(*schema.Resource); ok {
					if keyField := resourceSchema.Schema["key"]; keyField != nil {
						keyField.Sensitive = true
					}
					if secretField := resourceSchema.Schema["secret"]; secretField != nil {
						secretField.Sensitive = true
					}
				}
			}
		}

		r.UseAsync = true
		r.Kind = "Schema"

		// Add sensitive fields configuration
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if key, ok := attr["credentials"].(map[string]any)["key"].(string); ok {
				conn["key"] = []byte(key)
			}
			if secret, ok := attr["credentials"].(map[string]any)["secret"].(string); ok {
				conn["secret"] = []byte(secret)
			}
			return conn, nil
		}
	})
}
