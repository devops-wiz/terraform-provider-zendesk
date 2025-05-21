package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var AutomationSchema = schema.Schema{
	Version: 1,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"actions": GetActionsListObject("automation"),
		"conditions": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"all": schema.ListNestedAttribute{
					Description:  "Logical AND. All the conditions must be met",
					NestedObject: GetNestedConditionObject("automation"),
					Required:     true,
				},
				"any": schema.ListNestedAttribute{
					Description:  "Logical OR. Any condition can be met",
					NestedObject: GetNestedConditionObject("automation"),
					Optional:     true,
				},
			},
		},
		"title": schema.StringAttribute{
			Required: true,
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the automation is displayed or not.",
			Computed:    true,
			Optional:    true,
			Default:     booldefault.StaticBool(true),
		},
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The relative position of the ticket field on a ticket. Note that for accounts with ticket forms, positions are controlled by the different forms",
		},
		"description": schema.StringAttribute{
			Description: "The description of the automation.",
			Optional:    true,
			Computed:    true,
		},
		"created_at": schema.StringAttribute{
			Description: "The time the automation was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the automation.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the automation's details.",
			Computed:    true,
		},
	},
}

var AutomationSchemaV0 = schema.Schema{
	Version: 0,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"actions": GetActionsListObject("automation"),
		"conditions": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"all": schema.ListNestedAttribute{
					Description:  "Logical AND. All the conditions must be met",
					NestedObject: GetNestedConditionObjectV0("automation"),
					Required:     true,
				},
				"any": schema.ListNestedAttribute{
					Description:  "Logical OR. Any condition can be met",
					NestedObject: GetNestedConditionObjectV0("automation"),
					Optional:     true,
				},
			},
		},
		"title": schema.StringAttribute{
			Required: true,
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the automation is displayed or not.",
			Computed:    true,
			Optional:    true,
			Default:     booldefault.StaticBool(true),
		},
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The relative position of the ticket field on a ticket. Note that for accounts with ticket forms, positions are controlled by the different forms",
		},
		"description": schema.StringAttribute{
			Description: "The description of the automation.",
			Optional:    true,
			Computed:    true,
		},
		"created_at": schema.StringAttribute{
			Description: "The time the automation was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the automation.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the automation's details.",
			Computed:    true,
		},
	},
}
