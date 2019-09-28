# cfg - simple configuration library...
## which supports:

- Flags:

```
$ go run test/main.go -bool_arg 1 -string MY\ string -int 45 -default changed
```

- Environment variables:

```
$ BOOL_ARG=true STRING="some string" INT=45 go run test/main.go
```

- Defaults:

```golang
defaults := map[string]interface{}{}
defaults["somekey"] = "some value"
config := cfg.Config(defaults)
someKey, err := config.StringArg("somekey", "Some key description.", "somekey") // someKey with be "some value"
```


