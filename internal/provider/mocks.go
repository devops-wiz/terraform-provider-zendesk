package provider

import (
	"context"
	"fmt"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func getTestCfoList(list []models.CustomFieldOptionResourceBase) (types.List, error) {
	testCfoListInput, diags := types.ListValueFrom(context.Background(), types.ObjectType{AttrTypes: models.CustomFieldOptionResourceBase{}.AttributeTypes()}, list)

	if diags.HasError() {
		return types.ListNull(types.ObjectType{AttrTypes: models.CustomFieldOptionResourceBase{}.AttributeTypes()}), fmt.Errorf("error converting model slice to list: %v", diags.Errors())
	}
	return testCfoListInput, nil
}

var testCfosModelInput = []models.CustomFieldOptionResourceBase{
	{
		Name:  types.StringValue("Something O"),
		Value: types.StringValue("something_O"),
	},
	{
		Name:  types.StringValue("Something B"),
		Value: types.StringValue("something_B"),
	},
	{
		Name:  types.StringValue("Something T"),
		Value: types.StringValue("something_T"),
	},
	{
		Name:  types.StringValue("Something K"),
		Value: types.StringValue("something_K"),
	},
}

var testCfosModelExpected = []models.CustomFieldOptionResourceBase{
	{
		Name:  types.StringValue("Something B"),
		Value: types.StringValue("something_B"),
	},
	{
		Name:  types.StringValue("Something K"),
		Value: types.StringValue("something_K"),
	},
	{
		Name:  types.StringValue("Something O"),
		Value: types.StringValue("something_O"),
	},
	{
		Name:  types.StringValue("Something T"),
		Value: types.StringValue("something_T"),
	},
}
