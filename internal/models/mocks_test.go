package models

import (
	"context"
	"time"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	testId          int64 = 123
	testTitle             = "test title"
	testDescription       = "test desc"
	// testCatId             = "123"
	testUrl             = "https://example.org"
	testPosition  int64 = 2
	testField           = "status"
	testOperator        = "is"
	testValue           = "open"
	testCreatedAt       = time.Now()
	testUpdatedAt       = time.Now()
	// testCatIdInt, _       = strconv.ParseInt(testCatId, 10, 64)
	testActive = true

	// Ticket field mocks

	testTicketFieldId       int64 = 123
	testTicketFieldTitle          = "test ticket field title"
	testTicketFieldType           = "test ticket field type"
	testTicketFieldRequired       = false
	testTicketFieldDesc           = "test ticket field desc"
	testTicketFieldTag            = "test_tag"
	testTicketFieldRegex          = ".*"
	testTicketFieldActive         = true
	testTicketFieldVisible        = true
	testTicketFieldEditable       = true
	testTicketFieldPosition int64 = 2
	testCfoTitle                  = "Test custom option"
	testCfoTag                    = "test_tag"
	testTicketFieldTime           = time.Now().UTC()
	testTicketFieldUrl            = "http://example.com"
	testCfosGetApi                = []zendesk.CustomFieldOption{
		{
			ID:       &testTicketFieldId,
			Name:     testCfoTitle,
			Position: testTicketFieldPosition,
			RawName:  testCfoTitle,
			Value:    testCfoTag,
			URL:      testTicketFieldUrl,
		},
	}
	testCfosPostApi = []zendesk.CustomFieldOption{
		{
			ID:    &testId,
			Name:  testCfoTitle,
			Value: testCfoTag,
		},
	}
	testTicketFieldTf = TicketFieldResourceModel{
		ID:                  types.Int64Value(testTicketFieldId),
		Title:               types.StringValue(testTicketFieldTitle),
		TitleInPortal:       types.StringValue(testTicketFieldTitle),
		Type:                types.StringValue(testTicketFieldType),
		Required:            types.BoolValue(testTicketFieldRequired),
		RequiredInPortal:    types.BoolValue(testTicketFieldRequired),
		Description:         types.StringValue(testTicketFieldDesc),
		AgentDescription:    types.StringValue(testTicketFieldDesc),
		Tag:                 types.StringValue(testTicketFieldTag),
		RegexpForValidation: types.StringValue(testTicketFieldRegex),
		Active:              types.BoolValue(testTicketFieldActive),
		VisibleInPortal:     types.BoolValue(testTicketFieldVisible),
		EditableInPortal:    types.BoolValue(testTicketFieldEditable),
		Position:            types.Int64Value(testTicketFieldPosition),
		CustomFieldOptions: types.ListValueMust(types.ObjectType{AttrTypes: CustomFieldOptionResourceModel{}.AttributeTypes()}, []attr.Value{
			types.ObjectValueMust(CustomFieldOptionResourceModel{}.AttributeTypes(), map[string]attr.Value{
				"id":    types.Int64Value(testId),
				"name":  types.StringValue(testCfoTitle),
				"value": types.StringValue(testCfoTag),
			}),
		}),
		SystemFieldOptions: types.ListNull(types.ObjectType{AttrTypes: SystemFieldOption{}.AttributeTypes()}),
		CreatedAt:          types.StringValue(testTicketFieldTime.UTC().String()),
		UpdatedAt:          types.StringValue(testTicketFieldTime.UTC().String()),
		URL:                types.StringValue(testTicketFieldUrl),
	}

	testTicketFieldApiExpected = zendesk.TicketField{
		Title:               testTicketFieldTitle,
		TitleInPortal:       testTicketFieldTitle,
		Type:                testTicketFieldType,
		Required:            testTicketFieldRequired,
		RequiredInPortal:    testTicketFieldRequired,
		Description:         testTicketFieldDesc,
		AgentDescription:    testTicketFieldDesc,
		Tag:                 testTicketFieldTag,
		RegexpForValidation: testTicketFieldRegex,
		Active:              testTicketFieldActive,
		VisibleInPortal:     testTicketFieldVisible,
		EditableInPortal:    testTicketFieldEditable,
		Position:            testTicketFieldPosition,
		CustomFieldOptions:  testCfosPostApi,
	}

	testTicketFieldApiInput = zendesk.TicketField{
		ID:                  testTicketFieldId,
		Title:               testTicketFieldTitle,
		TitleInPortal:       testTicketFieldTitle,
		Type:                testTicketFieldType,
		Required:            testTicketFieldRequired,
		RequiredInPortal:    testTicketFieldRequired,
		Description:         testTicketFieldDesc,
		AgentDescription:    testTicketFieldDesc,
		Tag:                 testTicketFieldTag,
		RegexpForValidation: testTicketFieldRegex,
		Active:              testTicketFieldActive,
		VisibleInPortal:     testTicketFieldVisible,
		EditableInPortal:    testTicketFieldEditable,
		Position:            testTicketFieldPosition,
		CustomFieldOptions:  testCfosGetApi,
		CreatedAt:           &testTicketFieldTime,
		UpdatedAt:           &testTicketFieldTime,
		URL:                 testTicketFieldUrl,
	}

	testGroupColumn = "status"
	testSortColumn  = "assignee_id"
	testDirection   = "asc"
	testOutputRaw   = ViewOutputResourceModel{
		Columns: types.ListValueMust(types.StringType, []attr.Value{
			types.StringValue(testGroupColumn),
			types.StringValue(testSortColumn),
		}),
		GroupBy:    types.StringValue(testGroupColumn),
		GroupOrder: types.StringValue(testDirection),
		SortBy:     types.StringValue(testSortColumn),
		SortOrder:  types.StringValue(testDirection),
	}

	testKey = "test_key"
)

