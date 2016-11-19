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
	"fmt"
	"runtime"
	//"github.com/urfave/cli"
	//"github.com/Sirupsen/logrus"

	"github.com/lebarde/scrapix/scraplib"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	VERSION := "alpha 1"

	fmt.Println("Scrapix v.", VERSION)

	// TODO read command line and do actions

	// Read config file and populate data
	urls := scraplib.UrlsConfig(scraplib.ReadConfig())

	// TODO check for refresh times

	// Create channels
	//chUrls := make(chan *scraplib.Url)
	//chFinished := make(chan bool)

	// Retrieve the urls
	fmt.Println("Launching requests:")
	for _, u := range urls {
		fmt.Println("-", u.Address)
		go scraplib.Crawl(u /*, chUrls, chFinished*/)
	}

	// Subscribe to all the channels
	for _, u := range urls {
		//fmt.Println("debug: Range enter")
		select {
		case <-u.ChFinished:
			//fmt.Println("Found url", u.Address)
			if !u.Found {
				fmt.Println("debug: But url.Found = false…")
			}
		}
	}

	// We're done! Print the results...
	fmt.Println("\nFound some urls:")

	for _, u := range urls {
		//fmt.Println("(range:",u.Address,")")
		if u.Found {
			fmt.Println(" - " + u.Address)
			//fmt.Print(string(u.Body[:]))
		}
	}

	//close(chUrls)

	// TODO compare inside the database
	// then show the results.
}
