package main

import (
	"fmt"
	"io/ioutil"
	"os"

	billy "github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {

	// creating in-memory storage & filesystem
	var storage *memory.Storage
	var folders billy.Filesystem

	// initializing in-memory storage & filesystem
	storage = memory.NewStorage()
	folders = memfs.New()

	// set authentication
	auth := &http.BasicAuth{
		// Username: "username",
		// Password: "your password",
	}

	// set repo URL
	repo := "https://gitlab.alm.poste.it/feu/feu-configuration-cli"

	// cloning repo inside in-memory storage
	r, err := git.Clone(storage, folders, &git.CloneOptions{
		URL:  repo,
		Auth: auth,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// reading in-memory file
	file, err := folders.Open("README.md") // file's type is billy.File,
	if err != nil {                        // same as io.reader file
		fmt.Println(err)
		os.Exit(1)
	}

	// converting to []byte
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// saving file locally
	err = ioutil.WriteFile("README.md", fileContent, 0777)
	if err != nil {
		fmt.Println(err)
	}

	// do nothing, print response to execute app
	fmt.Println(r)
}

// WORKFLOW CLI

// 1 - Richiesta dati di accesso all'utente via UI

// 2 - Cloning del repository (branch master) in-memory

// 3 - Salvataggio in locale del file FCR desiderato

// 4 - Entrare nel main menu per acquisire le modifiche dell'utente

// 6 - Quando seleziona export, creazione nuovo branch per modifiche e checkout nel branch

// 7 - Aggiunta in-memory del file modificato correttamente

// 8 - Commit & push

// 9 - Creare una merge-request con le API GitLab
