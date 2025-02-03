package portia

import (
	"strings"

	"golang.org/x/net/html"

	"github.com/goodman-dev/portia/internal/wiki"
)

func ExtractWikiLinks(page *WebPage) ([]string, error) {

	doc, err := html.Parse(strings.NewReader(page.Content))
	if err != nil {
		return nil, err
	}

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if wiki.IsWikiPage(a.Val) {
						links = append(links, wiki.GetPageName(a.Val))
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links, nil
}
