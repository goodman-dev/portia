package wiki

import "strings"

/*
 * Types and functions to support wikipedia parsing
 */

func IsWikiPage(link string) bool {
	return strings.Contains(link, "/wiki/") && !strings.Contains(link, "https://")
}

func GetPageName(link string) string {
	return strings.TrimPrefix(link, "/wiki/")
}
