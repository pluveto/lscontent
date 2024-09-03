package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	ignore "github.com/sabhiram/go-gitignore"
)

var (
	recursive       bool
	useGitIgnore    bool
	customFormat    string
	copyToClipboard bool
)

type FileInfo struct {
	Path    string
	Content string
	Suffix  string
}

func main() {
	flag.BoolVar(&recursive, "r", false, "Also print subdirectories")
	flag.BoolVar(&useGitIgnore, "i", false, "Enable .gitignore rules")
	flag.StringVar(&customFormat, "f", "", "Use custom format")
	flag.BoolVar(&copyToClipboard, "c", false, "Copy output to clipboard")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: lscontent [-r] [-i] [-f format] [-c] <DIR>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	dir := flag.Arg(0)
	files, err := listFiles(dir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	var buffer bytes.Buffer
	for _, file := range files {
		output := formatFileContent(file)
		fmt.Print(output)
		if copyToClipboard {
			buffer.WriteString(output)
		}
	}

	if copyToClipboard {
		err := clipboard.WriteAll(buffer.String())
		if err != nil {
			fmt.Printf("Error copying to clipboard: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Output copied to clipboard.")
	}
}

func listFiles(dir string) ([]FileInfo, error) {
	var files []FileInfo
	var ignoreFile *ignore.GitIgnore
	var err error

	if useGitIgnore {
		ignoreFile, err = ignore.CompileIgnoreFile(filepath.Join(dir, ".gitignore"))
		if err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("error compiling .gitignore: %w", err)
		}
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if !recursive && path != dir {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		if useGitIgnore && ignoreFile != nil && ignoreFile.MatchesPath(relPath) {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		suffix := filepath.Ext(path)
		if suffix != "" {
			suffix = suffix[1:] // Remove the leading dot
		}

		files = append(files, FileInfo{
			Path:    relPath,
			Content: string(content),
			Suffix:  suffix,
		})

		return nil
	})

	return files, err
}

func formatFileContent(file FileInfo) string {
	format := customFormat
	if format == "" {
		format = "*{path}*:\n```{suffix}\n{content}\n```\n\n"
	}

	output := format
	output = strings.ReplaceAll(output, "{path}", file.Path)
	output = strings.ReplaceAll(output, "{suffix}", file.Suffix)
	output = strings.ReplaceAll(output, "{content}", file.Content)
	output = strings.ReplaceAll(output, "{linebreak}", "\n")

	return output
}
