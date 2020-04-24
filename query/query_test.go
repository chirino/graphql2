package query

import (
	"testing"

	"github.com/chirino/graphql/internal/introspection"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoLossOnQueryFormating(t *testing.T) {
	q1, err := Parse(introspection.Query)
	require.NoError(t, err)

	q2, err := Parse(q1.String())
	require.NoError(t, err)

	assert.Equal(t, q1.String(), q2.String())
}
