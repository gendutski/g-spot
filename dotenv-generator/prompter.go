package dotenvgenerator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type Prompter interface {
	PromptString(prompt string) string
	PromptPassword(prompt string) string
}

func newPrompter(reader *bufio.Reader) Prompter {
	return &prompter{reader}
}

type prompter struct {
	reader *bufio.Reader
}

func (p *prompter) PromptString(prompt string) string {
	fmt.Print(prompt)
	result, err := p.reader.ReadString('\n')
	if err != nil {
		fmt.Println("error getString:", err.Error())
		os.Exit(1)
	}
	result = strings.TrimSpace(result)
	return result
}

func (p *prompter) PromptPassword(prompt string) string {
	fmt.Print(prompt)
	result, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Println("error getString:", err.Error())
		os.Exit(1)
	}
	return strings.TrimSpace(string(result))
}
