package provider

import (
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

func GetNestedConditionObject(conditionType string) schema.NestedAttributeObject {

	var nestedConditionObject = schema.NestedAttributeObject{
		Validators: []validator.Object{
			&ConditionsValidator{ConfigType: conditionType},
		},
		Attributes: map[string]schema.Attribute{
			"field": schema.StringAttribute{
				MarkdownDescription: fmt.Sprintf(
					"Condition field to modify. Acceptable values: %s. See [Conditions Reference](https://developer.zendesk.com/documentation/ticketing/reference-guides/conditions-reference)",
					strings.Join(zendesk.ValidConditionOperatorValues.ValidKeys(), ", "),
				),
				Required: true,
			},
			"operator": schema.StringAttribute{
				Description: "A comparison operator",
				Optional:    true,
			},
			"value": schema.StringAttribute{
				Description: "The single value of the field",
				Optional:    true,
			},
			"values": schema.ListAttribute{
				Description: "A list of values for the field",
				ElementType: types.StringType,
				Optional:    true,
			},
			"custom_field_id": schema.Int64Attribute{
				Description: fmt.Sprintf("Required when field is set to 'custom_field' or 'ticket_field' for sla policys, ID of custom field to be modified by %s condition.", conditionType),
				Optional:    true,
			},
		},
	}

	return nestedConditionObject
}

func GetActionsListObject(actionType string) schema.ListNestedAttribute {

	var actionList = schema.ListNestedAttribute{
		MarkdownDescription: fmt.Sprintf("Required. An object describing what the %s will do. See [Actions reference](https://developer.zendesk.com/documentation/ticketing/reference-guides/actions-reference).", actionType),
		Required:            true,
		NestedObject: schema.NestedAttributeObject{
			Validators: []validator.Object{
				&ActionsValidator{
					ConfigType: actionType,
				},
			},
			Attributes: map[string]schema.Attribute{
				"field": schema.StringAttribute{
					Description: fmt.Sprintf(
						"The name of a ticket field to modify. Acceptable values: %s.",
						strings.Join(zendesk.ValidActionValuesMap.ValidKeys(), ","),
					),
					Required: true,
				},
				"target": schema.StringAttribute{
					MarkdownDescription: fmt.Sprintf("When field type is %s, %s, %s, or %s, this field is required as the target of the notification. *NOTE:* There are also some other cases where this is valid that are not documented on Zendesk side and need to be inferred from generated code", "notification_user",
						"notification_group",
						"notification_webhook", "notification_zis"),
					Optional: true,
				},
				"slack_workspace": schema.StringAttribute{
					Optional:    true,
					Description: fmt.Sprintf("When field type is %s, and target is 'slack', this is required.", zendesk.ActionNotificationZis),
				},
				"slack_channel": schema.StringAttribute{
					Optional:    true,
					Description: fmt.Sprintf("When field type is %s, and target is 'slack', this is required.", zendesk.ActionNotificationZis),
				},
				"slack_title": schema.StringAttribute{
					Optional:    true,
					Description: fmt.Sprintf("When field type is %s, and target is 'slack', this is required.", zendesk.ActionNotificationZis),
				},
				"notification_subject": schema.StringAttribute{
					Optional: true,
					MarkdownDescription: fmt.Sprintf("When field type is %s, or %s, this field is required as the target of the notification.", "notification_user",
						"notification_group"),
				},
				"content_type": schema.StringAttribute{
					Optional:    true,
					Description: fmt.Sprintf("When field type is %s, or %s, this field is required as the target of the notification.", zendesk.ActionSideConversationSlack, zendesk.ActionSideConversationTicket),
				},
				"value": schema.StringAttribute{
					Description: "The new value of the field, also the body of a notification for any of the notification action types.",
					Required:    true,
				},
				"custom_field_id": schema.Int64Attribute{
					Description: fmt.Sprintf("Required when field is set to 'custom_field', ID of custom field to be modified by %s action.", actionType),
					Optional:    true,
				},
			},
		},
	}

	return actionList
}

func GetNestedFormConditionObject(conditionType string) schema.NestedAttributeObject {

	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"field_value_map": schema.MapNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"child_field_conditions": schema.SetNestedAttribute{
							Required:    true,
							Description: "Set of child fields to show when the condition on the parent field is met",
							NestedObject: schema.NestedAttributeObject{
								Validators: []validator.Object{
									ConditionRequirementValidator{ConditionType: ConditionType(conditionType)},
								},
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{
									"id": schema.Int64Attribute{
										Required:    true,
										Description: "Id of the child field",
									},
									"is_required": schema.BoolAttribute{
										Required:    true,
										Description: "Is the child field required",
									},
									"required_on_statuses": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "How is the child field required",
										Attributes: map[string]schema.Attribute{
											"statuses": schema.ListAttribute{
												ElementType: types.StringType,
												Optional:    true,
												Description: fmt.Sprintf(
													"When type is set to SOME_STATUSES, list statuses child field is required. Valid statuses include: %s",
													strings.Join(zendesk.ValidRequirementStatuses, ", "),
												),
											},
											"type": schema.StringAttribute{
												Required: true,
												Description: fmt.Sprintf(
													"The type of required status, has values %v",
													zendesk.ValidRequirementTypes,
												),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
