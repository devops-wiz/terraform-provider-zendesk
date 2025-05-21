package models

import (
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
	"strings"
)

type ActionResourceModel struct {
	Field               types.String `tfsdk:"field"`
	Target              types.String `tfsdk:"target"`
	NotificationSubject types.String `tfsdk:"notification_subject"`
	ContentType         types.String `tfsdk:"content_type"`
	SlackWorkspace      types.String `tfsdk:"slack_workspace"`
	SlackChannel        types.String `tfsdk:"slack_channel"`
	SlackTitle          types.String `tfsdk:"slack_title"`
	Value               types.String `tfsdk:"value"`
	CustomFieldID       types.Int64  `tfsdk:"custom_field_id"`
}

func getApiActionsFromTf(actions []ActionResourceModel) ([]zendesk.Action, diag.Diagnostics) {
	apiActions := make([]zendesk.Action, len(actions))
	diags := diag.Diagnostics{}

	for index, action := range actions {
		fieldName := action.Field.ValueString()

		if action.Field.ValueString() == "custom_field" {
			fieldName = fmt.Sprintf("%s%d", zendesk.ActionFieldCustomField, action.CustomFieldID.ValueInt64())
		}

		switch {
		case action.Field.ValueString() == zendesk.ActionNotificationZis.String():
			if action.Target.ValueString() != "slack" {
				panic("Unsupported target value " + action.Target.ValueString())
			}
			apiActions[index] = zendesk.Action{
				Field: fieldName,
				Value: zendesk.ParsedValue{ListData: []string{
					action.Target.ValueString(),
					action.SlackWorkspace.ValueString(),
					action.SlackChannel.ValueString(),
					action.SlackTitle.ValueString(),
					action.Value.ValueString(),
				}},
			}
		case action.Field.ValueString() == zendesk.ActionFieldNotificationUser.String() || action.Field.ValueString() == zendesk.ActionFieldNotificationGroup.String():
			apiActions[index] = zendesk.Action{
				Field: fieldName,
				Value: zendesk.ParsedValue{ListData: []string{action.Target.ValueString(), action.NotificationSubject.ValueString(), action.Value.ValueString()}},
			}
		case action.Field.ValueString() == zendesk.ActionFieldNotificationWebhook.String():
			apiActions[index] = zendesk.Action{
				Field: fieldName,
				Value: zendesk.ParsedValue{ListData: []string{action.Target.ValueString(), action.Value.ValueString()}},
			}
		case action.Field.ValueString() == zendesk.ActionSideConversationTicket.String() || action.ContentType.ValueString() != "":
			apiActions[index] = zendesk.Action{
				Field: fieldName,
				Value: zendesk.ParsedValue{ListData: []string{action.NotificationSubject.ValueString(), action.Value.ValueString(), action.Target.ValueString(), action.ContentType.ValueString()}},
			}
		case action.Field.ValueString() == zendesk.ActionSideConversationSlack.String():
			apiActions[index] = zendesk.Action{
				Field: fieldName,
				Value: zendesk.ParsedValue{ListData: []string{action.Value.ValueString(), action.Target.ValueString()}},
			}
		case !action.Target.IsNull():
			diags.AddAttributeWarning(path.Root("target"), "Unknown usage of 'target' attribute", fmt.Sprintf("Unknown usage of 'target' attribute with field %s", action.Field.ValueString()))
		default:
			apiActions[index] = zendesk.Action{
				Field: fieldName,
				Value: zendesk.ParsedValue{Data: action.Value.ValueString()},
			}
		}

	}
	return apiActions, diags
}

