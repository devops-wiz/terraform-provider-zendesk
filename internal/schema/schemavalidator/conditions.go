package schemavalidator

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Object = &ConditionsValidator{}

type ConditionsValidator struct {
	ConfigType string
}

// Description implements validator.Object.
func (c *ConditionsValidator) Description(context.Context) string {
	return fmt.Sprintf(
		"Validates acceptable values for certain action types, acceptable values: %s",
		strings.Join(zendesk.ValidConditionOperatorValues.ValidKeys(), ", "),
	)
}

// MarkdownDescription implements validator.Object.
func (c *ConditionsValidator) MarkdownDescription(context.Context) string {
	return fmt.Sprintf(
		"Validates acceptable values for certain condition types, acceptable values: %s",
		strings.Join(zendesk.ValidConditionOperatorValues.ValidKeys(), ", "),
	)
}

// ValidateObject implements validator.Object.
func (c *ConditionsValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	attributes := req.ConfigValue.Attributes()

	if attributes["field"].Type(ctx) != types.StringType ||
		attributes["value"].Type(ctx) != types.StringType ||
		attributes["operator"].Type(ctx) != types.StringType {
		resp.Diagnostics.AddAttributeError(req.Path, "type error", "field and action must be string")
		return
	}

	field := attributes["field"].(types.String)
	operator := attributes["operator"].(types.String)
	value := attributes["value"].(types.String)
	values := attributes["values"].(types.List)

	if value.IsUnknown() || values.IsUnknown() {
		return
	}

	if value.IsNull() && values.IsNull() {
		resp.Diagnostics.AddAttributeError(req.Path, "value and values is null", "either value, or values needs to be set")
	}

	if !value.IsNull() && !values.IsNull() {
		resp.Diagnostics.AddAttributeError(req.Path, "value and values are both set", "only one value or values attribute needs to be set")
	}

	if field.IsUnknown() {
		return
	}

	if field.IsNull() || field.ValueString() == "" {
		resp.Diagnostics.AddAttributeWarning(req.Path, "type error", "field must not be empty, (may be just string interpolation issue")
		return
	}

	var newField string

	if field.ValueString() == "custom_field" || field.ValueString() == "ticket_field" {
		var cid string
		customFieldId := attributes["custom_field_id"].(types.Int64)
		if customFieldId.IsUnknown() {
			cid = "1234"
		} else {
			cid = customFieldId.String()
		}
		switch field.ValueString() {
		case "custom_field":
			newField = fmt.Sprintf("custom_fields_%s", cid)
		case "ticket_field":
			newField = fmt.Sprintf("ticket_fields_%s", cid)
		}
	} else {
		newField = field.ValueString()
	}

	apiCondition := zendesk.Condition{
		Field:    newField,
		Operator: operator.ValueString(),
	}

	if !value.IsNull() && !value.IsUnknown() {
		apiCondition.Value = zendesk.ParsedValue{Data: value.ValueString()}
	}

	if !values.IsNull() && !values.IsUnknown() {
		list := make([]string, len(values.Elements()))
		resp.Diagnostics.Append(values.ElementsAs(ctx, &list, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
		apiCondition.Value = zendesk.ParsedValue{ListData: list}
	}

	if err := apiCondition.Validate(zendesk.ConditionResourceType(c.ConfigType)); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"validation error",
			fmt.Sprintf(
				"error validating %s condition: %s",
				c.ConfigType,
				err.Error(),
			),
		)
	}
}
