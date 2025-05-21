package tfschema

import (
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/schemahelper"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var TriggerSchema = schema.Schema{
	Version: 1,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"actions": schemahelper.GetActionsListObject("trigger"),
		"conditions": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"all": schema.ListNestedAttribute{
					Description:  "Logical AND. All the conditions must be met",
					NestedObject: schemahelper.GetNestedConditionObject("trigger"),
					Optional:     true,
				},
				"any": schema.ListNestedAttribute{
					Description:  "Logical OR. Any condition can be met",
					NestedObject: schemahelper.GetNestedConditionObject("trigger"),
					Optional:     true,
				},
			},
		},
		"category_id": schema.Int64Attribute{
			Description: "The ID of the category the trigger belongs to",
			Required:    true,
		},
		"title": schema.StringAttribute{
			Required: true,
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the trigger is displayed or not.",
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
			Description: "The description of the trigger.",
			Optional:    true,
			Computed:    true,
		},
		"created_at": schema.StringAttribute{
			Description: "The time the trigger was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the trigger.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the trigger's details.",
			Computed:    true,
		},
	},
}

var TriggerSchemaV0 = schema.Schema{
	Version: 0,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"actions": schemahelper.GetActionsListObject("trigger"),
		"conditions": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"all": schema.ListNestedAttribute{
					Description:  "Logical AND. All the conditions must be met",
					NestedObject: schemahelper.GetNestedConditionObjectV0("trigger"),
					Optional:     true,
				},
				"any": schema.ListNestedAttribute{
					Description:  "Logical OR. Any condition can be met",
					NestedObject: schemahelper.GetNestedConditionObjectV0("trigger"),
					Optional:     true,
				},
			},
		},
		"category_id": schema.Int64Attribute{
			Description: "The ID of the category the trigger belongs to",
			Required:    true,
		},
		"title": schema.StringAttribute{
			Required: true,
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the trigger is displayed or not.",
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
			Description: "The description of the trigger.",
			Optional:    true,
			Computed:    true,
		},
		"created_at": schema.StringAttribute{
			Description: "The time the trigger was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the trigger.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the trigger's details.",
			Computed:    true,
		},
	},
}
