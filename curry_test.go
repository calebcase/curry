package curry_test

import (
	"fmt"
	"testing"

	"github.com/calebcase/curry"
	"github.com/stretchr/testify/require"
)

func ExampleA2R1() {
	add := func(a, b int) int {
		return a + b
	}

	cadd := curry.A2R1(add)

	plus1 := cadd(1)
	plus5 := cadd(5)

	fmt.Println(plus1(2))
	fmt.Println(plus5(10))

	// Output:
	// 3
	// 15
}

func TestCurry(t *testing.T) {
	t.Run("A2R1", func(t *testing.T) {
		add := func(a, b int) int {
			return a + b
		}

		require.Equal(t, 3, add(1, 2))

		plus1 := curry.A2R1(add)(1)
		require.Equal(t, 3, plus1(2))
		require.Equal(t, 4, plus1(3))
	})

	t.Run("A3R1", func(t *testing.T) {
		add := func(a, b, c int) int {
			return a + b + c
		}

		require.Equal(t, 6, add(1, 2, 3))

		plus1 := curry.A3R1(add)(1)
		require.Equal(t, 6, plus1(2)(3))

		plus1a2 := curry.A3R1(add)(1)(2)
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

		plus1 := curry.A4R2(add)(1)

		r, err = plus1(2)(3)(4)
		require.NoError(t, err)
		require.Equal(t, 10, r)

		plus1a2 := curry.A4R2(add)(1)(2)

		r, err = plus1a2(3)(4)
		require.NoError(t, err)
		require.Equal(t, 10, r)
	})
}
