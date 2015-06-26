# JSON templates
A simple [JSON] template rendering tool.
Templates are written as [GO templates][GOT].

## Usage
Simple rendering using a template:

```none
$ cat some.json | jt -template some.tmpl
```

## Building
#### Pre-requisites
- go
- coreutils
- make

#### Compile and install to default location (/usr/bin on Linux)

```none
$ make
$ sudo make install
```

#### Install to custom location

```none
$ make install prefix=/tmp/custom
```

[JSON]: http://json.org
[GOT]: http://golang.org/pkg/text/template/#pkg-overview
