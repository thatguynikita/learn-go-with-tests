package iteration

import (
	"testing"
	"fmt"
	"strings"
)

func TestRepeat(t *testing.T) {
	t.Run("Run test letter 'a' ten times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
	
		if repeated != expected {
			t.Errorf("got %q but expected %q", repeated, expected)
		}
	})

	t.Run("Compare test with strings.Repeat()", func(t *testing.T) {
		repeated := Repeat("asd", 10)
		expected := strings.Repeat("asd", 10)
	
		if repeated != expected {
			t.Errorf("got %q but expected %q", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("qwe", 3)
	fmt.Println(result)
	//Output: qweqweqwe
}
