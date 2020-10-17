package helpers

import "testing"

func TestStringInSlices(t *testing.T) {
	val := StringInSlice("test", []string{"test", "test2", "3"})

	if !val {
		t.Errorf("Boolean should be true")
	}

	val = StringInSlice("test", []string{"test2", "1"})

	if val {
		t.Errorf("Boolean should be false")
	}
}
