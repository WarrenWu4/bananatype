// properly find the resource path

package resourcepath

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetResourcePath() string {
	// TODO: fix this shit up cause it is not it and unreliable
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	path := strings.Split(filepath.Dir(filename), "/")
	textPath := strings.Join(path[0:len(path)-2], "/")
	if textPath == os.Getenv("HOME")+"/.local/bin" {
		fmt.Println("On local install")
		return os.Getenv("HOME") + "/.local/share/bananatype"
	} else if textPath == "/usr/bin" {
		return "/usr/share/bananatype"
	}
	return textPath + "/resources"
}
