package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/src-d/go-git.v4"
)

func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var answer string

	if scanner.Scan() {
		answer = scanner.Text()
	}
	return answer
}

func gitClone(repoURL string, folderPath string) {
	if !fileExists(folderPath) {
		_, err := git.PlainClone(folderPath, false, &git.CloneOptions{
			URL: repoURL,
		})
		try(err)
	}
}

func fileExists(folderPath string) bool {
	isExists := false
	_, err := os.Stat(folderPath)
	if !os.IsNotExist(err) {
		isExists = true
	}
	return isExists
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

func hashPass(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	try(err)

	return string(hash)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func saveConfig(folderPath string) {
	file, _ := json.MarshalIndent(p, "", " ")
	_ = ioutil.WriteFile(folderPath, file, 0644)
}

func loadConfig(filePath string) {
	file, _ := ioutil.ReadFile(filePath)
	_ = json.Unmarshal([]byte(file), &p)
}

func execCommand(command string) {
	args := strings.Fields(command)

	cmd := exec.Command(args[0], args[1:len(args)]...)
	reader, err := cmd.StdoutPipe()
	try(err)

	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
