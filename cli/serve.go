package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"

	"github.com/spf13/cobra"
)

func (r *Operations) ServeCmd() *cobra.Command {
	var port int
	var host string
	var hotreload bool
	var module string
	var remote bool

	cmd := &cobra.Command{
		Use:     "serve",
		Args:    cobra.MaximumNArgs(1),
		Aliases: []string{"s", "se"},
		Short:   "Serve a beamlit project",
		Long:    "Serve a beamlit project",
		Example: `  bl serve --remote --hotreload --port 1338`,
		Run: func(cmd *cobra.Command, args []string) {
			var activeProc *exec.Cmd

			// Check for pyproject.toml or package.json
			if _, err := os.Stat("pyproject.toml"); !os.IsNotExist(err) {
				activeProc = startUvicornServer(port, host, hotreload, module, remote)
			} else if _, err := os.Stat("package.json"); !os.IsNotExist(err) {
				activeProc = startTypescriptServer(port, host, hotreload, module, remote)
			} else {
				fmt.Println("Error: Neither pyproject.toml nor package.json found in current directory")
				os.Exit(1)
			}

			// Handle graceful shutdown on interrupt
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			go func() {
				<-c
				fmt.Println("\nShutting down server...")
				if err := activeProc.Process.Kill(); err != nil {
					fmt.Printf("Error killing process: %v\n", err)
				}
				os.Exit(0)
			}()

			// Wait for process to exit
			if err := activeProc.Wait(); err != nil {
				fmt.Printf("Server error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 1338, "Bind socket to this host")
	cmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "Bind socket to this port. If 0, an available port will be picked")
	cmd.Flags().StringVarP(&module, "module", "m", "agent.main", "Module to serve, can be an agent or a function")
	cmd.Flags().BoolVarP(&hotreload, "hotreload", "", false, "Watch for changes in the project")
	cmd.Flags().BoolVarP(&remote, "remote", "r", false, "Serve the project remotely. It will use functions deployed on beamlit cloud")
	return cmd
}

func startUvicornServer(port int, host string, hotreload bool, module string, remote bool) *exec.Cmd {
	uvicornCmd := "uvicorn"
	if _, err := os.Stat(".venv"); !os.IsNotExist(err) {
		uvicornCmd = ".venv/bin/uvicorn"
	}

	uvicorn := exec.Command(
		uvicornCmd,
		"beamlit.serve.app:app",
		"--port",
		fmt.Sprintf("%d", port),
		"--host",
		host,
	)
	if hotreload {
		uvicorn.Args = append(uvicorn.Args, "--reload")
	}

	uvicorn.Stdout = os.Stdout
	uvicorn.Stderr = os.Stderr

	// Set environment variables
	uvicorn.Env = getServerEnvironment(port, host, module, remote)

	err := uvicorn.Start()
	if err != nil {
		fmt.Printf("Error starting uvicorn server: %v\n", err)
		os.Exit(1)
	}

	return uvicorn
}

func startTypescriptServer(port int, host string, hotreload bool, module string, remote bool) *exec.Cmd {

	ts := exec.Command(
		"npx",
		"tsx",
		"--tsconfig",
		"./tsconfig.json",
	)
	if !hotreload {
		ts.Args = append(ts.Args, "--watch")
	}
	ts.Args = append(ts.Args, "node_modules/@beamlit/sdk/src/serve/index.ts")

	ts.Stdout = os.Stdout
	ts.Stderr = os.Stderr

	// Set environment variables
	ts.Env = getServerEnvironment(port, host, module, remote)
	if os.Getenv("BL_SERVER_MODULE") == "" {
		if _, err := os.Stat("src"); !os.IsNotExist(err) {
			ts.Env = append(ts.Env, fmt.Sprintf("BL_SERVER_MODULE=%s", "src.agent.agent"))
		} else {
			ts.Env = append(ts.Env, fmt.Sprintf("BL_SERVER_MODULE=%s", "agent.agent"))
		}
	}
	if os.Getenv("BL_AGENT_FUNCTIONS_DIRECTORY") == "" {
		if _, err := os.Stat("src/functions"); !os.IsNotExist(err) {
			ts.Env = append(ts.Env, fmt.Sprintf("BL_AGENT_FUNCTIONS_DIRECTORY=%s", "src/functions"))
		} else {
			ts.Env = append(ts.Env, fmt.Sprintf("BL_AGENT_FUNCTIONS_DIRECTORY=%s", "functions"))
		}
	}

	// Add src directory to NODE_PATH
	nodePath := "src"
	if currentPath := os.Getenv("NODE_PATH"); currentPath != "" {
		nodePath = fmt.Sprintf("%s:%s", nodePath, currentPath)
	}
	ts.Env = append(ts.Env, fmt.Sprintf("NODE_PATH=%s", nodePath))

	err := ts.Start()
	if err != nil {
		fmt.Printf("Error starting tsx server: %v\n", err)
		os.Exit(1)
	}

	return ts
}

func getServerEnvironment(port int, host string, module string, remote bool) []string {
	env := []string{}

	if environment == "" {
		environment = "production"
	}

	env = append(env, fmt.Sprintf("BL_ENVIRONMENT=%s", environment))
	env = append(env, fmt.Sprintf("BL_WORKSPACE=%s", workspace))
	env = append(env, fmt.Sprintf("BL_REMOTE=%t", remote))
	env = append(env, fmt.Sprintf("BL_SERVER_PORT=%d", port))
	env = append(env, fmt.Sprintf("BL_SERVER_HOST=%s", host))
	env = append(env, fmt.Sprintf("BL_SERVER_MODULE=%s", module))

	if os.Getenv("BL_ENV") != "" {
		env = append(env, fmt.Sprintf("BL_ENV=%s", os.Getenv("BL_ENV")))
	}

	// Add all current environment variables if not already set
	for _, envVar := range os.Environ() {
		found := false
		for _, existingVar := range env {
			if envVar == existingVar {
				found = true
				break
			}
		}
		if !found {
			env = append(env, envVar)
		}
	}

	return env
}
