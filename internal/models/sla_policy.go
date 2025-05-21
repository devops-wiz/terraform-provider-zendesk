package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"reflect"
)

var _ ResourceTransformWithID[zendesk.SLAPolicy] = &SLAPolicyResourceModel{}

type SLAPolicyResourceModel struct {
	ID              types.Int64                    `tfsdk:"id"`
	Title           types.String                   `tfsdk:"title"`
	Description     types.String                   `tfsdk:"description"`
	Position        types.Int64                    `tfsdk:"position"`
	Filter          ConditionsResourceModel        `tfsdk:"filter"`
	PolicyMetrics   []SLAPolicyMetricResourceModel `tfsdk:"policy_metrics"`
	MetricsSettings types.Object                   `tfsdk:"metrics_settings"`
	CreatedAt       types.String                   `tfsdk:"created_at"`
	UpdatedAt       types.String                   `tfsdk:"updated_at"`
}

type SLAPolicyResourceModelV0 struct {
	ID              types.Int64                    `tfsdk:"id"`
	Title           types.String                   `tfsdk:"title"`
	Description     types.String                   `tfsdk:"description"`
	Position        types.Int64                    `tfsdk:"position"`
	Filter          ConditionsResourceModelV0      `tfsdk:"filter"`
	PolicyMetrics   []SLAPolicyMetricResourceModel `tfsdk:"policy_metrics"`
	MetricsSettings types.Object                   `tfsdk:"metrics_settings"`
	CreatedAt       types.String                   `tfsdk:"created_at"`
	UpdatedAt       types.String                   `tfsdk:"updated_at"`
}

func (s *SLAPolicyResourceModel) GetID() int64 {
	return s.ID.ValueInt64()
}

type SLAPolicyMetricResourceModel struct {
	Priority      types.String `tfsdk:"priority"`
	Metric        types.String `tfsdk:"metric"`
	Target        types.Int64  `tfsdk:"target"`
	BusinessHours types.Bool   `tfsdk:"business_hours"`
}

type MetricSettingsResourceModel struct {
	FirstReplyTime     types.Object `tfsdk:"first_reply_time"`
	NextReplyTime      types.Object `tfsdk:"next_reply_time"`
	PeriodicUpdateTime types.Object `tfsdk:"periodic_update_time"`
}

func (m MetricSettingsResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"first_reply_time":     types.ObjectType{AttrTypes: FirstReplyTimeResourceModel{}.AttributeTypes()},
		"next_reply_time":      types.ObjectType{AttrTypes: NextReplyTimeResourceModel{}.AttributeTypes()},
		"periodic_update_time": types.ObjectType{AttrTypes: PeriodicUpdateTimeResourceModel{}.AttributeTypes()},
	}
}

type FirstReplyTimeResourceModel struct {
	ActivateOnTicketCreatedForEndUser                      types.Bool `tfsdk:"activate_on_ticket_created_for_end_user"`
	ActivateOnAgentTicketCreatedForEndUserWithInternalNote types.Bool `tfsdk:"activate_on_agent_ticket_created_for_end_user_with_internal_note"`
	ActivateOnLightAgentOnEmailForwardTicketFromEndUser    types.Bool `tfsdk:"activate_on_light_agent_on_email_forward_ticket_from_end_user"`
	ActivateOnAgentCreatedTicketForSelf                    types.Bool `tfsdk:"activate_on_agent_created_ticket_for_self"`
	FulfillOnAgentInternalNote                             types.Bool `tfsdk:"fulfill_on_agent_internal_note"`
}

func (f FirstReplyTimeResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_on_ticket_created_for_end_user":                          types.BoolType,
		"activate_on_agent_ticket_created_for_end_user_with_internal_note": types.BoolType,
		"activate_on_light_agent_on_email_forward_ticket_from_end_user":    types.BoolType,
		"activate_on_agent_created_ticket_for_self":                        types.BoolType,
		"fulfill_on_agent_internal_note":                                   types.BoolType,
	}
}

