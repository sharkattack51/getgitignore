package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

var (
	GITIGNORE_FILE       = ".gitignore"
	GITHUB_GITIGNORE_URL = "https://raw.githubusercontent.com/github/gitignore/main/%v.gitignore"
	APPEND_TO_IGNORES    = []string{".DS_Store"}

	_404 = "404: Not Found"
)

func main() {
	if len(os.Args) >= 2 {
		lang := strcase.ToCamel(os.Args[1])
		resp, err := http.Get(fmt.Sprintf(GITHUB_GITIGNORE_URL, lang))
		if err != nil {
			Println("fail", err.Error())
		} else {
			if !IsExist(GITIGNORE_FILE) {
				bytes, _ := io.ReadAll(resp.Body)
				if string(bytes) == _404 {
					Println("fail", _404)
				} else {
					strs := string(bytes)

					// append ignores
					b := false
					for _, ig := range APPEND_TO_IGNORES {
						if !strings.Contains(strs, ig) {
							if !b {
								b = true
								strs = strings.TrimRight(strs, " ")
								strs = strings.TrimRight(strs, "\n")
								strs += "\n\n# Append Ignores\n"
							}

							strs += (ig + "\n")
						}
					}

					err = os.WriteFile(GITIGNORE_FILE, []byte(strs), 0666)
					if err != nil {
						Println("fail", err.Error())
					} else {
						p, _ := filepath.Abs(GITIGNORE_FILE)
						Println("success", p)
					}
				}
			} else {
				Println("fail", ".gitignore already exist")
			}
		}
		defer resp.Body.Close()
	} else {
		Println("help", "enter development environment or programming language name.")
	}
}

func Println(stat string, arg string) {
	fmt.Println(fmt.Sprintf("getgitignore: [%v] %v", stat, arg))
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
