package replay

import "testing"

func TestBrowserPolicyRejectsUnknownValue(t *testing.T) {
	if err := BrowserPolicy("unexpected").Validate(); err == nil {
		t.Fatal("unknown browser policy was accepted")
	}
	for _, policy := range []BrowserPolicy{"", BrowserPolicyAllowed, BrowserPolicyForbidden} {
		if err := policy.Validate(); err != nil {
			t.Fatalf("browser policy %q was rejected: %v", policy, err)
		}
	}
}
