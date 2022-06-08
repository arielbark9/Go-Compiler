package Parser

import (
	"errors"
	"regexp"
	"strings"
)

var source string
var current int

var DIGIT_REGEX = regexp.MustCompile("[0-9]")
var CAPITAL_SMALL_LETTER_REGEX = regexp.MustCompile("[A-Za-z]")
var IDENTIFIER_REGEX = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*")
var INTEGER_REGEX = regexp.MustCompile("^(3276[0-7]|327[0-5]\\d|32[0-6]\\d{2}|3[01]\\d{3}|[12]\\d{4}|[1-9]\\d{3}|[1-9]\\d{2}|[1-9]\\d|\\d)$")
var STRING_REGEX = regexp.MustCompile("\".*\"")
var KEYWORD_REGEX = regexp.MustCompile("^(class|constructor|function|method|field|static|var|int|char|boolean|void|true|false|null|this|let|do|if|else|while|return)")
var SYMBOLS = [...]string{"{", "}", "(", ")", "[", "]", ".", ",", ";", "+", "-", "*", "/", "&", "|", "<", ">", "=", "~"}

func match(current int, regex *regexp.Regexp, _type string) (token, error) {
	currentLexeme := string(source[:current])
	if regex.MatchString(currentLexeme) {
		source = strings.TrimPrefix(source, currentLexeme)
		return token{tType: _type, tValue: strings.Replace(currentLexeme, "\"", "", -1)}, nil
	} else {
		return token{}, errors.New("invalid " + _type)
	}
}

func Q0() token {
	currChar := string(source[current])
	if CAPITAL_SMALL_LETTER_REGEX.MatchString(currChar) ||
		currChar == "_" {
		current++
		return Q1()
	} else if DIGIT_REGEX.MatchString(currChar) {
		current++
		return Q3()
	} else if contains(SYMBOLS[:], currChar) {
		value := currChar
		if currChar == "&" {
			value = "&amp;"
		} else if currChar == "<" {
			value = "&lt;"
		} else if currChar == ">" {
			value = "&gt;"
		} else if currChar == "\"" {
			value = "&quot;"
		}
		source = strings.TrimPrefix(source, currChar)
		current = 0
		return token{
			tType:  "symbol",
			tValue: value,
		}
	} else if currChar == "\"" {
		current++
		return Q4()
	} else {
		panic("invalid identifier")
	}
}

func Q1() token {
	currChar := string(source[current])
	if CAPITAL_SMALL_LETTER_REGEX.MatchString(currChar) {
		current++
		return Q1()
	} else if currChar == "_" || DIGIT_REGEX.MatchString(currChar) {
		current++
		return Q2()
	} else { // first try to match keywords.
		tok, err := match(current, KEYWORD_REGEX, "keyword")
		if err != nil {
			tok, err = match(current, IDENTIFIER_REGEX, "identifier")
			if err != nil {
				panic(err)
			}
		}
		current = 0
		return tok
	}
}

func Q2() token {
	currChar := string(source[current])
	if CAPITAL_SMALL_LETTER_REGEX.MatchString(currChar) ||
		currChar == "_" || DIGIT_REGEX.MatchString(currChar) {
		current++
		return Q2()
	} else {
		tok, err := match(current, IDENTIFIER_REGEX, "identifier")
		if err != nil {
			panic(err)
		}
		current = 0
		return tok
	}
}

func Q3() token {
	currChar := string(source[current])
	if DIGIT_REGEX.MatchString(currChar) {
		current++
		return Q3()
	} else {
		tok, err := match(current, INTEGER_REGEX, "integerConstant")
		if err != nil {
			panic(err)
		}
		current = 0
		return tok
	}
}

func Q4() token {
	currChar := string(source[current])
	if currChar == "\"" {
		current++
		tok, err := match(current, STRING_REGEX, "stringConstant")
		if err != nil {
			panic(err)
		}
		current = 0
		return tok
	} else if currChar == "\n" {
		panic("invalid string")
	} else {
		current++
		return Q4()
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetTokens(_source string) []token {
	source = _source
	current = 0
	var tokens []token
	for !(len(source) == 0) {
		source = strings.Trim(source, " \n\t\r")

		for strings.HasPrefix(source, "//") || strings.HasPrefix(source, "/*") {
			// if source start is //, remove until end-line
			if source[0:2] == "//" {
				source = strings.Join(strings.Split(source, "\n")[1:], "\n")
			}
			// if source start is /*, remove until */
			if source[0:2] == "/*" {
				source = strings.Join(strings.Split(source, "*/")[1:], "*/")
			}
			source = strings.Trim(source, " \n\t\r")
		}
		tokens = append(tokens, Q0())
	}

	return tokens
}
