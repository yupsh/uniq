package command

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
	IgnoreCase    IgnoreCaseFlag = true
	CaseSensitive IgnoreCaseFlag = false
)

type SkipFields int
type SkipChars int

type flags struct {
	Count          CountFlag
	DuplicatesOnly DuplicatesOnlyFlag
	UniqueOnly     UniqueOnlyFlag
	IgnoreCase     IgnoreCaseFlag
	SkipFields     SkipFields
	SkipChars      SkipChars
}

func (f CountFlag) Configure(flags *flags)          { flags.Count = f }
func (f DuplicatesOnlyFlag) Configure(flags *flags) { flags.DuplicatesOnly = f }
func (f UniqueOnlyFlag) Configure(flags *flags)     { flags.UniqueOnly = f }
func (f IgnoreCaseFlag) Configure(flags *flags)     { flags.IgnoreCase = f }
func (s SkipFields) Configure(flags *flags)         { flags.SkipFields = s }
func (s SkipChars) Configure(flags *flags)          { flags.SkipChars = s }
