package evaluation

import "testing"

func TestIsSignalCommandRequiresStandaloneShellWords(t *testing.T) {
	cliPath := "/opt/skill-issue"
	token := "abc123"
	stateRoot := "/tmp/state/.skill-issue"
	cases := []struct {
		name    string
		command string
		want    bool
	}{
		{name: "plain invocation", command: "/opt/skill-issue signal abc123 /tmp/state/.skill-issue", want: true},
		{name: "quoted invocation", command: `"/opt/skill-issue" signal "abc123" "/tmp/state/.skill-issue"`, want: true},
		{name: "chained invocation", command: "cd /tmp && /opt/skill-issue signal abc123 /tmp/state/.skill-issue; echo done", want: true},
		{name: "token embedded in larger word", command: "/opt/skill-issue signal xabc123x /tmp/state/.skill-issue", want: false},
		{name: "mentioned in echoed text", command: `echo "run /opt/skill-issue-later and signal abc123def"`, want: false},
		{name: "missing signal word", command: "/opt/skill-issue mark abc123 /tmp/state/.skill-issue", want: false},
		{name: "missing state root", command: "/opt/skill-issue signal abc123", want: false},
		{name: "cli path extended", command: "/opt/skill-issue2 signal abc123 /tmp/state/.skill-issue", want: false},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := isSignalCommand(testCase.command, cliPath, token, stateRoot); got != testCase.want {
				t.Fatalf("isSignalCommand(%q) = %v, want %v", testCase.command, got, testCase.want)
			}
		})
	}
}

func TestIsCodexSignalCommandAcceptsCapturedEscapedQuotes(t *testing.T) {
	token := "abc123"
	cases := []struct {
		name    string
		command string
		want    bool
	}{
		{name: "plain echo", command: `echo "abc123"`, want: true},
		{name: "captured zsh command", command: `/bin/zsh -lc "echo \"abc123\""`, want: true},
		{name: "token embedded in larger word", command: `/bin/zsh -lc "echo \"xabc123\""`, want: false},
		{name: "different command", command: `/bin/zsh -lc "printf \"abc123\""`, want: false},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := isCodexSignalCommand(testCase.command, token); got != testCase.want {
				t.Fatalf("isCodexSignalCommand(%q) = %v, want %v", testCase.command, got, testCase.want)
			}
		})
	}
}
