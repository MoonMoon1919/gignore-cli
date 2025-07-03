package builder

import (
	"context"
	"log"

	"github.com/MoonMoon1919/gignore"
	"github.com/urfave/cli/v3"
)

func logActionResults(results []gignore.Result) {
	for _, result := range results {
		log.Print(result.Log())
	}
}

func makePathCommand(name, usage string, flags []cli.Flag, handler func(path string, c *cli.Command) error) *cli.Command {
	pathFlag := cli.StringFlag{
		Name:  "path",
		Value: ".gitignore",
		Usage: "The path to which your ignore file will be saved",
	}

	return &cli.Command{
		Name:  name,
		Usage: usage,
		Flags: append([]cli.Flag{&pathFlag}, flags...),
		Action: func(ctx context.Context, c *cli.Command) error {
			path := c.String("path")

			return handler(path, c)
		},
	}
}

func makeRuleCommand(name, usage string, flags []cli.Flag, handler func(path string, parsedAction gignore.Action, c *cli.Command) error) *cli.Command {
	actionFlag := cli.StringFlag{
		Name:  "action",
		Value: "include",
		Usage: "if you would like to ignore or allow the path (useful for exclusion) - either 'include' or 'exclude'",
	}

	return makePathCommand(name, usage, append(flags, &actionFlag), func(path string, c *cli.Command) error {
		action := c.String("action")

		parsedAction, err := gignore.ActionFromString(action)
		if err != nil {
			return err
		}

		return handler(path, parsedAction, c)
	})
}

