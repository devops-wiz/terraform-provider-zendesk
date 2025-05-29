package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ validator.Object = &ActionsValidator{}

type ActionsValidator struct {
	ConfigType string
}

// ValidateObject implements validator.Object.
func (v *ActionsValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	attributes := req.ConfigValue.Attributes()

	if attributes["field"].Type(ctx) != types.StringType {
		resp.Diagnostics.AddAttributeError(req.Path, "type error", "field must be string")
		return
	}

	if attributes["value"].Type(ctx) != types.StringType {
		resp.Diagnostics.AddAttributeError(req.Path, "type error", "action must be string")
		return
	}

	if attributes["value"].IsUnknown() {
		return
	}

	field := attributes["field"].(types.String).ValueString()
	value := attributes["value"].(types.String).ValueString()

	if field == "custom_field" {
		var cid string
		customFieldId := attributes["custom_field_id"].(types.Int64)
		if customFieldId.IsUnknown() {
			cid = "1234"
		} else {
			cid = customFieldId.String()
		}
		field = fmt.Sprintf("custom_fields_%s", cid)
	}

	action := zendesk.Action{
		Field: field,
	}

	if attributes["target"].IsUnknown() {
		return
	}

	if !attributes["target"].IsNull() {
		switch target := attributes["target"].(type) {
		case types.String:
			if !target.IsUnknown() && !target.IsNull() {
				action.Value = zendesk.ParsedValue{ListData: []string{target.ValueString(), value}}
			}
		case types.Int64:
			if !target.IsUnknown() && !target.IsNull() {
				action.Value = zendesk.ParsedValue{ListData: []string{strconv.FormatInt(target.ValueInt64(), 10), value}}
			}
		}
	} else {
		action.Value = zendesk.ParsedValue{Data: value}
	}

	ctx = tflog.SetField(ctx, "fieldName", field)
	tflog.Info(ctx, "Getting attribute 'field'")

	if err := action.Validate(zendesk.ActionResourceType(v.ConfigType)); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"validation error",
			fmt.Sprintf(
				"error with validation: %s",
				err.Error(),
			),
		)
		return
	}

}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v *ActionsValidator) Description(context.Context) string {
	return fmt.Sprintf("Validates acceptable values for certain action types, acceptable values: %s", strings.Join(zendesk.ValidActionValuesMap.ValidKeys(), ", "))
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (v *ActionsValidator) MarkdownDescription(context.Context) string {
	return fmt.Sprintf("Validates acceptable values for certain action types, acceptable values: %s", strings.Join(zendesk.ValidActionValuesMap.ValidKeys(), ", "))
}
