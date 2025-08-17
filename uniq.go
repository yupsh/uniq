package uniq

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"
	localopt "github.com/yupsh/uniq/opt"
)

// Flags represents the configuration options for the uniq command
type Flags = localopt.Flags

// Command implementation
type command opt.Inputs[string, Flags]

// Uniq creates a new uniq command with the given parameters
func Uniq(parameters ...any) yup.Command {
	return command(opt.Args[string, Flags](parameters...))
}

func (c command) Execute(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
	return yup.ProcessFilesWithContext(
		ctx, c.Positional, input, output, stderr,
		yup.FileProcessorOptions{
			CommandName:     "uniq",
			ContinueOnError: true,
		},
		func(ctx context.Context, source yup.InputSource, output io.Writer) error {
			return c.processReader(ctx, source.Reader, output)
		},
	)
}

func (c command) processReader(ctx context.Context, reader io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(reader)

	var prevLine string
	var prevKey string
	var count int

	for yup.ScanWithContext(ctx, scanner) {
		line := scanner.Text()
		key := c.getComparisonKey(line)

		if key == prevKey {
			count++
		} else {
			// Output previous group if it exists
			if prevLine != "" {
				c.outputLine(output, prevLine, count)
			}

			// Start new group
			prevLine = line
			prevKey = key
			count = 1
		}
	}

	// Check if context was cancelled
	if err := yup.CheckContextCancellation(ctx); err != nil {
		return err
	}

	// Output final group
	if prevLine != "" {
		c.outputLine(output, prevLine, count)
	}

	return scanner.Err()
}

func (c command) getComparisonKey(line string) string {
	key := line

	// Skip fields if specified
	if c.Flags.SkipFields > 0 {
		fields := strings.Fields(line)
		skip := int(c.Flags.SkipFields)
		if skip < len(fields) {
			key = strings.Join(fields[skip:], " ")
		} else {
			key = ""
		}
	}

	// Skip characters if specified
	if c.Flags.SkipChars > 0 {
		runes := []rune(key)
		skip := int(c.Flags.SkipChars)
		if skip < len(runes) {
			key = string(runes[skip:])
		} else {
			key = ""
		}
	}

	// Apply case insensitive comparison if specified
	if bool(c.Flags.IgnoreCase) {
		key = strings.ToLower(key)
	}

	return key
}

func (c command) outputLine(output io.Writer, line string, count int) {
	// Apply filtering based on flags
	if bool(c.Flags.DuplicatesOnly) && count == 1 {
		return // Skip unique lines
	}

	if bool(c.Flags.UniqueOnly) && count > 1 {
		return // Skip duplicate lines
	}

	// Output with count if requested
	if bool(c.Flags.Count) {
		fmt.Fprintf(output, "%7d %s\n", count, line)
	} else {
		fmt.Fprintln(output, line)
	}
}
