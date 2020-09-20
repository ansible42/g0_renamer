package main

import (
	"fmt"
	//"io/ioutil"
	"os"
	//"path"
	//"path/filepath"
	"regexp"
)

/*
func DirList(Directory string) {
	var FilesFiltered map[int]string
	FilesFiltered = make(map[int]string)
	files, err := ioutil.ReadDir(Directory)
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, file := range files {
		if file.IsDir() != true {
			file.Name() = FilesFiltered[len(FilesFiltered)+1]
		}
	}
}
*/
func Replace(Find string, Replace string, Search string) string {
	var output string
	r, _ := regexp.Compile(Find)
	bReplace := []byte(Replace)
	bSearch := []byte(Search)
	if r.MatchString(Search) {
		out := r.ReplaceAllLiteral(bSearch, bReplace)
		output = string(out)
	} else {
		output = Search
	}

	return output
}
func main() {
	var searchstr string
	var replacestr string
	var regexstr string
	searchstr = os.Args[1]
	replacestr = os.Args[3]
	regexstr = os.Args[2]
	fmt.Println("Search string : %s and replace %s with %s", searchstr, regexstr, replacestr)
	fmt.Println("Return :: ")
	fmt.Printf(Replace(regexstr, replacestr, searchstr))
}
