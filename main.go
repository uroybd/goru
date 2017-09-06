package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var lines []string
	if len(os.Args) > 1 && os.Args[1] == "-i" {
		lines = append(lines, strings.TrimSpace(strings.Join(os.Args[2:], " ")))
	} else {
		scn := bufio.NewScanner(os.Stdin)
		for scn.Scan() {
			line := scn.Text()
			if len(line) == 1 {
				if line[0] == '\x1D' {
					break
				}
			}
			lines = append(lines, line)
		}
	}
	tmp := make([]string, len(lines))
	copy(tmp, lines)
	sort.Strings(tmp)
	largest := tmp[len(tmp)-1]
	length := len(largest)
	formattedCommentSlice := make([]string, len(lines))
	for _, line := range lines {
		spaceNeeded := strings.Repeat(" ", length-len(line))
		formattedCommentSlice = append(formattedCommentSlice, fmt.Sprintf("< %s %s>", line, spaceNeeded))
	}

	formattedComment := strings.TrimSpace(strings.Join(formattedCommentSlice, "\n"))
	border := strings.Repeat("-", length+2)
	bubble := fmt.Sprintf(" %s\n%s\n %s", border, formattedComment, border)
	cow := `        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`

	fmt.Printf("%s\n%s", bubble, cow)
}
