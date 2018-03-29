package services

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/CityOfZion/neo-local/cli/logger"
)

// StartComposeStack starts services within the Compose stack.
func StartComposeStack(verbose bool, saveState bool) error {
	command := "docker-compose"
	args := []string{
		"up", "-d", "--build", "--remove-orphans",
	}

	if !saveState {
		args = append(args, "--force-recreate")
	}

	spinner := logger.NewSpinner("Starting services in Docker containers")
	spinner.Start()

	out, err := exec.Command(command, args...).Output()
	spinner.Stop()
	if err != nil {
		return err
	}

	if verbose {
		output := strings.TrimSuffix(string(out), "\n")
		log.Printf("%s %s", command, strings.Join(args, " "))
		fmt.Println(output)
	}

	return nil
}
