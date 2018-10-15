package fake

var lowerLetters = []rune("abcdefghijklmnopqrstuvwxyz")
var upperLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numeric = []rune("0123456789")
var specialChars = []rune(`!'@#$%^&*()_+-=[]{};:",./?`)
var hexDigits = []rune("0123456789abcdef")

func (f *Faker) text(atLeast, atMost int, allowLower, allowUpper, allowNumeric, allowSpecial bool) string {
	allowedChars := []rune{}
	if allowLower {
		allowedChars = append(allowedChars, lowerLetters...)
	}
	if allowUpper {
		allowedChars = append(allowedChars, upperLetters...)
	}
	if allowNumeric {
		allowedChars = append(allowedChars, numeric...)
	}
	if allowSpecial {
		allowedChars = append(allowedChars, specialChars...)
	}

	result := []rune{}
	nTimes := f.r.Intn(atMost-atLeast+1) + atLeast
	for i := 0; i < nTimes; i++ {
		result = append(result, allowedChars[f.r.Intn(len(allowedChars))])
	}
	return string(result)
}

// Password generates password with the length from atLeast to atMOst charachers,
// allow* parameters specify whether corresponding symbols can be used
func Password(atLeast, atMost int, allowUpper, allowNumeric, allowSpecial bool) string {
	return f.Password(atLeast, atMost, allowUpper, allowNumeric, allowSpecial)
}

// SimplePassword is a wrapper around Password,
// it generates password with the length from 6 to 12 symbols, with upper characters and numeric symbols allowed
func SimplePassword() string {
	return f.SimplePassword()
}

// Color generates color name
func Color() string {
	return f.Color()
}

// DigitsN returns n digits as a string
func DigitsN(n int) string {
	return f.DigitsN(n)
}

// Digits returns from 1 to 5 digits as a string
func Digits() string {
	return f.Digits()
}

func hexDigitsStr(n int) string {
	var num []rune
	for i := 0; i < n; i++ {
		num = append(num, hexDigits[f.r.Intn(len(hexDigits))])
	}
	return string(num)
}

// HexColor generates hex color name
func HexColor() string {
	return f.HexColor()
}

// HexColorShort generates short hex color name
func HexColorShort() string {
	return f.HexColorShort()
}

// Password generates password with the length from atLeast to atMOst charachers,
// allow* parameters specify whether corresponding symbols can be used
func (f *Faker) Password(atLeast, atMost int, allowUpper, allowNumeric, allowSpecial bool) string {
	return f.text(atLeast, atMost, true, allowUpper, allowNumeric, allowSpecial)
}

// SimplePassword is a wrapper around Password,
// it generates password with the length from 6 to 12 symbols, with upper characters and numeric symbols allowed
func (f *Faker) SimplePassword() string {
	return Password(6, 12, true, true, false)
}

// Color generates color name
func (f *Faker) Color() string {
	return f.lookup(f.lang, "colors", true)
}

// DigitsN returns n digits as a string
func (f *Faker) DigitsN(n int) string {
	digits := make([]rune, n)
	for i := 0; i < n; i++ {
		digits[i] = numeric[f.r.Intn(len(numeric))]
	}
	return string(digits)
}

// Digits returns from 1 to 5 digits as a string
func (f *Faker) Digits() string {
	return DigitsN(f.r.Intn(5) + 1)
}

// HexColor generates hex color name
func (f *Faker) HexColor() string {
	return hexDigitsStr(6)
}

// HexColorShort generates short hex color name
func (f *Faker) HexColorShort() string {
	return hexDigitsStr(3)
}
