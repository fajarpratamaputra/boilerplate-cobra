package lineup

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"top-ranking-worker/lineup/domain"
)

var (
	contents = []domain.Content{
		{1, "title1"},
		{2, "title2"},
		{3, "title3"},
		{4, "title4"},
	}

	userInteractions = []domain.Interaction{
		{1, 1, "comment", "hot"},
		{2, 2, "love", "hot"},
		{1, 2, "view", "news"},
		{1, 3, "like", "news"},
		{1, 3, "share", "hot"},
	}
)

func Test_lineupCalculator_calculate(t *testing.T) {
	lc := newCalculator()
	result := lc.Calculate(contents, userInteractions)

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
