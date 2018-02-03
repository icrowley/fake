package fake

import (
	"strings"
)

// FileExtension generates random file extension
func FileExtension() string {
	return lookup(lang, "file_extensions", true)
}

// FileName generates random filename (including extension)
func FileName() string {
	return strings.Replace(WordsN(r.Intn(2)+1), " ", "-", -1) + "." + FileExtension()
}
