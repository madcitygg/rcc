package main

import (
	"testing"
)

func TestParseAddressNoPort(t *testing.T) {
	a, err := ParseAddress("csgo.steelseries.io")
	if err != nil {
		t.Error("Did not expect an error, got", err)
	}
	if a.Host != "csgo.steelseries.io" {
		t.Error("Expected host csgo.steelseries.io, got", a.Host)
	}
	if a.Port != 27015 {
		t.Error("Expected port 27015, got", a.Port)
	}
}

func TestParseAddressWithPort(t *testing.T) {
	a, err := ParseAddress("10.10.10.20:2345")
	if err != nil {
		t.Error("Did not expect an error, got", err)
	}
	if a.Host != "10.10.10.20" {
		t.Error("Expected host 10.10.10.20, got", a.Host)
	}
	if a.Port != 2345 {
		t.Error("Expected port 2345, got", a.Port)
	}
}

func TestParseAddressEmptyPort(t *testing.T) {
	_, err := ParseAddress("10.10.10.20:")
	if err == nil {
		t.Error("Expected error")
	}
}

func TestParseAddressMultipleColon(t *testing.T) {
	_, err := ParseAddress("csgo.steelseries.io:1234:12")
	if err == nil {
		t.Error("Expected error")
	}
}
