package main

import (
	"testing"
	"reflect"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name         	string
		inputSlice     	[]page
		expected      	[]page
	}{
		{
			name:     	"different names and visits",
			inputSlice:	[]page {
				{name: 	"c", visits: 1},
				{name: 	"b", visits: 1},
				{name: 	"a", visits: 1},
				{name: 	"d", visits: 2},
				{name: 	"e", visits: 3},
			},
			expected: []page {
				{name: 	"e", visits: 3},
				{name: 	"d", visits: 2},
				{name: 	"a", visits: 1},
				{name: 	"b", visits: 1},
				{name: 	"c", visits: 1},
			},
		},
		{
			name:     	"different names, same visits",
			inputSlice:	[]page {
				{name: 	"c", visits: 2},
				{name: 	"b", visits: 2},
				{name: 	"a", visits: 2},
				{name: 	"d", visits: 2},
				{name: 	"e", visits: 2},
			},
			expected: []page {
				{name: 	"a", visits: 2},
				{name: 	"b", visits: 2},
				{name: 	"c", visits: 2},
				{name: 	"d", visits: 2},
				{name: 	"e", visits: 2},
			},
		},
		{
			name:     	"same names, different visits",
			inputSlice:	[]page {
				{name: 	"a", visits: 2},
				{name: 	"a", visits: 12},
				{name: 	"a", visits: 15},
				{name: 	"a", visits: 82},
				{name: 	"a", visits: 20},
			},
			expected: []page {
				{name: 	"a", visits: 82},
				{name: 	"a", visits: 20},
				{name: 	"a", visits: 15},
				{name: 	"a", visits: 12},
				{name: 	"a", visits: 2},
			},
		},
		{
			name:     	"empty slice",
			inputSlice:	[]page {},
			expected: []page {},
		},
		{
			name:     	"one key",
			inputSlice:	[]page {
				{name: 	"a", visits: 1},
			},
			expected: []page {
				{name: 	"a", visits: 1},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.inputSlice)
			
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected Slice: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}