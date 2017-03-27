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
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cast"
)

type Page struct {
	gorm.Model
	
	//ID        int       // Exists by default
	Url       string `unique`
	Content   string
	Diff      string
	//CreatedAt time.Time // Exists by default
	//UpdatedAt time.Time // Exists by default
	//DeletedAt time.Time // Exists by default
}

func DbUpdateUrls(urls []*Url) {
	// Prepare database
	cacheFile := GetCacheLocation()+"cache.db"
	db, err := gorm.Open("sqlite3", cacheFile)
	checkErr(err)

	defer db.Close()

	log.Debug("Database location:", cacheFile)

	// Test if table exists in the database
	if !db.HasTable(&Page{}) {
		log.Warn("Database Pages did not exist. Creating table \"pages\" in cache.")
		db.CreateTable(&Page{})
	}
	
	// Actions on urls
	for _, u := range urls {
		if u.Found {
			dbCompare(db, u)
		}
	}
}

// The database has Pages, and we crawled Urls.
// This function compares a found Url with the corresponding
// Page in database.
func dbCompare(db *gorm.DB, u *Url) {
	var page Page
	db.Where(Page{Url: u.Address}).First(&page)

	// Body size
	if u.Body != nil {
		// Convert body to string
		body := cast.ToString(u.Body)

		if (page.Url == u.Address) {
			if (page.Content != body) {
				log.Info("==> Page updated : "+u.Address)
				page.Content = body
				db.Save(&page)
			} else {
				log.Debug("Page non modified : "+u.Address)
			}
		} else {
			log.Warn("New page "+u.Address+"; Recording it for later.")
			page = Page{Url: u.Address, Content: body}
			db.Create(&page)
		}
	} else {
		log.Fatal("ERROR : Url"+u.Address+"found but empty body.")
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
