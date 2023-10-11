package govalidator

import (
	"errors"
	"fmt"
	"html"
	"math"
	"path"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Contains checks if the string contains the substring.
func Contains[T ~string](str, substring T) bool {
	return strings.Contains(string(str), string(substring))
}

// Matches checks if string matches the pattern (pattern is regular expression)
// In case of error return false
func Matches[T ~string](str, pattern T) bool {
	match, _ := regexp.MatchString(string(pattern), string(str))
	return match
}

// LeftTrim trims characters from the left side of the input.
// If second argument is empty, it will remove leading spaces.
func LeftTrim[T ~string](str, chars T) string {
	if chars == "" {
		return strings.TrimLeftFunc(string(str), unicode.IsSpace)
	}
	r, _ := regexp.Compile("^[" + string(chars) + "]+")
	return r.ReplaceAllString(string(str), "")
}

// RightTrim trims characters from the right side of the input.
// If second argument is empty, it will remove trailing spaces.
func RightTrim[T ~string](str, chars T) string {
	if chars == "" {
		return strings.TrimRightFunc(string(str), unicode.IsSpace)
	}
	r, _ := regexp.Compile("[" + string(chars) + "]+$")
	return r.ReplaceAllString(string(str), "")
}

// Trim trims characters from both sides of the input.
// If second argument is empty, it will remove spaces.
func Trim[T ~string](str, chars T) string {
	return LeftTrim(RightTrim(str, chars), string(chars))
}

// WhiteList removes characters that do not appear in the whitelist.
func WhiteList[T ~string](str, chars T) string {
	pattern := "[^" + chars + "]+"
	r, _ := regexp.Compile(string(pattern))
	return r.ReplaceAllString(string(str), "")
}

// BlackList removes characters that appear in the blacklist.
func BlackList[T ~string](str, chars T) string {
	pattern := "[" + chars + "]+"
	r, _ := regexp.Compile(string(pattern))
	return r.ReplaceAllString(string(str), "")
}

// StripLow removes characters with a numerical value < 32 and 127, mostly control characters.
// If keep_new_lines is true, newline characters are preserved (\n and \r, hex 0xA and 0xD).
func StripLow[T ~string](str T, keepNewLines bool) string {
	chars := ""
	if keepNewLines {
		chars = "\x00-\x09\x0B\x0C\x0E-\x1F\x7F"
	} else {
		chars = "\x00-\x1F\x7F"
	}
	return BlackList(str, T(chars))
}

// ReplacePattern replaces regular expression pattern in string
func ReplacePattern[T ~string](str, pattern, replace T) string {
	r, _ := regexp.Compile(string(pattern))
	return r.ReplaceAllString(string(str), string(replace))
}

// Escape replaces <, >, & and " with HTML entities.
var Escape = html.EscapeString

func addSegment(inrune, segment []rune) []rune {
	if len(segment) == 0 {
		return inrune
	}
	if len(inrune) != 0 {
		inrune = append(inrune, '_')
	}
	inrune = append(inrune, segment...)
	return inrune
}

// UnderscoreToCamelCase converts from underscore separated form to camel case form.
// Ex.: my_func => MyFunc
func UnderscoreToCamelCase[T ~string](s T) string {
	caser := cases.Title(language.English)
	return strings.Replace(caser.String(strings.Replace(strings.ToLower(string(s)), "_", " ", -1)), " ", "", -1)
}

// CamelCaseToUnderscore converts from camel case form to underscore separated form.
// Ex.: MyFunc => my_func
func CamelCaseToUnderscore[T ~string](str T) string {
	var output []rune
	var segment []rune
	for _, r := range str {

		// not treat number as separate segment
		if !unicode.IsLower(r) && string(r) != "_" && !unicode.IsNumber(r) {
			output = addSegment(output, segment)
			segment = nil
		}
		segment = append(segment, unicode.ToLower(r))
	}
	output = addSegment(output, segment)
	return string(output)
}

// Reverse returns reversed string
func Reverse[T ~string](s T) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// GetLines splits string by "\n" and return array of lines
func GetLines[T ~string](s T) []string {
	return strings.Split(string(s), "\n")
}

// GetLine returns specified line of multiline string
func GetLine[T ~string](s T, index int) (string, error) {
	lines := GetLines(s)
	if index < 0 || index >= len(lines) {
		return "", errors.New("line index out of bounds")
	}
	return lines[index], nil
}

// RemoveTags removes all tags from HTML string
func RemoveTags[T ~string](s T) string {
	return ReplacePattern(s, "<[^>]*>", "")
}

// SafeFileName returns safe string that can be used in file names
func SafeFileName[T ~string](str T) string {
	name := strings.ToLower(string(str))
	name = path.Clean(path.Base(name))
	name = strings.Trim(name, " ")
	separators, err := regexp.Compile(`[ &_=+:]`)
	if err == nil {
		name = separators.ReplaceAllString(name, "-")
	}
	legal, err := regexp.Compile(`[^[:alnum:]-.]`)
	if err == nil {
		name = legal.ReplaceAllString(name, "")
	}
	for strings.Contains(name, "--") {
		name = strings.Replace(name, "--", "-", -1)
	}
	return name
}

// NormalizeEmail canonicalize an email address.
// The local part of the email address is lowercased for all domains; the hostname is always lowercased and
// the local part of the email address is always lowercased for hosts that are known to be case-insensitive (currently only GMail).
// Normalization follows special rules for known providers: currently, GMail addresses have dots removed in the local part and
// are stripped of tags (e.g. some.one+tag@gmail.com becomes someone@gmail.com) and all @googlemail.com addresses are
// normalized to @gmail.com.
func NormalizeEmail[T ~string](str T) (string, error) {
	if !IsEmail(string(str)) {
		return "", fmt.Errorf("%s is not an email", str)
	}
	parts := strings.Split(string(str), "@")
	parts[0] = strings.ToLower(parts[0])
	parts[1] = strings.ToLower(parts[1])
	if parts[1] == "gmail.com" || parts[1] == "googlemail.com" {
		parts[1] = "gmail.com"
		parts[0] = strings.Split(ReplacePattern(parts[0], `\.`, ""), "+")[0]
	}
	return strings.Join(parts, "@"), nil
}

// Truncate a string to the closest length without breaking words.
func Truncate[T ~string](str T, length int, ending string) T {
	var aftstr, befstr string
	if len(str) > length {
		words := strings.Fields(string(str))
		before, present := 0, 0
		for i := range words {
			befstr = aftstr
			before = present
			aftstr = aftstr + words[i] + " "
			present = len(aftstr)
			if present > length && i != 0 {
				if (length - before) < (present - length) {
					return T(Trim(befstr, " /\\.,\"'#!?&@+-") + ending)
				}
				return T(Trim(aftstr, " /\\.,\"'#!?&@+-") + ending)
			}
		}
	}

	return str
}

// PadLeft pads left side of a string if size of string is less then indicated pad length
func PadLeft[T ~string](str T, padStr T, padLen int) T {
	return buildPadStr(str, padStr, padLen, true, false)
}

// PadRight pads right side of a string if size of string is less then indicated pad length
func PadRight[T ~string](str T, padStr T, padLen int) T {
	return buildPadStr(str, padStr, padLen, false, true)
}

// PadBoth pads both sides of a string if size of string is less then indicated pad length
func PadBoth[T ~string](str T, padStr T, padLen int) T {
	return buildPadStr(str, padStr, padLen, true, true)
}

// PadString either left, right or both sides.
// Note that padding string can be unicode and more then one character
func buildPadStr[T ~string](str T, padStr T, padLen int, padLeft bool, padRight bool) T {

	// When padded length is less then the current string size
	if padLen < utf8.RuneCountInString(string(str)) {
		return str
	}

	padLen -= utf8.RuneCountInString(string(str))

	targetLen := padLen

	targetLenLeft := targetLen
	targetLenRight := targetLen
	if padLeft && padRight {
		targetLenLeft = padLen / 2
		targetLenRight = padLen - targetLenLeft
	}

	strToRepeatLen := utf8.RuneCountInString(string(padStr))

	repeatTimes := int(math.Ceil(float64(targetLen) / float64(strToRepeatLen)))
	repeatedString := strings.Repeat(string(padStr), repeatTimes)

	var leftSide T
	if padLeft {
		leftSide = T(repeatedString[0:targetLenLeft])
	}

	var rightSide T
	if padRight {
		rightSide = T(repeatedString[0:targetLenRight])
	}

	return leftSide + str + rightSide
}

// TruncatingErrorf removes extra args from fmt.Errorf if not formatted in the str object
func TruncatingErrorf[T ~string](str T, args ...interface{}) error {
	n := strings.Count(string(str), "%s")
	return fmt.Errorf(string(str), args[:n]...)
}
