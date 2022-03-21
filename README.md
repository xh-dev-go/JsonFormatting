## Installation
```shell
go install github.com/xh-dev-go/JsonFormatting@latest
```

## Commands options
```
-i string indentation with
-indent string indentation with (default "  ")
-p	      pretty print json (default true)
-pretty   pretty print json
-u	      packed json
-ugly     packed json
-version
show application version
```

## Commands

```
# pretty printing the json in clipboard with indentation and save back to clipboard
JsonFormatting -p
## or
JsonFormatting -pretty

# printing json in clipboard without indentation and save back to clipboard
JsonFormatting -u
## or
JsonFormatting -ugly
```