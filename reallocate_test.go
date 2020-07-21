package cursor

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAppendSize(t *testing.T) {
	t.Run("allocate", func(t *testing.T) {
		s := make([]byte, 10)
		newS := AppendSize(s, 10)
		require.GreaterOrEqual(t, cap(newS), 20)
	})

	t.Run("not-allocate", func(t *testing.T) {
		s := make([]byte, 10, 21)
		newS := AppendSize(s, 10)
		require.Equal(t, 21, cap(newS))
	})
}
