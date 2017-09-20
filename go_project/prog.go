package main
import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type Cat struct{
	XMLName xml.Name `xml:"cat"`
	Health int		`xml:"health"`
	Name string		`xml:"name"`
	Color string	`xml:"color"`
}

func(c Cat) String() string{
	return fmt.Sprintf("%s\n%v\n%s",c.Name, c.Health, c.Color )
}

func catch(err error){
	if err != nil{
		panic(err)
	}	
}

func main() {
	content, err := ioutil.ReadFile("configXML.xml")
	catch(err)

	var flafy = Cat{Health:100, Name:"flafy", Color:"orange"}
	err = xml.Unmarshal(content, &flafy)
	catch(err)
	fmt.Println(flafy)
}