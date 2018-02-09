package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstConfig(t *testing.T) {
	cfg, err := Config()

	assert.NoError(t, err)
	assert.Equal(t, FirstStrategy, cfg.Strategy, "should be equal")
}
