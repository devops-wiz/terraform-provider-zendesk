package tfschema

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var GroupSchema = schema.Schema{
	Version: 0,
	MarkdownDescription: "Group for Zendesk, " +
		"see [Documentation](https://developer.zendesk.com/api-reference/ticketing/groups/groups/) " +
		"for more information on configuration",
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"description": schema.StringAttribute{
			Description: "The description of the group.",
			Optional:    true,
			Computed:    true,
		},
		"default": schema.BoolAttribute{
			Computed:    true,
			Description: "If the group is the default one for the account",
		},
		"is_public": schema.BoolAttribute{
			Computed: true,
			Optional: true,
			Description: "If true, the group is public. If false, the group is private." +
				" Changing a private group to a public group will recreate the resource. Default value for provider is false",
			Default: booldefault.StaticBool(false),
		},
		"deleted": schema.BoolAttribute{
			Computed: true,
		},
		"created_at": schema.StringAttribute{
			Description: "The time the group was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the group.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Description: "The URL for this resource",
		},
	},
}
