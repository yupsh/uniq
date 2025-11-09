package command_test

import (
	"errors"
	"testing"

	"github.com/gloo-foo/testable/assertion"
	"github.com/gloo-foo/testable/run"
	command "github.com/yupsh/uniq"
)

// Test basic duplicate removal
func TestUniq_Basic(t *testing.T) {
	result := run.Command(command.Uniq()).
		WithStdinLines("a", "a", "b", "b", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"a", "b", "c"})
}

// Test count flag
func TestUniq_Count(t *testing.T) {
	result := run.Command(command.Uniq(command.Count)).
		WithStdinLines("a", "a", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 3)
	assertion.Contains(t, result.Stdout, "2 a")
	assertion.Contains(t, result.Stdout, "1 b")
}

// Test duplicates only
func TestUniq_DuplicatesOnly(t *testing.T) {
	result := run.Command(command.Uniq(command.DuplicatesOnly)).
		WithStdinLines("a", "a", "b", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"a", "b"})
}

// Test unique only
func TestUniq_UniqueOnly(t *testing.T) {
	result := run.Command(command.Uniq(command.UniqueOnly)).
		WithStdinLines("a", "a", "b", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"c"})
}

// Test ignore case
func TestUniq_IgnoreCase(t *testing.T) {
	result := run.Command(command.Uniq(command.IgnoreCase)).
		WithStdinLines("a", "A", "b", "B").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"a", "b"})
}

// Test skip fields
func TestUniq_SkipFields(t *testing.T) {
	result := run.Command(command.Uniq(command.SkipFields(1))).
		WithStdinLines("x a", "y a", "z b").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"x a", "z b"})
}

// Test skip chars
func TestUniq_SkipChars(t *testing.T) {
	result := run.Command(command.Uniq(command.SkipChars(2))).
		WithStdinLines("xxa", "yya", "zzb").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"xxa", "zzb"})
}

// Test empty input
func TestUniq_EmptyInput(t *testing.T) {
	result := run.Quick(command.Uniq())
	assertion.NoError(t, result.Err)
	assertion.Empty(t, result.Stdout)
}

// Test single line
func TestUniq_SingleLine(t *testing.T) {
	result := run.Command(command.Uniq()).
		WithStdinLines("only").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"only"})
}

// Test all unique
func TestUniq_AllUnique(t *testing.T) {
	result := run.Command(command.Uniq()).
		WithStdinLines("a", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"a", "b", "c"})
}

// Test all duplicates
func TestUniq_AllDuplicates(t *testing.T) {
	result := run.Command(command.Uniq()).
		WithStdinLines("a", "a", "a").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"a"})
}

// Test errors
func TestUniq_InputError(t *testing.T) {
	result := run.Command(command.Uniq()).
		WithStdinError(errors.New("read failed")).Run()
	assertion.ErrorContains(t, result.Err, "read failed")
}

// Table-driven tests
func TestUniq_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"basic", []string{"a", "a", "b"}, []string{"a", "b"}},
		{"three groups", []string{"a", "a", "b", "b", "c", "c"}, []string{"a", "b", "c"}},
		{"no dups", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"all same", []string{"x", "x", "x"}, []string{"x"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := run.Command(command.Uniq()).
				WithStdinLines(tt.input...).Run()
			assertion.NoError(t, result.Err)
			assertion.Lines(t, result.Stdout, tt.expected)
		})
	}
}

func TestUniq_SkipFieldsExceedsLineLength(t *testing.T) {
	// When skip fields exceeds line length, comparison treats as empty
	result := run.Command(command.Uniq(command.SkipFields(10))).
		WithStdinLines("a b", "c d", "e f").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1) // All become empty after skipping
}

func TestUniq_SkipCharsExceedsLineLength(t *testing.T) {
	// When skip chars exceeds line length, comparison treats as empty
	result := run.Command(command.Uniq(command.SkipChars(100))).
		WithStdinLines("short", "line", "text").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1) // All become empty after skipping
}

