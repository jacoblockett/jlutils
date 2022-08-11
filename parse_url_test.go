package jlutils

import "testing"

func TestParseURL(t *testing.T) {
	var tests = []struct {
		arga     string
		expected URL
	}{
		{"abc.com", URL{
			Protocol:  "",
			Hostname:  "abc.com",
			Subdomain: "",
			Domain:    "abc.com",
			Port:      0,
			Path:      "",
			Query:     "",
			Fragment:  "",
		}},
		{"abc.co.uk", URL{
			Protocol:  "",
			Hostname:  "abc.co.uk",
			Subdomain: "",
			Domain:    "abc.co.uk",
			Port:      0,
			Path:      "",
			Query:     "",
			Fragment:  "",
		}},
		{"abc.de", URL{
			Protocol:  "",
			Hostname:  "abc.de",
			Subdomain: "",
			Domain:    "abc.de",
			Port:      0,
			Path:      "",
			Query:     "",
			Fragment:  "",
		}},
		{"https://abc.com", URL{
			Protocol:  "https",
			Hostname:  "abc.com",
			Subdomain: "",
			Domain:    "abc.com",
			Port:      0,
			Path:      "",
			Query:     "",
			Fragment:  "",
		}},
		{"abc.com:5000", URL{
			Protocol:  "",
			Hostname:  "abc.com",
			Subdomain: "",
			Domain:    "abc.com",
			Port:      5000,
			Path:      "",
			Query:     "",
			Fragment:  "",
		}},
		{"test.abc.com", URL{
			Protocol:  "",
			Hostname:  "test.abc.com",
			Subdomain: "test",
			Domain:    "abc.com",
			Port:      0,
			Path:      "",
			Query:     "",
			Fragment:  "",
		}},
		{"abc.com/path", URL{
			Protocol:  "",
			Hostname:  "abc.com",
			Subdomain: "",
			Domain:    "abc.com",
			Port:      0,
			Path:      "path",
			Query:     "",
			Fragment:  "",
		}},
		{"abc.com/double/path", URL{
			Protocol:  "",
			Hostname:  "abc.com",
			Subdomain: "",
			Domain:    "abc.com",
			Port:      0,
			Path:      "double/path",
			Query:     "",
			Fragment:  "",
		}},
		{"abc.com?lang=go", URL{
			Protocol:  "",
			Hostname:  "abc.com",
			Subdomain: "",
			Domain:    "abc.com",
			Port:      0,
			Path:      "",
			Query:     "lang=go",
			Fragment:  "",
		}},
		{"abc.com?test1=1&test2=2", URL{
			Protocol:  "",
			Hostname:  "abc.com",
			Subdomain: "",
			Domain:    "abc.com",
			Port:      0,
			Path:      "",
			Query:     "test1=1&test2=2",
			Fragment:  "",
		}},
		{"abc.com#header", URL{
			Protocol:  "",
			Hostname:  "abc.com",
			Subdomain: "",
			Domain:    "abc.com",
			Port:      0,
			Path:      "",
			Query:     "",
			Fragment:  "header",
		}},
	}

	for _, test := range tests {
		output := ParseURL(test.arga)

		if output.Protocol != test.expected.Protocol {
			t.Errorf("Test failed (Protocol):\n1st arg: %s\nexpected: %s\nreceived: %s", test.arga, test.expected.Protocol, output.Protocol)
		}
		if output.Hostname != test.expected.Hostname {
			t.Errorf("Test failed (Hostname):\n1st arg: %s\nexpected: %s\nreceived: %s", test.arga, test.expected.Hostname, output.Hostname)
		}
		if output.Subdomain != test.expected.Subdomain {
			t.Errorf("Test failed (Subdomain):\n1st arg: %s\nexpected: %s\nreceived: %s", test.arga, test.expected.Subdomain, output.Subdomain)
		}
		if output.Domain != test.expected.Domain {
			t.Errorf("Test failed (Domain):\n1st arg: %s\nexpected: %s\nreceived: %s", test.arga, test.expected.Domain, output.Domain)
		}
		if output.Port != test.expected.Port {
			t.Errorf("Test failed (Port):\n1st arg: %s\nexpected: %d\nreceived: %d", test.arga, test.expected.Port, output.Port)
		}
		if output.Path != test.expected.Path {
			t.Errorf("Test failed (Path):\n1st arg: %s\nexpected: %s\nreceived: %s", test.arga, test.expected.Path, output.Path)
		}
		if output.Query != test.expected.Query {
			t.Errorf("Test failed (Query):\n1st arg: %s\nexpected: %s\nreceived: %s", test.arga, test.expected.Query, output.Query)
		}
		if output.Fragment != test.expected.Fragment {
			t.Errorf("Test failed (Fragment):\n1st arg: %s\nexpected: %s\nreceived: %s", test.arga, test.expected.Fragment, output.Fragment)
		}
	}
}
