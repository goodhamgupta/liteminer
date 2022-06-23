package test

import (
	liteminer "liteminer/pkg"
    "testing"
)

func TestGenerateIntervals(t *testing.T) {
    var intervals = liteminer.GenerateIntervals(uint64(4), 2)
    var got = len(intervals)
    var want = 2
    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}
