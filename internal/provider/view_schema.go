package provider

import (
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var ViewSchema = schema.Schema{
	Version: 1,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"conditions": schema.SingleNestedAttribute{
			Required:    true,
			Description: "Conditions for a view",
			Attributes: map[string]schema.Attribute{
				"all": schema.ListNestedAttribute{
					Description:  "Logical AND. All the conditions must be met",
					NestedObject: GetNestedConditionObject("view"),
					Optional:     true,
				},
				"any": schema.ListNestedAttribute{
					Description:  "Logical OR. Any condition can be met",
					NestedObject: GetNestedConditionObject("view"),
					Optional:     true,
				},
			},
		},
		"output": schema.SingleNestedAttribute{
			Description: "Columns to view and grouping/sorting rules for view",
			Required:    true,
			Attributes: map[string]schema.Attribute{
				"columns": schema.ListAttribute{
					Description: "Columns to use in the view",
					ElementType: types.StringType,
					Required:    true,
					Validators: []validator.List{
						listvalidator.ValueStringsAre(stringvalidator.OneOf(zendesk.ValidViewColumns.StringsSlice()...)),
					},
				},
				"group_by": schema.StringAttribute{
					Description: "",
					Required:    true,
					Validators: []validator.String{
						stringvalidator.OneOf(zendesk.ValidViewColumns.StringsSlice()...),
					},
				},
				"group_order": schema.StringAttribute{
					Description: "",
					Required:    true,
				},
				"sort_by": schema.StringAttribute{
					Description: "",
					Required:    true,
					Validators: []validator.String{
						stringvalidator.OneOf(zendesk.ValidViewColumns.StringsSlice()...),
					},
				},
				"sort_order": schema.StringAttribute{
					Description: "",
					Required:    true,
				},
			},
		},
		"title": schema.StringAttribute{
			Required: true,
		},
		"restriction": schema.SingleNestedAttribute{
			Description: "An object that describes who can access the view. To give all agents access to the view, omit this property.",
			Optional:    true,
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					Description: "Required. Allowed values are 'Group' or 'User'.",
					Required:    true,
					Validators: []validator.String{
						stringvalidator.OneOf([]string{"Group", "User"}...),
					},
				},
				"ids": schema.SetAttribute{
					Description: "The numeric IDs of the groups or users.",
					ElementType: types.Int64Type,
					Required:    true,
				},
			},
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the view is displayed or not.",
			Computed:    true,
			Optional:    true,
			Default:     booldefault.StaticBool(true),
		},
		"description": schema.StringAttribute{
			Description: "The description of the view.",
			Optional:    true,
			Computed:    true,
		},
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The relative position of the view",
		},
		"created_at": schema.StringAttribute{
			Description: "The time the view was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the view.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the view's details.",
			Computed:    true,
		},
	},
}

var ViewSchemaV0 = schema.Schema{
	Version: 0,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"conditions": schema.SingleNestedAttribute{
			Required:    true,
			Description: "Conditions for a view",
			Attributes: map[string]schema.Attribute{
				"all": schema.ListNestedAttribute{
					Description:  "Logical AND. All the conditions must be met",
					NestedObject: GetNestedConditionObjectV0("view"),
					Optional:     true,
				},
				"any": schema.ListNestedAttribute{
					Description:  "Logical OR. Any condition can be met",
					NestedObject: GetNestedConditionObjectV0("view"),
					Optional:     true,
				},
			},
		},
		"output": schema.SingleNestedAttribute{
			Description: "Columns to view and grouping/sorting rules for view",
			Required:    true,
			Attributes: map[string]schema.Attribute{
				"columns": schema.ListAttribute{
					Description: "Columns to use in the view",
					ElementType: types.StringType,
					Required:    true,
					Validators: []validator.List{
						listvalidator.ValueStringsAre(stringvalidator.OneOf(zendesk.ValidViewColumns.StringsSlice()...)),
					},
				},
				"group_by": schema.StringAttribute{
					Description: "",
					Required:    true,
					Validators: []validator.String{
						stringvalidator.OneOf(zendesk.ValidViewColumns.StringsSlice()...),
					},
				},
				"group_order": schema.StringAttribute{
					Description: "",
					Required:    true,
				},
				"sort_by": schema.StringAttribute{
					Description: "",
					Required:    true,
					Validators: []validator.String{
						stringvalidator.OneOf(zendesk.ValidViewColumns.StringsSlice()...),
					},
				},
				"sort_order": schema.StringAttribute{
					Description: "",
					Required:    true,
				},
			},
		},
		"title": schema.StringAttribute{
			Required: true,
		},
		"restriction": schema.SingleNestedAttribute{
			Description: "An object that describes who can access the view. To give all agents access to the view, omit this property.",
			Optional:    true,
			Attributes: map[string]schema.Attribute{
				"type": schema.StringAttribute{
					Description: "Required. Allowed values are 'Group' or 'User'.",
					Required:    true,
					Validators: []validator.String{
						stringvalidator.OneOf([]string{"Group", "User"}...),
					},
				},
				"ids": schema.SetAttribute{
					Description: "The numeric IDs of the groups or users.",
					ElementType: types.Int64Type,
					Required:    true,
				},
			},
		},
		"active": schema.BoolAttribute{
			Description: "Allowed values are true or false. Determines if the view is displayed or not.",
			Computed:    true,
			Optional:    true,
			Default:     booldefault.StaticBool(true),
		},
		"description": schema.StringAttribute{
			Description: "The description of the view.",
			Optional:    true,
			Computed:    true,
		},
		"position": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The relative position of the view",
		},
		"created_at": schema.StringAttribute{
			Description: "The time the view was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the view.",
			Computed:    true,
		},
		"url": schema.StringAttribute{
			Description: "A URL to access the view's details.",
			Computed:    true,
		},
	},
}
