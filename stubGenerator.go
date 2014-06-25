package main

import (
	"crypto/aes"
	"crypto/sha256"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/template"
)

var inFile = flag.String("in", "", "which file to read")
var name = flag.String("name", "somename", "which name to give the stub")
var outDir = flag.String("out", "./stubber", "stubber package directory")
var encryptKey = flag.String("key", "unsecure", "which key to use for encryption")

var gotemplate = `package stubber
import "crypto/aes"
func init(){
var data = {{.Data}}
var key = {{.Key}}
cipher,_ := aes.NewCipher(key)
cipher.Decrypt(data,data)
registerStub("{{.Name}}",data)
}
`

func main() {
	flag.Parse()
	if *inFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	f, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	key := sha256.New().Sum([]byte(*encryptKey))[:32]
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	cipher.Encrypt(data, data)

	dataArray := dataToGoArray(data)
	keyArray := dataToGoArray(key)

	t := template.Must(template.New("goprog").Parse(gotemplate))

	f, err = os.Create(*outDir + "/" + *name + "_stub.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	t.Execute(f, map[string]interface{}{
		"Name": *name,
		"Key":  keyArray,
		"Data": dataArray,
	})

}

func dataToGoArray(data []byte) string {
	result := "[]byte{\n"
	for idx, b := range data {
		bStr := strconv.FormatInt(int64(b), 16)
		if len(bStr) == 1 {
			bStr = "0" + bStr
		}
		bStr = "0x" + bStr + ", "
		if idx%20 == 19 {
			bStr += "\n"
		}
		result += bStr
	}
	result = result[0 : len(result)-1]
	result += "\n}"
	return result
}
