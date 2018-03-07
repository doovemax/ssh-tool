package main

import (
	"flag"

	"github.com/doovemax/ssh-tool/module"
)

var (
	username *string
	password *string
	ip       *string
	port     *int
	cmd      *[]string
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

	flag.Parse()

	command := flag.Args()
	cmd = &command

	session, err := module.Sshconnect(*username, *password, *ip, *port)
	if err !=nil {
		log.Println(err)
		return
	}
	def session.Close()
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(cmd)
	log.Println(session.Stdout)

}
