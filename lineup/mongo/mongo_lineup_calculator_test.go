package mongo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"top-ranking-worker/infra"
)

var (
	ctx = context.Background()
)

func TestCalculator_Calculate(t *testing.T) {
	m, err := infra.NewMongoDatabase(ctx)
	assert.Nil(t, err)

	cal := Calculator{Database: m}

	result, err := cal.Calculate(ctx, nil, nil)
	assert.Nil(t, err)

	expected := &map[int]float64{
		189321: 4,
	}

	assert.NotNil(t, result)
	assert.Equal(t, expected, result)
}
