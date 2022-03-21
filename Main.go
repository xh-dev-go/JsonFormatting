package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/xh-dev-go/xhUtils/flagUtils"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagBool"
	"github.com/xh-dev-go/xhUtils/flagUtils/flagString"
	"os"
)

func main() {
	prettyPrintVar := flagBool.
		NewDefault("p", "pretty print json", true).
		BindCmd().Share("pretty", "pretty print json").BindCmd()
	uglyPrintVar := flagBool.
		New("u", "packed json").
		BindCmd().Share("ugly", "packed json").BindCmd()

	indentWithVar := flagString.NewDefault("indent", "  ", "indentation with").BindCmd().
		Share("i", "indentation with").BindCmd()

	versionVar := flagUtils.Version().BindCmd()

	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if versionVar.Value() {
		fmt.Printf("Version: %s", "1.0.0")
	}

	msg, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	err = json.Unmarshal([]byte(msg), &m)
	if err != nil {
		panic(err)
	}

	if uglyPrintVar.Value() {
		uglyString, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		err = clipboard.WriteAll(string(uglyString))
		if err != nil {
			return
		}
		err = clipboard.WriteAll(string(uglyString))
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}
	if prettyPrintVar.Value() {
		indent, err := json.MarshalIndent(m, "", indentWithVar.Value())
		if err != nil {
			panic(err)
		}
		err = clipboard.WriteAll(string(indent))
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}
}
