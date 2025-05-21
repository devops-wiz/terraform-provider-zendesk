package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ResourceTransformWithID[zendesk.Brand] = &BrandResourceModel{}

type BrandResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	URL               types.String `tfsdk:"url"`
	Name              types.String `tfsdk:"name"`
	BrandURL          types.String `tfsdk:"brand_url"`
	HasHelpCenter     types.Bool   `tfsdk:"has_help_center"`
	HelpCenterState   types.String `tfsdk:"help_center_state"`
	Active            types.Bool   `tfsdk:"active"`
	Default           types.Bool   `tfsdk:"default"`
	IsDeleted         types.Bool   `tfsdk:"is_deleted"`
	TicketFormIDs     types.List   `tfsdk:"ticket_form_ids"`
	Subdomain         types.String `tfsdk:"subdomain"`
	HostMapping       types.String `tfsdk:"host_mapping"`
	SignatureTemplate types.String `tfsdk:"signature_template"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
}

func (b *BrandResourceModel) GetID() int64 {
	return b.ID.ValueInt64()
}

func (b *BrandResourceModel) GetApiModelFromTfModel(ctx context.Context) (newBrand zendesk.Brand, diags diag.Diagnostics) {
	ticketFormIds := make([]int64, len(b.TicketFormIDs.Elements()))

	diags = b.TicketFormIDs.ElementsAs(ctx, &ticketFormIds, true)

	if diags.HasError() {
		return newBrand, diags
	}

	newBrand = zendesk.Brand{
		Name:              b.Name.ValueString(),
		BrandURL:          b.BrandURL.ValueString(),
		HasHelpCenter:     b.HasHelpCenter.ValueBool(),
		HelpCenterState:   b.HelpCenterState.ValueString(),
		Active:            b.Active.ValueBool(),
		TicketFormIDs:     ticketFormIds,
		Subdomain:         b.Subdomain.ValueString(),
		HostMapping:       b.HostMapping.ValueString(),
		SignatureTemplate: b.SignatureTemplate.ValueString(),
	}

	return newBrand, diags

}

func (b *BrandResourceModel) GetTfModelFromApiModel(ctx context.Context, brand zendesk.Brand) (diags diag.Diagnostics) {
	tfTicketFormIds, diags := types.ListValueFrom(ctx, types.Int64Type, brand.TicketFormIDs)

	if diags.HasError() {
		return diags
	}

	*b = BrandResourceModel{
		ID:                types.Int64Value(brand.ID),
		URL:               types.StringValue(brand.URL),
		Name:              types.StringValue(brand.Name),
		BrandURL:          types.StringValue(brand.BrandURL),
		HasHelpCenter:     types.BoolValue(brand.HasHelpCenter),
		HelpCenterState:   types.StringValue(brand.HelpCenterState),
		Active:            types.BoolValue(brand.Active),
		Default:           types.BoolValue(brand.Default),
		TicketFormIDs:     tfTicketFormIds,
		Subdomain:         types.StringValue(brand.Subdomain),
		HostMapping:       types.StringValue(brand.HostMapping),
		SignatureTemplate: types.StringValue(brand.SignatureTemplate),
		CreatedAt:         types.StringValue(brand.CreatedAt.UTC().String()),
		UpdatedAt:         types.StringValue(brand.UpdatedAt.UTC().String()),
	}

	return diags
}
