package gowhere

import (
	"testing"
)

func TestName(t *testing.T) {
	if Name("AU") != "Australia" {
		t.Error("Name is not working correctly... ")
	}
	if Name("invalid") != "" {
		t.Error("Name is not working correctly... ")
	}
}

func TestPhone(t *testing.T) {
	if Phone("AU") != "+61" {
		t.Error("Phone is not working correctly... ")
	}
	if Phone("invalid") != "" {
		t.Error("Phone is not working correctly... ")
	}
}

func TestCurrency(t *testing.T) {
	if Currency("AU") != "AUD" {
		t.Error("Currency is not working correctly... ")
	}
	if Currency("invalid") != "" {
		t.Error("Currency is not working correctly... ")
	}
}

func TestCurrencySymbol(t *testing.T) {
	if CurrencySymbol("AU") != "$" {
		t.Error("CurrencySymbol is not working correctly... ")
	}
	if CurrencySymbol("invalid") != "" {
		t.Error("CurrencySymbol is not working correctly... ")
	}
}

func TestCountryCodes(t *testing.T) {
	codes := CountryCodes()
	if len(codes) != len(Countries) || len(codes) == 0 {
		t.Error("country codes is not working")
	}

	if codes[0] != "AD" {
		t.Error("country codes is not working", codes[0])
	}
}
func TestCountryNames(t *testing.T) {
	names := CountryNames()
	if len(names) != len(Countries) || len(names) == 0 {
		t.Error("country names is not working")
	}

	if names[0] != "Afghanistan" {
		t.Error("country codes is not working: ", names[0])
	}
}
