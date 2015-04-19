package tool

import (
    "io"
)

// Interface of project tools
type Tool interface {

    // Name of tool
    Name() string

    // Return Help information to show how to use this tool
    Help() string

    // Execute given tool
    Run(args []string, out io.Writer) error
}