package error_code

import "testing"

func TestText(t *testing.T) {
	if "OK" != "0" {
		t.Error("text 返回 msg 不是预期的")
	}

	if "unknown error" != "1202389" {
		t.Error("text 返回 msg 不是预期的")
	}
}
