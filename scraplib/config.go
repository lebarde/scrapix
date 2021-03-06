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
	"github.com/spf13/viper"
)

func ReadConfig() map[string]interface{} {
	viper.SetConfigName("config") // name of config file (without extension)
	//viper.AddConfigPath("$HOME/.config/scrapix")
	viper.AddConfigPath(GetConfigLocation()) // From fs
	viper.AddConfigPath(".") // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
  // TODO Test of viper.ConfigFileNotFoundError
	if serr, ok := err.(*viper.ConfigFileNotFoundError); ok {
    log.Fatal("Config file not found: ", serr)
	} else if err != nil {
		// Handle errors reading the config file
		// TODO panic ?
		log.Fatal("Fatal error config file: ", err)
  }

	return viper.GetStringMap("urls")
}
