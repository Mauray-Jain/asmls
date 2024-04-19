package parse
// An idea:
// run this whole lexer cum parser in one goroutine
// all incremental data will be passed to this lexer via channel
// it will give tokens to another goroutine via channel which will generate
// (this go routine shld also maintain a whole list of possible completions
// at all times including labels etc.)
// a list of possible completion options and pass this list
// via channel to main goroutine (the lsp server) which will display it

// parse whole doc on didOpen notif
// then parse only current lines on didChange

import "fmt"

type itemType int
type stateFn func(*lexer) stateFn

const (
	itemErr itemType = iota
	itemDriective
	itemConstant
	itemRegister
	itemInstruction
	itemLabel
	itemSymbol
	itemEOF
)

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemErr:
		return i.val
	}
	if len(i.val) > 15 {
		return fmt.Sprintf("%15q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

type lexer struct {
	name string
	input string

}

func run() {
	for state := startState; state != nil; {
		state = state(lexer)
	}
}