var testOutput, _ = types.ObjectValueFrom(context.Background(), testOutputRaw.AttributeTypes(), testOutputRaw)

var automationModelInput = AutomationResourceModel{
	Actions:     testActionModels,
	Active:      types.BoolValue(testActive),
	Description: types.StringValue(testDescription),
	Conditions:  testConditionsModel,
	Position:    types.Int64Value(testPosition),
	Title:       types.StringValue(testTitle),
}

var automationModelNoPositionInput = AutomationResourceModel{
	Actions:     testActionModels,
	Active:      types.BoolValue(testActive),
	Description: types.StringValue(testDescription),
	Conditions:  testConditionsModel,
	Title:       types.StringValue(testTitle),
}

var apiAutomationModelExpected = zendesk.Automation{
	Title:       testTitle,
	Description: testDescription,
	Active:      testActive,
	Position:    testPosition,
	Conditions:  testApiConditionsModel,
	Actions:     testApiActionModels,
}

var apiAutomationModelNoPositionExpected = zendesk.Automation{
	Title:       testTitle,
	Description: testDescription,
	Active:      testActive,
	Conditions:  testApiConditionsModel,
	Actions:     testApiActionModels,
}

var apiAutomationModelInput = zendesk.Automation{
	ID:          testId,
	Title:       testTitle,
	Description: testDescription,
	Active:      testActive,
	Position:    testPosition,
	Conditions:  testApiConditionsModel,
	Actions:     testApiActionModels,
	CreatedAt:   &testCreatedAt,
	UpdatedAt:   &testUpdatedAt,
	URL:         testUrl,
}

var automationModelExpected = AutomationResourceModel{
	ID:          types.Int64Value(testId),
	Actions:     testActionModels,
	Active:      types.BoolValue(testActive),
	Description: types.StringValue(testDescription),
	Conditions:  testConditionsModel,
	CreatedAt:   types.StringValue(testCreatedAt.UTC().String()),
	Position:    types.Int64Value(testPosition),
	Title:       types.StringValue(testTitle),
	UpdatedAt:   types.StringValue(testUpdatedAt.UTC().String()),
	URL:         types.StringValue(testUrl),
}

