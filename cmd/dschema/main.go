package main

import (
	"fmt"

	"gitgub.com/marzdog/dschema/internal/model"
)

func main() {
	fmt.Println("Howdy")

	meth := model.AxMethod{Name: "fnItemId", Text: "fnItemId", Source: "source code"}

	fmt.Println(meth)

	fmt.Println("Done")
}
