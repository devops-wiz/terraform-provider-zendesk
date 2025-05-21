package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ResourceTransformWithID[zendesk.Group] = &GroupResourceModel{}

// GroupResourceModel is struct for group payload
// https://developer.zendesk.com/rest_api/docs/support/groups
type GroupResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	URL         types.String `tfsdk:"url"`
	Name        types.String `tfsdk:"name"`
	Default     types.Bool   `tfsdk:"default"`
	Deleted     types.Bool   `tfsdk:"deleted"`
	IsPublic    types.Bool   `tfsdk:"is_public"`
	Description types.String `tfsdk:"description"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

func (g *GroupResourceModel) GetID() int64 {
	return g.ID.ValueInt64()
}

func (g *GroupResourceModel) GetApiModelFromTfModel(_ context.Context) (group zendesk.Group, diags diag.Diagnostics) {
	group = zendesk.Group{
		Description: g.Description.ValueString(),
		Name:        g.Name.ValueString(),
		IsPublic:    g.IsPublic.ValueBool(),
	}

	return group, diags
}

func (g *GroupResourceModel) GetTfModelFromApiModel(_ context.Context, group zendesk.Group) (diags diag.Diagnostics) {
	*g = GroupResourceModel{
		ID:          types.Int64Value(group.ID),
		URL:         types.StringValue(group.URL),
		Name:        types.StringValue(group.Name),
		Default:     types.BoolValue(group.Default),
		Deleted:     types.BoolValue(group.Deleted),
		IsPublic:    types.BoolValue(group.IsPublic),
		Description: types.StringValue(group.Description),
		CreatedAt:   types.StringValue(group.CreatedAt.UTC().String()),
		UpdatedAt:   types.StringValue(group.UpdatedAt.UTC().String()),
	}

	return diags
}
