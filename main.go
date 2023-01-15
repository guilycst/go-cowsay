package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	maxCharPerline = 21
	flag_borg      = flag.Bool("b", false, "Borg")
	flag_dead      = flag.Bool("d", false, "Dead")
	flag_greedy    = flag.Bool("g", false, "Greedy")
	flag_paranoid  = flag.Bool("p", false, "Paranoid")
	flag_stoned    = flag.Bool("s", false, "Stoned")
	flag_tired     = flag.Bool("t", false, "Tired")
	flag_wired     = flag.Bool("w", false, "Wired")
	flag_youthful  = flag.Bool("y", false, "Youthful")
	flag_eyes      = flag.String("e", "", "Eye string")
	flag_file      = flag.String("f", "cow", "cow or tux")
)

func main() {
	flag.Parse()
	args := os.Args[1:]
	argsLen := len(args)
	hasArgs := argsLen > 0
	if !hasArgs {
		cowsay(nil)
		return
	}

	text := args[len(args)-1]
	text = fill(text)
	wrappedText := wrap(text)

	cowsay(wrappedText)
}

func cowsay(wrappedText []string) {
	t, err := template.ParseFiles("./templates/" + *flag_file + ".tmpl")
	if err != nil {
		log.Print(err)
	}

	t.Execute(os.Stdout, map[string]interface{}{
		"text":           wrappedText,
		"borg":           *flag_borg,
		"dead":           *flag_dead,
		"greedy":         *flag_greedy,
		"paranoid":       *flag_paranoid,
		"stoned":         *flag_stoned,
		"tired":          *flag_tired,
		"wired":          *flag_wired,
		"youthful":       *flag_youthful,
		"has_eye_string": len(*flag_eyes) > 0,
		"eye_string":     *flag_eyes,
	})
}

func wrap(text string) []string {
	textLen := len(text)
	if textLen <= maxCharPerline {
		return []string{text}
	}

	splitText := strings.Split(text, " ")
	hasWhiteSpaces := len(splitText) > 0
	if hasWhiteSpaces {
		lines := []string{""}
		currentLine := 0
		for _, v := range splitText {
			newLineText := lines[currentLine] + v + " "
			accept := len(newLineText) < maxCharPerline
			if accept {
				lines[currentLine] = newLineText
				continue
			}
			currentLine++
			lines = append(lines, v+" ")
		}

		for i, v := range lines {
			lines[i] = fill(v)
		}
		return lines
	}

	return []string{text}
}

func fill(text string) string {
	textLen := len(text)
	for textLen < maxCharPerline {
		text = " " + text
		textLen = len(text)

		if textLen == maxCharPerline {
			break
		}

		text = text + " "
		textLen = len(text)
	}
	return text
}
