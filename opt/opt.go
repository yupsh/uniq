package opt

// Boolean flag types with constants
type CountFlag bool
const (
	Count   CountFlag = true
	NoCount CountFlag = false
)

type DuplicatesOnlyFlag bool
const (
	DuplicatesOnly   DuplicatesOnlyFlag = true
	NoDuplicatesOnly DuplicatesOnlyFlag = false
)

type UniqueOnlyFlag bool
const (
	UniqueOnly   UniqueOnlyFlag = true
	NoUniqueOnly UniqueOnlyFlag = false
)

type IgnoreCaseFlag bool
const (
	IgnoreCase   IgnoreCaseFlag = true
	CaseSensitive IgnoreCaseFlag = false
)

// Custom types for parameters
type SkipFields int
type SkipChars int

// Flags represents the configuration options for the uniq command
type Flags struct {
	Count          CountFlag          // Prefix lines by number of occurrences
	DuplicatesOnly DuplicatesOnlyFlag // Only print duplicate lines
	UniqueOnly     UniqueOnlyFlag     // Only print unique lines
	IgnoreCase     IgnoreCaseFlag     // Case insensitive comparison
	SkipFields     SkipFields         // Skip first N fields
	SkipChars      SkipChars          // Skip first N characters
}

// Configure methods for the opt system
func (f CountFlag) Configure(flags *Flags) { flags.Count = f }
func (f DuplicatesOnlyFlag) Configure(flags *Flags) { flags.DuplicatesOnly = f }
func (f UniqueOnlyFlag) Configure(flags *Flags) { flags.UniqueOnly = f }
func (f IgnoreCaseFlag) Configure(flags *Flags) { flags.IgnoreCase = f }
func (s SkipFields) Configure(flags *Flags) { flags.SkipFields = s }
func (s SkipChars) Configure(flags *Flags) { flags.SkipChars = s }
