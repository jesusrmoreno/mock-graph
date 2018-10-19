package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/icrowley/fake"
)

// Values object that keeps track of stuff
type Values struct {
	Count   int           `json:"count"`
	Options []interface{} `json:"options"`
}

// Property is a custom type that allows us to support
// unmarshaling properties as either an Property
// either Values or Value will be populated so consumers should check for
// the length of Values before proceeding with Value
type Property struct {
	Values Values `json:"values"`
	Value  string `json:"value"`
}

// UnmarshalJSON fulfills the interface so we can use json.Unmarshal
func (p *Property) UnmarshalJSON(b []byte) error {
	if b[0] == '{' {
		return json.Unmarshal(b, &p.Values)
	}
	return json.Unmarshal(b, &p.Value)
}

func shuffleValues(vals []interface{}) []interface{} {
	v := []interface{}{}
	rand.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})
	return append(v, vals...)
}

func sampleValues(v []interface{}, count int) []interface{} {
	sv := shuffleValues(v)
	vv := []interface{}{}
	for i := 0; i < count && i < len(v); i++ {
		vv = append(vv, sv[i])
	}
	return vv
}

// Populate will return either a value from the array of options if the property
// is an array, a mock value if the property string matches one of the supported
// generators, or the property string if it is not supported.
func (p *Property) Populate() interface{} {
	if p.Values.Count > 0 && len(p.Values.Options) > 0 {
		return sampleValues(p.Values.Options, p.Values.Count)
	}
	casedProp := strings.Title(p.Value)
	switch casedProp {
	case "Name":
		v := []string{fake.MaleFullName(), fake.FemaleFullName()}
		return v[random(0, 2)]
	case "Brand":
		return fake.Brand()
	case "Character":
		return fake.Character()
	case "Characters":
		return fake.Characters()
	case "City":
		return fake.City()
	case "Color":
		return fake.Color()
	case "Company":
		return fake.Company()
	case "Continent":
		return fake.Continent()
	case "Country":
		return fake.Country()
	case "CreditCardNum":
		return fmt.Sprintf("0000-%d-%d-%d", random(1000, 10000), random(1000, 10000), random(1000, 10000))
	case "CreditCardType":
		return fake.CreditCardType()
	case "Currency":
		return fake.Currency()
	case "CurrencyCode":
		return fake.CurrencyCode()
	case "Day":
		return fake.Day()
	case "Digits":
		return fake.Digits()
	case "DomainName":
		return fake.DomainName()
	case "DomainZone":
		return fake.DomainZone()
	case "EmailAddress":
		return fake.EmailAddress()
	case "EmailBody":
		return fake.EmailBody()
	case "EmailSubject":
		return fake.EmailSubject()
	case "FemaleFirstName":
		return fake.FemaleFirstName()
	case "FemaleFullName":
		return fake.FemaleFullName()
	case "FemaleFullNameWithPrefix":
		return fake.FemaleFullNameWithPrefix()
	case "FemaleFullNameWithSuffix":
		return fake.FemaleFullNameWithSuffix()
	case "FemaleLastName":
		return fake.FemaleLastName()
	case "FemalePatronymic":
		return fake.FemalePatronymic()
	case "FirstName":
		return fake.FirstName()
	case "FullName":
		return fake.FullName()
	case "FullNameWithPrefix":
		return fake.FullNameWithPrefix()
	case "FullNameWithSuffix":
		return fake.FullNameWithSuffix()
	case "Gender":
		return fake.Gender()
	case "GenderAbbrev":
		return fake.GenderAbbrev()
	case "GetLangs":
		return fake.GetLangs()
	case "HexColor":
		return fake.HexColor()
	case "HexColorShort":
		return fake.HexColorShort()
	case "IPv4":
		return fake.IPv4()
	case "IPv6":
		return fake.IPv6()
	case "Industry":
		return fake.Industry()
	case "JobTitle":
		return fake.JobTitle()
	case "Language":
		return fake.Language()
	case "LastName":
		return fake.LastName()
	case "Latitude":
		return fake.Latitude()
	case "LatitudeDegrees":
		return fake.LatitudeDegrees()
	case "LatitudeDirection":
		return fake.LatitudeDirection()
	case "LatitudeMinutes":
		return fake.LatitudeMinutes()
	case "LatitudeSeconds":
		return fake.LatitudeSeconds()
	case "Longitude":
		return fake.Longitude()
	case "LongitudeDegrees":
		return fake.LongitudeDegrees()
	case "LongitudeDirection":
		return fake.LongitudeDirection()
	case "LongitudeMinutes":
		return fake.LongitudeMinutes()
	case "LongitudeSeconds":
		return fake.LongitudeSeconds()
	case "MaleFirstName":
		return fake.MaleFirstName()
	case "MaleFullName":
		return fake.MaleFullName()
	case "MaleFullNameWithPrefix":
		return fake.MaleFullNameWithPrefix()
	case "MaleFullNameWithSuffix":
		return fake.MaleFullNameWithSuffix()
	case "MaleLastName":
		return fake.MaleLastName()
	case "MalePatronymic":
		return fake.MalePatronymic()
	case "Model":
		return fake.Model()
	case "Month":
		return fake.Month()
	case "MonthNum":
		return fake.MonthNum()
	case "MonthShort":
		return fake.MonthShort()
	case "Paragraph":
		return fake.Paragraph()
	case "Paragraphs":
		return fake.Paragraphs()
	case "Patronymic":
		return fake.Patronymic()
	case "Phone":
		return fake.Phone()
	case "Product":
		return fake.Product()
	case "ProductName":
		return fake.ProductName()
	case "Sentence":
		return fake.Sentence()
	case "Sentences":
		return fake.Sentences()
	case "SimplePassword":
		return fake.SimplePassword()
	case "State":
		return fake.State()
	case "StateAbbrev":
		return fake.StateAbbrev()
	case "Street":
		return fake.Street()
	case "StreetAddress":
		return fake.StreetAddress()
	case "Title":
		return fake.Title()
	case "TopLevelDomain":
		return fake.TopLevelDomain()
	case "UserAgent":
		return fake.UserAgent()
	case "UserName":
		return fake.UserName()
	case "WeekDay":
		return fake.WeekDay()
	case "WeekDayShort":
		return fake.WeekDayShort()
	case "WeekdayNum":
		return fake.WeekdayNum()
	case "Word":
		return fake.Word()
	case "Words":
		return fake.Words()
	case "Year":
		return fake.Year(0, time.Now().Year())
	case "Zip":
		return fake.Zip()
	}
	return p
}
