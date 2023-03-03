package ansible_test

import (
	"fmt"
	ansible "installTools/pkg/ansible"
	"log"
	"strings"
	"testing"
)

func TestConnectPrivateKey(t *testing.T) {
	// s := ansible.SshConfig{
	// 	SshUser:    "tidb",
	// 	SshPort:    22,
	// 	SshIp:      "127.0.0.1",
	// 	PrivateKey: "id_rsa.test",
	// }
	s := ansible.InitSshConfig()
	s.SshUser = "tidb"
	s.PrivateKey = "id_rsa.test"
	fmt.Println(s)
	client := s.ConnectPrivateKey()
	// defer client.Close()
	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	// Run a command
	var b strings.Builder
	session.Stdout = &b
	session.Run("ls")
	fmt.Print(b.String())
}
