package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var WebhookSchema = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"description": schema.StringAttribute{
			Description: "The description of the webhook.",
			Optional:    true,
			Computed:    true,
		},
		"authentication": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					Description: "Allowed values are 'api_key', 'basic_auth', or 'bearer_token'. Determines what authentication type is used.",
					Required:    true,
				},
				"add_position": schema.StringAttribute{
					Description: "Allowed value is 'header'. Determines where the authentication is added in the request.",
					Optional:    true,
					Computed:    true,
					Default:     stringdefault.StaticString("header"),
				},
				"credentials": schema.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "The name of the header that api authentication should use. Must use with 'value' attribute.",
							Optional:    true,
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "The value of the header that api authentication should use. Must use with 'name' attribute.",
							Optional:    true,
							Sensitive:   true,
						},
						"username": schema.StringAttribute{
							Description: "The username for basic authentication that should be used. Must use with 'password' attribute.",
							Optional:    true,
						},
						"password": schema.StringAttribute{
							Description: "The password for authentication that should be used. Must use with 'username' attribute.",
							Optional:    true,
							Sensitive:   true,
						},
						"token": schema.StringAttribute{
							Description: "The token for bearer authentication that should be use.",
							Optional:    true,
							Sensitive:   true,
						},
					},
				},
			},
		},
		"endpoint": schema.StringAttribute{
			Required:    true,
			Description: "The destination URL that the webhook notifies.",
		},
		"http_method": schema.StringAttribute{
			Required:    true,
			Description: "Allowed values are 'GET', 'POST', 'PUT', 'PATCH', or 'DELETE'. Determines what HTTP method to use in the webhook request.",
		},
		"custom_headers": schema.MapAttribute{
			Optional:    true,
			ElementType: types.StringType,
		},
		"request_format": schema.StringAttribute{
			Required:    true,
			Description: "Allowed values are 'json', 'xml', or 'form_encoded'. Determines what request format to use in webhook request payload.",
		},
		"status": schema.StringAttribute{
			Description: "Allowed values are 'active' or 'inactive'. Determines if the webhook is displayed or not.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("active"),
		},
		"subscriptions": schema.ListAttribute{
			Description: "List of events that the webhook is subscribed to. 'conditional_ticket_events' entry for Triggers.",
			Optional:    true,
			Computed:    true,
			ElementType: types.StringType,
			Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("conditional_ticket_events"),
			})),
		},
		"secret": schema.StringAttribute{
			Description:   "",
			Computed:      true,
			Sensitive:     true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		},
		"created_by": schema.StringAttribute{
			Description: "The user ID the webhook was originally created by.",
			Computed:    true,
		},
		"created_at": schema.StringAttribute{
			Description: "The time the webhook was created.",
			Computed:    true,
		},
		"updated_by": schema.StringAttribute{
			Description: "The user ID the webhook was last updated by.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the webhook.",
			Computed:    true,
		},
	},
}
