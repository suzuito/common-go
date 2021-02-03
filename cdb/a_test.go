package cdb

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPMySQLURLString(t *testing.T) {
	testCases := []struct {
		desc     string
		input    url.URL
		expected string
	}{
		{
			desc: "",
			input: url.URL{
				Scheme:   "mysql",
				User:     url.UserPassword("suzuito", "aiueo"),
				Host:     "example.com",
				Path:     "database1",
				RawQuery: "k1=v1&k2=v2",
			},
			expected: "suzuito:aiueo@tcp(example.com)/database1?k1=v1&k2=v2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			real := NewTCPMySQLURLString(&tC.input)
			assert.Equal(t, tC.expected, real)
		})
	}
}

func TestNewLocalhostMySQLURLString(t *testing.T) {
	testCases := []struct {
		desc     string
		input    url.URL
		expected string
	}{
		{
			desc: "",
			input: url.URL{
				Scheme:   "mysql",
				User:     url.UserPassword("suzuito", "aiueo"),
				Host:     "example.com",
				Path:     "database1",
				RawQuery: "k1=v1&k2=v2",
			},
			expected: "suzuito:aiueo@tcp/database1?k1=v1&k2=v2",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			real := NewLocalhostMySQLURLString(&tC.input)
			assert.Equal(t, tC.expected, real)
		})
	}
}
