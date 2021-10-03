package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTheseArgs(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		expectedPaths []string
		expectedErr   string
	}{
		{
			name:          "path not found",
			input:         []string{"progname", "missing-file.txt"},
			expectedPaths: nil,
			expectedErr:   "missing-file.txt: path not found\n",
		},
		{
			name:          "path is dir",
			input:         []string{"progname", "."},
			expectedPaths: nil,
			expectedErr:   ".: is a directory\n",
		},
		{
			name:          "found files",
			input:         []string{"progname", "main.go", "main_test.go"},
			expectedPaths: []string{"main.go", "main_test.go"},
			expectedErr:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotPaths, gotErr := parseTheseArgs(test.input)
			assert.Equal(t, test.expectedPaths, gotPaths, "paths")
			if test.expectedErr == "" {
				assert.NoError(t, gotErr)
			} else {
				assert.Equal(t, test.expectedErr, gotErr.Error(), "err")
			}
		})
	}
}

func TestWriteAllMMDBNetworks(t *testing.T) {
	testMMDB := "../../test/GeoIP2-Enterprise-Test.mmdb"

	var buf []byte
	b := bytes.NewBuffer(buf)

	expected := `::27d:a0d8/125
::432b:9c00/120
::4ad1:1000/116
::5102:45a0/123
::59a0:1470/124
::af10:c700/120
::cac4:e000/116
::d42f:eb51/128
::d42f:eb52/128
::d8a0:5338/125
2.125.160.216/29
67.43.156.0/24
74.209.16.0/20
81.2.69.160/27
89.160.20.112/28
175.16.199.0/24
202.196.224.0/20
212.47.235.81/32
212.47.235.82/32
216.160.83.56/29
2001:0:27d:a0d8::/61
2001:0:432b:9c00::/56
2001:0:4ad1:1000::/52
2001:0:5102:45a0::/59
2001:0:59a0:1470::/60
2001:0:af10:c700::/56
2001:0:cac4:e000::/52
2001:0:d42f:eb51::/64
2001:0:d42f:eb52::/64
2001:0:d8a0:5338::/61
2002:27d:a0d8::/45
2002:432b:9c00::/40
2002:4ad1:1000::/36
2002:5102:45a0::/43
2002:59a0:1470::/44
2002:af10:c700::/40
2002:cac4:e000::/36
2002:d42f:eb51::/48
2002:d42f:eb52::/48
2002:d8a0:5338::/45
`

	err := writeAllMMDBNetworks(testMMDB, b)
	require.NoError(t, err)
	assert.Equal(
		t,
		expected,
		b.String(),
		"expected output for test DB",
	)
}