func getTfActionsFromApi(actions []zendesk.Action) ([]ActionResourceModel, diag.Diagnostics) {
	newTfActions := make([]ActionResourceModel, len(actions))
	diags := diag.Diagnostics{}

	for index, action := range actions {
		fieldName := action.Field
		cid := types.Int64Unknown()
		if strings.HasPrefix(action.Field, zendesk.ActionFieldCustomField.String()) {
			fieldName = "custom_field"
			id, _ := strings.CutPrefix(action.Field, zendesk.ActionFieldCustomField.String())
			convertedId, _ := strconv.ParseInt(id, 10, 64)
			cid = types.Int64Value(convertedId)
		}

		slackWorkspace := types.StringUnknown()
		slackChannel := types.StringUnknown()
		slackTitle := types.StringUnknown()
		target := types.StringUnknown()
		var newValue types.String
		contentType := types.StringUnknown()
		subject := types.StringUnknown()

		switch {
		case action.Field == zendesk.ActionNotificationZis.String():
			values := action.Value.ListData
			target = types.StringValue(values[0])
			slackWorkspace = types.StringValue(values[1])
			slackChannel = types.StringValue(values[2])
			slackTitle = types.StringValue(values[3])
			newValue = types.StringValue(values[4])
		case action.Field == zendesk.ActionFieldNotificationWebhook.String():
			values := action.Value.ListData
			target = types.StringValue(values[0])
			newValue = types.StringValue(values[1])
		case action.Field == zendesk.ActionFieldNotificationUser.String() || action.Field == zendesk.ActionFieldNotificationGroup.String():
			values := action.Value.ListData
			target = types.StringValue(values[0])
			subject = types.StringValue(values[1])
			if len(values) > 2 {
				newValue = types.StringValue(values[2])
			} else {
				newValue = types.StringNull()
			}
		case action.Field == zendesk.ActionSideConversationTicket.String():
			values := action.Value.ListData
			subject = types.StringValue(values[0])
			newValue = types.StringValue(values[1])
			if len(values) > 2 {
				if values[2] == "" {
					target = types.StringNull()
				} else {
					target = types.StringValue(values[2])
				}
				if len(values) > 3 {
					if values[3] == "" {
						contentType = types.StringNull()
					} else {
						contentType = types.StringValue(values[3])
					}
				}
			}
		case action.Field == zendesk.ActionSideConversationSlack.String():
			values := action.Value.ListData
			newValue = types.StringValue(values[0])
			target = types.StringValue(values[1])
		case len(action.Value.ListData) > 0:
			values := action.Value.ListData
			switch values[0] {
			case "days_from_now":
				target = types.StringValue(values[0])
				newValue = types.StringValue(values[1])
			default:
				diags.AddAttributeError(path.Root("actions"), "Unknown value type", fmt.Sprintf("Unknown value type: %s,\n full value obj %+v\n, field: %s", values[0], values, action.Field))
				return newTfActions, diags
			}
		case action.Value.Data == "" && len(action.Value.ListData) == 0:
			diags.AddAttributeWarning(path.Root("actions"), "Potential data issue, see details for data", fmt.Sprintf("%+v", action.Value))
			fallthrough
		default:
			newValue = types.StringValue(action.Value.Data)
		}

		newTfActions[index] = ActionResourceModel{
			Field: types.StringValue(fieldName),
			Value: newValue,
		}

		if !cid.IsUnknown() && !cid.IsNull() {
			newTfActions[index].CustomFieldID = cid
		}
		if !target.IsUnknown() && !target.IsNull() {
			newTfActions[index].Target = target
		}
		if !subject.IsUnknown() && !subject.IsNull() {
			newTfActions[index].NotificationSubject = subject
		}
		if !slackWorkspace.IsUnknown() && !slackWorkspace.IsNull() {
			newTfActions[index].SlackWorkspace = slackWorkspace
		}
		if !slackChannel.IsUnknown() && !slackChannel.IsNull() {
			newTfActions[index].SlackChannel = slackChannel
		}
		if !slackTitle.IsUnknown() && !slackTitle.IsNull() {
			newTfActions[index].SlackTitle = slackTitle
		}
		if !contentType.IsUnknown() && !contentType.IsNull() {
			newTfActions[index].ContentType = contentType
		}

	}
	return newTfActions, diags
}
