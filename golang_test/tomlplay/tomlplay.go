package main

import (
	"github.com/BurntSushi/toml"
	"fmt"
)

type Header struct{
	header1 string
	header2 int
}

type Param struct {
	param1 string
	param2 int
}

type Tab struct {
	Header
	Param
}

type TomlConfig struct{
	Tabs map[string]Tab
}


func main(){

	var conf TomlConfig

	tomlData := `
[tabs]
	[tabs.iron]
		[tabs.iron.header]
			header1 = "1"
			header2 = 2
		
		[tabs.iron.param]
			param1 = "1"
			param2 = 2
	[tabs.spc]
		[tabs.spc.header]
			header1 = "1"
			header2 = 2
		
		[tabs.spc.param]
			param1 = "1"
			param2 = 2
`

fmt.Println(tomlData)

	if _, err := toml.Decode(tomlData, &conf); err != nil {
		fmt.Println(err)
		// handle error
	}
	fmt.Printf("%#v",conf)

}