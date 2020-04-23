package cmd

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(upCmd)
}

var upCmd = &cobra.Command{
	Use: "up",
	Short: "Up",
	Long: "Up",
	Run: func (cmd *cobra.Command, args []string) {
		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}

		containerConfig := &container.Config{
			Hostname:        "test",
			Domainname:      "",
			User:            "0",
			AttachStdin:     false,
			AttachStdout:    false,
			AttachStderr:    false,
			ExposedPorts:    nil,
			Tty:             false,
			OpenStdin:       false,
			StdinOnce:       false,
			Env:             nil,
			Cmd:             []string{"tail", "-f", "/dev/null"},
			Healthcheck:     nil,
			ArgsEscaped:     false,
			Image:           "ubuntu:latest",
			Volumes:         nil,
			WorkingDir:      "",
			Entrypoint:      nil,
			NetworkDisabled: false,
			MacAddress:      "",
			OnBuild:         nil,
			Labels:          nil,
			StopSignal:      "",
			StopTimeout:     nil,
			Shell:           nil,
		}
		ctr, err := cli.ContainerCreate(context.Background(), containerConfig, nil, nil, "tail-null")
		if err != nil {
			panic(err)
		}

		fmt.Printf("created container: %v\n", ctr.ID[:10])

		err = cli.ContainerStart(context.Background(), ctr.ID, types.ContainerStartOptions{})
		if err != nil {
			panic(err)
		}

		fmt.Printf("started container: %v\n", ctr.ID[:10])
	},
}
