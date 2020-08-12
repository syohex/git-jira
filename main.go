package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pkg/browser"
)

func main() {
	os.Exit(_main())
}

func gitBranch() (string, error) {
	var buf bytes.Buffer
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}

func jiraBaseURL() (string, error) {
	var buf bytes.Buffer
	cmd := exec.Command("git", "config", "jira.baseURL")
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}

func jiraProjects() ([]string, error) {
	var buf bytes.Buffer
	cmd := exec.Command("git", "config", "jira.projects")
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(buf.String()), ","), nil
}

func _main() int {
	showURL := flag.Bool("url", false, "Show only URL")

	branch, err := gitBranch()
	if err != nil {
		fmt.Println("Could not get git branch name. Please check")
		return 1
	}

	baseURL, err := jiraBaseURL()
	if err != nil {
		fmt.Println("Please set jira.baseURL in git config")
		return 1
	}

	projects, err := jiraProjects()
	if err != nil {
		fmt.Println("Please set jira.projects in git config")
		return 1
	}

	matched := false
	for _, project := range projects {
		// Ticket name is like SOMEPROJ-1234
		re := regexp.MustCompile(fmt.Sprintf(`%s-\d+`, project))
		ticket := re.FindString(branch)
		if ticket == "" {
			continue
		}

		matched = true

		// URL is like https://jira.example.com/browse/SOMEPROJ-1234
		url := fmt.Sprintf("%s/%s", baseURL, ticket)
		if *showURL {
			fmt.Println(url)
		} else {
			if err := browser.OpenURL(url); err != nil {
				fmt.Printf("Could not open URL: %s", url)
			}
		}

		break
	}

	if !matched {
		fmt.Printf("%s is not matched any projects(%s)\n", branch, projects)
		return 1
	}

	return 0
}
