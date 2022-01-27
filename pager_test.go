package pager

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPager(t *testing.T) {
	tests := []struct {
		name     string
		pager    chan []string
		expected [][]string
	}{
		{
			name:     "no data",
			pager:    New([]string{}, 3),
			expected: nil,
		},
		{
			name:  "small quantity of data than page size",
			pager: New([]string{"a", "b"}, 3),
			expected: [][]string{
				[]string{"a", "b"},
			},
		},
		{
			name:  "input same length as page size",
			pager: New([]string{"a", "b", "c"}, 3),
			expected: [][]string{
				[]string{"a", "b", "c"},
			},
		},
		{
			name:  "two pages",
			pager: New([]string{"a", "b", "c"}, 2),
			expected: [][]string{
				[]string{"a", "b"},
				[]string{"c"},
			},
		},
		{
			name:  "three pages",
			pager: New([]string{"a", "b", "c"}, 1),
			expected: [][]string{
				[]string{"a"},
				[]string{"b"},
				[]string{"c"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual [][]string
			for page := range test.pager {
				actual = append(actual, page)
			}
			if diff := cmp.Diff(test.expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
