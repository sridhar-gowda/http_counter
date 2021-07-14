package utils

import (
	"reflect"
	"testing"
)

func TestGetLatestWindow(t *testing.T) {
	tables := []struct {
		now  int64
		size int
		out  int64
	}{
		{now: 1620643425, size: 10, out: 1620643415},
		{now: 1620643425, size: 100, out: 1620643325},
		{now: 1620643425, size: 0, out: 16206434251},
	}
	for _, table := range tables {
		output := GetLatestWindow(table.now, table.size)
		if !reflect.DeepEqual(output, table.out) {
			t.Errorf("test case failed for input %d, %d", table.now, table.size)
		}
	}
}
