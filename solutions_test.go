package golangtest

import (
	"reflect"
	"testing"
)

func TestMultipleValues(t *testing.T) {
	a, b := MultipleValues()
	c, d := 3, 5
	if a != c && b != d {
		t.Errorf("got %d, %d want %d, %d", a, b, c, d)
	}
}

func TestMarksMap(t *testing.T) {
	got := MarksMap(100, 95)
	want := make(map[string]int)
	want["Maths"] = 100
	want["English"] = 95
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", reflect.TypeOf(got).String(), reflect.TypeOf(want).String())
	}
}

func TestIntSeq(t *testing.T) {
	NextInt := IntSeq()
	got := NextInt()
	want := 1
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPointerToPointer(t *testing.T) {
	got := reflect.TypeOf(PointerToPointer())
	var pt **int
	want := reflect.TypeOf(pt)
	if got != want {
		t.Errorf("got %q, want %q", got.String(), want.String())
	}
}

func TestTypeCast(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		got := IntToString(12)
		want := "12"
		if got != want {
			t.Errorf("got %T, want %T", got, want)
		}
	})

	t.Run("string to int", func(t *testing.T) {
		got, _ := StringToInt("19")
		want := 19
		if got != want {
			t.Errorf("got %T, want %T", got, want)
		}
	})

	t.Run("float to int", func(t *testing.T) {
		got := FloatToInt(7.4)
		want := 7
		if got != want {
			t.Errorf("got %T, want %T", got, want)
		}
	})

	t.Run("int to float", func(t *testing.T) {
		got := IntToFloat(7)
		want := 7.0
		if got != want {
			t.Errorf("got %T, want %T", got, want)
		}
	})
}

func TestIncrementOne(t *testing.T) {
	got := 10
	IncrementOne(&got)
	want := 11
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