type NextReplyTimeResourceModel struct {
	FulfillOnNonRequestingAgentInternalNoteAfterActivation        types.Bool `tfsdk:"fulfill_on_non_requesting_agent_internal_note_after_activation"`
	ActivateOnEndUserAddedInternalNote                            types.Bool `tfsdk:"activate_on_end_user_added_internal_note"`
	ActivateOnAgentRequestedTicketWithPublicCommentOrInternalNote types.Bool `tfsdk:"activate_on_agent_requested_ticket_with_public_comment_or_internal_note"`
	ActivateOnLightAgentInternalNote                              types.Bool `tfsdk:"activate_on_light_agent_internal_note"`
}

func (n NextReplyTimeResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"fulfill_on_non_requesting_agent_internal_note_after_activation":          types.BoolType,
		"activate_on_end_user_added_internal_note":                                types.BoolType,
		"activate_on_agent_requested_ticket_with_public_comment_or_internal_note": types.BoolType,
		"activate_on_light_agent_internal_note":                                   types.BoolType,
	}
}

type PeriodicUpdateTimeResourceModel struct {
	ActivateOnAgentInternalNote types.Bool `tfsdk:"activate_on_agent_internal_note"`
}

func (p PeriodicUpdateTimeResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"activate_on_agent_internal_note": types.BoolType,
	}
}

func (s *SLAPolicyResourceModel) GetApiModelFromTfModel(ctx context.Context) (newSLA zendesk.SLAPolicy, diags diag.Diagnostics) {
	newConditions, diags := getApiConditionsFromTf(ctx, s.Filter)

	if diags.HasError() {
		return zendesk.SLAPolicy{}, diags
	}

	var newMetricsSettings zendesk.MetricSettings

	if !s.MetricsSettings.IsNull() && !s.MetricsSettings.IsUnknown() {
		newMetricsSettings, diags = getApiMetricsSettingsFromTf(ctx, s.MetricsSettings)
		if diags.HasError() {
			return newSLA, diags
		}
	}

	var newPolicyMetrics = getApiPolicyMetricFromTf(s.PolicyMetrics)

	newSLA = zendesk.SLAPolicy{
		Title:          s.Title.ValueString(),
		Description:    s.Description.ValueString(),
		MetricSettings: newMetricsSettings,
		Filter:         newConditions,
		PolicyMetrics:  newPolicyMetrics,
	}

	if !s.Position.IsNull() && !s.Position.IsUnknown() {
		newSLA.Position = s.Position.ValueInt64()
	}

	return newSLA, diags
}

func getApiMetricsSettingsFromTf(ctx context.Context, tfSettingsObj types.Object) (settings zendesk.MetricSettings, diags diag.Diagnostics) {
	var tfSettings MetricSettingsResourceModel

	asOptions := basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	}
	diags = tfSettingsObj.As(ctx, &tfSettings, asOptions)

	if diags.HasError() {
		return settings, diags
	}

	var tfSettingsFirstReply FirstReplyTimeResourceModel

	diags.Append(tfSettings.FirstReplyTime.As(ctx, &tfSettingsFirstReply, asOptions)...)

	if diags.HasError() {
		return settings, diags
	}

	var tfSettingsPeriodicUpdate PeriodicUpdateTimeResourceModel

	diags.Append(tfSettings.PeriodicUpdateTime.As(ctx, &tfSettingsPeriodicUpdate, asOptions)...)

	if diags.HasError() {
		return settings, diags
	}
	var tfSettingsNextReply NextReplyTimeResourceModel

	diags.Append(tfSettings.NextReplyTime.As(ctx, &tfSettingsNextReply, asOptions)...)

	if diags.HasError() {
		return settings, diags
	}

	settings = zendesk.MetricSettings{
		FirstReplyTime: zendesk.FirstReplyTime{
			ActivateOnTicketCreatedForEndUser:                      tfSettingsFirstReply.ActivateOnTicketCreatedForEndUser.ValueBool(),
			ActivateOnAgentCreatedTicketForSelf:                    tfSettingsFirstReply.ActivateOnAgentCreatedTicketForSelf.ValueBool(),
			FulfillOnAgentInternalNote:                             tfSettingsFirstReply.FulfillOnAgentInternalNote.ValueBool(),
			ActivateOnLightAgentOnEmailForwardTicketFromEndUser:    tfSettingsFirstReply.ActivateOnLightAgentOnEmailForwardTicketFromEndUser.ValueBool(),
			ActivateOnAgentTicketCreatedForEndUserWithInternalNote: tfSettingsFirstReply.ActivateOnAgentTicketCreatedForEndUserWithInternalNote.ValueBool(),
		},
		NextReplyTime: zendesk.NextReplyTime{
			ActivateOnAgentRequestedTicketWithPublicCommentOrInternalNote: tfSettingsNextReply.ActivateOnAgentRequestedTicketWithPublicCommentOrInternalNote.ValueBool(),
			ActivateOnEndUserAddedInternalNote:                            tfSettingsNextReply.ActivateOnEndUserAddedInternalNote.ValueBool(),
			ActivateOnLightAgentInternalNote:                              tfSettingsNextReply.ActivateOnLightAgentInternalNote.ValueBool(),
			FulfillOnNonRequestingAgentInternalNoteAfterActivation:        tfSettingsNextReply.FulfillOnNonRequestingAgentInternalNoteAfterActivation.ValueBool(),
		},
		PeriodicUpdateTime: zendesk.PeriodicUpdateTime{
			ActivateOnAgentInternalNote: tfSettingsPeriodicUpdate.ActivateOnAgentInternalNote.ValueBool(),
		},
	}

	return settings, diags
}

