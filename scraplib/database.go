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
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Page struct {
	Id           int
	Url          string `sql:unique`
	LastModified time.Time
	Content      string
}

func DbSearch() {
	db, err := gorm.Open("sqlite3", "cache.db")
	checkErr(err)

	defer db.Close()

	db.LogMode(true)

	if !db.HasTable(new(Page)) {
		fmt.Println("Database Pages did not exist. Creating table \"pages\" in cache.")
		db.CreateTable(new(Page))
	}

	pageOne := Page{Url: "http://agreg.org/", Content: "Hello World!"}
	pageTwo := Page{Url: "http://www.topologix.fr/sujets/"}

	db.Save(&pageOne)
	db.Save(&pageTwo)

	var somePage Page

	db.Where("url = ?", "http://agreg.org/").Find(&somePage)

	fmt.Println("Content of Page:", somePage.Content)

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
