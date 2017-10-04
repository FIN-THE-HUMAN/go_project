package main
import (
	"fmt"
	"encoding/xml"
	"flag"
	"go_project/stream"
	"go_project/cats"
	"go_project/validator"
	"go_project/colors"
)

func flagReadCommandLine(key string, defaultValue string, helpInfo string) string{
	var config = flag.String(key, defaultValue, helpInfo)
	flag.Parse()
	return *config
}

func main() {
//	config:=flagReadCommandLine("fileName","configXML.xml","set config file name")
	var config = flag.String("fileName","configXML.xml","set config file name")
	flag.Parse()
	content := stream.ReadFile(*config)
	var data Cats.Cats
	xml.Unmarshal(content, &data)
	fmt.Println(data)

	validator.TestAgeAllCats(data)

	fmt.Println()
	var colorFlag = flag.String("colorFileName","colors.xml","set color config file name")
	colorsFile := stream.ReadFile(*colorFlag)
	var ProgrammColors colors.Colors
	xml.Unmarshal(colorsFile, &ProgrammColors)

	validator.TestColorAllCats(data ,ProgrammColors)
//	stream.Pause()
}