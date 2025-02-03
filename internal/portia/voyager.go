package portia

import (
	"fmt"
	"os"
	"slices"
)

/*
 * Voyager is in charge of charting paths from a starting point to a destination
 */

func ChartPath(start string, destination string) ([]string, error) {

	fmt.Println("Charting path from", start, "to", destination)

	// Starting point
	startPage, err := NewWebPage(start, "https://en.wikipedia.org/wiki/"+start)
	if err != nil {
		return nil, err
	}

	route := []string{}
	pagesChecked := []string{}

	var f func(*WebPage, []string)
	f = func(page *WebPage, journey []string) {

		// Land at the new page and update our route
		route = append(journey, page.name)

		// If this is our destination, we're done
		if page.name == destination {
			route = journey
			return
		}

		pagesChecked = append(pagesChecked, page.name)

		// Otherwise, keep exploring - get all links and recurse
		links, err := ExtractWikiLinks(page)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Found %d links on page %s\n", len(links), page.name)

		for _, link := range links {
			if slices.Contains(pagesChecked, link) {
				continue
			}
			nextPage, err := NewWebPage(link, "https://en.wikipedia.org/wiki/"+link)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			f(nextPage, journey)
		}

	}

	f(startPage, []string{})

	return route, nil

}
