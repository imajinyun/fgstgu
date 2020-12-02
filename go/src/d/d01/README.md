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
