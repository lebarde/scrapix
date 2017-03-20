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
	//"github.com/jinzhu/gorm"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Page struct {
	ID        int    //xorm.Model // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	Url       string `sql:unique`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func DbCheckUrls(urls []*Url) {
	// Prepare database
	//db, err := gorm.Open("sqlite3", "$HOME/.cache/scrapix/cache.db")
	db, err := xorm.NewEngine("sqlite3", GetCacheLocation()+"cache.db")
	checkErr(err)

	defer db.Close()

	//db.LogMode(true)
	db.ShowSQL(true)

	fmt.Println("Database location:", GetCacheLocation()+"cache.db")

	// Actions on urls
	for _, u := range urls {
		dbCompare(db, u)
	}
}

// The database has Pages, and we crawled Urls.
// This function compares a found Url with the corresponding
// Page in database.
func dbCompare(db *xorm.Engine, u *Url) {
	/*
	// Test if table exists in the database
	if !db.HasTable(new(Page)) {
		fmt.Println("Database Pages did not exist. Creating table \"pages\" in cache.")
		db.CreateTable(new(Page))
	}*/
	db.Sync(new(Page))

	fmt.Println("beep")
	//pageOne := Page{Url: "http://agreg.org/", Content: "Hello World!"}
	//pageTwo := Page{Url: "http://www.topologix.fr/sujets/"}

	//db.Save(&pageOne)
	//db.Save(&pageTwo)

	//db.Where("url = ?", "http://agreg.org/").Find(&somePage)
	//somePage := Page{Content: "Hello World!"}
	//has, err := db.Where(&somePage).Get(&somePage)

	var somePages []Page
	//err = db.Asc("id").Find(&somePages)
	err := db.Sql("select * from pages").Find(&somePages)
	
	if err != nil {
		fmt.Println("Found", len(somePages), "pages:")
		fmt.Printf("%#v\n", somePages)
		for _, p := range somePages {
			fmt.Println("Content of Page:", p.Url)
		}
	} else {
		fmt.Println("Content not found. Insert things.")
		page := Page{Url: "helloworld", Content:time.Now().String()}
		db.Insert(page)
		fmt.Println("Inserted page:", page.ID)
	}


	/*

		db, err := sql.Open("sqlite3", "./foo.db")
		checkErr(err)

		err = checkDB(db)
		checkErr(err)

		// insert
		stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
		checkErr(err)

		res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
		// update
		stmt, err = db.Prepare("update userinfo set username=? where uid=?")
		checkErr(err)

		res, err = stmt.Exec("astaxieupdate", id)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		// query
		rows, err := db.Query("SELECT * FROM userinfo")
		checkErr(err)
		var uid int
		var username string
		var department string
		var created time.Time

		for rows.Next() {
			err = rows.Scan(&uid, &username, &department, &created)
			checkErr(err)
			fmt.Println(uid)
			fmt.Println(username)
			fmt.Println(department)
			fmt.Println(created)
		}

		rows.Close() //good habit to close

		// delete
		stmt, err = db.Prepare("delete from userinfo where uid=?")
		checkErr(err)

		res, err = stmt.Exec(id)
		checkErr(err)

		affect, err = res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		db.Close()
	*/

}

/*
// Checks if the database is well structured, and if not,
// (re)generate it
func checkDB(db *sql.DB) (err error) {
	// Check if the rows exist
	rows, err :=
	rows, err := db.Query(`CREATE TABLE "userinfo" (
    "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
    "username" VARCHAR(64) NULL,
    "departname" VARCHAR(64) NULL,
    "created" DATE NULL
    );`)
	fmt.Println("Rows:", rows)

	return nil
}
*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
