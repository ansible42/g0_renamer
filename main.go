package main

import (
	"fmt"
	//"io/ioutil"
	"os"
	//"path"
	"path/filepath"
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

func RenameFolderFiles(path string, PatternToReplace string, ReplacemntPattern string) {

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			// rename iff is file not directory
			fmt.Printf("File %s renamed to %s\n", info.Name(), Replace(PatternToReplace, ReplacemntPattern, info.Name()))
			/*os.Rename(info.Name(), Replace(PatternToReplace, ReplacemntPattern, info.Name())) */
		}
		return nil
	})

}

func main() {

	var SearchDirectory = os.Args[1]
	var SearchString = os.Args[2]
	var ReplaceString = os.Args[3]

	fmt.Printf("Search directory : %s\n", SearchDirectory)
	fmt.Printf("Search String : %s to be replaced with : %s \n", SearchString, ReplaceString)

	RenameFolderFiles(SearchDirectory, SearchString, ReplaceString)
}