// VIEWS

var apiViewInput = zendesk.View{
	URL:         testUrl,
	ID:          testId,
	Title:       testTitle,
	Active:      testActive,
	UpdatedAt:   testUpdatedAt.UTC().String(),
	CreatedAt:   testCreatedAt.UTC().String(),
	Default:     false,
	Position:    testPosition,
	Description: testDescription,
	Execution: map[string]interface{}{
		"columns": []interface{}{
			map[string]interface{}{
				"id": testGroupColumn,
			},
			map[string]interface{}{
				"id": testSortColumn,
			},
		},
		"group_by":    testGroupColumn,
		"group_order": testDirection,
		"sort_by":     testSortColumn,
		"sort_order":  testDirection,
	},
	Conditions:  testApiConditionsModel,
	Restriction: nil,
}

var testViewModelExpected = ViewResourceModel{
	URL:         types.StringValue(testUrl),
	ID:          types.Int64Value(testId),
	Title:       types.StringValue(testTitle),
	Description: types.StringValue(testDescription),
	UpdatedAt:   types.StringValue(testUpdatedAt.UTC().String()),
	CreatedAt:   types.StringValue(testCreatedAt.UTC().String()),
	Active:      types.BoolValue(testActive),
	Position:    types.Int64Value(testPosition),
	Conditions: &ConditionsResourceModel{
		All: testConditionsModel.All,
		Any: []ConditionResourceModel(nil),
	},
	Output:      testOutput,
	Restriction: types.ObjectNull(RestrictionResourceModel{}.AttributeTypes()),
}

var testActionModels = []ActionResourceModel{
	{
		Field:         types.StringValue(testField),
		Value:         types.StringValue(testValue),
		Target:        types.StringNull(),
		CustomFieldID: types.Int64Null(),
	},
}

var testApiActionModels = []zendesk.Action{
	{
		Field: testField,
		Value: zendesk.ParsedValue{Data: testValue},
	},
}

var testConditionsModel = ConditionsResourceModel{
	All: []ConditionResourceModel{
		{
			Field:         types.StringValue(testField),
			Operator:      types.StringValue(testOperator),
			Value:         types.StringValue(testValue),
			Values:        types.ListNull(types.StringType),
			CustomFieldID: types.Int64Null(),
		},
	},
}

var testApiConditionsModel = zendesk.Conditions{
	All: []zendesk.Condition{
		{
			Field:    testField,
			Operator: testOperator,
			Value:    zendesk.ParsedValue{Data: testValue},
		},
	},
	Any: []zendesk.Condition{},
}

var testMacroWithoutPositionRestriction = zendesk.Macro{
	Actions:     testApiActionModels,
	Active:      true,
	Description: testDescription,
	Title:       testTitle,
}

var testMacroWithPositionWithoutRestriction = zendesk.Macro{
	Actions:     testApiActionModels,
	Active:      true,
	Position:    int(testPosition),
	Description: testDescription,
	Title:       testTitle,
}

var testRestriction = zendesk.Restriction{
	Type: "Group",
	ID:   0,
	IDS:  []int64{1234},
}

var testMacroWithPositionWithRestriction = zendesk.Macro{
	Actions:     testApiActionModels,
	Active:      true,
	Position:    int(testPosition),
	Description: testDescription,
	Title:       testTitle,
	Restriction: testRestriction,
}

var testMacroWithoutPositionWithRestriction = zendesk.Macro{
	Actions:     testApiActionModels,
	Active:      true,
	Description: testDescription,
	Title:       testTitle,
	Restriction: testRestriction,
}

var testMacroResourceModelWithPositionWithoutRestriction = MacroResourceModel{
	Actions:     testActionModels,
	Title:       types.StringValue(testTitle),
	Active:      types.BoolValue(true),
	Description: types.StringValue(testDescription),
	Position:    types.Int64Value(testPosition),
}

var testMacroResourceModelWithoutPositionRestriction = MacroResourceModel{
	Actions:     testActionModels,
	Title:       types.StringValue(testTitle),
	Active:      types.BoolValue(true),
	Description: types.StringValue(testDescription),
	Position:    types.Int64Null(),
}

