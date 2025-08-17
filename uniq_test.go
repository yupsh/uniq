package uniq_test

import (
	"context"
	"os"
	"strings"

	"github.com/yupsh/uniq"
	"github.com/yupsh/uniq/opt"
)

func ExampleUniq() {
	ctx := context.Background()
	input := strings.NewReader("apple\napple\nbanana\nbanana\nbanana\ncherry\n")

	cmd := uniq.Uniq()
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output: apple
	// banana
	// cherry
}

func ExampleUniq_withCount() {
	ctx := context.Background()
	input := strings.NewReader("apple\napple\nbanana\nbanana\nbanana\n")

	cmd := uniq.Uniq(opt.Count)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:       2 apple
	//       3 banana
}
