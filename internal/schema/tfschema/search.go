package tfschema

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var SearchSchema = schema.Schema{
	MarkdownDescription: `The Search API is a unified search API that returns tickets, users, and organizations. 
You can define filters to narrow your search results according to resource type, dates, and object properties, such as 
ticket requester or tag. 

See [Search API Docs](https://developer.zendesk.com/api-reference/ticketing/ticket-management/search/) and
[Search Reference](https://support.zendesk.com/hc/en-us/articles/4408886879258-Zendesk-Support-search-reference)`,
	Attributes: map[string]schema.Attribute{
		"query": schema.StringAttribute{
			Required: true,
			MarkdownDescription: `Query to run a search for. 
See [Search Reference](https://support.zendesk.com/hc/en-us/articles/4408886879258-Zendesk-Support-search-reference)

*Make sure* to set 'type' search attribute to be either 'user' or 'organization'. If not, unexpected things will happen`,
		},
		"results": schema.SingleNestedAttribute{

			Computed: true,
			Attributes: map[string]schema.Attribute{
				"users": schema.ListNestedAttribute{
					Computed: true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"email": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"organization_id": schema.Int64Attribute{
								Computed: true,
							},
							"external_id": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
				"organizations": schema.ListNestedAttribute{
					Computed: true,
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
							"external_id": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	},
}
