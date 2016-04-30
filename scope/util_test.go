package scope

import "testing"

func TestInclude(t *testing.T) {

	origin := "openid email custom"
	if !Include(origin, OpenID) {
		t.Error("'openid' should be found")
	}

	if !IncludeOpenID(origin) {
		t.Error("'openid' should be found")
	}

	if !Include(origin, Email) {
		t.Error("'email' should be found")
	}

	if Include(origin, "unknown") {
		t.Error("'unknown' should be found")
	}

	origin = "custom offline_access"
	if Include(origin, OpenID) {
		t.Error("'openid' shouldn't be found")
	}

	if !IncludeOfflineAccess(origin) {
		t.Error("'offline_access' should be found")
	}
}