package pager_test

import (
	"io"
	"testing"

	"github.com/a-h/pager"
	"github.com/google/go-cmp/cmp"
)

var tests = []struct {
	name     string
	data     []string
	pageSize int
	expected [][]string
}{
	{
		name:     "no data",
		data:     []string{},
		pageSize: 3,
		expected: nil,
	},
	{
		name:     "quantity of data less than page size",
		data:     []string{"a", "b"},
		pageSize: 3,
		expected: [][]string{
			{"a", "b"},
		},
	},
	{
		name:     "input same length as page size",
		data:     []string{"a", "b", "c"},
		pageSize: 3,
		expected: [][]string{
			{"a", "b", "c"},
		},
	},
	{
		name:     "two pages",
		data:     []string{"a", "b", "c"},
		pageSize: 2,
		expected: [][]string{
			{"a", "b"},
			{"c"},
		},
	},
	{
		name:     "three pages",
		data:     []string{"a", "b", "c"},
		pageSize: 1,
		expected: [][]string{
			{"a"},
			{"b"},
			{"c"},
		},
	},
}

func TestChannel(t *testing.T) {
	for _, test := range tests {
		t.Run("Channel: "+test.name, func(t *testing.T) {
			var actual [][]string
			for page := range pager.Channel(test.data, test.pageSize) {
				actual = append(actual, page)
			}
			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestFunc(t *testing.T) {
	for _, test := range tests {
		t.Run("Func: "+test.name, func(t *testing.T) {
			var actual [][]string
			err := pager.Func(test.data, test.pageSize, func(page []string) error {
				actual = append(actual, page)
				return nil
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func BenchmarkChannel(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for page := range pager.Channel(tests[3].data, tests[3].pageSize) {
			io.Discard.Write([]byte(page[0]))
		}
	}
}

func BenchmarkFunc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pager.Func(tests[3].data, tests[3].pageSize, func(page []string) error {
			io.Discard.Write([]byte(page[0]))
			return nil
		})
	}
}
