package cmd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInfix2Postfix(t *testing.T) {
	t.Run("a+b", func(t *testing.T) {
		input := "a+b"
		want := "ab+"
		got := Infix2Postfix(input)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("a+b*c", func(t *testing.T) {
		input := "a+b*c"
		want := "abc*+"
		got := Infix2Postfix(input)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("(a+b)*c+d", func(t *testing.T) {
		input := "(a+b)*c+d"
		want := "ab+c*d+"
		got := Infix2Postfix(input)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("a+b*c*(d+e)", func(t *testing.T) {
		input := "a+b*c*(d+e)"
		want := "abc*de+*+"
		got := Infix2Postfix(input)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("((a+(b*c))+d)", func(t *testing.T) {
		input := "((a+(b*c))+d)"
		want := "abc*+d+"
		got := Infix2Postfix(input)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	})
}
