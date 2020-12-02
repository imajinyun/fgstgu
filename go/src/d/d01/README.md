# 词法和语法分析

## 词法分析

```bash
$ lex example.l
$ cat lex.yy.c
```

```bash
$ cc lex.yy.c -o example -ll
$ cat main.go | ./example
PACKAGE  IDENT

IMPORT  QUOTE IDENT QUOTE

IDENT  IDENT LPAREN RPAREN  LBRACE
  IDENT DOT IDENT LPAREN QUOTE IDENT QUOTE RPAREN
RBRACE
```

Go 语言的词法解析是通过 [src/cmd/compile/internal/syntax/scanner.go](https://github.com/golang/go/blob/dev.boringcrypto.go1.15/src/cmd/compile/internal/syntax/scanner.go) 文件中的 [src/cmd/compile/internal/syntax/scanner.scanner](https://github.com/golang/go/blob/dev.boringcrypto.go1.15/src/cmd/compile/internal/syntax/scanner.go#L30~L44) 结构体实现的。
