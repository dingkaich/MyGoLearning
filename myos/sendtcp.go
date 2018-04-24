package myos

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func sedntcp() {

	var soapfile string
	if len(os.Args) == 1 {
		soapfile = "soap.xml"
	} else {
		soapfile = os.Args[1]
	}
	if !IsFileExist(soapfile) {
		fmt.Println(soapfile, "not exits")
		os.Exit(1)
	}

	destipport := readjsonfile("config.json")
	fmt.Printf("==Send Soap Message to : %s ==\n", destipport)
	filestr := readsoapfile(soapfile)
	conn, err := net.Dial("tcp", destipport)
	if err != nil {
		log.Fatal(err)
	}
	//发送数据
	fmt.Fprintf(conn, filestr)
	defer conn.Close()
	mustCopy(os.Stdout, conn)
	fmt.Println("")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func readsoapfile(filePath string) string {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("read file: %v error: %v", filePath, err)
		return "read file failed"
	}
	s := string(b)
	fmt.Println(s)
	return s
}

func readjsonfile(filePath string) string {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("read file: %v error: %v", filePath, err)
		return "read file failed"
	}
	var ipport struct {
		Ip   string `json:"ip"`
		Port string `json:"port"`
	}
	if err := json.Unmarshal(data, &ipport); err != nil {
		fmt.Println(filePath, "Unmarshal failed")
		return "Unmarshal failed"
	}
	return ipport.Ip + ":" + ipport.Port
}

//!-
