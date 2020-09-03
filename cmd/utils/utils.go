package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/eiannone/keyboard"
	"golang.org/x/crypto/bcrypt"
)

// ReadInput is ...
func ReadInput() string {
	var answer string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer = scanner.Text()
	}
	return answer
}

// FileExists is ...
func FileExists(folderPath string) bool {
	isExists := false
	_, err := os.Stat(folderPath)
	if !os.IsNotExist(err) {
		isExists = true
	}
	return isExists
}

// ParseFile is ...
func ParseFile(fileName string, vars interface{}) string {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return processTemplate(tmpl, vars)
}

func processTemplate(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer
	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		log.Fatal(err)
	}
	return tmplBytes.String()
}

// WriteToFile is ...
func WriteToFile(content string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// HashPass is ...
func HashPass(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	rawpw := string(hash)
	replacedpw := strings.ReplaceAll(rawpw, "$", "$$")
	return replacedpw
}

// ExecShellCommand is ...
func ExecShellCommand(command string, hideOutput bool) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if hideOutput {
		reader, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(reader)
		go func() {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}()
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error: %v", stderr.String())
	}
}

// ExitOnKeyStroke is ...
func ExitOnKeyStroke(menu func()) {
loop:
	for {
		fmt.Print("----------------------------------------------\n")
		fmt.Println("Press q to exit!")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		switch char {
		case 'q':
			menu()
			break loop
		}
	}
}
