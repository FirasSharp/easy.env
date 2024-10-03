package main

import (
	easyenv "github.com/FriscPlusPlus/easy.env/pkg/easyenvlib"
)

func main() {
	easy := easyenv.NewEasyEnv()
	easy.CreateNewDB("mydb.db")
	prj, err := easy.AddProject("mmmm", "tes/data/")
	if err != nil {
		print("123")
	}
	prj, err = easy.GetProject(prj.GetProjectID())
	if err != nil {
		print("123")
	}
}
