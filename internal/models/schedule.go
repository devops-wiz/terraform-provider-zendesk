package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type ScheduleIntervalModel struct {
	StartTime types.Int64 `tfsdk:"start_time"`
	EndTime   types.Int64 `tfsdk:"end_time"`
}

func (m ScheduleIntervalModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"start_time": types.Int64Type,
		"end_time":   types.Int64Type,
	}
}

type ScheduleIntervalObjectModel struct {
	Sunday    types.Object `tfsdk:"sunday"`
	Monday    types.Object `tfsdk:"monday"`
	Tuesday   types.Object `tfsdk:"tuesday"`
	Wednesday types.Object `tfsdk:"wednesday"`
	Thursday  types.Object `tfsdk:"thursday"`
	Friday    types.Object `tfsdk:"friday"`
	Saturday  types.Object `tfsdk:"saturday"`
}

func (s ScheduleIntervalObjectModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"sunday":    types.ObjectType{AttrTypes: ScheduleIntervalModel{}.AttributeTypes()},
		"monday":    types.ObjectType{AttrTypes: ScheduleIntervalModel{}.AttributeTypes()},
		"tuesday":   types.ObjectType{AttrTypes: ScheduleIntervalModel{}.AttributeTypes()},
		"wednesday": types.ObjectType{AttrTypes: ScheduleIntervalModel{}.AttributeTypes()},
		"thursday":  types.ObjectType{AttrTypes: ScheduleIntervalModel{}.AttributeTypes()},
		"friday":    types.ObjectType{AttrTypes: ScheduleIntervalModel{}.AttributeTypes()},
		"saturday":  types.ObjectType{AttrTypes: ScheduleIntervalModel{}.AttributeTypes()},
	}
}

