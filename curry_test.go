package curry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurry(t *testing.T) {
	t.Run("A2R1", func(t *testing.T) {
		add := func(a, b int) int {
			return a + b
		}

		require.Equal(t, 3, add(1, 2))

		plus1 := A2R1(add)(1)
		require.Equal(t, 3, plus1(2))
		require.Equal(t, 4, plus1(3))
	})

	t.Run("A3R1", func(t *testing.T) {
		add := func(a, b, c int) int {
			return a + b + c
		}

		require.Equal(t, 6, add(1, 2, 3))

		plus1 := A3R1(add)(1)
		require.Equal(t, 6, plus1(2)(3))

		plus1a2 := A3R1(add)(1)(2)
		require.Equal(t, 6, plus1a2(3))
	})

	t.Run("A4R2", func(t *testing.T) {
		var err error
		var r int

		add := func(a, b, c, d int) (int, error) {
			return a + b + c + d, nil
		}

		r, err = add(1, 2, 3, 4)
		require.NoError(t, err)
		require.Equal(t, 10, r)

		plus1 := A4R2(add)(1)

		r, err = plus1(2)(3)(4)
		require.NoError(t, err)
		require.Equal(t, 10, r)

		plus1a2 := A4R2(add)(1)(2)

		r, err = plus1a2(3)(4)
		require.NoError(t, err)
		require.Equal(t, 10, r)
	})
}
