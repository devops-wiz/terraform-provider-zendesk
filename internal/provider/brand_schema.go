package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var BrandSchema = schema.Schema{
	Version: 0,
	MarkdownDescription: "Brands are your customer-facing identities. " +
		"They might represent multiple products or services, or they might literally be multiple brands owned and " +
		"represented by your company.\n\nThe default brand is the one that tickets get assigned to if the ticket is " +
		"generated from a non-branded channel. You can update the default brand using the " +
		"[Update Account Settings](https://developer.zendesk.com/api-reference/ticketing/account-configuration/account_settings/#update-account-settings) endpoint.",
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Description: "The ID automatically assigned when the brand is created",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:    true,
			Description: "The name of the brand",
		},
		"subdomain": schema.StringAttribute{
			Required:    true,
			Description: "The subdomain of the brand",
		},
		"ticket_form_ids": schema.ListAttribute{
			ElementType: types.Int64Type,
			Computed:    true,
			Description: "The ids of ticket forms that are available for use by a brand",
		},
		"brand_url": schema.StringAttribute{
			Computed:    true,
			Description: "The url of the brand",
		},
		"has_help_center": schema.BoolAttribute{
			Optional:    true,
			Computed:    true,
			Description: "If the brand has a Help Center",
		},
		"help_center_state": schema.StringAttribute{
			Computed:    true,
			Description: "The state of the Help Center. Allowed values are \"enabled\", \"disabled\", or \"restricted\".",
		},
		"active": schema.BoolAttribute{
			Computed:    true,
			Optional:    true,
			Description: "If the brand is set as active",
			Default:     booldefault.StaticBool(true),
		},
		"default": schema.BoolAttribute{
			Computed:    true,
			Optional:    true,
			Description: "Is the brand the default brand for this account",
		},
		"is_deleted": schema.BoolAttribute{
			Computed:    true,
			Optional:    true,
			Description: "If the brand object is deleted or not",
		},
		"host_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The hostmapping to this brand, if any. Only admins view this property.",
		},
		"signature_template": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The signature template for a brand",
		},
		"url": schema.StringAttribute{
			Computed:    true,
			Description: "The API url of this brand",
		},
		"created_at": schema.StringAttribute{
			Computed:    true,
			Description: "The time the brand was created",
		},
		"updated_at": schema.StringAttribute{
			Computed:    true,
			Description: "The time of the last update of the brand",
		},
	},
}
