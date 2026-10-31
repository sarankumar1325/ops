package onprem

import "testing"

func TestFormatOctet(t *testing.T) {
	tests := []struct {
		mac  string
		want string
	}{
		{"d6:3f:9b:0f:0c:c8", "d6:3f:9b:f:c:c8"},
		{"52:54:00:12:34:56", "52:54:0:12:34:56"},
		{"00:00:00:00:00:01", "0:0:0:0:0:1"},
	}
	for _, tt := range tests {
		got, err := formatOctet(tt.mac)
		if err != nil {
			t.Fatalf("formatOctet(%q) returned error: %v", tt.mac, err)
		}
		if got != tt.want {
			t.Errorf("formatOctet(%q) = %q, want %q", tt.mac, got, tt.want)
		}
	}

	if _, err := formatOctet(""); err == nil {
		t.Error("formatOctet(\"\") should return an error")
	}
}
