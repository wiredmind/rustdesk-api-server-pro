package admin

import "testing"

// TestResolveTfaSecret locks down which secret an admin login code is
// checked against: the already-enrolled secret when one exists, or the
// pending enrollment secret otherwise - never both, never neither.
func TestResolveTfaSecret(t *testing.T) {
	secret, enrolling := resolveTfaSecret("USERSECRET", "PENDINGSECRET")
	if secret != "USERSECRET" || enrolling {
		t.Errorf("resolveTfaSecret(enrolled) = (%q, %v), want (USERSECRET, false)", secret, enrolling)
	}

	secret, enrolling = resolveTfaSecret("", "PENDINGSECRET")
	if secret != "PENDINGSECRET" || !enrolling {
		t.Errorf("resolveTfaSecret(unenrolled) = (%q, %v), want (PENDINGSECRET, true)", secret, enrolling)
	}

	secret, enrolling = resolveTfaSecret("", "")
	if secret != "" || !enrolling {
		t.Errorf("resolveTfaSecret(nothing pending) = (%q, %v), want (\"\", true)", secret, enrolling)
	}
}
