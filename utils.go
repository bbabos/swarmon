package main

import (
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func gitClone(repoURL string) {
	_, err := git.PlainClone("./tmp", false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	try(err)
}

func try(err error) {
	if err != nil {
		panic(err)
	}
}
