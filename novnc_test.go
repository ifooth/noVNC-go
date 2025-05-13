package novnc

import (
	"testing"
)

func TestAssets(t *testing.T) {
	tests := []string{
		"vnc.html",
		"my_vendor/pako/README.md",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			f, err := assets.Open(tt)
			if err != nil {
				t.Fatalf("failed to open file %s: %v", tt, err)
			}

			body := make([]byte, 10)
			length, err := f.Read(body)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", tt, err)
			}

			if length != 10 {
				t.Errorf("expected length 10, got %d", length)
			}
		})
	}
}

func TestVendor(t *testing.T) {
	tests := []string{
		"vnc.html",
		"vendor/pako/README.md",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			f, err := FS().Open(tt)
			if err != nil {
				t.Fatalf("failed to open file %s: %v", tt, err)
			}

			body := make([]byte, 10)
			length, err := f.Read(body)
			if err != nil {
				t.Fatalf("failed to read file %s: %v", tt, err)
			}

			if length != 10 {
				t.Errorf("expected length 10, got %d", length)
			}
		})
	}
}
