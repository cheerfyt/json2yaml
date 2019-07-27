### json2yaml

Usage:

```bash
$ git clone github.com/
$ make build
$ 
$ # normal use
$ ./json2yaml -i test.json -o test.yaml
$
$ # read from stdin
$ cat test.json | ./json2yaml -o test.yaml 
$
$ # out to stdin
$ ./json2yaml -i test.json  > test.yaml
```

install to path:

```bash
$ make install
```