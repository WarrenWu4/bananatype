// properly find the resource path

package resourcepath

import (
	"bananas/pkg/logger"
)

var (
	Build = "dev"
)

func GetResourcePath() string {
	// FIX: not entirely reliable since where the pkg manager
	// installs the machine is varied
	if Build == "prod" {
		logger.Log(logger.DEBUG, "Using prod resource path: /usr/share/bananatype")
		return "/usr/share/bananatype"
	} else {
		logger.Log(logger.DEBUG, "Using dev resource path: ./resources")
		return "./resources"
	}
	// TODO: fix this shit up cause it is not it and unreliable
	// FIX: in prod -trimpath strips absolute path of runtime.Caller(0)
	// thus making this a piece of shit

	// _, filename, _, ok := runtime.Caller(0)
	// if !ok {
	// 	panic("No caller information")
	// }
	// logger.Log(logger.DEBUG, "Using filename: "+filename)
	// path := strings.Split(filepath.Dir(filename), "/")
	// textPath := strings.Join(path[0:len(path)-2], "/")
	// if textPath == os.Getenv("HOME")+"/.local/bin" {
	// 	fmt.Println("On local install")
	// 	return os.Getenv("HOME") + "/.local/share/bananatype"
	// } else if textPath == "/usr/bin" {
	// 	return "/usr/share/bananatype"
	// }
	// return textPath + "/resources"
}
