package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
	errAtoi      = "failed to convert to int: %w"
	byrMin       = 1920
	byrMax       = 2002
	iyrMin       = 2010
	iyrMax       = 2020
	eyrMin       = 2020
	eyrMax       = 2030
)

type passportField struct {
	value string
	valid func(string) (bool, error)
}

func (p *passportField) present() bool {
	return p.value != ""
}

func newIntegerPassportField(min, max int) *passportField {
	passportField := &passportField{
		valid: func(value string) (bool, error) {
			i, err := strconv.Atoi(value)
			if err != nil {
				return false, fmt.Errorf(errAtoi, err)
			}
			return i >= min && i <= max, nil
		},
	}
	return passportField
}

type passport struct {
	fields map[string]*passportField
}

func newPassport() *passport {
	fields := map[string]*passportField{
		"byr": newIntegerPassportField(byrMin, byrMax),
		"iyr": newIntegerPassportField(iyrMin, iyrMax),
		"eyr": newIntegerPassportField(eyrMin, eyrMax),
		"hgt": {
			valid: func(value string) (bool, error) {
				if strings.Contains(value, "cm") {
					cIndex := strings.Index(value, "c")
					number, err := strconv.Atoi(value[:cIndex])
					if err != nil {
						return false, fmt.Errorf(errAtoi, err)
					}
					return number >= 150 && number <= 193, nil
				} else if strings.Contains(value, "in") {
					iIndex := strings.Index(value, "i")
					number, err := strconv.Atoi(value[:iIndex])
					if err != nil {
						return false, fmt.Errorf(errAtoi, err)
					}
					return number >= 59 && number <= 76, nil
				}
				return false, nil
			},
		},
		"hcl": {
			valid: func(value string) (bool, error) {
				return len(value) > 0 && value[0] == '#' && len(value[1:]) == 6, nil
			},
		},
		"ecl": {
			valid: func(value string) (bool, error) {
				return util.Exists(value, "amb", "blu", "brn", "gry", "grn", "hzl", "oth"), nil
			},
		},
		"pid": {
			valid: func(value string) (bool, error) {
				return len(value) == 9, nil
			},
		},
	}
	return &passport{
		fields: fields,
	}
}

func (p *passport) allPresent() bool {
	for _, field := range p.fields {
		if !field.present() {
			return false
		}
	}
	return true
}

func (p *passport) valid() (bool, error) {
	for _, field := range p.fields {
		if !field.present() {
			return false, nil
		}
		valid, err := field.valid(field.value)
		if err != nil {
			return false, err
		}
		if !valid {
			return false, nil
		}
	}
	return true, nil
}

func (p *passport) ingest(input string) {
	colon := strings.Index(input, ":")
	name := input[:colon]
	value := input[colon+1:]
	if field, ok := p.fields[name]; ok {
		field.value = value
	}
}

func Part1() (int, error) {
	lines, err := util.ReadLines()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	var valid int
	passport := newPassport()
	for _, line := range lines {
		if line == "" {
			if passport.allPresent() {
				valid++
			}
			passport = newPassport()
		} else {
			splits := strings.Split(line, " ")
			for _, split := range splits {
				passport.ingest(split)
			}
		}
	}
	if passport.allPresent() {
		valid++
	}
	return valid, nil
}

func Part2() (int, error) {
	lines, err := util.ReadLines()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	var valid int
	passport := newPassport()
	for _, line := range lines {
		if line == "" {
			validPassport, err := passport.valid()
			if err != nil {
				return 0, err
			}
			if validPassport {
				valid++
			}
			passport = newPassport()
		} else {
			splits := strings.Split(line, " ")
			for _, split := range splits {
				passport.ingest(split)
			}
		}
	}
	validPassport, err := passport.valid()
	if err != nil {
		return 0, err
	}
	if validPassport {
		valid++
	}
	return valid, nil
}
