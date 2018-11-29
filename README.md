# Gowhere

[![Go Report Card](https://goreportcard.com/badge/ae0000/gowhere)](https://goreportcard.com/report/ae0000/gowhere)
[![Build Status](https://travis-ci.org/ae0000/gowhere.svg?branch=master)](https://travis-ci.org/ae0000/gowhere)
[![cover.run](https://cover.run/go/github.com/ae0000/gowhere.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fae0000%2Fgowhere)
[![GoDoc](https://godoc.org/github.com/ae0000/gowhere?status.svg)](https://godoc.org/github.com/ae0000/gowhere)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/ae0000/gowhere/blob/master/LICENSE)

Gowhere provides an easy way to get country data. See below for usage. Each country consists of the following struct:

```
Country{
    Code:           "AU", // iso2 code
    Currency:       "AUD",
    CurrencySymbol: "$",
    Phone:          "+61",
    Name:           "Australia",
},
```

## Install

```
go get -u github.com/ae0000/gowhere
```

## Usage

#### Get a list of countries (to iterate over, etc.)

```
for _, c := range gowhere.Countries {
    fmt.Println(c.Name, c.Code, c.Currency)
}
```

#### Get a countries phone, name, currency symbol or currency

```
countryCode := "AU"
fmt.Printf("In %s where the phone ext is %s, the price is: %s10.00 (%s)",
    gowhere.Name(countryCode),
    gowhere.Phone(countryCode),
    gowhere.CurrencySymbol(countryCode),
    gowhere.Currency(countryCode))

// In Australia where the phone ext is +61, the price is: $10.00 (AUD)
```

#### Get a country

```
c := gowhere.GetCountry(users.CountryCode)
fmt.Println(c.Name)
fmt.Println(c.Phone)
fmt.Println(c.Currency)
```

#### Get a list of names

```
names := gowhere.CountryNames()
for _, c := range names{
    fmt.Println(c)
}
```

#### Get a list of codes

```
codes := gowhere.CountryCodes()
for _, c := range codes{
    fmt.Println(c)
}
```
