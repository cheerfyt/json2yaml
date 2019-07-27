### json2yaml

install:
```bash
$ go get -u -v github.com/cheerfyt/json2yaml
```

build-install:

```bash
$ git clone git@github.com:cheerfyt/json2yaml.git
$ cd json2yaml && make build
$ # or install to /usr/local/bin (cd json2yaml && make install)
```



Usage:

```bash
$ # normal use
$ ./json2yaml -i test.json -o test.yaml
$
$ # read from stdin
$ cat test.json | ./json2yaml -o test.yaml 
$
$ # out to stdin
$ ./json2yaml -i test.json  > test.yaml
```
