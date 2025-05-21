package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ DatasourceTransform[zendesk.SearchResults] = &SearchDatasourceModel{}

type SearchDatasourceModel struct {
	Query   types.String `tfsdk:"query"`
	Results types.Object `tfsdk:"results"`
}

// GetTfModelFromApiModel generates computed API response from search api
func (s *SearchDatasourceModel) GetTfModelFromApiModel(ctx context.Context, results zendesk.SearchResults) (diags diag.Diagnostics) {
	diags = make(diag.Diagnostics, 0)

	var resultsList types.List

	var resultsObject types.Object

	if len(results.List()) == 0 {
		diags.AddError("No search results", "There are no search results from this query")
		return diags
	} else {
		switch results.List()[0].(type) {
		case zendesk.User:
			tfUsers := make([]SearchResultsUserDatasourceModel, len(results.List()))
			for i, result := range results.List() {
				user := result.(zendesk.User)
				tfUsers[i] = SearchResultsUserDatasourceModel{
					ID:             types.Int64Value(user.ID),
					Email:          types.StringValue(user.Email),
					Name:           types.StringValue(user.Name),
					OrganizationID: types.Int64Value(user.OrganizationID),
					ExternalID:     types.StringValue(user.ExternalID),
				}
			}

			resultsList, diags = types.ListValueFrom(
				ctx,
				types.ObjectType{AttrTypes: SearchResultsUserDatasourceModel{}.AttributeTypes()},
				tfUsers,
			)

			if diags.HasError() {
				return diags
			}

			resultsDatasourceModel := SearchResultsDatasourceModel{
				Users:         resultsList,
				Organizations: types.ListNull(types.ObjectType{AttrTypes: SearchResultsOrganizationDatasourceModel{}.AttributeTypes()}),
			}

			resultsObject, diags = types.ObjectValueFrom(ctx, resultsDatasourceModel.AttributeTypes(), resultsDatasourceModel)

			if diags.HasError() {
				return diags
			}

			s.Results = resultsObject

		case zendesk.Organization:
			tfOrgs := make([]SearchResultsOrganizationDatasourceModel, len(results.List()))
			for i, result := range results.List() {
				org := result.(zendesk.Organization)
				tfOrgs[i] = SearchResultsOrganizationDatasourceModel{
					ID:         types.Int64Value(org.ID),
					Name:       types.StringValue(org.Name),
					ExternalID: types.StringValue(org.ExternalID),
				}
			}
			resultsList, diags = types.ListValueFrom(
				ctx,
				types.ObjectType{AttrTypes: SearchResultsOrganizationDatasourceModel{}.AttributeTypes()},
				tfOrgs,
			)
			if diags.HasError() {
				return diags
			}
			resultsDatasourceModel := SearchResultsDatasourceModel{
				Organizations: resultsList,
				Users:         types.ListNull(types.ObjectType{AttrTypes: SearchResultsUserDatasourceModel{}.AttributeTypes()}),
			}

			resultsObject, diags = types.ObjectValueFrom(ctx, resultsDatasourceModel.AttributeTypes(), resultsDatasourceModel)

			if diags.HasError() {
				return diags
			}

			s.Results = resultsObject
		default:
			diags.AddError("Unsupported result type", "Currently, only User and Organization results are supported.")
			return diags
		}
	}

	return diags
}

func (s *SearchDatasourceModel) GetApiQueryOptionsFromTf() (options zendesk.SearchOptions) {
	query := s.Query.ValueString()

	options = zendesk.SearchOptions{
		PageOptions: zendesk.PageOptions{
			PerPage: 1000,
			Page:    1,
		},
		Query:     query,
		SortBy:    "created_at",
		SortOrder: "asc",
	}

	return options
}

type SearchResultsDatasourceModel struct {
	Users         types.List `tfsdk:"users"`
	Organizations types.List `tfsdk:"organizations"`
}

func (s *SearchResultsDatasourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"users":         types.ListType{ElemType: types.ObjectType{AttrTypes: SearchResultsUserDatasourceModel{}.AttributeTypes()}},
		"organizations": types.ListType{ElemType: types.ObjectType{AttrTypes: SearchResultsOrganizationDatasourceModel{}.AttributeTypes()}},
	}
}

type SearchResultsUserDatasourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	Email          types.String `tfsdk:"email"`
	Name           types.String `tfsdk:"name"`
	OrganizationID types.Int64  `tfsdk:"organization_id"`
	ExternalID     types.String `tfsdk:"external_id"`
}

func (m SearchResultsUserDatasourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":              types.Int64Type,
		"email":           types.StringType,
		"name":            types.StringType,
		"organization_id": types.Int64Type,
		"external_id":     types.StringType,
	}
}

type SearchResultsOrganizationDatasourceModel struct {
	ID         types.Int64  `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	ExternalID types.String `tfsdk:"external_id"`
}

func (m SearchResultsOrganizationDatasourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":          types.Int64Type,
		"name":        types.StringType,
		"external_id": types.StringType,
	}
}
