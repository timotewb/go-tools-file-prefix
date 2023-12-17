package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ncruces/zenity"
)

func main(){
	inDir, err := zenity.SelectFile(
		zenity.Filename(""),
		zenity.Directory(),
		zenity.DisallowEmpty(),
		zenity.Title("Select input directory"),
	)
	if err != nil {
		zenity.Error(
			err.Error(),
			zenity.Title("Error"),
			zenity.ErrorIcon,
		)
		log.Fatal(err)
	}

	dlm, err := zenity.Entry("Enter delimitier",
		zenity.Title("Delimiter"),
	)

	if dlm != "" {
		files, err := os.ReadDir(inDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files{
			newName := strings.TrimSpace(file.Name()[0:strings.Index(file.Name(), dlm)])+filepath.Ext(file.Name())
			e := os.Rename(filepath.Join(inDir, file.Name()), filepath.Join(inDir, newName))
			// Check for Error
			if e != nil {
			log.Fatal(e)
			}
		}
		zenity.Info("Files renamed!",
			zenity.Title("Complete"),
			zenity.InfoIcon,
		)
	} else {
		zenity.Info("No delimiter provided!",
			zenity.Title("Attention"),
			zenity.InfoIcon,
		)
	}
}