var testRestrictionModelRaw = RestrictionResourceModel{
	Type: types.StringValue("Group"),
	IDS:  types.SetValueMust(types.Int64Type, []attr.Value{types.Int64Value(1234)}),
}

var testRestrictionModel, _ = types.ObjectValueFrom(
	context.Background(),
	testRestrictionModelRaw.AttributeTypes(),
	testRestrictionModelRaw,
)

var testMacroResourceModelWithRestrictionWithoutPosition = MacroResourceModel{
	Actions:     testActionModels,
	Title:       types.StringValue(testTitle),
	Active:      types.BoolValue(true),
	Description: types.StringValue(testDescription),
	Restriction: testRestrictionModel,
}
var testMacroResourceModelWithRestrictionWithPosition = MacroResourceModel{
	Actions:     testActionModels,
	Title:       types.StringValue(testTitle),
	Active:      types.BoolValue(true),
	Description: types.StringValue(testDescription),
	Restriction: testRestrictionModel,
	Position:    types.Int64Value(testPosition),
}

var (
	testWebhookId               = "123"
	testWebhookName             = "test name"
	testWebhookDesc             = "test desc"
	testWebhookEndpoint         = "https://example.org"
	testWebhookHttpMethod       = "GET"
	testWebhookRequestFormat    = "json"
	testWebhookStatus           = "active"
	testWebhookBasicAuth        = "basic_auth"
	testWebhookApiKey           = "api_key"
	testWebhookBearerToken      = "bearer_token"
	testWebhookAddPosition      = "header"
	testWebhookUsername         = "username"
	testWebhookPassword         = "password"
	testWebhookApiHeaderKey     = "api-header-key"
	testWebhookApiHeaderValue   = "apiHeaderValue"
	testWebhookBearerTokenValue = "abc.12345"
	testWebhookSigningSecret    = "adfkljsldkfJlkdsdklDK="
	testWebhookSigningAlgorithm = "SHA256"
	testWebhookUser             = "5859"
	testWebhookTime             = time.Now()
)

var testWebhookCustomHeadersTf = map[string]attr.Value{
	"header-key": types.StringValue("headerValue"),
}
var testWebhookCustomHeadersMap = types.MapValueMust(types.StringType, testWebhookCustomHeadersTf)
var testWebhookCustomHeadersApi = map[string]string{
	"header-key": "headerValue",
}

var testMetricSettingsFirstReplyRaw = FirstReplyTimeResourceModel{
	ActivateOnTicketCreatedForEndUser:                      types.BoolValue(true),
	ActivateOnAgentTicketCreatedForEndUserWithInternalNote: types.BoolValue(false),
	ActivateOnLightAgentOnEmailForwardTicketFromEndUser:    types.BoolNull(),
	ActivateOnAgentCreatedTicketForSelf:                    types.BoolNull(),
	FulfillOnAgentInternalNote:                             types.BoolNull(),
}

var testMetricsSettingsFirstReplyObj, _ = types.ObjectValueFrom(
	context.Background(),
	testMetricSettingsFirstReplyRaw.AttributeTypes(),
	testMetricSettingsFirstReplyRaw,
)

var testMetricSettingsRaw = MetricSettingsResourceModel{
	FirstReplyTime:     testMetricsSettingsFirstReplyObj,
	NextReplyTime:      types.ObjectNull(NextReplyTimeResourceModel{}.AttributeTypes()),
	PeriodicUpdateTime: types.ObjectNull(PeriodicUpdateTimeResourceModel{}.AttributeTypes()),
}

var testMetricSettingsObj, _ = types.ObjectValueFrom(context.Background(), testMetricSettingsRaw.AttributeTypes(), testMetricSettingsRaw)

