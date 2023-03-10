package ansible

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

type SshConfig struct {
	TimeOut     int
	PrivateKey  string
	SshUser     string
	SshPort     int
	SshPassword string
	SshIp       string
}

// init struct SshConfig default
func InitSshConfig() *SshConfig {
	return &SshConfig{
		TimeOut:     5,
		PrivateKey:  "~/.installTools/ssh/id_rsa",
		SshUser:     "root",
		SshPort:     22,
		SshPassword: "",
		SshIp:       "127.0.0.1",
	}
}

// Ssh by PrivateKey
func (s *SshConfig) ConnectPrivateKey() *ssh.Client {
	// Load the private key
	privateKey, err := os.ReadFile(s.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// Set up the config
	config := &ssh.ClientConfig{
		User: s.SshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote host
	addr := fmt.Sprintf("%s:%d", s.SshIp, s.SshPort)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// Ssh by password
func (s *SshConfig) ConnectPassword() *ssh.Client {
	// Set up the config
	config := &ssh.ClientConfig{
		User: s.SshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.SshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote host
	addr := fmt.Sprintf("%s:%d", s.SshIp, s.SshPort)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
