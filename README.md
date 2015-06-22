# JSON templates
A simple [JSON] template rendering tool.
Templates are written as [GO templates][GOT].

## Usage
Simple rendering using a template:

```none
$ cat some.json | jt -template some.tmpl
```

## Building
** Pre-requisites **
- go
- makeutils

** Standard binary **

```none
$ make
```

** Statically linked binary **

```none
$ make static
```

[JSON]: http://json.org
[GOT]: http://golang.org/pkg/text/template/#pkg-overview
