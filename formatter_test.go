package main

import "testing"

func TestSimpleStringFormatOk(t *testing.T) {
	r := result{
		accessible: true,
		host: host{
			id:       "foo",
			hostname: "foo",
			port:     42,
		},
	}

	str := simpleStringFormatter(r)

	if str != "Ok - foo:42" {
		t.Error("Unexpected formatted string", str)
	}
}

func TestSimpleStringFormatOkDifferentHostname(t *testing.T) {
	r := result{
		accessible: true,
		host: host{
			id:       "foo",
			hostname: "foobar",
			port:     42,
		},
	}

	str := simpleStringFormatter(r)

	if str != "Ok - foo (foobar:42)" {
		t.Error("Unexpected formatted string", str)
	}
}

func TestSimpleStringFormatError(t *testing.T) {
	r := result{
		accessible: false,
		host: host{
			id:       "foo",
			hostname: "foo",
			port:     44,
		},
	}

	str := simpleStringFormatter(r)

	if str != "Error - foo:44" {
		t.Error("Unexpected formatted string", str)
	}
}
