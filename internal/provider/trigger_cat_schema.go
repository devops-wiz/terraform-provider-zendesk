package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var TriggerCategorySchema = schema.Schema{
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
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The relative position of the ticket field on a ticket. Note that for accounts with ticket forms, positions are controlled by the different forms",
		},
		"created_at": schema.StringAttribute{
			Description: "The time the macro was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the macro.",
			Computed:    true,
		},
	},
}