func getApiPolicyMetricFromTf(metrics []SLAPolicyMetricResourceModel) []zendesk.SLAPolicyMetric {
	slaMetrics := make([]zendesk.SLAPolicyMetric, len(metrics))

	for i, metric := range metrics {
		slaMetrics[i] = zendesk.SLAPolicyMetric{
			Priority:      metric.Priority.ValueString(),
			Metric:        metric.Metric.ValueString(),
			Target:        int(metric.Target.ValueInt64()),
			BusinessHours: metric.BusinessHours.ValueBool(),
		}
	}

	return slaMetrics
}

func (s *SLAPolicyResourceModel) GetTfModelFromApiModel(ctx context.Context, sla zendesk.SLAPolicy) (diags diag.Diagnostics) {
	newTfConditions, diags := getTfConditionsFromApi(ctx, sla.Filter)

	if diags.HasError() {
		return diags
	}

	newTfMetricsSettings, diags := getTfMetricsSettingsFromApi(ctx, sla.MetricSettings)

	if diags.HasError() {
		return diags
	}

	newTfPolicyMetrics := getTfPolicyMetricsFromApi(sla.PolicyMetrics)

	*s = SLAPolicyResourceModel{
		ID:              types.Int64Value(sla.ID),
		Title:           types.StringValue(sla.Title),
		Description:     types.StringValue(sla.Description),
		Position:        types.Int64Value(sla.Position),
		Filter:          newTfConditions,
		MetricsSettings: newTfMetricsSettings,
		PolicyMetrics:   newTfPolicyMetrics,
		CreatedAt:       types.StringValue(sla.CreatedAt.UTC().String()),
		UpdatedAt:       types.StringValue(sla.UpdatedAt.UTC().String()),
	}

	return diags
}

