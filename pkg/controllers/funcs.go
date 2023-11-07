package controllers

import (
	"fmt"
	"os"
	path2 "path"
)

func StringsContain(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func IntsContain(array []int, value int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

var Path = execPath()

type path string

func execPath() path {
	ph, _ := os.Executable()
	//if _, fl := path2.Split(ph); fl != "bff" {
	//	return "/home/marat/variag/prmObjects/"
	//}
	return path(path2.Dir(ph) + "/")
}

func (p *path) String() string {
	return fmt.Sprintf("%s", *p)
}

func (p *path) ToFile(filepath string) string {
	return fmt.Sprintf("%s%s", *p, filepath)
}

func (p *path) ToDir(dirpath string) string {
	return fmt.Sprintf("%s%s", *p, dirpath)
}
