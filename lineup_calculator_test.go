package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	contents = []content{
		{1, "title1"},
		{2, "title2"},
		{3, "title3"},
		{4, "title4"},
	}

	userInteractions = []interaction{
		{1, 1, "comment"},
		{2, 2, "like"},
		{1, 2, "view"},
		{1, 3, "like"},
	}
)

func Test_lineupCalculator_calculate(t *testing.T) {
	lc := newCalculator()
	result := lc.calculate(contents, userInteractions)

	expected := &map[int]int{
		1: 3,
		2: 1,
	}

	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}

func Test_newLineupCalculator(t *testing.T) {
	t.SkipNow()
}
