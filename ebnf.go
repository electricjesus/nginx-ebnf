package ebnf

import (
	"io"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
)

type Parser struct {
	Options []participle.Option
	Ctx     *participle.Parser
}

// Block - Block recursive structure
type Block struct {
	// TODO: Pos       lexer.Position
	Children  []*Block `parser:"(\"{\" @@* \"}\")?" json:"children,omitempty"`
	Directive string   `parser:"@Ident?" json:"directive,omitempty"`
	Args      []string `parser:"(@(Ident | Number | String | KeyVal)* \";\")?" json:"args,omitempty"`
	Comment   string   `parser:"(@Comment)?" json:"comment,omitempty"`
}

var (
	notation = `
	Ident = (alpha | "_") { "_" | alpha | digit } 					.
	Comment = ("#") { "\u0000"…"\uffff"-"\n" } ("\n")				.
	Number = ("." | digit) {"." | digit} 							.
	String = ("\"") { "\u0000"…"\uffff"-"\n"-"\"" } ("\"")			.
	KeyVal = { Ident } ("=") { alpha | digit }						.
	Whitespace = " " | "\t" | "\n" | "\r" 							.
	Punct = "!"…"/" | ":"…"@" | "["…` + "\"`\"" + ` | "{"…"~" 		.
	alpha = "a"…"z" | "A"…"Z" 										.
	digit = "0"…"9" 												.
	`
	lexicon = lexer.Must(ebnf.New(notation))
)

func NewParser(comments bool) (p *Parser) {
	p = &Parser{Options: []participle.Option{
		participle.Lexer(lexicon),
		participle.Elide("Whitespace"),
		participle.Unquote("String", "Comment"),
	}}

	if !comments {
		p.Options = append(p.Options, participle.Elide("Comment"))
	}
	p.Ctx = participle.MustBuild(&Block{}, p.Options...)
	return p
}

func (p *Parser) Parse(r io.Reader) (ast *Block, err error) {
	ast = &Block{}
	err = p.Ctx.Parse(r, ast)
	return ast, err
}
