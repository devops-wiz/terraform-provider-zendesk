package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
	"testing"
)

var testScheduleName = "test Schedule"

var testTimeZone = "Pacific Time (US & Canada)"

var testIntervalTf = ScheduleIntervalModel{
	StartTime: types.Int64Value(5),
	EndTime:   types.Int64Value(13),
}

var testDayObj, _ = types.ObjectValueFrom(context.Background(), testIntervalTf.AttributeTypes(), testIntervalTf)

var emptyDayObj = types.ObjectNull(ScheduleIntervalModel{}.AttributeTypes())

var testTfIntervalsSundayOnly = ScheduleIntervalObjectModel{
	Sunday:    testDayObj,
	Monday:    emptyDayObj,
	Tuesday:   emptyDayObj,
	Wednesday: emptyDayObj,
	Thursday:  emptyDayObj,
	Friday:    emptyDayObj,
	Saturday:  emptyDayObj,
}

var testTfIntervalsAllDays = ScheduleIntervalObjectModel{
	Sunday:    testDayObj,
	Monday:    testDayObj,
	Tuesday:   testDayObj,
	Wednesday: testDayObj,
	Thursday:  testDayObj,
	Friday:    testDayObj,
	Saturday:  testDayObj,
}

var intervalsObjOneDay, _ = types.ObjectValueFrom(context.Background(), testTfIntervalsSundayOnly.AttributeTypes(), testTfIntervalsSundayOnly)
var intervalsObjAllDays, _ = types.ObjectValueFrom(context.Background(), testTfIntervalsAllDays.AttributeTypes(), testTfIntervalsAllDays)

var testApiScheduleExpectedOneDay = zendesk.Schedule{
	Intervals: []zendesk.ScheduleInterval{
		{
			StartTime: 5 * 60,
			EndTime:   13 * 60,
		},
	},
	Name:     testScheduleName,
	TimeZone: testTimeZone,
}
var testApiScheduleInputOneDay = zendesk.Schedule{
	Id: testId,
	Intervals: []zendesk.ScheduleInterval{
		{
			StartTime: 5 * 60,
			EndTime:   13 * 60,
		},
	},
	Name:      testScheduleName,
	TimeZone:  testTimeZone,
	CreatedAt: &testCreatedAt,
	UpdatedAt: &testUpdatedAt,
}

var testApiScheduleExpectedAllDays = zendesk.Schedule{
	Intervals: []zendesk.ScheduleInterval{
		{
			StartTime: 5 * 60,
			EndTime:   13 * 60,
		},
		{
			StartTime: 5*60 + (1 * 60 * 24),
			EndTime:   13*60 + (1 * 60 * 24),
		},
		{
			StartTime: 5*60 + (2 * 60 * 24),
			EndTime:   13*60 + (2 * 60 * 24),
		},
		{
			StartTime: 5*60 + (3 * 60 * 24),
			EndTime:   13*60 + (3 * 60 * 24),
		},
		{
			StartTime: 5*60 + (4 * 60 * 24),
			EndTime:   13*60 + (4 * 60 * 24),
		},
		{
			StartTime: 5*60 + (5 * 60 * 24),
			EndTime:   13*60 + (5 * 60 * 24),
		},
		{
			StartTime: 5*60 + (6 * 60 * 24),
			EndTime:   13*60 + (6 * 60 * 24),
		},
	},
	Name:     testScheduleName,
	TimeZone: testTimeZone,
}
var testApiScheduleInputAllDays = zendesk.Schedule{
	Id: testId,
	Intervals: []zendesk.ScheduleInterval{
		{
			StartTime: 5 * 60,
			EndTime:   13 * 60,
		},
		{
			StartTime: 5*60 + (1 * 60 * 24),
			EndTime:   13*60 + (1 * 60 * 24),
		},
		{
			StartTime: 5*60 + (2 * 60 * 24),
			EndTime:   13*60 + (2 * 60 * 24),
		},
		{
			StartTime: 5*60 + (3 * 60 * 24),
			EndTime:   13*60 + (3 * 60 * 24),
		},
		{
			StartTime: 5*60 + (4 * 60 * 24),
			EndTime:   13*60 + (4 * 60 * 24),
		},
		{
			StartTime: 5*60 + (5 * 60 * 24),
			EndTime:   13*60 + (5 * 60 * 24),
		},
		{
			StartTime: 5*60 + (6 * 60 * 24),
			EndTime:   13*60 + (6 * 60 * 24),
		},
	},
	Name:      testScheduleName,
	TimeZone:  testTimeZone,
	CreatedAt: &testCreatedAt,
	UpdatedAt: &testUpdatedAt,
}

