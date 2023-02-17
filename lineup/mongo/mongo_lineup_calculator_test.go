package mongo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
)

var (
	ctx = context.Background()
)

// TestCalculator_Calculate is a test function for Calculate
func TestCalculator_Calculate(t *testing.T) {
	m, err := infra.NewMongoDatabase(ctx)
	assert.Nil(t, err)

	cal := NewCalculator(m)

	result, err := cal.Calculate(ctx, "views", nil)
	assert.Nil(t, err)

	expected := &domain.Lineup{
		123: &domain.Content{
			ContentId:   0,
			Service:     "hot",
			ContentType: "video",
			Score:       1,
		},
		1999: &domain.Content{
			ContentId:   0,
			Service:     "hot",
			ContentType: "video",
			Score:       1,
		},
	}

	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}
