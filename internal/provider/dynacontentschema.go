package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var DynamicContentSchema = schema.Schema{
	Version: 0,
	MarkdownDescription: `
Dynamic content is a combination of a default version of some text and variants of the text in other languages. The combined content is represented by a placeholder such as {{dc.my_placeholder}}. Dynamic content is available only on the Professional plan and above. [Learn more](https://support.zendesk.com/hc/en-us/articles/203663356) in the Support Help Center.

This page contains the API reference for dynamic content items. See [Dynamic Content Item Variants](https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content_item_variants/) for the reference for variants.
`,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed:    true,
			Description: "Automatically assigned when creating items",
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:    true,
			Description: "The unique name of the item",
		},
		"placeholder": schema.StringAttribute{
			Computed:    true,
			Description: "Automatically generated placeholder for the item, derived from name",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"default_locale_id": schema.Int64Attribute{
			Required:            true,
			MarkdownDescription: `The default locale for the item. Must be one of the [locales the account has active.](https://developer.zendesk.com/api-reference/ticketing/account-configuration/locales/#list-locales)`,
		},
		"variants": schema.ListNestedAttribute{
			MarkdownDescription: `All variants within this item. See [Dynamic Content Item Variants](https://developer.zendesk.com/api-reference/ticketing/ticket-management/dynamic_content_item_variants/)`,
			Required:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "Automatically assigned when the variant is created",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"content": schema.StringAttribute{
						Description: "The content of the variant",
						Required:    true,
					},
					"locale_id": schema.Int64Attribute{
						Description: "An active locale",
						Required:    true,
					},
					"active": schema.BoolAttribute{
						Description: "If the variant is active and useable",
						Computed:    true,
						Optional:    true,
					},
					"default": schema.BoolAttribute{
						Description: "If the variant is the default for the item it belongs to",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	},
}
