package module

import (
	"fmt"
	"net"
	"time"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	auth         []ssh.AuthMethod
	addr         string
	clientConfig *ssh.ClientConfig
	client       *ssh.Client
	session      *ssh.Session
	sftpClient   *sftp.Client
	err          error
)


func Sshconnect(user, password, host string, port int) (*ssh.Session, error) {

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func Sftpconnect(user, password, host string, port int) (*sftp.Client, error) {

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(client); err != nil {
		return nil, err
	}

	return sftpClient, nil
}
