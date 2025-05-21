package tfschema

import (
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var CustomRoleSchema = schema.Schema{
	Version: 0,
	MarkdownDescription: `
Zendesk Support accounts on the Enterprise plan or above can provide more granular access to their agents by defining custom agent roles. For more information, see [Creating custom roles and assigning agents in the Support Help Center](https://support.zendesk.com/hc/en-us/articles/203662026).
`,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"description": schema.StringAttribute{
			Description: "A description of the role",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("Role Description"),
		},
		"name": schema.StringAttribute{
			Description: "Name of the custom role",
			Required:    true,
		},
		"role_type": schema.Int64Attribute{
			Description: "The user's role. 0 stands for a custom agent, 1 for a light agent, 2 for a chat agent, 3 for a contributor, 4 for an admin and 5 for a billing admin.",
			Required:    true,
		},
		"configuration": schema.SingleNestedAttribute{
			MarkdownDescription: `
Configuration settings for the role. See [Configuration](https://developer.zendesk.com/api-reference/ticketing/account-configuration/custom_roles/#configuration)
`,
			Computed: true,
			Optional: true,
			Default: objectdefault.StaticValue(
				types.ObjectValueMust(models.ConfigurationResourceModel{}.AttributeTypes(), map[string]attr.Value{
					"assign_tickets_to_any_brand":     types.BoolValue(false),
					"assign_tickets_to_any_group":     types.BoolValue(false),
					"chat_access":                     types.BoolUnknown(),
					"custom_objects":                  types.MapNull(types.ObjectType{AttrTypes: models.ScopesResourceModel{}.AttributeTypes()}),
					"end_user_list_access":            types.StringValue("none"),
					"end_user_profile_access":         types.StringValue("readonly"),
					"explore_access":                  types.StringValue("none"),
					"forum_access":                    types.StringValue("readonly"),
					"forum_access_restricted_content": types.BoolValue(false),
					"light_agent":                     types.BoolUnknown(),
					"macro_access":                    types.StringValue("readonly"),
					"manage_automations":              types.BoolValue(false),
					"manage_business_rules":           types.BoolValue(false),
					"manage_contextual_workspaces":    types.BoolValue(false),
					"manage_dynamic_content":          types.BoolValue(false),
					"manage_extensions_and_channels":  types.BoolValue(false),
					"manage_facebook":                 types.BoolValue(false),
					"manage_group_memberships":        types.BoolValue(false),
					"manage_groups":                   types.BoolValue(false),
					"manage_organization_fields":      types.BoolValue(false),
					"manage_organizations":            types.BoolValue(false),
					"manage_roles":                    types.StringValue("none"),
					"manage_skills":                   types.BoolValue(false),
					"manage_slas":                     types.BoolValue(false),
					"manage_suspended_tickets":        types.BoolValue(false),
					"manage_team_members":             types.StringValue("readonly"),
					"manage_ticket_fields":            types.BoolValue(false),
					"manage_ticket_forms":             types.BoolValue(false),
					"manage_triggers":                 types.BoolValue(false),
					"manage_user_fields":              types.BoolValue(false),
					"organization_editing":            types.BoolValue(false),
					"organization_notes_editing":      types.BoolValue(false),
					"report_access":                   types.StringValue("none"),
					"side_conversation_create":        types.BoolValue(false),
					"ticket_access":                   types.StringValue("within-groups"),
					"ticket_comment_access":           types.StringValue("none"),
					"ticket_deletion":                 types.BoolValue(false),
					"ticket_redaction":                types.BoolValue(false),
					"view_deleted_tickets":            types.BoolValue(false),
					"ticket_editing":                  types.BoolValue(false),
					"ticket_merge":                    types.BoolValue(false),
					"ticket_tag_editing":              types.BoolValue(false),
					"twitter_search_access":           types.BoolValue(false),
					"view_access":                     types.StringValue("readonly"),
					"voice_access":                    types.BoolValue(false),
					"voice_dashboard_access":          types.BoolValue(false),
				}),
			),
			Attributes: map[string]schema.Attribute{
				"assign_tickets_to_any_brand": schema.BoolAttribute{
					Description: "Whether or not the agent can assign tickets to any brand and list all brands associated with an account sorted by name",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"assign_tickets_to_any_group": schema.BoolAttribute{
					Description: "Whether or not the agent can assign tickets to any group",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"chat_access": schema.BoolAttribute{
					Description: "Whether or not the agent has access to Chat",
					Computed:    true,
					Optional:    true,
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"custom_objects": schema.MapAttribute{
					ElementType: types.ObjectType{AttrTypes: models.ScopesResourceModel{}.AttributeTypes()},
					Computed:    true,
					Optional:    true,
					Default:     mapdefault.StaticValue(types.MapNull(types.ObjectType{AttrTypes: models.ScopesResourceModel{}.AttributeTypes()})),
					Description: "A list of custom object keys mapped to JSON objects that define the agent's permissions (scopes) for each object. Allowed values: \"read\", \"update\", \"delete\", \"create\". The \"read\" permission is required if any other scopes are specified. Example: { \"shipment\": { \"scopes\": [\"read\", \"update\"] } }",
				},
				"end_user_list_access": schema.StringAttribute{
					Description: "Whether or not the agent can view lists of user profiles. Allowed values: \"full\", \"none\"",
					Default:     stringdefault.StaticString("none"),
					Computed:    true,
					Optional:    true,
				},
				"end_user_profile_access": schema.StringAttribute{
					Description: "What the agent can do with end-user profiles. Allowed values: \"edit\", \"edit-within-org\", \"full\", \"readonly\"",
					Default:     stringdefault.StaticString("readonly"),
					Computed:    true,
					Optional:    true,
				},
				"explore_access": schema.StringAttribute{
					Description: "Allowed values: \"edit\", \"full\", \"none\", \"readonly\"",
					Default:     stringdefault.StaticString("readonly"),
					Computed:    true,
					Optional:    true,
				},
				"forum_access": schema.StringAttribute{
					Description: "The kind of access the agent has to Guide. Allowed values: \"edit-topics\", \"full\", \"readonly\"",
					Default:     stringdefault.StaticString("readonly"),
					Computed:    true,
					Optional:    true,
				},
				"forum_access_restricted_content": schema.BoolAttribute{
					Description: "Can access restricted content in Guide",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"light_agent": schema.BoolAttribute{
					Description: "Has Light Agent Permissions",
					Computed:    true,
					Optional:    true,
					PlanModifiers: []planmodifier.Bool{
						boolplanmodifier.UseStateForUnknown(),
					},
				},
				"macro_access": schema.StringAttribute{
					Description: "What the agent can do with macros. Allowed values: \"full\", \"manage-group\", \"manage-personal\", \"readonly\"",
					Default:     stringdefault.StaticString("readonly"),
					Computed:    true,
					Optional:    true,
				},
				"manage_automations": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage automations\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_business_rules": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage schedules and view rules analysis\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_contextual_workspaces": schema.BoolAttribute{
					Description: "Whether or not the agent can view, add, and edit contextual workspaces",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_dynamic_content": schema.BoolAttribute{
					Description: "Whether or not the agent can access dynamic content\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_extensions_and_channels": schema.BoolAttribute{
					Description: "Whether or not the agent can manage channels and extensions\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_facebook": schema.BoolAttribute{
					Description: "Whether or not the agent can manage facebook pages\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_group_memberships": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage group memberships\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_groups": schema.BoolAttribute{
					Description: "Whether or not the agent can create and modify groups\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_organization_fields": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage organization fields\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_organizations": schema.BoolAttribute{
					Description: "Whether or not the agent can create and modify organizations\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_roles": schema.StringAttribute{
					Description: "Whether or not the agent can create and manage custom roles with the exception of the role they're currently assigned. Doesn't allow agents to update role assignments for other agents. Allowed values: \"all-except-self\", \"none\"",
					Default:     stringdefault.StaticString("none"),
					Computed:    true,
					Optional:    true,
				},
				"manage_skills": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage skills\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_slas": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage SLAs\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_suspended_tickets": schema.BoolAttribute{
					Description: "Whether or not the agent can manage suspended tickets\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_team_members": schema.StringAttribute{
					Description: "Whether or not the agent can manage team members. Allows agents to update role assignments for other agents. Allowed values: \"all-with-self-restriction\", \"readonly\", \"none\"\n",
					Default:     stringdefault.StaticString("readonly"),
					Computed:    true,
					Optional:    true,
				},
				"manage_ticket_fields": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage ticket fields\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_ticket_forms": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage ticket forms\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_triggers": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage triggers\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"manage_user_fields": schema.BoolAttribute{
					Description: "Whether or not the agent can create and manage user fields\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"organization_editing": schema.BoolAttribute{
					Description: "Whether or not the agent can add or modify organizations\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"organization_notes_editing": schema.BoolAttribute{
					Description: "Whether or not the agent can add or modify organization notes\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"report_access": schema.StringAttribute{
					Description: "What the agent can do with reports. Allowed values: \"full\", \"none\", \"readonly\"\n",
					Default:     stringdefault.StaticString("none"),
					Computed:    true,
					Optional:    true,
				},
				"side_conversation_create": schema.BoolAttribute{
					Description: "Whether or not the agent can contribute to side conversations\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"ticket_access": schema.StringAttribute{
					Description: "What kind of tickets the agent can access. Allowed values: \"all\", \"assigned-only\", \"within-groups\", \"within-groups-and-public-groups\", \"within-organization\". Agents must have \"all\" access to create or edit end users from the Agent Workspace. However, the ability to create or edit end users through the API is determined by the user's role, not by ticket_access.\n",
					Default:     stringdefault.StaticString("within-groups"),
					Computed:    true,
					Optional:    true,
				},
				"ticket_comment_access": schema.StringAttribute{
					Description: "What type of comments the agent can make. Allowed values: \"public\", \"none\"",
					Default:     stringdefault.StaticString("none"),
					Computed:    true,
					Optional:    true,
				},
				"ticket_deletion": schema.BoolAttribute{
					Description: "Whether or not the agent can delete tickets\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"ticket_redaction": schema.BoolAttribute{
					Description: "Whether or not the agent can redact content from tickets. Only applicable to tickets permitted by ticket_access\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"view_deleted_tickets": schema.BoolAttribute{
					Description: "Whether or not the agent can view deleted tickets\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"ticket_editing": schema.BoolAttribute{
					Description: "Whether or not the agent can edit ticket properties\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"ticket_merge": schema.BoolAttribute{
					Description: "Whether or not the agent can merge tickets\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"ticket_tag_editing": schema.BoolAttribute{
					Description: "Whether or not the agent can edit ticket tags\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"twitter_search_access": schema.BoolAttribute{
					Description: "can access twitter search",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"view_access": schema.StringAttribute{
					Description: "What the agent can do with views. Allowed values: \"full\", \"manage-group\", \"manage-personal\", \"playonly\", \"readonly\"\n",
					Default:     stringdefault.StaticString("readonly"),
					Computed:    true,
					Optional:    true,
				},
				"voice_access": schema.BoolAttribute{
					Description: "Whether or not the agent can answer and place calls to end users\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
				"voice_dashboard_access": schema.BoolAttribute{
					Description: "Whether or not the agent can view details about calls on the Talk dashboard\n",
					Default:     booldefault.StaticBool(false),
					Computed:    true,
					Optional:    true,
				},
			},
		},
	},
}
