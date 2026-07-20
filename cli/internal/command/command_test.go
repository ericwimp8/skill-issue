package command

import (
	"os"
	"testing"
	"time"
)

func menuKeyFromBytes(t *testing.T, input []byte) string {
	t.Helper()
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("create pipe: %v", err)
	}
	defer reader.Close()
	defer writer.Close()
	if _, err := writer.Write(input); err != nil {
		t.Fatalf("write pipe: %v", err)
	}
	type result struct {
		key string
		err error
	}
	results := make(chan result, 1)
	go func() {
		key, err := readMenuKey(reader)
		results <- result{key: key, err: err}
	}()
	select {
	case got := <-results:
		if got.err != nil {
			t.Fatalf("readMenuKey: %v", got.err)
		}
		return got.key
	case <-time.After(2 * time.Second):
		t.Fatal("readMenuKey did not return")
		return ""
	}
}

func TestReadMenuKeyParsesArrowAndSelection(t *testing.T) {
	cases := []struct {
		name  string
		input []byte
		want  string
	}{
		{name: "enter", input: []byte{'\r'}, want: "select"},
		{name: "ctrl-c", input: []byte{3}, want: "cancel"},
		{name: "arrow up", input: []byte{0x1b, '[', 'A'}, want: "up"},
		{name: "arrow down", input: []byte{0x1b, '[', 'B'}, want: "down"},
		{name: "application arrow up", input: []byte{0x1b, 'O', 'A'}, want: "up"},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := menuKeyFromBytes(t, testCase.input); got != testCase.want {
				t.Fatalf("readMenuKey(%v) = %q, want %q", testCase.input, got, testCase.want)
			}
		})
	}
}

func TestReadMenuKeyTreatsLoneEscapeAsCancel(t *testing.T) {
	// A bare ESC press sends a single byte; nothing else ever arrives, so the
	// reader must give up on the escape sequence instead of blocking.
	if got := menuKeyFromBytes(t, []byte{0x1b}); got != "cancel" {
		t.Fatalf("lone escape = %q, want cancel", got)
	}
}