var testSlaPolicyModelInput = SLAPolicyResourceModel{
	Title:           types.StringValue(testTitle),
	Description:     types.StringValue(testDescription),
	Position:        types.Int64Value(testPosition),
	Filter:          testConditionsModel,
	MetricsSettings: testMetricSettingsObj,
	PolicyMetrics: []SLAPolicyMetricResourceModel{
		{
			Priority:      types.StringValue("low"),
			Metric:        types.StringValue(zendesk.AgentWorkTimeMetric),
			Target:        types.Int64Value(60),
			BusinessHours: types.BoolValue(false),
		},
		{
			Priority:      types.StringValue("high"),
			Metric:        types.StringValue(zendesk.AgentWorkTimeMetric),
			Target:        types.Int64Value(10),
			BusinessHours: types.BoolValue(false),
		},
	},
}

var testSlaPolicyExpected = zendesk.SLAPolicy{
	Title:       testTitle,
	Description: testDescription,
	Position:    testPosition,
	Filter:      testApiConditionsModel,
	MetricSettings: zendesk.MetricSettings{
		FirstReplyTime: zendesk.FirstReplyTime{
			ActivateOnTicketCreatedForEndUser:                      true,
			ActivateOnAgentTicketCreatedForEndUserWithInternalNote: false,
		},
	},
	PolicyMetrics: []zendesk.SLAPolicyMetric{
		{
			Priority:      "low",
			Metric:        zendesk.AgentWorkTimeMetric,
			Target:        60,
			BusinessHours: false,
		},
		{
			Priority:      "high",
			Metric:        zendesk.AgentWorkTimeMetric,
			Target:        10,
			BusinessHours: false,
		},
	},
}
var testSlaPolicyInput = zendesk.SLAPolicy{
	ID:          testId,
	Title:       testTitle,
	Description: testDescription,
	Position:    testPosition,
	Filter:      testApiConditionsModel,
	MetricSettings: zendesk.MetricSettings{
		FirstReplyTime: zendesk.FirstReplyTime{
			ActivateOnTicketCreatedForEndUser:                      true,
			ActivateOnAgentTicketCreatedForEndUserWithInternalNote: false,
		},
	},
	PolicyMetrics: []zendesk.SLAPolicyMetric{
		{
			Priority:      "low",
			Metric:        zendesk.AgentWorkTimeMetric,
			Target:        60,
			BusinessHours: false,
		},
		{
			Priority:      "high",
			Metric:        zendesk.AgentWorkTimeMetric,
			Target:        10,
			BusinessHours: false,
		},
	},
	CreatedAt: &testCreatedAt,
	UpdatedAt: &testUpdatedAt,
}

var testMetricSettingsFirstReplyRawExpected = FirstReplyTimeResourceModel{
	ActivateOnTicketCreatedForEndUser:                      types.BoolValue(true),
	ActivateOnAgentTicketCreatedForEndUserWithInternalNote: types.BoolValue(false),
	ActivateOnLightAgentOnEmailForwardTicketFromEndUser:    types.BoolValue(false),
	ActivateOnAgentCreatedTicketForSelf:                    types.BoolValue(false),
	FulfillOnAgentInternalNote:                             types.BoolValue(false),
}

var testMetricsSettingsFirstReplyObjExpected, _ = types.ObjectValueFrom(
	context.Background(),
	testMetricSettingsFirstReplyRawExpected.AttributeTypes(),
	testMetricSettingsFirstReplyRawExpected,
)

var testMetricSettingsRawExpected = MetricSettingsResourceModel{
	FirstReplyTime:     testMetricsSettingsFirstReplyObjExpected,
	NextReplyTime:      types.ObjectNull(NextReplyTimeResourceModel{}.AttributeTypes()),
	PeriodicUpdateTime: types.ObjectNull(PeriodicUpdateTimeResourceModel{}.AttributeTypes()),
}

var testMetricSettingsObjExpected, _ = types.ObjectValueFrom(context.Background(), testMetricSettingsRawExpected.AttributeTypes(), testMetricSettingsRawExpected)

