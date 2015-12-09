# tt

tt (template [cli] tool) is a Go text/html template command line tool

usage:

```
tt <(echo 'hello {{.data}}') <(echo '{"data": "world"}')
tt -t h <(echo 'data is {{.a}}') <(echo '{"a": "<tag>"}')
```
