package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Start Apps",
		Usage: "A simple CLI app with urfave/cli to run apps nodeJS",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Prints a greeting",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "project",
						Usage: "Your project name",
					},
				},
				Action: func(c *cli.Context) error {
					// Get the current directory
					dir, err := os.Getwd()
					if err != nil {
						log.Fatal(err)
					}

					// Define the possible entry files
					entryFiles := []string{"server.js", "app.js", "index.js"}

					// Check if the "project" flag is provided
					projectName := c.String("project")

					// If a project name is provided, run only that specific project
					if projectName != "" {
						projectPath := filepath.Join(dir, projectName)

						// Check if the project directory exists
						info, err := os.Stat(projectPath)
						if err != nil {
							if os.IsNotExist(err) {
								fmt.Printf("Project '%s' does not exist in the current directory.\n", projectName)
							} else {
								log.Fatal(err)
							}
							return nil
						}

						// Check if it's a directory
						if !info.IsDir() {
							fmt.Printf("'%s' is not a directory.\n", projectName)
							return nil
						}

						// Try to find one of the entry files in this project
						for _, entryFile := range entryFiles {
							appFilePath := filepath.Join(projectPath, entryFile)
							if _, err := os.Stat(appFilePath); err == nil {
								// Run `node <entryFile>` in this directory
								fmt.Printf("Running project in: %s\n", projectPath)

								cmd := exec.Command("node", entryFile)
								cmd.Dir = projectPath // Set the directory to the project folder
								cmd.Stdout = os.Stdout
								cmd.Stderr = os.Stderr

								// Run the command
								if err := cmd.Run(); err != nil {
									log.Printf("Error running node for project %s: %v\n", projectPath, err)
								}
								return nil // Exit after running the specified project
							}
						}

						// If no valid entry file is found
						fmt.Printf("No valid entry file found in project '%s'.\n", projectName)
						return nil
					}

					// If no project flag is provided, iterate over all subdirectories
					err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
						if err != nil {
							log.Println(err)
							return err
						}

						// Check if the path is a directory and is not the current directory
						if info.IsDir() && path != dir {
							// Try to find one of the entry files (server.js, app.js, or index.js)
							for _, entryFile := range entryFiles {
								appFilePath := filepath.Join(path, entryFile)
								if _, err := os.Stat(appFilePath); err == nil {
									// Run `node <entryFile>` in this directory
									fmt.Printf("Running project in: %s\n", path)

									cmd := exec.Command("node", entryFile)
									cmd.Dir = path // Set the directory to the project folder
									cmd.Stdout = os.Stdout
									cmd.Stderr = os.Stderr

									// Run the command
									if err := cmd.Run(); err != nil {
										log.Printf("Error running node for project %s: %v\n", path, err)
									}
									break // If one entry file is found, no need to check further
								}
							}
						}
						return nil
					})

					if err != nil {
						log.Fatal(err)
					}

					return nil
				},
			},
		},
	}

	// Run the app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
