package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var logPattern = regexp.MustCompile(`^(?P<Indent> *)(?P<Direction>>|<) [0-9]+:[0-9]+:[0-9]+: parseRule (?P<Function>[A-Za-z0-9_]+) \[U\+[0-9A-F]+(?: '.')?\] ?(?P<TS>[0-9]+)?`)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	calls := make(map[string]*callCount)

	scanner := bufio.NewScanner(file)
	scan(scanner, calls, "", 0)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rules := make(map[string]*ruleCollection)
	var rcs []*ruleCollection
	for path, count := range calls {
		//fmt.Printf("%s: %d\n", path, count.count)
		rule, ok := rules[count.rule]
		if !ok {
			rule = &ruleCollection{rule: count.rule, calls: make(map[string]*callCount)}
			rules[count.rule] = rule
			rcs = append(rcs, rule)
		}
		rule.count += count.count
		rule.elapsed += count.elapsed
		cc, ok := rule.calls[path]
		if !ok {
			cc = &callCount{}
			rule.calls[path] = cc
		}
		cc.count += count.count
		cc.elapsed += count.elapsed
	}

	slices.SortStableFunc(rcs, func(a, b *ruleCollection) int {
		if a.elapsed < b.elapsed {
			return -1
		} else if a.elapsed > b.elapsed {
			return 1
		}
		return 0
	})

	for _, rc := range rcs {
		fmt.Printf("Rule: %s: Count: %d\n", rc.rule, rc.count)
		var calls []*callCount
		totalCount := 0
		var totalElapsed time.Duration
		for call, count := range rc.calls {
			calls = append(calls, &callCount{rule: call, count: count.count, elapsed: count.elapsed})
			totalCount += count.count
			totalElapsed += count.elapsed

		}
		fmt.Printf("\tPaths: %d: Count: %d Elapsed: %s (avg. %s\n", len(rc.calls), totalCount, totalElapsed.String(), (totalElapsed / time.Duration(totalCount)).String())
		slices.SortStableFunc(calls, func(a, b *callCount) int {
			if a.elapsed < b.elapsed {
				return 1
			} else if a.elapsed > b.elapsed {
				return -1
			}
			return 0
		})
		for i := 0; i < len(calls) && i < 5; i++ {
			cc := calls[i]
			fmt.Printf("\t%s: %d %s (avg. %s)\n", cc.rule, cc.count, cc.elapsed.String(), (cc.elapsed / time.Duration(cc.count)).String())
		}
	}
}

type callCount struct {
	rule    string
	count   int
	elapsed time.Duration
}

type ruleCollection struct {
	rule    string
	count   int
	elapsed time.Duration
	calls   map[string]*callCount
}

func scan(scanner *bufio.Scanner, calls map[string]*callCount, path string, start int64) {
	for scanner.Scan() {
		text := scanner.Text()
		matches := logPattern.FindStringSubmatch(text)
		if matches == nil {
			if strings.Contains(text, "parseRule ") {
				panic(fmt.Errorf("unrecognized pattern: %s", text))
			}
			continue
		}
		ts, _ := strconv.ParseInt(matches[4], 10, 64)
		rule := matches[3]
		switch matches[2] {
		case ">":
			scan(scanner, calls, path+"/"+rule, ts)
		case "<":
			if ts < start {
				ts += 1_000_000_000
			}
			cc, ok := calls[path]
			if !ok {
				cc = &callCount{rule: rule, count: 1, elapsed: time.Duration(ts - start)}
				calls[path] = cc
			} else {
				cc.count++
				cc.elapsed += time.Duration(ts - start)
			}
			return
		}
	}
}
