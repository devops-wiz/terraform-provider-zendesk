package tfschema

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func GetUserOrgFieldSchema(fieldType string) schema.Schema {
	var fieldMarkdownDescription string

	switch fieldType {
	case "org":
		fieldMarkdownDescription = `You can use this API to add fields to the Organization page in the Zendesk user interface. 
Basic text fields, date fields, as well as customizable drop-down and number fields are available. 
The fields correspond to the organization fields that admins can add using the Zendesk admin interface.
See [Adding custom fields to organizations](https://support.zendesk.com/hc/en-us/articles/203662076) in Zendesk help.`
	case "user":
		fieldMarkdownDescription = `You can use this API to add fields to the user profile page in the Zendesk user interface. 
Basic text fields, date fields, as well as customizable drop-down and number fields are available. 
The fields correspond to the user fields that admins can add using the Zendesk admin interface. 
See [Adding custom fields to users](https://support.zendesk.com/hc/en-us/articles/203662066) in Zendesk help.`
	}
	return schema.Schema{
		MarkdownDescription: fieldMarkdownDescription,
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Description: "If true, this field is available for use",
				Optional:    true,
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: fmt.Sprintf("The time of the last update of the %s field.", fieldType),
				Computed:    true,
			},
			"custom_field_options": schema.ListNestedAttribute{
				MarkdownDescription: "Required and presented for a custom field of type \"dropdown\". Each option is represented by an object with a `name` and `value` property.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
							Description: "ID of the custom field option.",
						},
						"name": schema.StringAttribute{
							Required:    true,
							Description: "Display name of the custom field.",
						},
						"value": schema.StringAttribute{
							Required:    true,
							Description: "Tag value of the custom field.",
						},
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "User-defined description of this field's purpose",
				Optional:    true,
				Computed:    true,
			},
			"id": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "Automatically assigned upon creation.",
			},
			"key": schema.StringAttribute{
				Required: true,
				Description: "A unique key that identifies this custom field. " +
					"This is used for updating the field and referencing in placeholders. " +
					"The key must consist of only letters, numbers, and underscores. It can't be only numbers.",
			},
			"position": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Ordering of the field relative to other fields",
			},
			"regexp_for_validation": schema.StringAttribute{
				Optional:    true,
				Description: "Regular expression field only. The validation pattern for a field value to be deemed valid",
			},
			"relationship_filter": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"all": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"field": schema.StringAttribute{
									Required: true,
								},
								"operator": schema.StringAttribute{
									Required: true,
								},
								"value": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
					"any": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"field": schema.StringAttribute{
									Required: true,
								},
								"operator": schema.StringAttribute{
									Required: true,
								},
								"value": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
				},
				MarkdownDescription: `A filter definition that allows your autocomplete to filter down results. 

A condition that defines a subset of records as the options in your lookup relationship field. 
See [Filtering the field's options](https://support.zendesk.com/hc/en-us/articles/4591924111770#topic_t14_w3l_5tb) 
in Zendesk help and [Conditions reference](https://developer.zendesk.com/documentation/ticketing/reference-guides/conditions-reference/)`,
				Optional: true,
			},
			"relationship_target_type": schema.StringAttribute{
				Optional: true,
				Description: "A representation of what type of object the field references. " +
					"Options are \"zen:user\", \"zen:organization\", \"zen:ticket\", and \"zen:custom_object:{key}\" where " +
					"key is a custom object key. For example \"zen:custom_object:apartment\".",
			},
			"system": schema.BoolAttribute{
				Computed:    true,
				Description: "If true, only active and position values of this field can be changed",
			},
			"tag": schema.StringAttribute{
				Optional:    true,
				Description: "Optional for custom field of type \"checkbox\"; not presented otherwise.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The title of the custom field",
			},
			"type": schema.StringAttribute{
				Required: true,
				MarkdownDescription: "The custom field type: \"checkbox\", \"date\", \"decimal\", \"dropdown\", \"integer\", " +
					"[\"lookup\"](https://developer.zendesk.com/api-reference/ticketing/lookup_relationships/lookup_relationships/), " +
					"\"multiselect\", \"regexp\", \"text\", or \"textarea\"",
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: "The time of the last update of the ticket field",
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: "The URL for this resource",
			},
		},
	}
}
