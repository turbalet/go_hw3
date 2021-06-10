package task5

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)


func SortImports(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var imports[] string
	isImports := true
	// appends imports to "imports"
	for scanner.Scan() && isImports {
		if strings.Contains(scanner.Text(), "import (") {
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), ")") {
					isImports = false
					break
				} else {
					imports = append(imports, scanner.Text())
				}
			}
		}
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// copy "imports"
	oldImports := make([]string, len(imports))
	copy(oldImports, imports)
	// sort "imports"
	sort.Strings(imports)

	// replace old imports with sorted
	newContents := strings.Replace(string(content), strings.Join(oldImports, "\n"),strings.Join(imports, "\n") , -1)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		log.Fatal(err)
	}
}




//func SortImports(path string) {
//	file, err := os.Open(path)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//	scanner := bufio.NewScanner(file)
//	var imports[] string
//	isImports := true
//	for scanner.Scan() && isImports {
//		if strings.Contains(scanner.Text(), "import (") {
//			for scanner.Scan() {
//				if strings.Contains(scanner.Text(), ")") {
//					isImports = false
//					break
//				} else {
//					imports = append(imports, scanner.Text())
//				}
//			}
//		}
//	}
//
//	for _, s:= range imports {
//		fmt.Println(s)
//	}
//
//	sort.Strings(imports)
//	for _, s:= range imports {
//		fmt.Println(s)
//	}
//
//	newContents := strings.Replace(string(file), "old", "new", -1)
//
//	fmt.Println(newContents)
//
//	err = ioutil.WriteFile(path, []byte(newContents), 0)
//	if err != nil {
//		panic(err)
//	}
//}
