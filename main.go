package main

import (
	"fmt"
	"io/ioutil"
	"log"
	//"os"
)

//see
//http://golang.org/src/pkg/text/template/parse/lex.go

/*
	Top Level

	First Pass:
	- split code into tokens
	- find functions signatures and code blocks

	Second Pass:
	- parse statements
	-
*/

var (
	BGN_BLK   = 0 // {
	END_BLK   = 1 // }
	BGN_PAREN = 2
	END_PAREN = 3

	VERB_FUNC = 4 // func keyword
	STATEMENT = 5 // statement until ;
)

//depth
//tokens are trees
type Token struct {
	File   string //use int
	Line   int    //line number
	Offset int    //offset
}

type Parser struct {
	TokenList []Token
}

func (self *Parser) Next() {

}

/*

*/

type Cursor struct {
	Idx  int
	Line int
}

//
func extractBlock(in []byte, cursor Cursor) (string, []byte) {
	if in[cursor.Idx] != byte('{') {
		log.Panic()
	}
	d := 1
	max := len(in)

	for i := cursor.Idx; i < max; i++ {
		if in[i] == byte('{') {
			d++ //increment depth
			continue
		}
		if in[i] == byte('}') {
			d--
			if d == 0 {
				break //end bracket
			}
		}

		if in[i] == byte('\n') {
			cursor.Line++
		}
		continue
	}
	if d != 0 {
		log.Panic("error unterminated block")
	}

	//do something
	return "", nil
}

func getNextTok(in []byte, cursor *Cursor) (string, []byte) {

	for i, c := range in {
		if c == byte(' ') {
			return string(in[:i]), in[i:]
		}
	}

	return "", nil
}

/*
func getNextTok(in []byte) (string, []byte) {

	for i, c := range in {
		if c == byte(' ') {
			return string(in[:i]), in[i:]
		}
	}

	return "", nil
}
*/

func main() {

	data, err := ioutil.ReadFile("testfile.c")
	if err != nil {
		log.Fatal(err)
	}

	var cursor Cursor

	fmt.Printf("read %d bytes: %q\n", len(data), data)
	tok, data := getNextTok(data, &cursor)
	fmt.Printf("split: %s, %s \n", tok, data)

	switch {
	case tok == "import":
		//start import block
		fmt.Printf("tok, start import= %s \n", tok)
	case tok == "var":
		//start var definition
		fmt.Printf("tok, start var= %s \n", tok)
	case tok == "func":
		//start function definit
		fmt.Printf("tok, start func= %s \n", tok)
	case tok == "type":
		//start struct definition
		fmt.Printf("tok, start type= %s \n", tok)
	default:
		fmt.Printf("error: line, char ... expecting valid symbol \n")
	}

}
