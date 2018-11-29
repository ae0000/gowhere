package gowhere

import (
	"fmt"
	"sort"
)

// Country holds the counties details
type Country struct {
	Code           string // AU
	Currency       string // AUD
	CurrencySymbol string // $
	Name           string // Australia
	Phone          string // +61
}

// GetCountry returns a country based on the code
func GetCountry(code string) (Country, error) {
	for _, c := range Countries {
		if c.Code == code {
			return c, nil
		}
	}

	return Country{}, fmt.Errorf("could not find country: %s", code)
}

// CountryNames returns a list of country names, sorted by alpha
func CountryNames() []string {
	s := []string{}

	for _, c := range Countries {
		s = append(s, c.Name)
	}

	sort.Strings(s)

	return s
}

// CountryCodes returns a list of country codes, sorted by alpha
func CountryCodes() []string {
	s := []string{}

	for _, c := range Countries {
		s = append(s, c.Code)
	}

	sort.Strings(s)

	return s
}

// Currency returns a countries currency. Note that it defaults to an empty
// string if country not found
func Currency(code string) string {
	c, err := GetCountry(code)
	if err != nil {
		return ""
	}
	return c.Currency
}

// CurrencySymbol returns a countries currency symbol. Note that it defaults to
// an empty string if country not found
func CurrencySymbol(code string) string {
	c, err := GetCountry(code)
	if err != nil {
		return ""
	}
	return c.CurrencySymbol
}

// Phone returns a countries phone ext. Note that it defaults to an empty string
// if country not found
func Phone(code string) string {
	c, err := GetCountry(code)
	if err != nil {
		return ""
	}
	return c.Phone
}

// Name returns a countries name. Note that it defaults to an empty string if
// country not found
func Name(code string) string {
	c, err := GetCountry(code)
	if err != nil {
		return ""
	}
	return c.Name
}
