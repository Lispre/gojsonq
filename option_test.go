package gojsonq

import (
	"strings"
	"testing"
)

func TestSetDebug(t *testing.T) {
	out := catchStdOut(func() {
		New(SetDebug(true)).log("gojsonq: debug mode on")
	})
	expected := "gojsonq: debug mode on"
	if !strings.Contains(out, expected) {
		t.Errorf("failed to enable SetDebug")
	}
}

func TestSetLogger(t *testing.T) {
	out := catchStdOut(func() {
		New(SetDebug(true), SetLogger(&cLogger{})).log("gojsonq: debug mode on")
	})
	expected := "gojsonq: debug mode on"
	if !strings.Contains(out, expected) ||
		!strings.Contains(out, "custom logger") {
		t.Errorf("failed to SetLogger")
	}
}

func TestSetLogger_with_nil_expecting_error(t *testing.T) {
	jq := New(SetLogger(nil))
	if jq.Error() == nil {
		t.Error("failed to catch nil SetLogger error")
	}
}

func TestSetDecoder(t *testing.T) {
	jq := New(SetDecoder(&cDecoder{}))
	if jq.option.decoder == nil {
		t.Error("failed to set decoder as option")
	}
}

func TestSetDecoder_with_nil_expecting_an_error(t *testing.T) {
	jq := New(SetDecoder(nil))
	if jq.Error() == nil {
		t.Error("failed to catch nil in SetDecoder")
	}
}
