package log

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func testNew(t *testing.T) {
	logger := New()
	require.NotNil(t, logger)
}
