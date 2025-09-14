package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		name      string
		keyString string
		wantErr   bool
		wantKey   string
	}{
		{
			name:      "Valid key",
			keyString: "ApiKey THISISANAPIKEY",
			wantErr:   false,
			wantKey:   "THISISANAPIKEY",
		},
		{
			name:      "Valid token w/ no ApiKey",
			keyString: "Key THISISAKEY",
			wantErr:   false,
			wantKey:   "",
		},
		{
			name:      "Invalid key",
			keyString: "ApiKey ",
			wantErr:   false,
			wantKey:   "",
		},
		{
			name:      "Invalid key w/ spaces",
			keyString: " ApiKey  ",
			wantErr:   true,
			wantKey:   "",
		},
		{
			name:      "Empty key",
			keyString: "",
			wantErr:   true,
			wantKey:   "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req, err := http.NewRequest("", "", http.NoBody)
			if err != nil {
				t.Fatalf("failed to create test request: %v", err)
			}

			req.Header.Set("Authorization", c.keyString)

			keyStr, err := GetAPIKey(req.Header)
			if (err != nil) != c.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, c.wantErr)
				return
			}

			if keyStr != c.wantKey {
				t.Errorf("GetAPIKey() = %v, want %v", keyStr, c.wantKey)
			}
		})
	}
}
