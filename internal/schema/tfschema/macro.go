package tfschema

import (
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/schemahelper"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var MacroSchema = schema.Schema{
	Version: 0,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"actions": schemahelper.GetActionsListObject("macro"),
		"title": schema.StringAttribute{
			Required: true,
		},
		"restriction": schema.SingleNestedAttribute{
			Description: "An object that describes who can access the macro. To give all agents access to the macro, omit this property.",
			Optional:    true,
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					Description: "Required. Allowed values are 'Group' or 'User'.",
					Required:    true,
					Validators: []validator.String{
						stringvalidator.OneOf([]string{"Group", "User"}...),
					},
				},
				"ids": schema.SetAttribute{
					Description: "The numeric IDs of the groups or users.",
					ElementType: types.Int64Type,
					Required:    true,
				},
			},
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the macro is displayed or not.",
			Computed:    true,
			Optional:    true,
			Default:     booldefault.StaticBool(true),
		},
		"description": schema.StringAttribute{
			Description: "The description of the macro.",
			Optional:    true,
			Computed:    true,
		},
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The position of a macro.",
		},
		"created_at": schema.StringAttribute{
			Description: "The time the macro was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the macro.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the macro's details.",
			Computed:    true,
		},
	},
}
