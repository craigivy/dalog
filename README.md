# dalog
Logger abstraction allowing JSON via Zap logger and string via Go's logger.  Also supports a key value context.
Features
* Both JSON and column std output logger
* Context based loggers
* Support for [pkg errors](https://github.com/pkg/errors) stack traces
* DebugContext allows debug out put to be related and annotated
* Environment variable for configuration

## Options
* DALOG_LOGGER=[ZAP|GO] - Define which logger to use
* DALOG_DEBUG=[TRUE|FALSE] - include debug and stack log statements
* DALOG_STACK=[TRUE|FALSE] - include stack traces in the error log statements

## See it in action
Running ```make``` will compile, lint, vet and run tests
dalog contains a lame test that is really just sample usage code (log_test.go).

### JSON via Zap

```go
os.Setenv("DALOG_LOGGER", "ZAP")
dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
```
```json
{"level":"info","ts":1512254914.971346,"msg":"hello world","ID":"A123","Hostname":"MacBook-Pro.local"}
```

### String via go log

```go
os.Setenv("DALOG_LOGGER", "GO")
dalog.WithContext(dalog.WithID("A123"), dalog.WithHostname()).Infof("%s %s", "hello", "world")
```
```
2017/12/02 00:43:28 INFO hello world, ID=A123, Hostname=MacBook-Pro.local
```

## Stack Traces
Stack traces are supported using [pkg errors](https://github.com/pkg/errors) and error log message.

```go
	e := errors.New("This is an error using pkg error")
	dalog.NoContext().Error(e)
```
### JSON output
```json
{"level":"error","ts":1514921204.09196,"msg":"This is an error using pkg error","stack":"\ngithub.com/craigivy/dalog_test.TestStack\n\t/Users/civerson/dev/go/src/github.com/craigivy/dalog/log_test.go:61\ntesting.tRunner\n\t/usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746\nruntime.goexit\n\t/usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337"}
```


### Standard output
```
2018/01/02 12:26:44 ERROR This is an error using pkg error, stack=
github.com/craigivy/dalog_test.TestStack
        /Users/civerson/dev/go/src/github.com/craigivy/dalog/log_test.go:61
testing.tRunner
        /usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746
runtime.goexit
        /usr/local/Cellar/go/1.9.2/libexec/src/runtime/asm_amd64.s:2337
```

 