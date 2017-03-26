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
)

type Url struct {
	Address    string
	Refresh    string // = "1d" // TODO default values??
	Watch      string //= nil
	Found      bool
	Body       []byte
	ChFinished chan bool
}

func NewUrl() *Url {
	u := Url{Refresh: "1d", Found: false, ChFinished: make(chan bool)}

	return &u
}

func UrlsConfig(urlsConfig map[string]interface{}) []*Url {
	var urls []*Url

	for address, _ := range urlsConfig { // TODO second parameter = opt
		u := NewUrl()
		u.Address = address

		// TODO refresh, watch (from config file)

		urls = append(urls, u)
	}
	return urls
}
