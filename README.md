<p align="center"><img src="./gopher.png" width="360"></p>
<p align="center">
<a href="https://goreportcard.com/badge/github.com/hysios/gonginx"><img src="https://goreportcard.com/badge/github.com/hysios/gonginx" alt="Report Card" /></a>
<a href="https://github.com/tufanbarisyildirim/gonging/actions"><img src="https://github.com/hysios/gonginx/workflows/Go/badge.svg" alt="Actions Status" /></a>
</p>

# Gonginx
TBH, I would like to rewrite the parser next time I need it again :)) but it still does its job. 

Gonginx is an Nginx configuration parser helps you to parse, edit, regenerate your nginx config files in your go applications. It makes managing your balancer configurations easier. 

## Basic grammar of an nginx config file
```yacc

%token Keyword Variable BlockStart BlockEnd Semicolon Regex

%%

config      :  /* empty */ 
            | config directives
            ;
block       : BlockStart directives BlockEnd
            ;
directives  : directives directive
            ;
directive   : Keyword [parameters] (semicolon|block)
            ;
parameters  : parameters keyword
            ;
keyword     : Keyword 
            | Variable 
            | Regex
            ;
```

## API 

## Core Components
- ### [Parser](/parser/parser.go) 
  Parser is the main package that analyzes and turns nginx structred files into objects. It basically has 3 libraries, `lexer` explodes it into `token`s and `parser` turns tokens into config objects which are in their own package, 
- ### [Config](/config/config.go)
  Config package is representation of any context, directive or their parameters in golang. So basically they are models and also AST
- ### [Dumper](/dumper/dumper.go)
  Dumper is the package that holds styling configuration only. 

## Examples
- [Formatting](/examples/formatting/main.go)
- [Adding a Server to upstream block](/examples/adding-server/main.go)
- [add-custom-directive](/examples/add-custom-directive/main.go)
- [dump-nginx-config](/examples/dump-nginx-config/main.go)
- [update-directive](/examples/update-directive/main.go)
- [update-server-listen-port](/examples/update-server-listen-port/main.go)

### [Examples and Library Reference](/GUIDE.md)
#### TODO
- [x]  associate comments with config objects to print them on config generation and make it configurable with `dumper.Style`
- [x]  move any context wrapper into their own file (remove from parser)
- [x]  Parse included files recusively, keep relative path on load, save all in a related structure and make that optional in dumper.Style
- [ ]  Implement specific searches, like finding servers by server_name (domain) or any upstream by target etc.
- [ ]  add more examples
- [x]  link the parent directive to any directive for easier manipulation

## Limitations
There is no known limitations yet. PRs are more than welcome if you want to implement a specific directive / block, please read [Contributing](CONTRIBUTING.md) before your first PR.

## License
[MIT License](LICENSE)
