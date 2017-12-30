// +build ignore

package main

import (
	"fmt"
	"github.com/shurcooL/vfsgen"
	"net/http"
)

func main() {
	var fs http.FileSystem = http.Dir("../front/dist")

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "main",
		BuildTags:    "!dev",
		VariableName: "assets",
	})
	if err != nil {
		fmt.Println("Cannot build assets", err)
	}
}
