// https://golang.org/doc/tutorial/add-a-test

package utils

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestParseLinkHeader calls utils.TestParseLinkHeader with a Link header string, checking
// for a valid return value.
func TestParseLinkHeader(t *testing.T) {
	fmt.Println("INFO >>> Handles good link string")
	var pagingExpected = PagingInfo{
		Next:  15,
		Last:  34,
		First: 1,
		Prev:  13,
	}
	linkHeaderFull := `<https://api.github.com/search/code?q=addClass+user%3Amozilla&page=15>; rel="next",
  <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=34>; rel="last",
  <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=1>; rel="first",
  <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=13>; rel="prev"`
	result, err := GetPagingFromLinkHeader(linkHeaderFull)
	if err != nil {
		fmt.Println("ERROR: could not get paging from link header", err)

	}
	assert.Equal(t, result, pagingExpected)
}
func TestParseLinkHeaderEmptyString(t *testing.T) {
	fmt.Println("INFO >>> Handles empty link string")
	var pagingExpected = PagingInfo{
		Next:  0,
		Last:  0,
		First: 0,
		Prev:  0,
	}
	linkHeaderFull := ``
	result, _ := GetPagingFromLinkHeader(linkHeaderFull)

	assert.Equal(t, result, pagingExpected)
}

// This test is incomplete.
func TestGetRateLimitPerSecond(t *testing.T) {
	date := "Mon, 20 Sep 2021 18:08:38 GMT"
	resetTime := "1632164442"
	remaining := "3689"

	rateLimitInfo, err := ConvertRateLimitInfo(date, resetTime, remaining)
	if err != nil {
		fmt.Println("INFO >>> err", err)
	}
	rateLimitPerSecond := GetRateLimitPerSecond(rateLimitInfo)

	fmt.Println("INFO >>> rateLimitPerSecond", rateLimitPerSecond)
}
