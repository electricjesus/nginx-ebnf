# NGINX EBNF

Simple EBNF Grammar for NGINX

## Usage

```golang
package main

import (
    "log"

    ebnf "github.com/electricjesus/nginx-ebnf"
)

func main () {
    p := ebnf.NewParser(false)

    f, err := os.Open("/etc/nginx/nginx.conf")
    if err != nil {
        log.Fatal(err)
    }

    ast, err := p.Parse(f)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(ast)
}
```
