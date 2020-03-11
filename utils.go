package main

import (
	"bytes"
	"os"
	"text/template"

	"gopkg.in/src-d/go-git.v4"
)

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func gitClone(repoURL string) {
	_, err := git.PlainClone("./tmp", false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	try(err)
}

func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	try(err)

	return tmplBytes.String()
}

func processFile(fileName string, vars interface{}) string {
	tmpl, err := template.ParseFiles(fileName)
	try(err)

	return process(tmpl, vars)
}