var testSlaPolicyModelExpected = SLAPolicyResourceModel{
	ID:              types.Int64Value(testId),
	Title:           types.StringValue(testTitle),
	Description:     types.StringValue(testDescription),
	Position:        types.Int64Value(testPosition),
	Filter:          testConditionsModel,
	MetricsSettings: testMetricSettingsObjExpected,
	PolicyMetrics: []SLAPolicyMetricResourceModel{
		{
			Priority:      types.StringValue("low"),
			Metric:        types.StringValue(zendesk.AgentWorkTimeMetric),
			Target:        types.Int64Value(60),
			BusinessHours: types.BoolValue(false),
		},
		{
			Priority:      types.StringValue("high"),
			Metric:        types.StringValue(zendesk.AgentWorkTimeMetric),
			Target:        types.Int64Value(10),
			BusinessHours: types.BoolValue(false),
		},
	},
	CreatedAt: types.StringValue(testCreatedAt.UTC().String()),
	UpdatedAt: types.StringValue(testUpdatedAt.UTC().String()),
}

var testCredentialsModelPasswordAuth = CredentialsResourceModel{
	HeaderName:  types.StringNull(),
	HeaderValue: types.StringNull(),
	Username:    types.StringValue(testWebhookUsername),
	Password:    types.StringValue(testWebhookPassword),
	Token:       types.StringNull(),
}

var testCredsObjPassAuth, _ = types.ObjectValueFrom(context.Background(), testCredentialsModelPasswordAuth.AttributeTypes(), testCredentialsModelPasswordAuth)

var testAuthModel = AuthenticationResourceModel{
	Type:        types.StringValue(testWebhookBasicAuth),
	AddPosition: types.StringValue(testWebhookAddPosition),
	Credentials: testCredsObjPassAuth,
}
var testAuthObj, _ = types.ObjectValueFrom(context.Background(), testAuthModel.AttributeTypes(), testAuthModel)

var testCredentialsHeaderAuth = CredentialsResourceModel{
	HeaderName:  types.StringValue(testWebhookApiHeaderKey),
	HeaderValue: types.StringValue(testWebhookApiHeaderValue),
	Username:    types.StringNull(),
	Password:    types.StringNull(),
	Token:       types.StringNull(),
}

var testCredsObjHeaderAuth, _ = types.ObjectValueFrom(context.Background(), testCredentialsHeaderAuth.AttributeTypes(), testCredentialsHeaderAuth)

var testCredentialsModelTokenAuth = CredentialsResourceModel{
	HeaderName:  types.StringNull(),
	HeaderValue: types.StringNull(),
	Username:    types.StringNull(),
	Password:    types.StringNull(),
	Token:       types.StringValue(testWebhookBearerTokenValue),
}

var testCredsObjTokenAuth, _ = types.ObjectValueFrom(context.Background(), testCredentialsModelTokenAuth.AttributeTypes(), testCredentialsModelTokenAuth)

var testAuthModelInputPass = AuthenticationResourceModel{
	Type:        types.StringValue(testWebhookBasicAuth),
	AddPosition: types.StringValue(testWebhookAddPosition),
	Credentials: testCredsObjPassAuth,
}

var testAuthObjPass, _ = types.ObjectValueFrom(context.Background(), testAuthModelInputPass.AttributeTypes(), testAuthModelInputPass)

var testAuthModelHeaderAuth = AuthenticationResourceModel{
	Type:        types.StringValue(testWebhookApiKey),
	AddPosition: types.StringValue(testWebhookAddPosition),
	Credentials: testCredsObjHeaderAuth,
}

var testAuthObjHeader, _ = types.ObjectValueFrom(context.Background(), testAuthModelHeaderAuth.AttributeTypes(), testAuthModelHeaderAuth)

var testAuthModelTokenAuth = AuthenticationResourceModel{
	Type:        types.StringValue(testWebhookBearerToken),
	AddPosition: types.StringValue(testWebhookAddPosition),
	Credentials: testCredsObjTokenAuth,
}

var testAuthObjToken, _ = types.ObjectValueFrom(context.Background(), testAuthModelTokenAuth.AttributeTypes(), testAuthModelTokenAuth)