func getTfMetricsSettingsFromApi(ctx context.Context, settings zendesk.MetricSettings) (metricsSettingsObj types.Object, diags diag.Diagnostics) {

	var metricsSettingsFirstReplyObj types.Object

	var metricsSettingsNextReplyObj types.Object

	var metricsSettingsPeriodicUpdateObj types.Object

	if !reflect.DeepEqual(settings.FirstReplyTime, zendesk.FirstReplyTime{}) {
		metricsSettingsFirstReply := FirstReplyTimeResourceModel{
			ActivateOnTicketCreatedForEndUser:                      types.BoolValue(settings.FirstReplyTime.ActivateOnTicketCreatedForEndUser),
			ActivateOnAgentTicketCreatedForEndUserWithInternalNote: types.BoolValue(settings.FirstReplyTime.ActivateOnAgentTicketCreatedForEndUserWithInternalNote),
			ActivateOnLightAgentOnEmailForwardTicketFromEndUser:    types.BoolValue(settings.FirstReplyTime.ActivateOnLightAgentOnEmailForwardTicketFromEndUser),
			ActivateOnAgentCreatedTicketForSelf:                    types.BoolValue(settings.FirstReplyTime.ActivateOnAgentCreatedTicketForSelf),
			FulfillOnAgentInternalNote:                             types.BoolValue(settings.FirstReplyTime.FulfillOnAgentInternalNote),
		}

		metricsSettingsFirstReplyObj, diags = types.ObjectValueFrom(ctx, metricsSettingsFirstReply.AttributeTypes(), metricsSettingsFirstReply)
		if diags.HasError() {
			return metricsSettingsObj, diags
		}
	} else {
		metricsSettingsFirstReplyObj = types.ObjectNull(FirstReplyTimeResourceModel{}.AttributeTypes())
	}

	if !reflect.DeepEqual(settings.NextReplyTime, zendesk.NextReplyTime{}) {
		metricsSettingsNextReply := NextReplyTimeResourceModel{
			FulfillOnNonRequestingAgentInternalNoteAfterActivation:        types.BoolValue(settings.NextReplyTime.FulfillOnNonRequestingAgentInternalNoteAfterActivation),
			ActivateOnEndUserAddedInternalNote:                            types.BoolValue(settings.NextReplyTime.ActivateOnEndUserAddedInternalNote),
			ActivateOnAgentRequestedTicketWithPublicCommentOrInternalNote: types.BoolValue(settings.NextReplyTime.ActivateOnAgentRequestedTicketWithPublicCommentOrInternalNote),
			ActivateOnLightAgentInternalNote:                              types.BoolValue(settings.NextReplyTime.ActivateOnLightAgentInternalNote),
		}

		metricsSettingsNextReplyObj, diags = types.ObjectValueFrom(ctx, metricsSettingsNextReply.AttributeTypes(), metricsSettingsNextReply)
		if diags.HasError() {
			return metricsSettingsObj, diags
		}
	} else {
		metricsSettingsNextReplyObj = types.ObjectNull(NextReplyTimeResourceModel{}.AttributeTypes())
	}

	if !reflect.DeepEqual(settings.PeriodicUpdateTime, zendesk.PeriodicUpdateTime{}) {
		metricsSettingsPeriodicUpdate := PeriodicUpdateTimeResourceModel{
			ActivateOnAgentInternalNote: types.BoolValue(settings.PeriodicUpdateTime.ActivateOnAgentInternalNote),
		}

		metricsSettingsPeriodicUpdateObj, diags = types.ObjectValueFrom(ctx, metricsSettingsPeriodicUpdate.AttributeTypes(), metricsSettingsPeriodicUpdate)
		if diags.HasError() {
			return metricsSettingsObj, diags
		}
	} else {
		metricsSettingsPeriodicUpdateObj = types.ObjectNull(PeriodicUpdateTimeResourceModel{}.AttributeTypes())
	}

	metricsSettings := MetricSettingsResourceModel{
		FirstReplyTime:     metricsSettingsFirstReplyObj,
		NextReplyTime:      metricsSettingsNextReplyObj,
		PeriodicUpdateTime: metricsSettingsPeriodicUpdateObj,
	}

	metricsSettingsObj, diags = types.ObjectValueFrom(ctx, metricsSettings.AttributeTypes(), metricsSettings)

	return metricsSettingsObj, diags
}

func getTfPolicyMetricsFromApi(metrics []zendesk.SLAPolicyMetric) []SLAPolicyMetricResourceModel {
	tfMetrics := make([]SLAPolicyMetricResourceModel, len(metrics))

	for i, metric := range metrics {
		tfMetrics[i] = SLAPolicyMetricResourceModel{
			Priority:      types.StringValue(metric.Priority),
			Metric:        types.StringValue(metric.Metric),
			Target:        types.Int64Value(int64(metric.Target)),
			BusinessHours: types.BoolValue(metric.BusinessHours),
		}
	}

	return tfMetrics
}
