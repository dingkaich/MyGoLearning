package myos

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

//测试osexec的部分命令
func Myosexec() {
	stri, err := exec.LookPath("base")
	log.Println(stri, err)

	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("in all caps: %s\n", out.String())

	// cmd = exec.Command("sleep", "5")
	// err = cmd.Start()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Waiting for command to finish...")
	// err = cmd.Wait()
	// log.Printf("Command finished with error: %v", err)

	out1, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The date is %s\n", out1)

}
