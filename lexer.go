package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	LBrace = "["
	RBrace = "]"

	OpEq = "EQ"
	OpLe = "LE"
	OpGr = "GR"
)

type LexerCtx struct {
	root    *Expr
	current *Expr
}

func (lexer *LexerCtx) Root() *Expr {
	return lexer.root
}

func (lexer *LexerCtx) Parse(tokens []string) {
	for i, token := range tokens {
		switch token {
		case LBrace:
			exp := new(Expr)

			if lexer.root == nil {
				lexer.root = exp
				lexer.current = exp
				continue
			}

			lexer.current.Child = exp
			lexer.current = exp

			lexer.Parse(tokens[i+1:])

		case RBrace:
			return

		case OpEq, OpLe, OpGr:
			lexer.current.Operator = token

		default:
			if isInteger(token) {
				val, err := strconv.ParseInt(token, 10, 64)
				if err != nil {
					log.Fatal("invalid number: " + token)
				}

				lexer.current.N = val
				lexer.current.Sets = make([][]int64, 0, val)
				continue
			}

			if isFile(token) {
				lexer.current.Sets = append(lexer.current.Sets, extractSetFromFile(token))
				continue
			}

			log.Fatal("unknown token: " + token)
		}
	}
}

func isFile(token string) bool {
	_, err := os.Stat(token)
	return err == nil
}

func isInteger(token string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(token)
}

func extractSetFromFile(filename string) []int64 {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	var set []int64
	for _, line := range lines {
		if !isInteger(line) {
			continue
		}

		val, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal("invalid number in file: " + line)
		}

		set = append(set, val)
	}
	return set
}
