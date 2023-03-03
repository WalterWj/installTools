package ansible

func (s *SshConfig) ExecuteCommand(command string) (string, error) {
	// Connect to remote host
	client := s.ConnectPrivateKey()
	// if err != nil {
	// 	return "", err
	// }
	defer client.Close()

	// Create session
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// Execute command
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func ParallelExecute(configs []*SshConfig, commands []string) map[string]string {
	results := make(map[string]string)

	// Create channel to receive results
	resultChan := make(chan struct {
		ip     string
		result string
	})

	// Start goroutine for each remote host
	for _, config := range configs {
		go func(config *SshConfig) {
			// Connect to remote host and execute command
			for _, command := range commands {
				output, err := config.ExecuteCommand(command)
				if err != nil {
					resultChan <- struct {
						ip     string
						result string
					}{ip: config.SshIp, result: err.Error()}
					break
				}
				resultChan <- struct {
					ip     string
					result string
				}{ip: config.SshIp, result: output}
			}
		}(config)
	}

	// Collect results from channel
	for range configs {
		result := <-resultChan
		results[result.ip] += result.result
	}

	return results
}
