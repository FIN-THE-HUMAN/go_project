package stream
import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
)

func Catch(err error){
	if err != nil{
		panic(err)
	}	
}

func ReadFile(config string) []byte {
	content, err := ioutil.ReadFile(config)
	Catch(err)
	return content
}

func Pause(){
	fmt.Print("Print enter to exit")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
