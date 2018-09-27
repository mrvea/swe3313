package constants

import "time"

var ValidFormats = []string{
	time.RFC1123,
	time.RFC3339,
	"01/02/2006",
	"01/2/06",
	"2006-01-02",
}
