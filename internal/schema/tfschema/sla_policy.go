package tfschema

import (
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/schemahelper"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var SLASchema = schema.Schema{
	Version: 1,
	MarkdownDescription: "SLA policy for Zendesk, " +
		"see [Documentation](https://developer.zendesk.com/api-reference/ticketing/business-rules/sla_policies) " +
		"for more information on configuration",
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"filter": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"all": schema.ListNestedAttribute{
					Description:  "Logical AND. All the conditions must be met",
					NestedObject: schemahelper.GetNestedConditionObject("sla"),
					Optional:     true,
				},
				"any": schema.ListNestedAttribute{
					Description:  "Logical OR. Any condition can be met",
					NestedObject: schemahelper.GetNestedConditionObject("sla"),
					Optional:     true,
				},
			},
		},
		"policy_metrics": schema.ListNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"priority": schema.StringAttribute{
						Required:    true,
						Description: "Priority that a ticket must match\n",
					},
					"metric": schema.StringAttribute{
						Required: true,
						MarkdownDescription: "The definition of the time that is being measured. " +
							"See [Metrics](https://developer.zendesk.com/api-reference/ticketing/business-rules/sla_policies/#metrics)",
					},
					"target": schema.Int64Attribute{
						Required:    true,
						Description: "The total time within which the end-state for a metric should be met, measured in minutes",
					},
					"business_hours": schema.BoolAttribute{
						Required:    true,
						Description: "Whether the metric targets are being measured in business hours or calendar hours",
					},
				},
			},
			MarkdownDescription: "Array of SLA Policy Metrics See " +
				"[Policy Metrics](https://developer.zendesk.com/api-reference/ticketing/business-rules/sla_policies/#policy-metric)",
		},
		"metrics_settings": schema.SingleNestedAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Settings for SLA metrics",
			Attributes: map[string]schema.Attribute{
				"first_reply_time": schema.SingleNestedAttribute{
					Optional: true,
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"activate_on_ticket_created_for_end_user": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_agent_ticket_created_for_end_user_with_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_light_agent_on_email_forward_ticket_from_end_user": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_agent_created_ticket_for_self": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"fulfill_on_agent_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
					},
				},
				"next_reply_time": schema.SingleNestedAttribute{
					Optional: true,
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"fulfill_on_non_requesting_agent_internal_note_after_activation": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_end_user_added_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_agent_requested_ticket_with_public_comment_or_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_light_agent_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
					},
				},
				"periodic_update_time": schema.SingleNestedAttribute{
					Optional: true,
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"activate_on_agent_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
					},
				},
			},
		},
		"title": schema.StringAttribute{
			Required: true,
		},
		"position": schema.Int64Attribute{
			Optional: true,
			Computed: true,
			Description: "Position of the SLA policy that determines the order they will be matched. " +
				"If not specified, the SLA policy is added as the last position",
		},
		"description": schema.StringAttribute{
			Description: "The description of the sla policy.",
			Optional:    true,
			Computed:    true,
		},
		"created_at": schema.StringAttribute{
			Description: "The time the sla policy was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the sla policy.",
			Computed:    true,
		},
	},
}

var SLASchemaV0 = schema.Schema{
	Version: 0,
	MarkdownDescription: "SLA policy for Zendesk, " +
		"see [Documentation](https://developer.zendesk.com/api-reference/ticketing/business-rules/sla_policies) " +
		"for more information on configuration",
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"filter": schema.SingleNestedAttribute{
			Required: true,
			Attributes: map[string]schema.Attribute{
				"all": schema.ListNestedAttribute{
					Description:  "Logical AND. All the conditions must be met",
					NestedObject: schemahelper.GetNestedConditionObjectV0("sla"),
					Optional:     true,
				},
				"any": schema.ListNestedAttribute{
					Description:  "Logical OR. Any condition can be met",
					NestedObject: schemahelper.GetNestedConditionObjectV0("sla"),
					Optional:     true,
				},
			},
		},
		"policy_metrics": schema.ListNestedAttribute{
			Required: true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"priority": schema.StringAttribute{
						Required:    true,
						Description: "Priority that a ticket must match\n",
					},
					"metric": schema.StringAttribute{
						Required: true,
						MarkdownDescription: "The definition of the time that is being measured. " +
							"See [Metrics](https://developer.zendesk.com/api-reference/ticketing/business-rules/sla_policies/#metrics)",
					},
					"target": schema.Int64Attribute{
						Required:    true,
						Description: "The total time within which the end-state for a metric should be met, measured in minutes",
					},
					"business_hours": schema.BoolAttribute{
						Required:    true,
						Description: "Whether the metric targets are being measured in business hours or calendar hours",
					},
				},
			},
			MarkdownDescription: "Array of SLA Policy Metrics See " +
				"[Policy Metrics](https://developer.zendesk.com/api-reference/ticketing/business-rules/sla_policies/#policy-metric)",
		},
		"metrics_settings": schema.SingleNestedAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Settings for SLA metrics",
			Attributes: map[string]schema.Attribute{
				"first_reply_time": schema.SingleNestedAttribute{
					Optional: true,
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"activate_on_ticket_created_for_end_user": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_agent_ticket_created_for_end_user_with_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_light_agent_on_email_forward_ticket_from_end_user": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_agent_created_ticket_for_self": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"fulfill_on_agent_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
					},
				},
				"next_reply_time": schema.SingleNestedAttribute{
					Optional: true,
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"fulfill_on_non_requesting_agent_internal_note_after_activation": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_end_user_added_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_agent_requested_ticket_with_public_comment_or_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"activate_on_light_agent_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
					},
				},
				"periodic_update_time": schema.SingleNestedAttribute{
					Optional: true,
					Computed: true,
					Attributes: map[string]schema.Attribute{
						"activate_on_agent_internal_note": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
					},
				},
			},
		},
		"title": schema.StringAttribute{
			Required: true,
		},
		"position": schema.Int64Attribute{
			Optional: true,
			Computed: true,
			Description: "Position of the SLA policy that determines the order they will be matched. " +
				"If not specified, the SLA policy is added as the last position",
		},
		"description": schema.StringAttribute{
			Description: "The description of the sla policy.",
			Optional:    true,
			Computed:    true,
		},
		"created_at": schema.StringAttribute{
			Description: "The time the sla policy was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the sla policy.",
			Computed:    true,
		},
	},
}
