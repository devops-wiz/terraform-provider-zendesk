package provider

import (
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"strings"
)

var TicketFieldSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
			Description: "Ticket Field ID. Automatically assigned when created",
		},
		"title": schema.StringAttribute{
			Required:    true,
			Description: "The title of the ticket field",
		},
		"type": schema.StringAttribute{
			Required: true,
			Description: fmt.Sprintf(
				"Ticket Field Type, acceptable values include %s.\n",
				strings.Join(zendesk.ValidTicketFieldsTypes.StringSlice(), ", "),
			),
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Validators: []validator.String{
				stringvalidator.Any(
					stringvalidator.All(
						stringvalidator.OneOf(utils.SliceFilter(zendesk.ValidTicketFieldsTypes.StringSlice(), func(s string) bool {
							return s != zendesk.Tagger.String() && s != zendesk.Checkbox.String() && s != zendesk.Multiselect.String()
						})...),
					),
					stringvalidator.All(
						stringvalidator.OneOf(zendesk.Checkbox.String()),
						stringvalidator.AlsoRequires(path.MatchRoot("tag")),
					),
					stringvalidator.All(
						stringvalidator.OneOf(zendesk.Tagger.String(), zendesk.Multiselect.String()),
						stringvalidator.AlsoRequires(path.MatchRoot("custom_field_options")),
					),
				),
			},
		},
		"active": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Whether this field is available",
			Default:     booldefault.StaticBool(true),
		},
		"agent_description": schema.StringAttribute{
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Description: "A description of the ticket field that only agents can see",
		},
		"created_at": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Description: "The time the custom ticket field was created",
		},
		"portal_description": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Describes the purpose of the ticket field to users",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"editable_in_portal": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Whether this field is editable by end users in Help Center",
			Default:     booldefault.StaticBool(false),
		},
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The relative position of the ticket field on a ticket. Note that for accounts with ticket forms, positions are controlled by the different forms",
		},
		"regexp_for_validation": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "For 'regexp' fields only. The validation pattern for a field value to be deemed valid",
		},
		"required": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Description: "If true, agents must enter a value in the field to change the ticket status to solved",
			Default:     booldefault.StaticBool(false),
		},
		"required_in_portal": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Description: "If true, end users must enter a value in the field to create the request",
			Default:     booldefault.StaticBool(false),
		},
		"tag": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "For 'checkbox' fields only. A tag added to tickets when the checkbox field is selected",
		},
		"title_in_portal": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The title of the ticket field for end users in Help Center",
		},
		"updated_at": schema.StringAttribute{
			Computed:    true,
			Description: "The time the custom ticket field was last updated",
		},
		"url": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			Description: "The URL for this resource",
		},
		"visible_in_portal": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Whether this field is visible to end users in Help Center",
			Default:     booldefault.StaticBool(false),
		},
		"custom_field_options": schema.ListNestedAttribute{
			Optional:    true,
			Description: "Required and presented for a custom ticket field of type 'multiselect' or 'tagger'",
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"value": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
		"system_field_options": schema.ListNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "The name of the system field value",
						Computed:    true,
					},
					"value": schema.StringAttribute{
						Description: "The value of the system field value",
						Computed:    true,
					},
				},
			},
		},
	},
}
