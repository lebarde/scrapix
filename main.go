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

package main

import (
	log "github.com/Sirupsen/logrus"
	"runtime"

	"github.com/lebarde/scrapix/scraplib"
	//"github.com/spf13/cobra"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	VERSION := "alpha 1"
	log.Debug("Scrapix v.", VERSION)

	// TODO read command line and do actions

	// Read config file and populate data
	// Data structure is a map of Url - see url.go
	urls := scraplib.UrlsConfig(scraplib.ReadConfig())

	// Retrieve the urls
	log.Debug("Launching ", len(urls), " requests...")
	for _, u := range urls {
		go scraplib.Crawl(u /*, chUrls, chFinished*/)
	}

	// Subscribe to all the channels
	for _, u := range urls {
		select {
		case <-u.ChFinished:
			// Do nothing?
		}
	}

	// We're done! Print the results...
	log.Debug("Found some urls:")

	for _, u := range urls {
		if u.Found {
			log.Debug(" - " + u.Address)
		}
	}

	db, err := scraplib.InitDatabase()
	defer db.Close()
	checkErr(err)
	// TODO compare inside the database
	// then show the results.
	db.DbUpdateUrls(urls)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
