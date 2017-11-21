//пакет включает набор методов для работы с файлами, читать из файла.
package stream

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}

func CatchF(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(config string) []byte {
	content, err := ioutil.ReadFile(config)
	Catch(err)
	return content
}

func Pause() {
	fmt.Print("Print enter to exit")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func WriteStringToFile(info string, way string) {
	var result = info
	result += "\n"
	byteResult := []byte(result)
	err := ioutil.WriteFile(way, byteResult, 0644)
	Catch(err)
}

func WriteStringArrayToFile(info []string, way string) {
	for i := 0; i < len(info); i++ {
		WriteStringToFile(info[i], way)
	}
}

func flagReadCommandLine(key string, defaultValue string, helpInfo string) string {
	var config = flag.String(key, defaultValue, helpInfo)
	flag.Parse()
	return *config
}
