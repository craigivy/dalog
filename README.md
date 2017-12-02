# dalog
Logger abstraction allowing JSON via Zap logger and string via Go's logger.  Also supports a key value context.

## See it in action
Running ```make``` will compile, lint, vet and run tests
dalog contains a lame test that is really just sample usage code.

### JSON

```go
os.Setenv("DALOG", "ZAP")
dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
```
```json
{"level":"info","ts":1512200608.536454,"caller":"dalog/zapLog.go:28","msg":"hello world","ID":"A123","Hostname":"MacBook-Pro.local"}
```

### String


```go
os.Setenv("DALOG", "ZAP")
dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
```
```
2017/12/02 00:43:28 INFO hello world, ID=A123, Hostname=MacBook-Pro.local
```

