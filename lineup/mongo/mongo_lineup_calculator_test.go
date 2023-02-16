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

	cal := Calculator{Database: m}

	result, err := cal.Calculate(ctx, nil, nil)
	assert.Nil(t, err)

	expected := domain.LineupMap{
		123: &domain.LineupContent{
			ContentId:   0,
			Service:     "hot",
			ContentType: "video",
			Score:       3,
		},
		456: &domain.LineupContent{
			ContentId:   0,
			Service:     "hot",
			ContentType: "video",
			Score:       2,
		},
		789: &domain.LineupContent{
			ContentId:   0,
			Service:     "hot",
			ContentType: "video",
			Score:       2,
		},
		1999: &domain.LineupContent{
			ContentId:   0,
			Service:     "hot",
			ContentType: "video",
			Score:       1,
		},
	}

	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}