var testScheduleResourceInputOneDay = ScheduleResourceModel{
	Name:      types.StringValue(testScheduleName),
	TimeZone:  types.StringValue(testTimeZone),
	Intervals: intervalsObjOneDay,
}

var testScheduleResourceExpectedOneDay = ScheduleResourceModel{
	ID:        types.Int64Value(testId),
	Name:      types.StringValue(testScheduleName),
	TimeZone:  types.StringValue(testTimeZone),
	Intervals: intervalsObjOneDay,
	CreatedAt: types.StringValue(testCreatedAt.UTC().String()),
	UpdatedAt: types.StringValue(testUpdatedAt.UTC().String()),
}

var testScheduleResourceInputAllDays = ScheduleResourceModel{
	Name:      types.StringValue(testScheduleName),
	TimeZone:  types.StringValue(testTimeZone),
	Intervals: intervalsObjAllDays,
}

var testScheduleResourceExpectedAllDays = ScheduleResourceModel{
	ID:        types.Int64Value(testId),
	Name:      types.StringValue(testScheduleName),
	TimeZone:  types.StringValue(testTimeZone),
	Intervals: intervalsObjAllDays,
	CreatedAt: types.StringValue(testCreatedAt.UTC().String()),
	UpdatedAt: types.StringValue(testUpdatedAt.UTC().String()),
}

func TestScheduleResourceModel_GetApiModelFromTfModel(t *testing.T) {
	cases := []struct {
		testName string
		input    ScheduleResourceModel
		expected zendesk.Schedule
	}{
		{
			"should get api model with one day",
			testScheduleResourceInputOneDay,
			testApiScheduleExpectedOneDay,
		}, {
			"should get api model with all days",
			testScheduleResourceInputAllDays,
			testApiScheduleExpectedAllDays,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			out, diags := tc.input.GetApiModelFromTfModel(t.Context())
			if diags.HasError() {
				t.Fatalf("GetApiModelFromTfModel() got error: %v", diags.Errors())
			}
			if !reflect.DeepEqual(out, tc.expected) {
				t.Fatalf(errorOutputMismatch, tc.testName, out, tc.expected)
			}
		})
	}
}

func TestScheduleResourceModel_GetTfModelFromApiModel(t *testing.T) {
	cases := []struct {
		testName string
		target   ScheduleResourceModel
		input    zendesk.Schedule
		expected ScheduleResourceModel
	}{
		{
			"should get tf model with one day",
			ScheduleResourceModel{},
			testApiScheduleInputOneDay,
			testScheduleResourceExpectedOneDay,
		},
		{
			"should get tf model with all days",
			ScheduleResourceModel{},
			testApiScheduleInputAllDays,
			testScheduleResourceExpectedAllDays,
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			diags := tc.target.GetTfModelFromApiModel(t.Context(), tc.input)
			if diags.HasError() {
				t.Fatalf("GetTfModelFromApiModel() got error: %v", diags.Errors())
			}
			if !reflect.DeepEqual(tc.target, tc.expected) {
				t.Fatalf(errorOutputMismatch, tc.testName, tc.target, tc.expected)
			}
		})
	}
}
