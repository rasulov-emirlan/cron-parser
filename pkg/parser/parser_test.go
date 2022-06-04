package parser_test

import (
	"reflect"
	"testing"

	"github.com/rasulov-emirlan/cron-parser/pkg/parser"
)

func TestParseMinute(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  []int
	}{
		{
			desc:  "single -",
			input: "1-5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "single / with number before it",
			input: "1/5",
			want:  []int{1, 6, 11, 16, 21, 26, 31, 36, 41, 46, 51, 56},
		},
		{
			desc:  "commas",
			input: "1,2,3,4,5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "commas with -",
			input: "1,2-5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "commas with /",
			input: "1,2/5",
			want:  []int{1, 2, 7, 12, 17, 22, 27, 32, 37, 42, 47, 52, 57},
		},
		{
			desc:  "commas with / and -",
			input: "1,2-5/5",
			want:  []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := parser.NewParser()

			err := service.ParseMinute(tC.input)
			if err != nil {
				t.Errorf("parseMinute(%s) error = %v", tC.input, err)
			}
			if !reflect.DeepEqual(service.Minute, tC.want) {
				t.Errorf("parseMinute(%s) = %v, want %v", tC.input, service.Minute, tC.want)
			}
		})
	}
}

func TestParseHours(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  []int
	}{
		{
			desc:  "single -",
			input: "1-5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "single / with number before it",
			input: "1/5",
			want:  []int{1, 6, 11, 16, 21, 26, 31, 36, 41, 46, 51, 56},
		},
		{
			desc:  "commas",
			input: "1,2,3,4,5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "commas with -",
			input: "1,2-5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "commas with /",
			input: "1,2/5",
			want:  []int{1, 2, 7, 12, 17, 22, 27, 32, 37, 42, 47, 52, 57},
		},
		{
			desc:  "commas with / and -",
			input: "1,2-5/5",
			want:  []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := parser.NewParser()

			err := service.ParseHour(tC.input)
			if err != nil {
				t.Errorf("parseHour(%s) error = %v", tC.input, err)
			}
			if !reflect.DeepEqual(service.Hour, tC.want) {
				t.Errorf("parseHour(%s) = %v, want %v", tC.input, service.Hour, tC.want)
			}
		})
	}
}

func TestParseMonth(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  []int
	}{
		{
			desc:  "single -",
			input: "1-5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "single / with number before it",
			input: "1/5",
			want:  []int{1, 6, 11},
		},
		{
			desc:  "commas",
			input: "1,2,3,4,5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "commas with -",
			input: "1,2-5",
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			desc:  "commas with /",
			input: "1,2/5",
			want:  []int{1, 2, 7, 12},
		},
		{
			desc:  "commas with / and -",
			input: "1,2-5/5",
			want:  []int{1, 2, 3, 4, 5, 10},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := parser.NewParser()

			err := service.ParseMonth(tC.input)
			if err != nil {
				t.Errorf("parseMonth(%s) error = %v", tC.input, err)
			}
			if !reflect.DeepEqual(service.Month, tC.want) {
				t.Errorf("parseMonth(%s) = %v, want %v", tC.input, service.Month, tC.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  parser.CronCommand
	}{
		// {
		// 	desc:  "all",
		// 	input: `*/15 0 1,15 * 1-5 /usr/bin/find`,
		// 	want: parser.CronCommand{
		// 		Minute:     []int{0, 15, 30, 45},
		// 		Hour:       []int{0},
		// 		DayOfMonth: []int{1, 15},
		// 		Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		// 		DayOfWeek:  []int{1, 2, 3, 4, 5},
		// 		Command:    "/usr/bin/find",
		// 	},
		// },
		// {
		// 	desc:  "second try",
		// 	input: `*/15 1-5 1,15 * 1-5 /usr/bin/find`,
		// 	want: parser.CronCommand{
		// 		Minute:     []int{0, 15, 30, 45},
		// 		Hour:       []int{1, 2, 3, 4, 5},
		// 		DayOfMonth: []int{1, 15},
		// 		Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		// 		DayOfWeek:  []int{1, 2, 3, 4, 5},
		// 		Command:    "/usr/bin/find",
		// 	},
		// },
		{
			desc:  "third try",
			input: `*/15 1-5/21 1,15 * 1-5 /usr/bin/find`,
			want: parser.CronCommand{
				Minute:     []int{0, 15, 30, 45},
				Hour:       []int{1, 2, 3, 4, 5},
				DayOfMonth: []int{1, 15},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5},
				Command:    "/usr/bin/find",
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			service := parser.NewParser()

			err := service.ParseAll(tC.input)
			if err != nil {
				t.Errorf("parse(%s) error = %v", tC.input, err)
			}
			if !reflect.DeepEqual(service.Minute, tC.want.Minute) {
				t.Errorf("parse(%s) = %v, want %v", tC.input, service.Minute, tC.want.Minute)
			}
			if !reflect.DeepEqual(service.Hour, tC.want.Hour) {
				t.Errorf("parse(%s) = %v, want %v", tC.input, service.Hour, tC.want.Hour)
			}
			if !reflect.DeepEqual(service.DayOfMonth, tC.want.DayOfMonth) {
				t.Errorf("parse(%s) = %v, want %v", tC.input, service.DayOfMonth, tC.want.DayOfMonth)
			}
			if !reflect.DeepEqual(service.Month, tC.want.Month) {
				t.Errorf("parse(%s) = %v, want %v", tC.input, service.Month, tC.want.Month)
			}
			if !reflect.DeepEqual(service.DayOfWeek, tC.want.DayOfWeek) {
				t.Errorf("parse(%s) = %v, want %v", tC.input, service.DayOfWeek, tC.want.DayOfWeek)
			}
			if !reflect.DeepEqual(service.Command, tC.want.Command) {
				t.Errorf("parse(%s) = %v, want %v", tC.input, service.Command, tC.want.Command)
			}
		})
	}
}
