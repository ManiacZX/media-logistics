package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	command := flag.String("command", "help", "Command to execute")
	directory := flag.String("directory", ".", "Directory to execute command in")
	title := flag.String("title", "", "Title of the series")
	season := flag.Int("season", 1, "Season of the series")
	episode := flag.Int("episode", 1, "Episode of the series")
	pattern := flag.String("pattern", "", "Pattern to match")
	flag.Parse()

	switch *command {
	case "rename":
		rename(*directory, *title, *season, *episode, *pattern)
	}
}

func rename(directory string, title string, season int, episode int, pattern string) error {
	// files := getFiles(directory, pattern)
	files, _ := getFiles(directory, pattern)

	extension, _ := regexp.Compile(`\.[a-zA-Z0-9]+$`)

	for _, file := range files {
		target := fmt.Sprintf("%s_S%02d_E%02d%s", title, season, episode, extension.FindString(file))
		fmt.Println(file, "rename to", target)
		err := os.Rename(fmt.Sprint(directory, file), fmt.Sprint(directory, target))
		if err != nil {
			fmt.Println("Error renaming file", err)
		}
		episode++
	}

	return nil
}

func getFiles(directory string, pattern string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory", err)
		return nil, err
	}

	match, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling pattern", err)
	}

	fileNames := []string{}
	for _, file := range files {
		if pattern == "" || match.MatchString(file.Name()) {
			fileNames = append(fileNames, file.Name())
		}
	}

	// sort files by name or pattern
	sort.Slice(fileNames, func(i, j int) bool {
		if pattern != "" {
			return match.FindString(fileNames[i]) < match.FindString(fileNames[j])
		}
		return fileNames[i] < fileNames[j]
	})

	return fileNames, nil
}
