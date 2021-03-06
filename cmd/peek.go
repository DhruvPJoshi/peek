/*
Copyright © 2022 Dhruv Joshi joshidhruv997@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
  "fmt"
	"github.com/DhruvPJoshi/peek/ops"
	"github.com/spf13/cobra"
	"os"
)

var verbose bool

var peekCmd = &cobra.Command{
	Use:   "peek STRING FILE",
	Short: "Displays lines of a file beginning with a given string.",
	Long: `The peek utility will display the lines in a specified file that matches a given string as prefix.

This utility is somewhat similar to the Unix's look utility.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
		  fmt.Println("Verbose logging active.")
		}

		search := args[0]
		if verbose {
  		fmt.Printf("Search string: %s\n", search)
		}
		file := args[1]
		if verbose {
		  fmt.Printf("File to search: %s\n", file)
		}

		ops.PeekIn(search, &file)
	},
}

func Execute() {
	err := peekCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	peekCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose messaging")
}