func AppBuilder(svc gignore.Service) *cli.Command {
	filePathFlag := cli.StringFlag{
		Name:     "filepath",
		Usage:    "the filepath to file you would like to ignore",
		Required: true,
	}
	extensionFlag := cli.StringFlag{
		Name:     "extension",
		Usage:    "The extension to ignore, e.g., 'txt'",
		Required: true,
	}
	patternFlag := cli.StringFlag{
		Name:     "pattern",
		Usage:    "The glob pattern for the rule",
		Required: true,
	}
	dirNameFlag := cli.StringFlag{
		Name:     "name",
		Usage:    "The name of the directory",
		Required: true,
	}
	dirModeFlag := cli.StringFlag{
		Name:     "mode",
		Usage:    "The mode of the directory - directory, recursive, children, anywhere, root",
		Required: true,
	}
	sourcePatternFlag := cli.StringFlag{
		Name:     "source-pattern",
		Usage:    "The pattern of the rule you're moving",
		Required: true,
	}
	destinationPatternFlag := cli.StringFlag{
		Name:     "destination-pattern",
		Usage:    "The pattern of the rule you'd like to move a rule before or after",
		Required: true,
	}
	directionFlag := cli.StringFlag{
		Name:     "direction",
		Usage:    "The direction of the move - before or after",
		Required: true,
	}
	fixFlag := cli.BoolFlag{
		Name:  "fix",
		Usage: "Informs the tool to automatically fix found conflicts and optimize the file",
	}
	maxFixes := cli.IntFlag{
		Name:  "max",
		Value: 20,
		Usage: "The number of attempted fixes the autofixer will perform before exiting",
	}

	cmd := &cli.Command{
		Name:  "gignore-cli",
		Usage: "Manage your ignore files with ease",
		Commands: []*cli.Command{
			makePathCommand("create", "create a new ignore file", []cli.Flag{},
				func(path string, c *cli.Command) error {
					return svc.Init(path)
				},
			),
			{
				Name:  "add",
				Usage: "Add a new rule",
				Commands: []*cli.Command{
					makeRuleCommand("file", "Add a new file rule", []cli.Flag{&filePathFlag},
						func(path string, parsedAction gignore.Action, c *cli.Command) error {
							filePath := c.String("filepath")

							results, err := svc.AddFileRule(path, filePath, parsedAction)
							logActionResults(results)

							return err
						},
					),
					makeRuleCommand("directory", "Add a new directory rule", []cli.Flag{&dirNameFlag, &dirModeFlag},
						func(path string, parsedAction gignore.Action, c *cli.Command) error {
							name := c.String("name")
							mode := c.String("mode")

							parsedMode, err := gignore.DirectoryModeFromString(mode)
							if err != nil {
								return err
							}

							results, err := svc.AddDirectoryRule(path, name, parsedMode, parsedAction)
							logActionResults(results)

							return err
						},
					),
					makeRuleCommand("extension", "Add a new extension rule, e.g., 'txt'", []cli.Flag{&extensionFlag},
						func(path string, parsedAction gignore.Action, c *cli.Command) error {
							ext := c.String("extension")

							results, err := svc.AddExtensionRule(path, ext, parsedAction)
							logActionResults(results)

							return err
						},
					),
					makeRuleCommand("glob", "Add a glob rule - e.g., '.coverage.*'", []cli.Flag{&patternFlag},
						func(path string, parsedAction gignore.Action, c *cli.Command) error {
							pattern := c.String("pattern")

							results, err := svc.AddGlobRule(path, pattern, parsedAction)
							logActionResults(results)

							return err
						},
					),
				},
			},
			{
				Name:  "delete",
				Usage: "Delete an existing rule",
				Commands: []*cli.Command{
					makeRuleCommand("file", "Remove a file rule", []cli.Flag{&filePathFlag},
						func(path string, parsedAction gignore.Action, c *cli.Command) error {
							filePath := c.String("filepath")

							results, err := svc.DeleteFileRule(path, filePath, parsedAction)
							logActionResults([]gignore.Result{results})

							return err
						},
					),
					makeRuleCommand("directory", "Delete a directory rule", []cli.Flag{&dirNameFlag, &dirModeFlag},
						func(path string, parsedAction gignore.Action, c *cli.Command) error {
							name := c.String("name")
							mode := c.String("mode")

							parsedMode, err := gignore.DirectoryModeFromString(mode)
							if err != nil {
								return err
							}

							results, err := svc.DeleteDirectoryRule(path, name, parsedMode, parsedAction)
							logActionResults([]gignore.Result{results})

							return err
						},
					),
					makeRuleCommand("extension", "Delete an extension rule, e.g., 'txt'", []cli.Flag{&extensionFlag},
						func(path string, parsedAction gignore.Action, c *cli.Command) error {
							ext := c.String("extension")

							results, err := svc.DeleteExtensionRule(path, ext, parsedAction)
							logActionResults([]gignore.Result{results})

							return err
						},
					),
					makeRuleCommand("glob", "Delete a glob rule - e.g., '.coverage.*'", []cli.Flag{&patternFlag},
						func(path string, parsedAction gignore.Action, c *cli.Command) error {
							pattern := c.String("pattern")

							results, err := svc.DeleteGlobRule(path, pattern, parsedAction)
							logActionResults([]gignore.Result{results})

							return err
						},
					),
				},
			},
			makePathCommand("move", "manually move a rule", []cli.Flag{&sourcePatternFlag, &destinationPatternFlag, &directionFlag},
				func(path string, c *cli.Command) error {
					source := c.String("source-pattern")
					destination := c.String("destination-pattern")
					direction := c.String("direction")

					parsedDirection, err := gignore.MoveDirectionFromString(direction)
					if err != nil {
						return err
					}

					results, err := svc.MoveRule(path, source, destination, parsedDirection)
					logActionResults([]gignore.Result{results})

					return err
				},
			),
			makePathCommand("analyze", "Check if your ignorefile has any conflicts, optionally fix them", []cli.Flag{&fixFlag, &maxFixes},
				func(path string, c *cli.Command) error {
					fix := c.Bool("fix")
					maxFixes := c.Int("max")

					conflicts, err := svc.AnalyzeConflicts(path)
					if err != nil {
						return err
					}

					if len(conflicts) == 0 {
						log.Print("No conflicts found")
					}

					for _, conflict := range conflicts {
						log.Printf("FOUND CONFLICT: Left: %s, Right: %s, Type: %s \n", conflict.Left.Render(), conflict.Right.Render(), conflict.ConflictType)
					}

					if fix && len(conflicts) > 0 {
						results, err := svc.AutoFix(path, maxFixes)
						logActionResults(results)

						return err
					}

					return nil
				},
			),
		},
	}

	return cmd
}
