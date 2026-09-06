package admin

import "testing"

// TestBuildConnectURL locks down the RustDesk uni-link scheme
// (rustdesk://<authority>/<id>) as sourced from the RustDesk client itself
// (core_main.rs core_main_invoke_new_connection: --connect, --play,
// --file-transfer map to authorities "connect", "play", "file-transfer";
// "terminal" is the maintainer-confirmed authority for the client's
// built-in terminal). There is no shell/command parameter in this scheme -
// this test guards against ever reintroducing one.
func TestBuildConnectURL(t *testing.T) {
	cases := []struct {
		action     string
		id         string
		wantAction string
		wantURL    string
	}{
		{"remote", "123456789", "remote", "rustdesk://connect/123456789"},
		{"file", "123456789", "file", "rustdesk://file-transfer/123456789"},
		{"mirror", "123456789", "mirror", "rustdesk://play/123456789"},
		{"terminal", "123456789", "terminal", "rustdesk://terminal/123456789"},
		// Unknown/empty actions must fall back to a safe full-control connect,
		// never pass through attacker-controlled text into the URL scheme.
		{"", "123456789", "remote", "rustdesk://connect/123456789"},
		{"shell", "123456789", "remote", "rustdesk://connect/123456789"},
	}

	for _, tc := range cases {
		gotAction, gotURL := buildConnectURL(tc.action, tc.id)
		if gotAction != tc.wantAction {
			t.Errorf("buildConnectURL(%q, %q) action = %q, want %q", tc.action, tc.id, gotAction, tc.wantAction)
		}
		if gotURL != tc.wantURL {
			t.Errorf("buildConnectURL(%q, %q) url = %q, want %q", tc.action, tc.id, gotURL, tc.wantURL)
		}
	}
}

// TestValidateDeviceIds guards against a malformed/empty delete request
// resolving to an unscoped bulk delete.
func TestValidateDeviceIds(t *testing.T) {
	if err := validateDeviceIds(nil); err == nil {
		t.Error("validateDeviceIds(nil) = nil, want error")
	}
	if err := validateDeviceIds([]int{}); err == nil {
		t.Error("validateDeviceIds([]int{}) = nil, want error")
	}
	if err := validateDeviceIds([]int{1, 2, 3}); err != nil {
		t.Errorf("validateDeviceIds([1,2,3]) = %v, want nil", err)
	}
}
