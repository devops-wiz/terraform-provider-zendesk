package provider

import (
	"cmp"
	"context"
	"fmt"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"slices"
)

var _ function.Function = &SortCustomFieldOptions{}

type SortCustomFieldOptions struct{}

func NewSortCustomFieldOptions() function.Function {
	return &SortCustomFieldOptions{}
}

func (s *SortCustomFieldOptions) Metadata(ctx context.Context, request function.MetadataRequest, response *function.MetadataResponse) {
	response.Name = "sort_custom_field_options"
}

func (s *SortCustomFieldOptions) Definition(ctx context.Context, request function.DefinitionRequest, response *function.DefinitionResponse) {
	response.Definition = function.Definition{
		Summary:     "Sort custom field options",
		Description: "Sort custom field options by display name",
		Parameters: []function.Parameter{
			function.ListParameter{
				Name:        "custom_field_options",
				Description: "Custom field options to sort",
				ElementType: types.ObjectType{AttrTypes: models.CustomFieldOptionResourceBase{}.AttributeTypes()},
			},
		},
		Return: function.ListReturn{ElementType: types.ObjectType{AttrTypes: models.CustomFieldOptionResourceBase{}.AttributeTypes()}},
	}
}

func (s *SortCustomFieldOptions) Run(ctx context.Context, request function.RunRequest, response *function.RunResponse) {
	var customFieldOptionsList types.List

	response.Error = function.ConcatFuncErrors(response.Error, request.Arguments.Get(ctx, &customFieldOptionsList))

	var customFieldOptions []models.CustomFieldOptionResourceBase

	diags := customFieldOptionsList.ElementsAs(ctx, &customFieldOptions, true)

	if diags.HasError() {
		response.Error = function.ConcatFuncErrors(response.Error, function.NewFuncError(fmt.Sprintf("Error converting list to custom field options: %v", diags.Errors())))
		return
	}

	slices.SortFunc(customFieldOptions, func(a, b models.CustomFieldOptionResourceBase) int {
		return cmp.Compare(a.Name.ValueString(), b.Name.ValueString())
	})

	customFieldOptionsList, diags = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CustomFieldOptionResourceBase{}.AttributeTypes()}, customFieldOptions)

	if diags.HasError() {
		response.Error = function.ConcatFuncErrors(response.Error, function.NewFuncError("Error converting custom field options to list"))
		return
	}

	response.Error = function.ConcatFuncErrors(response.Error, response.Result.Set(ctx, customFieldOptionsList))
}
