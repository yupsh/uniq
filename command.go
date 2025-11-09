package command

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	gloo "github.com/gloo-foo/framework"
)

type command gloo.Inputs[gloo.File, flags]

func Uniq(parameters ...any) gloo.Command {
	return command(gloo.Initialize[gloo.File, flags](parameters...))
}

func (p command) Executor() gloo.CommandExecutor {
	return gloo.Inputs[gloo.File, flags](p).Wrap(func(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
		scanner := bufio.NewScanner(stdin)
		var lastLine string
		var lastOriginal string
		var lastCount int64
		var isFirst = true

		// Process each line and track state
		for scanner.Scan() {
			line := scanner.Text()

			// Prepare comparison line (skip fields and chars if specified)
			compareLine := line

			// Skip fields
			if p.Flags.SkipFields > 0 {
				fields := strings.Fields(line)
				skipCount := int(p.Flags.SkipFields)
				if skipCount < len(fields) {
					compareLine = strings.Join(fields[skipCount:], " ")
				} else {
					compareLine = ""
				}
			}

			// Skip chars
			if p.Flags.SkipChars > 0 {
				skipCount := int(p.Flags.SkipChars)
				if skipCount < len(compareLine) {
					compareLine = compareLine[skipCount:]
				} else {
					compareLine = ""
				}
			}

			// Ignore case if flag is set
			checkLine := compareLine
			checkLast := lastLine
			if bool(p.Flags.IgnoreCase) {
				checkLine = strings.ToLower(compareLine)
				checkLast = strings.ToLower(lastLine)
			}

			// Check if this line is different from the last
			if isFirst {
				isFirst = false
				lastLine = checkLine
				lastOriginal = line
				lastCount = 1
				continue
			}

			if checkLine == checkLast {
				// Duplicate line
				lastCount++
				continue
			}

			// Different line - emit the previous line
			shouldEmit := true

			// Apply filters
			if bool(p.Flags.DuplicatesOnly) && lastCount == 1 {
				shouldEmit = false
			}
			if bool(p.Flags.UniqueOnly) && lastCount > 1 {
				shouldEmit = false
			}

			if shouldEmit {
				if bool(p.Flags.Count) {
					fmt.Fprintf(stdout, "%7d %s\n", lastCount, lastOriginal)
				} else {
					fmt.Fprintln(stdout, lastOriginal)
				}
			}

			// Update for next iteration
			lastLine = checkLine
			lastOriginal = line
			lastCount = 1
		}

		// Emit final line
		if !isFirst {
			shouldEmit := true
			if bool(p.Flags.DuplicatesOnly) && lastCount == 1 {
				shouldEmit = false
			}
			if bool(p.Flags.UniqueOnly) && lastCount > 1 {
				shouldEmit = false
			}

			if shouldEmit {
				if bool(p.Flags.Count) {
					fmt.Fprintf(stdout, "%7d %s\n", lastCount, lastOriginal)
				} else {
					fmt.Fprintln(stdout, lastOriginal)
				}
			}
		}

		return scanner.Err()
	})
}
