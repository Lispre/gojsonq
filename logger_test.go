package gojsonq

import (
	"strings"
	"testing"
)

func Test_DefaultLogger(t *testing.T) {
	out := catchStdOut(func() {
		New(SetDebug(true), SetLogger(&DefaultLogger{})).log("hello to default logger")
	})
	expected := "gojsonq: Debug mode on"
	if !strings.Contains(out, expected) ||
		!strings.Contains(out, "gojsonq: hello to default logger") {
		t.Errorf("default logger failed to write stdout")
	}
}
