package integers

import "testing"

import "fmt"

func TestAdd(t *testing.T) {
	t.Run("2 + 2 = 4 check", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected %d, got %d", expected, sum)
		}
	})

	t.Run("Zero check", func(t *testing.T) {
		sum := Add(0, 0)
		expected := 0

		if sum != expected {
			t.Errorf("expected %d, got %d", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
