package uniq

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Options struct {
	Count      bool
	Repeated   bool
	Unique     bool
	SkipFields int
	SkipChars  int
	IgnoreCase bool
}

func Uniq(r io.Reader, w io.Writer, opt Options) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return
	}
	current := scanner.Text()
	cnt := 1

	for scanner.Scan() {
		next := scanner.Text()
		if LinesIsEqual(&current, &next, &opt) {
			cnt++
		} else {
			PrintLine(w, &current, cnt, &opt)
			cnt = 1
			current = next
		}
	}
	PrintLine(w, &current, cnt, &opt)
}

func LinesIsEqual(current *string, next *string, opt *Options) bool {
	var curr, nxt string

	if opt.SkipFields != 0 || opt.SkipChars != 0 {

		fields := strings.Fields(*current)
		if opt.SkipFields >= len(fields) {
			fields = []string{}
		} else {
			fields = fields[opt.SkipFields:]
		}
		curr = strings.Join(fields, " ")

		fields = strings.Fields(*next)
		if opt.SkipFields >= len(fields) {
			fields = []string{}
		} else {
			fields = fields[opt.SkipFields:]
		}
		nxt = strings.Join(fields, " ")

		if opt.SkipChars >= len(curr) {
			curr = ""
		} else {
			curr = curr[opt.SkipChars:]
		}

		if opt.SkipChars >= len(nxt) {
			nxt = ""
		} else {
			nxt = nxt[opt.SkipChars:]
		}

	} else {
		curr = *current
		nxt = *next
	}

	if opt.IgnoreCase {
		curr = strings.ToLower(curr)
		nxt = strings.ToLower(nxt)
	}

	return curr == nxt
}

func PrintLine(w io.Writer, current *string, cnt int, opt *Options) {
	if opt.Count {
		fmt.Fprintf(w, "%d %s\n", cnt, *current)
	} else if opt.Repeated && cnt != 1 {
		fmt.Fprintln(w, *current)
	} else if opt.Unique && cnt == 1 {
		fmt.Fprintln(w, *current)
	} else if !opt.Count && !opt.Repeated && !opt.Unique {
		fmt.Fprintln(w, *current)
	}
}


