package ebnf_test

import (
	"encoding/json"
	"strings"
	"testing"

	ebnf "github.com/electricjesus/nginx-ebnf"
)

func TestParse(t *testing.T) {
	conf := `
	{
		# a comment
		# comment2
		http {
			# another comment
			server {
				listen 127.0.0.1;
			}
		}
	}`

	p := ebnf.NewParser(false)

	ast, err := p.Parse(strings.NewReader(conf))
	if err != nil {
		t.Error(err)
	}
	b, err := json.Marshal(ast.Children)
	if err != nil {
		t.Error(err)
	}

	expected := `[{"directive":"http"},{"children":[{"directive":"server"},{"children":[{"directive":"listen","args":["127.0.0.1"]}]}]}]`
	if expected != string(b) {
		t.Errorf("Result not equal: %v", string(b))
	}
}
