package tfschema

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var intervalSchema = schema.SingleNestedAttribute{
	Optional: true,
	Attributes: map[string]schema.Attribute{
		"start_time": schema.Int64Attribute{
			Required:    true,
			Description: "Start time, offset from beginning of day in hours",
		},
		"end_time": schema.Int64Attribute{
			Required:    true,
			Description: "End time, offset from beginning of day in hours",
		},
	},
}

var ScheduleSchema = schema.Schema{
	Version: 0,
	MarkdownDescription: `
You can set a schedule in Zendesk to acknowledge your support team's availability and give customers a better sense of when they can expect a personal response to their support requests.

You can use this API to create multiple schedules with different business hours and holidays.

To learn more about schedules, see [Setting your schedule with business hours and holidays in Zendesk help](https://support.zendesk.com/hc/en-us/articles/203662206).
`,
	Attributes: map[string]schema.Attribute{
		"id": schema.Int64Attribute{
			Computed: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			Required:    true,
			Description: "Name of the schedule",
		},
		"time_zone": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: `Time zone of the schedule, see [Time Zones](https://developer.zendesk.com/api-reference/introduction/data-types/#time-zones)`,
		},
		"intervals": schema.SingleNestedAttribute{
			Optional:    true,
			Description: "Schedule intervals divided by day of week.",
			Attributes: map[string]schema.Attribute{
				"sunday":    intervalSchema,
				"monday":    intervalSchema,
				"tuesday":   intervalSchema,
				"wednesday": intervalSchema,
				"thursday":  intervalSchema,
				"friday":    intervalSchema,
				"saturday":  intervalSchema,
			},
		},
		"created_at": schema.StringAttribute{
			Description: "The time the schedule was created.",
			Computed:    true,
		},
		"updated_at": schema.StringAttribute{
			Description: "The time of the last update of the schedule.",
			Computed:    true,
		},
	},
}
