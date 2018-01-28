package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Pwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir + "/"
}

func AssetsDir() string {
	return Pwd() + "assets/"
}

func FilesDir() string {
	return Pwd() + "assets/files/"
}

// func ReadFile(arg string) string {
// content, err := ioutil.ReadFile(AssetsDir() + arg)
// if err != nil {
// fmt.Println("ERROR ReadFile()")
// }
// lines := strings.Split(string(content), "\n")
// return lines[0]
// }

func TruncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num]
	}
	return bnoden
}

func FixTitle(arg string) (string, string) {
	semifixed := strings.Replace(arg, "_", " ", -1)
	fixed := strings.Split(semifixed, ".")
	if len(fixed) == 1 {
		return fixed[0], ""
	} else {
		return fixed[0], fixed[1]
	}
}

func StripPath(arg string) string {
	for n := len(arg) - 1; n >= 0; n-- {
		if arg[n:n+1] == "/" {
			return arg[n+1:]
		}
	}
	return ""
}

func TestingAssetsDir() string {
	return "/Users/andrewlee/Documents/School/Interim2018/MCA/pause/backends/assets/"
}

func TestingFiles() string {
	return "/Users/andrewlee/Documents/School/Interim2018/MCA/pause/backends/assets/files/"
}
