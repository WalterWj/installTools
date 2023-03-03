package ansible_test

import (
	"fmt"
	ansible "installTools/pkg/ansible"
	"testing"
)

func TestExecute(t *testing.T) {
	// Create ssh configs
	config1 := ansible.InitSshConfig()
	config2 := ansible.InitSshConfig()
	config1.SshIp, config2.SshIp = "127.0.0.1", "172.16.201.210"
	config2.SshUser, config1.SshUser = "tidb", "tidb"
	config1.PrivateKey, config2.PrivateKey = "id_rsa.test", "id_rsa.test"
	configs := []*ansible.SshConfig{config1, config2}

	// Execute commands in parallel
	commands := []string{"whoami", "ls"}
	results := ansible.ParallelExecute(configs, commands)

	// Print results
	for ip, output := range results {
		fmt.Printf("%s:\n%s\n", ip, output)
	}
}
