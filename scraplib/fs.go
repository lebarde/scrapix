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
	//"github.com/spf13/afero"
	"os/user"
	"runtime"
)

func GetCacheLocation() string {
	return getHomeLocation()+ ".cache/" + getName() + "/"
}

func GetConfigLocation() string {
	return getHomeLocation() + ".config/" + getName() + "/"
}

func getHomeLocation() string {
	//var AppFs afero.Fs = afero.NewOsFs()

	// Get the user's directory
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Fatal: Could not get the user's directory.", err)
	}

	usrDir := usr.HomeDir

	// The config directory depends on the system
	switch os := runtime.GOOS; os {
	case "darwin":
		// Everything from apple
		fallthrough
	case "dragonfly":
		fallthrough
	case "freebsd":
		fallthrough
	case "netbsd":
		fallthrough
	case "openbsd":
		fallthrough
	case "android":
		fallthrough
	case "plan9":
		fallthrough
	case "solaris":
		fallthrough
	case "linux":
		return usrDir + "/"
	case "windows":
		return usrDir + `\`
	default:
		return ""
	}
}

func getName() string {
	return "scrapix"
}
