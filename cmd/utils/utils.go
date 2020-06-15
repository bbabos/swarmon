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
		panic(err)
	}
	return processTemplate(tmpl, vars)
}

func processTemplate(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer
	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}

// WriteToFile is ...
func WriteToFile(content string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

// HashPass is ...
func HashPass(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	rawpw := string(hash)
	replacedpw := strings.ReplaceAll(rawpw, "$", "$$")
	return replacedpw
}

// ExecShellCommand is ...
func ExecShellCommand(command string) {
	args := strings.Fields(command)
	cmd := exec.Command(args[0], args[1:]...)
	reader, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
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
		fmt.Printf("Error: %v", stderr.String())
	}
}

// ExitOnKeyStroke is ...
func ExitOnKeyStroke(menu func()) {
loop:
	for {
		fmt.Println("-------------------------------")
		fmt.Println("Press q to exit!")
		fmt.Println("-------------------------------")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		switch char {
		case 'q':
			menu()
			break loop
		}
	}
}
