package novnc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssets(t *testing.T) {
	tests := []string{
		"vnc.html",
		"my_vendor/pako/README.md",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			f, err := assets.Open(tt)
			require.NoError(t, err)

			body := make([]byte, 10)
			length, err := f.Read(body)
			assert.NoError(t, err)
			assert.Equal(t, length, 10)
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
			require.NoError(t, err)

			body := make([]byte, 10)
			length, err := f.Read(body)
			assert.NoError(t, err)
			assert.Equal(t, length, 10)
		})
	}
}
