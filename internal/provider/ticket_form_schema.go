package provider

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var TicketFormSchema = schema.Schema{
	Version:     1,
	Description: "Ticket form attributes",
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"form_name": schema.StringAttribute{
			Required:    true,
			Description: "The name of the form.",
		},
		"end_user_display_name": schema.StringAttribute{
			Optional:    true,
			Description: "The name of the form that is displayed to an end user.",
		},
		"ticket_field_ids": schema.ListAttribute{
			Optional:    true,
			ElementType: types.Int64Type,
			MarkdownDescription: `
ids of all ticket fields which are in this ticket form. 
The products use the order of the ids to show the field values in the tickets. 

***IMPORTANT*** Zendesk applies default fields when creating a ticket form via API.
These field ids need to be included here, and the fields are 
* Subject
* Description
* Status
* Group
* Assignee
* Ticket Status (custom, need to find id via API)

then any other fields to include in the form can be added after these.
If there are less than 6 fields in this attribute, terraform will throw a validation error,
to help ensure that the default fields are included in the definition. If these values are not included,
but the total fields are 6+, terraform will throw an error trying to create the resource.`,
			Validators: []validator.List{
				listvalidator.SizeAtLeast(6),
			},
		},
		"agent_conditions": schema.MapNestedAttribute{
			Optional:     true,
			Description:  "Map of condition sets for agent workspaces. Key is the name of the parent ticket field of the conditions",
			NestedObject: GetNestedFormConditionObject("agent"),
		},
		"end_user_conditions": schema.MapNestedAttribute{
			Optional:     true,
			Description:  "Map of condition sets for end user products. Key is the name of the parent ticket field of the conditions",
			NestedObject: GetNestedFormConditionObject("end_user"),
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the ticket form is displayed or not.",
			Computed:    true,
			Optional:    true,
			Default:     booldefault.StaticBool(true),
		},
		"end_user_visible": schema.BoolAttribute{
			Description: "Is the form visible to the end user",
			Optional:    true,
			Computed:    true,
		},
		"default": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Is the form the default form for this account",
		},
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The position of this form among other forms in the account, i.e. dropdown",
		},
		"created_at": schema.StringAttribute{
			Description: "The time the ticket form was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the ticket form.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the ticket form's details.",
			Computed:    true,
		},
	},
}

var TicketFormSchemaV0 = schema.Schema{
	Description: "Ticket form attributes",
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"form_name": schema.StringAttribute{
			Required:    true,
			Description: "The name of the form.",
		},
		"end_user_display_name": schema.StringAttribute{
			Optional:    true,
			Description: "The name of the form that is displayed to an end user.",
		},
		"ticket_field_ids": schema.ListAttribute{
			Optional:    true,
			ElementType: types.Int64Type,
			MarkdownDescription: `
ids of all ticket fields which are in this ticket form. 
The products use the order of the ids to show the field values in the tickets. 

***IMPORTANT*** Zendesk applies default fields when creating a ticket form via API.
These field ids need to be included here, and the fields are 
* Subject
* Description
* Status
* Group
* Assignee
* Ticket Status (custom, need to find id via API)

then any other fields to include in the form can be added after these.
If there are less than 6 fields in this attribute, terraform will throw a validation error,
to help ensure that the default fields are included in the definition. If these values are not included,
but the total fields are 6+, terraform will throw an error trying to create the resource.`,
			Validators: []validator.List{
				listvalidator.SizeAtLeast(6),
			},
		},
		"agent_conditions": schema.SetNestedAttribute{
			Optional:     true,
			Description:  "Array of condition sets for agent workspaces.",
			NestedObject: GetNestedFormConditionObjectV0("agent"),
		},
		"end_user_conditions": schema.SetNestedAttribute{
			Optional:     true,
			Description:  "Array of condition sets for end user products",
			NestedObject: GetNestedFormConditionObjectV0("end_user"),
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the ticket form is displayed or not.",
			Computed:    true,
			Optional:    true,
			Default:     booldefault.StaticBool(true),
		},
		"end_user_visible": schema.BoolAttribute{
			Description: "Is the form visible to the end user",
			Optional:    true,
			Computed:    true,
		},
		"default": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Is the form the default form for this account",
		},
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The position of this form among other forms in the account, i.e. dropdown",
		},
		"created_at": schema.StringAttribute{
			Description: "The time the ticket form was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the ticket form.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the ticket form's details.",
			Computed:    true,
		},
	},
}
