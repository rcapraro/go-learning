package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FindReplaceFile replace old string with new string on the src file
func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer dstFile.Close()

	//only words
	old = old + " "
	new = new + " "

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	
	lineNumber := 1
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if (found) {
			occ += o
			lines = append(lines, lineNumber)
		}
		lineNumber ++
		fmt.Fprintf(writer, res+"\n")
	}
	return occ, lines, nil
}

// ProcessLine process each line of the file
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

func main() {
	occ, lines, err := FindReplaceFile("./python.txt", "go.txt", "Python", "Go")
	if err != nil {
		fmt.Printf("error during find replace: %v", err)
		return
	}
	fmt.Println("===== SUMMARY =====")
	defer fmt.Println("===== END OF SUMMARY =====")
	fmt.Printf("Number of occurrences of Python: %d\n", occ)
	fmt.Printf("Number of lines: %d\n", len(lines))
	fmt.Print("Lines: [ ")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len - 1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")
}
