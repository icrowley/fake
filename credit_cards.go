package fake

import (
	"strings"

	"strconv"
)

type creditCard struct {
	vendor   string
	length   int
	prefixes []int
}

var creditCards = map[string]creditCard{
	"visa":       {"VISA", 16, []int{4539, 4556, 4916, 4532, 4929, 40240071, 4485, 4716, 4}},
	"mastercard": {"MasterCard", 16, []int{51, 52, 53, 54, 55}},
	"amex":       {"American Express", 15, []int{34, 37}},
	"discover":   {"Discover", 16, []int{6011}},
}

// CreditCardType returns one of the following credit values:
// VISA, MasterCard, American Express and Discover
func CreditCardType() string {
	n := len(creditCards)
	var vendors []string
	for _, cc := range creditCards {
		vendors = append(vendors, cc.vendor)
	}

	return vendors[r.Intn(n)]
}

// CreditCardNum generated credit card number according to the card number rules
func CreditCardNum(vendor string) string {
	if vendor != "" {
		vendor = strings.ToLower(vendor)
	} else {
		var vendors []string
		for v := range creditCards {
			vendors = append(vendors, v)
		}
		vendor = vendors[r.Intn(len(vendors))]
	}
	card := creditCards[vendor]
	prefix := strconv.Itoa(card.prefixes[r.Intn(len(card.prefixes))])
	num := []rune(prefix)

	// -1 to leave room for the checksum
	for i := 0; i < card.length-len(prefix)-1; i++ {
		num = append(num, random_int_rune())
	}

	// and now add the checksum
	num = append(num, genCCDigit(num))
	return string(num)
}

func random_int_rune() rune {
	// random int [0,9) => string => array of runes => single rune
	return []rune(strconv.Itoa(r.Intn(10)))[0]
}

func genCCDigit(num []rune) rune {
	// See https://en.wikipedia.org/wiki/Luhn_algorithm
	// But be sure to also read the sordid "talk" page --
	// the example implementations have errors.
	sum := 0
	pos := 0
	for i := len(num) - 1; i >= 0; i-- {
		// The next line is equivalent to `n, _ := strconv.Atoi(string(num[i]))`
		n := int(num[i] - '0')
		if pos % 2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		pos++
	}

	// https://en.wikipedia.org/wiki/Talk:Luhn_algorithm#Formula_error
	checksum := (10 - (sum % 10) % 10)

	// int => string => array of runes => single rune
	return []rune(strconv.Itoa(checksum))[0]
}
