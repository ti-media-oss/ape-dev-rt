package ui

import (
	"bytes"
	"testing"
)

func TestStreamedUi_Error(t *testing.T) {
	outErr := new(bytes.Buffer)
	ui := &StreamedUi{
		ErrorWriter: outErr,
	}
	ui.Error("HELLO")
	bufErr := ui.ErrorBuffer
	if bufErr.String() != "HELLO\n" {
		t.Fatalf("bad: %s", bufErr.String())
	}
	if outErr.String() != "HELLO\n" {
		t.Fatalf("bad: %s", outErr.String())
	}
}

func TestStreamedUi_Output(t *testing.T) {
	outWriter := new(bytes.Buffer)
	ui := &StreamedUi{
		OutputWriter: outWriter,
	}
	ui.Output("HELLO")
	bufOut := ui.OutputBuffer
	if bufOut.String() != "HELLO\n" {
		t.Fatalf("bad: %s", bufOut.String())
	}
	if outWriter.String() != "HELLO\n" {
		t.Fatalf("bad: %s", outWriter.String())
	}
}
