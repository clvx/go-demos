package src

import (
	"os"
	"path/filepath"
)

//GetPath returns executable path
func GetPath() (p string) {
	if len(os.Args) > 1{
		abspath, err := filepath.Abs(os.Args[1])
		Check(err)
		return filepath.Join(abspath, "dump")
	} 
	abspath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	Check(err)
	return filepath.Join(abspath, "dump")
}

//Check error if it's nil
func Check(e error){
	if e !=  nil {
		os.Exit(1)			
	}
}
