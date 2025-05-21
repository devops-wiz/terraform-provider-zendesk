package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func getApiOutputFromTf(ctx context.Context, tfOutput types.Object) (zendesk.ViewOutput, diag.Diagnostics) {
	var output ViewOutputResourceModel
	diags := tfOutput.As(ctx, &output, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: false, UnhandledUnknownAsEmpty: false})

	newColumns := make([]string, len(output.Columns.Elements()))

	for index, column := range output.Columns.Elements() {
		newColumns[index] = column.(types.String).ValueString()
	}

	newOutput := zendesk.ViewOutput{
		Columns:    newColumns,
		GroupBy:    output.GroupBy.ValueString(),
		GroupOrder: output.GroupOrder.ValueString(),
		SortBy:     output.SortBy.ValueString(),
		SortOrder:  output.SortOrder.ValueString(),
	}

	return newOutput, diags
}

func getTfOutputFromApi(ctx context.Context, viewExecution zendesk.ViewExecution, tfOutput *types.Object) diag.Diagnostics {
	outputColumns := make([]attr.Value, len(viewExecution.Columns))

	for i, column := range viewExecution.Columns {
		outputColumns[i] = types.StringValue(column.ID)
	}

	output := ViewOutputResourceModel{
		Columns:    types.ListValueMust(types.StringType, outputColumns),
		GroupBy:    types.StringValue(viewExecution.GroupBy),
		GroupOrder: types.StringValue(viewExecution.GroupOrder),
		SortBy:     types.StringValue(viewExecution.SortBy),
		SortOrder:  types.StringValue(viewExecution.SortOrder),
	}

	var diags diag.Diagnostics

	*tfOutput, diags = types.ObjectValueFrom(ctx, output.AttributeTypes(), output)

	return diags
}
