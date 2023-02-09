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
		{1, 1, "comment", "hot"},
		{2, 2, "love", "hot"},
		{1, 2, "view", "news"},
		{1, 3, "like", "news"},
		{1, 3, "share", "hot"},
	}
)

func Test_lineupCalculator_calculate(t *testing.T) {
	lc := newCalculator()
	result := lc.calculate(contents, userInteractions)

	expected := &map[int]float64{
		1: 8.0,
		2: 2.0,
	}

	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}

func Test_newLineupCalculator(t *testing.T) {
	t.SkipNow()
}
