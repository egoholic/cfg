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




