package main

import (
	"github.com/mchirico/harvest/configure"
	"github.com/mchirico/harvest/pkg"
	"os/user"
)

func main() {

	a := pkg.App{}
	a.Initilize()
	usr, _ := user.Current()
	file := usr.HomeDir + "/.secretHarvest"
	s, _ := configure.GetSecret(file)
	a.InitSS(&s)

	a.Run("4571", 15, 15)

}
