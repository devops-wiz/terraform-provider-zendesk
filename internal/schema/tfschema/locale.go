package tfschema

import "github.com/hashicorp/terraform-plugin-framework/datasource/schema"

var LocaleSchema = schema.Schema{
	Description: "Datasource to get a single locale based on locale code. Ex: English (United States) would be 'en-us'",
	Attributes: map[string]schema.Attribute{
		"code": schema.StringAttribute{
			Description: "Locale code to filter locales from Zendesk",
			Required:    true,
		},
		"locale": schema.SingleNestedAttribute{
			Computed:    true,
			Description: "Locale returned based on inputted code",
			Attributes: map[string]schema.Attribute{
				"id": schema.Int64Attribute{
					Computed:    true,
					Description: "Locale ID",
				},
				"locale_code": schema.StringAttribute{
					Computed:    true,
					Description: "Locale Code",
				},
				"name": schema.StringAttribute{
					Computed:    true,
					Description: "Locale name",
				},
			},
		},
	},
}
