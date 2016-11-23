// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package scraplib

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	//"strings"
	"io/ioutil"
)

// Got from https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

// Extract all http** links from a given webpage
func Crawl(url *Url /*ch chan *Url, chFinished chan bool*/) {
	resp, err := http.Get(url.Address)

	defer func() {
		// Notify that we're done after this function
		url.ChFinished <- true
	}()

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url.Address + "\". Please check the configuration file for mistakes.")
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns

	//defer io.Copy(url.Body, b) // Copy that to the Url buffer
	body, err := ioutil.ReadAll(resp.Body)
	url.Body = body
	if err != nil {
		fmt.Println("ERROR: Failed to read \"" + url.Address + "\"")
	}
	url.Found = true
	/*
		z := html.NewTokenizer(b)

		for {
			tt := z.Next()

			switch {
			case tt == html.ErrorToken:
				// End of the document, we're done
				return
			case tt == html.StartTagToken:
				t := z.Token()

				// Check if the token is an <a> tag
				isAnchor := t.Data == "a"
				if !isAnchor {
					continue
				}

				// Extract the href value, if there is one
				ok, uAdd := getHref(t)
				if !ok {
					continue
				}
				url.Address = uAdd

				// Make sure the url begines in http**
				hasProto := strings.Index(url.Address, "http") == 0
				if hasProto {
					ch <- url
				}
			}
		}*/
	//ch <- url
}
