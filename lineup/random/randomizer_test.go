package randomizer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ctx = context.Background()
)

// TestRandom_Randomizer is a test function for Random Generator
func TestRandom_Randomizer(t *testing.T) {
	sm := NewRandomizer()

	slice, err := sm.Randomizer(ctx, "hot")
	assert.Nil(t, err)
	result := slice

	expected := []map[string]interface{}{
		{"content_id": 6134, "content_type": "video", "service": "hot"},
		{"content_id": 6146, "content_type": "video", "service": "hot"},
		{"content_id": 6150, "content_type": "video", "service": "hot"},
	}
	assert.NotNil(t, result)
	assert.Equal(
		t,
		expected,
		result)
}
