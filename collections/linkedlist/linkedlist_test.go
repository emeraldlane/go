package collections

import "testing"

func TestEmptyString(t *testing.T) {
	l := new(LinkedList)
	s := l.String()
	if len(s) > 0 {
		t.Error("Should be empty")
	}
}

func TestString(t *testing.T) {
	l := new(LinkedList)
	first := new(Element)
	first.Value = "test"
	first.Next = nil
	l.root = first
	s := l.String()
	if s != "test" {
		t.Errorf("expected 'test', got '%s'", s)
	}
}