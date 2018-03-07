package main

import (
	"flag"
	"log"

	"bytes"

	"strings"

	"fmt"
	"os"

	"github.com/doovemax/ssh-tool/module"
	"github.com/pkg/sftp"
)

var (
	username *string
	password *string
	ip       *string
	port     *int
	action   *string
)

func main() {
	// ciphers := []string{}
	// session, err := sshconnect(username, password, ip, port)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer session.Close()
	// var stdoutBuf bytes.Buffer
	// session.Stdout = &stdoutBuf
	// session.Run(cmd)
	// log.Println(session.Stdout)
	username = flag.String("u", "", "username")
	password = flag.String("P", "", "user passwd")
	ip = flag.String("h", "", "Host")
	port = flag.Int("p", 22, "ssh port")
	action = flag.String("action", "shell", `shell or get  or put  ?
		get src local
		put local src
		full path `)
	//*username = "root"
	//*password = "Wawjj80290608$?"
	//*ip = "192.168.50.73"
	//*port = 2222
	//*action = "put"

	flag.Parse()

	command := flag.Args()
	//fmt.Println(command)
	//command := []string{"./test.go", "/root/test.go"}
	switch *action {
	case "get":

	case "put":
		putFile(username, password, ip, port, command)
	default:
		runCmd(username, password, ip, port, command)
	}

}

func runCmd(username *string, password *string, ip *string, port *int, command []string) {
	session, err := module.Sshconnect(*username, *password, *ip, *port)
	//fmt.Println(*password)
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(strings.Join(command, " "))
	log.Println(session.Stdout)
}

func putFile(username *string, password *string, ip *string, port *int, command []string) {

	var err error
	var sftpClient *sftp.Client

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	fmt.Println()
	sftpClient, err = module.Sftpconnect(*username, *password, *ip, *port)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	// 用来测试的本地文件路径 和 远程机器上的文件夹
	var localFilePath = command[0]
	var remotePath = command[1]
	//fmt.Println(command)
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		//fmt.Println(localFilePath)
		log.Fatal(err)

	}
	defer srcFile.Close()
	//fmt.Println(remotePath)
	dstFile, err := sftpClient.Create(remotePath)
	if err != nil {
		//fmt.Println(remotePath)
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}

	fmt.Println("copy file to remote server finished!")
}
