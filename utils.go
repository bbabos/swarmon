package main

import (
	"bufio"
	"bytes"
	"os"
	"text/template"

	"golang.org/x/crypto/bcrypt"
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

func gitClone(repoURL string, folderPath string) {
	_, err := git.PlainClone(folderPath, false, &git.CloneOptions{
		URL: repoURL,
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

func hashPass(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	try(err)

	return string(hash)
}
