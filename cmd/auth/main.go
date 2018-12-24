package main

import (
"github.com/mchirico/harvest/pkg"
)

func main() {

	a := pkg.App{}
	a.Initilize()
	a.Run("4571", 15, 15)

}

