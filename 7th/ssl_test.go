package main

import "testing"

func TestSupportSslLWithAbaBeforeBab(t *testing.T) {
	input := ip("xax[axa]xsd")
	if !input.supportsSSL() {
		t.Fatalf("%s was expected to support SSL but did not!", input)
	}
}

func TestDoesNotSupportSsl(t *testing.T) {
	input := ip("xbx[axa]xsd")
	if input.supportsSSL() {
		t.Fatalf("%s was NOT expected to support SSL but did not!", input)
	}

	input = ip("xax[bxa]xsd")
	if input.supportsSSL() {
		t.Fatalf("%s was NOT expected to support SSL but did not!", input)
	}
}

func TestSupportSslLWithAbaAfterBab(t *testing.T) {
	input := ip("xbx[axa]xax")
	if !input.supportsSSL() {
		t.Fatalf("%s was expected to support SSL but did not!", input)
	}
}
