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
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/html"
	"net/http"
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
		log.Error("Failed to crawl \"" + url.Address + "\". Please check the configuration file for mistakes.")
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns

	body, err := ioutil.ReadAll(resp.Body)
	url.Body = body
	if err != nil {
		log.Error("Failed to read \"" + url.Address + "\"")
	}
	url.Found = true
}
