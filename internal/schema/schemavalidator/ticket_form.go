package schemavalidator

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ConditionType string

const (
	AgentType   ConditionType = "agent"
	EndUserType ConditionType = "end_user"
)

type ConditionRequirementValidator struct {
	ConditionType ConditionType
}

var _ validator.Object = &ConditionRequirementValidator{}

func (c ConditionRequirementValidator) Description(_ context.Context) string {
	return "When required_on_statuses.type is set to ALL_STATUSES, is_required needs to be set to true"
}

func (c ConditionRequirementValidator) MarkdownDescription(_ context.Context) string {
	return "When `required_on_statuses.type` is set to ALL_STATUSES, `is_required` needs to be set to true"
}

func (c ConditionRequirementValidator) ValidateObject(ctx context.Context, request validator.ObjectRequest, response *validator.ObjectResponse) {
	var childFieldModel models.FormChildFieldConditions

	if c.ConditionType == EndUserType {
		return
	}

	response.Diagnostics.Append(request.ConfigValue.As(ctx, &childFieldModel, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if response.Diagnostics.HasError() {
		return
	}

	var requiredOnStatusesModel models.RequiredOnStatusesResourceModel

	response.Diagnostics.Append(childFieldModel.RequiredOnStatuses.As(ctx, &requiredOnStatusesModel, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	requirementTypeModel := requiredOnStatusesModel.Type

	requirementType := zendesk.RequirementType(requirementTypeModel.ValueString())

	err := requirementType.Validate()

	if err != nil {
		response.Diagnostics.AddAttributeError(request.Path, "Error validating status requirement", fmt.Sprintf("Error from lib: %s", err.Error()))
		return
	}

	required := childFieldModel.IsRequired.ValueBool()

	switch {
	case !required && requirementType == zendesk.SomeStatuses:
		response.Diagnostics.AddAttributeWarning(request.Path, "Potential Invalid combination of attributes", "When required_on_statuses.type "+
			"is set to 'SOME_STATUSES', is_required needs to be set to true, unless the referenced child ticket field id is already required. Please verify this.")
		return
	case !required && requirementType == zendesk.AllStatuses:
		response.Diagnostics.AddAttributeError(request.Path, "Invalid combination of attributes", "When required_on_statuses.type "+
			"is set to 'ALL_STATUSES', is_required needs to be set to true")
		return
	case required && requirementType == zendesk.NoStatuses:
		response.Diagnostics.AddAttributeError(request.Path, "Invalid combination of attributes", "When required_on_statuses.type "+
			"is set to 'NO_STATUSES', is_required needs to be set to false")
	default:
		return
	}

}
