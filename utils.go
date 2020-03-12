package main

import (
	"bufio"
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

func readInput(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var answer string

	if scanner.Scan() {
		answer = scanner.Text()
	}
	return answer
}

func gitClone(repoURL string) {
	_, err := git.PlainClone("./tmp", false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	try(err)
}

func processTemplate(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	try(err)

	return tmplBytes.String()
}

func parseFile(fileName string, vars interface{}) string {
	tmpl, err := template.ParseFiles(fileName)
	try(err)

	return processTemplate(tmpl, vars)
}

func writeToFile(content string, filename string) {
	f, err := os.Create(filename)
	try(err)

	_, err = f.WriteString(content)
	try(err)

	err = f.Close()
	try(err)
}

func (i *input) execute(action func(text string) string) {
	action(i.Answer)
}
