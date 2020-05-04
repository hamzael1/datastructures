package lnklist

import "testing"

func TestCount(t *testing.T) {
	l := New("1", "2", "3", "4", "5", "2", "4", "2");
	if l.Count("2") != 3 {
		t.Error("Error Counting times element appears in list");
	}
}

func TestGetNth(t *testing.T) {
	l := New("1", "2", "3", "4", "5", "2", "4", "2");
	ret, found := l.GetNth(4);
	if ret != "5" || found != true {
		t.Error("Error Exercise GetNth")
	}
}