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

type DB struct {
	*gorm.DB
}

// Page entity describes one URL.
type Page struct {
	gorm.Model
	//PageStates  []PageState // One Page has many PageState(s)

	Url         string `unique`
	Type        string
}

// PageState entity is the state of an URL when one Update is launched.
type PageState struct {
	gorm.Model
	UpdateID  uint
	Page      Page
	
	Content   string
	Diff      string
}

// The Update entity describes one call to `scrapix update`.
type Update struct {
	gorm.Model
	//PageStates  []PageState
}

func InitDatabase() (*DB, error) {
	// Prepare database
	cacheFile := GetCacheLocation()+"cache.db"
	db, err := gorm.Open("sqlite3", cacheFile)
	checkErr(err)
	log.Debug("Database location:", cacheFile)

	// Test if tables exist in the database
	if !db.HasTable(&Page{}) {
		log.Warn("Database Pages did not exist. Creating table \"pages\" in cache.")
		db.CreateTable(&Page{})
	}
	if !db.HasTable(&PageState{}) {
		log.Warn("Database Pages did not exist. Creating table \"pagestate\" in cache.")
		db.CreateTable(&PageState{})
	}
	if !db.HasTable(&Update{}) {
		log.Warn("Database Pages did not exist. Creating table \"update\" in cache.")
		db.CreateTable(&Update{})
	}

	
	return &DB{db}, err
}

func (db *DB)DbUpdateUrls(urls []*Url) {
	var	update Update
	

	// Actions on urls
	for _, u := range urls {
		if u.Found {
			db.dbCompare(update, u)
		}
	}
}

// The database has Pages, PageStates and Updates, and
// we crawled Urls.
// This function compares a found Url with the corresponding
// Page in database.
func (db *DB) dbCompare(update Update, u *Url) {
	var (
		page Page
		pageState PageState
		pageStates []PageState
	)
	
	//db.Where(PageState{Page.Url: u.Address}).Order("created_at desc").First(&pageState)
	db.Model(Page{Url: u.Address}).Related(&pageStates).Order("created_at desc").First(&pageState)

	// Body size
	if u.Body != nil {
		// Convert body to string
		body := cast.ToString(u.Body)

		if (pageState.Page.Url == u.Address) {
			if (pageState.Content != body) {
				log.Info("==> Page updated : "+u.Address)
				pageState.Content = body
				db.Save(&pageState)
			} else {
				log.Debug("Page non modified : "+u.Address)
			}
		} else {
			// TODO s'assurer que page est bien le bon objet
			log.Warn("New page "+u.Address+"; Recording it for later.")
			pageState = PageState{Page: page, Content: body}
			db.Create(&pageState)
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
