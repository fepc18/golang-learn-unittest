//https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func TestFooerTableDriven(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		// the table itself
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

//Errors and Logs

func TestFooer2(t *testing.T) {
	input := 3
	result := Fooer(3)
	t.Logf("The input was %d", input)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
	// t.Error("This will stop the test and report a failure")
	//t.Fatalf("Stop the test now, we have seen enough") // This will stop the test and report a failure
	t.Error("This won't be executed")
}

func TestFooerParallel(t *testing.T) {
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(3)
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	})
	t.Run("Test 7 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(7)
		if result != "7" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "7")
		}
	})
}

func TestFooerSkiped(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

//go test -v -test.short

// Test Teardown and Cleanup
func Test_With_Cleanup(t *testing.T) {

	// Some test code here

	t.Cleanup(func() {
		// cleanup logic
	})

	// more test code here

}
func helper(t *testing.T) {
	t.Helper()
	// do something
}

func TestFooerTempDir(t *testing.T) {
	//tmpDir := t.TempDir() // create a temp dir

	// your tests
}

func BenchmarkFooer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fooer(i)
	}
}

func FuzzFooer(f *testing.F) {
	f.Add(3)
	f.Fuzz(func(t *testing.T, a int) {
		Fooer(a)
	})
}

func TestFooerWithTestify(t *testing.T) {

	// assert equality
	assert.Equal(t, "Foo", Fooer(0), "0 is divisible by 3, should return Foo")

	// assert inequality
	assert.NotEqual(t, "Foo", Fooer(1), "1 is not divisible by 3, should not return Foo")
}