type ScheduleResourceModel struct {
	ID        types.Int64  `tfsdk:"id"`
	Intervals types.Object `tfsdk:"intervals"`
	Name      types.String `tfsdk:"name"`
	TimeZone  types.String `tfsdk:"time_zone"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (s *ScheduleResourceModel) GetID() int64 {
	return s.ID.ValueInt64()
}

func (s *ScheduleResourceModel) GetApiModelFromTfModel(ctx context.Context) (schedule zendesk.Schedule, diags diag.Diagnostics) {

	var apiIntervals []zendesk.ScheduleInterval

	if !s.Intervals.IsNull() && !s.Intervals.IsUnknown() {
		apiIntervals, diags = getApiIntervalsFromTf(ctx, s.Intervals)
		if diags.HasError() {
			return schedule, diags
		}
	}

	schedule = zendesk.Schedule{
		Intervals: apiIntervals,
		Name:      s.Name.ValueString(),
		TimeZone:  s.TimeZone.ValueString(),
	}

	return schedule, diags
}

func getApiIntervalsFromTf(ctx context.Context, tfIntervalsObj types.Object) (intervals []zendesk.ScheduleInterval, diags diag.Diagnostics) {
	var intervalsModel ScheduleIntervalObjectModel

	diags = tfIntervalsObj.As(ctx, &intervalsModel, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    false,
		UnhandledUnknownAsEmpty: false,
	})

	if diags.HasError() {
		return intervals, diags
	}

	if !intervalsModel.Sunday.IsNull() && !intervalsModel.Sunday.IsUnknown() {

		var interval zendesk.ScheduleInterval
		interval, diags = getApiIntervalFromTf(ctx, intervalsModel.Sunday, 0)
		intervals = append(intervals, interval)
	}
	if !intervalsModel.Monday.IsNull() && !intervalsModel.Monday.IsUnknown() {

		var interval zendesk.ScheduleInterval
		interval, diags = getApiIntervalFromTf(ctx, intervalsModel.Monday, 1)
		intervals = append(intervals, interval)
	}
	if !intervalsModel.Tuesday.IsNull() && !intervalsModel.Tuesday.IsUnknown() {

		var interval zendesk.ScheduleInterval
		interval, diags = getApiIntervalFromTf(ctx, intervalsModel.Tuesday, 2)
		intervals = append(intervals, interval)
	}
	if !intervalsModel.Wednesday.IsNull() && !intervalsModel.Wednesday.IsUnknown() {

		var interval zendesk.ScheduleInterval
		interval, diags = getApiIntervalFromTf(ctx, intervalsModel.Wednesday, 3)
		intervals = append(intervals, interval)
	}
	if !intervalsModel.Thursday.IsNull() && !intervalsModel.Thursday.IsUnknown() {

		var interval zendesk.ScheduleInterval
		interval, diags = getApiIntervalFromTf(ctx, intervalsModel.Thursday, 4)
		intervals = append(intervals, interval)
	}
	if !intervalsModel.Friday.IsNull() && !intervalsModel.Friday.IsUnknown() {

		var interval zendesk.ScheduleInterval
		interval, diags = getApiIntervalFromTf(ctx, intervalsModel.Friday, 5)
		intervals = append(intervals, interval)
	}
	if !intervalsModel.Saturday.IsNull() && !intervalsModel.Saturday.IsUnknown() {

		var interval zendesk.ScheduleInterval
		interval, diags = getApiIntervalFromTf(ctx, intervalsModel.Saturday, 6)
		intervals = append(intervals, interval)
	}

	return intervals, diags

}

func getApiIntervalFromTf(ctx context.Context, dayIntervalObj types.Object, dayOffset int) (interval zendesk.ScheduleInterval, diags diag.Diagnostics) {
	var intervalModel ScheduleIntervalModel

	diags = dayIntervalObj.As(ctx, &intervalModel, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    false,
		UnhandledUnknownAsEmpty: false,
	})

	if diags.HasError() {
		return interval, diags
	}

	startTime := intervalModel.StartTime.ValueInt64()
	endTime := intervalModel.EndTime.ValueInt64()

	interval = zendesk.ScheduleInterval{
		StartTime: int(startTime)*60 + (dayOffset * 24 * 60),
		EndTime:   int(endTime)*60 + (dayOffset * 24 * 60),
	}

	return interval, diags
}

func (s *ScheduleResourceModel) GetTfModelFromApiModel(ctx context.Context, schedule zendesk.Schedule) (diags diag.Diagnostics) {
	var tfIntervals types.Object

	if len(schedule.Intervals) > 0 {
		tfIntervals, diags = getTfIntervalsFromApi(ctx, schedule.Intervals)
	} else {
		tfIntervals = types.ObjectNull(ScheduleIntervalObjectModel{}.AttributeTypes())
	}

	*s = ScheduleResourceModel{
		ID:        types.Int64Value(int64(schedule.Id)),
		Intervals: tfIntervals,
		Name:      types.StringValue(schedule.Name),
		TimeZone:  types.StringValue(schedule.TimeZone),
		CreatedAt: types.StringValue(schedule.CreatedAt.UTC().String()),
		UpdatedAt: types.StringValue(schedule.UpdatedAt.UTC().String()),
	}

	return diags
}

func getTfIntervalsFromApi(ctx context.Context, intervals []zendesk.ScheduleInterval) (intervalsObj types.Object, diags diag.Diagnostics) {
	var sundayObj types.Object
	var mondayObj types.Object
	var tuesdayObj types.Object
	var wednesdayObj types.Object
	var thursdayObj types.Object
	var fridayObj types.Object
	var saturdayObj types.Object

	for _, interval := range intervals {
		if sundayObj.IsNull() || sundayObj.IsUnknown() {
			sundayObj, diags = getTfIntervalFromApi(ctx, interval, 0)
			if diags.HasError() {
				return intervalsObj, diags
			}
		}

		if mondayObj.IsNull() || mondayObj.IsUnknown() {
			mondayObj, diags = getTfIntervalFromApi(ctx, interval, 1)
			if diags.HasError() {
				return intervalsObj, diags
			}
		}

		if tuesdayObj.IsNull() || tuesdayObj.IsUnknown() {
			tuesdayObj, diags = getTfIntervalFromApi(ctx, interval, 2)
			if diags.HasError() {
				return intervalsObj, diags
			}
		}

		if wednesdayObj.IsNull() || wednesdayObj.IsUnknown() {
			wednesdayObj, diags = getTfIntervalFromApi(ctx, interval, 3)
			if diags.HasError() {
				return intervalsObj, diags
			}
		}

		if thursdayObj.IsNull() || thursdayObj.IsUnknown() {
			thursdayObj, diags = getTfIntervalFromApi(ctx, interval, 4)
			if diags.HasError() {
				return intervalsObj, diags
			}
		}

		if fridayObj.IsNull() || fridayObj.IsUnknown() {
			fridayObj, diags = getTfIntervalFromApi(ctx, interval, 5)
			if diags.HasError() {
				return intervalsObj, diags
			}
		}

		if saturdayObj.IsNull() || saturdayObj.IsUnknown() {
			saturdayObj, diags = getTfIntervalFromApi(ctx, interval, 6)
			if diags.HasError() {
				return intervalsObj, diags
			}
		}

	}

	scheduleIntervalObjModel := ScheduleIntervalObjectModel{
		Sunday:    sundayObj,
		Monday:    mondayObj,
		Tuesday:   tuesdayObj,
		Wednesday: wednesdayObj,
		Thursday:  thursdayObj,
		Friday:    fridayObj,
		Saturday:  saturdayObj,
	}

	intervalsObj, diags = types.ObjectValueFrom(ctx, scheduleIntervalObjModel.AttributeTypes(), scheduleIntervalObjModel)

	return intervalsObj, diags
}

func getTfIntervalFromApi(ctx context.Context, interval zendesk.ScheduleInterval, dayOffset int) (intervalObj types.Object, diags diag.Diagnostics) {

	if interval.StartTime >= dayOffset*24*60 && interval.StartTime < 24*60+(dayOffset*24*60) {
		intervalModel := ScheduleIntervalModel{
			StartTime: types.Int64Value(int64((interval.StartTime / 60) - (dayOffset * 24))),
			EndTime:   types.Int64Value(int64((interval.EndTime / 60) - (dayOffset * 24))),
		}
		intervalObj, diags = types.ObjectValueFrom(ctx, intervalModel.AttributeTypes(), intervalModel)

		if diags.HasError() {
			return intervalObj, diags
		}
	} else {
		intervalObj = types.ObjectNull(ScheduleIntervalModel{}.AttributeTypes())
	}

	return intervalObj, diags
}
