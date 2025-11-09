# Uniq Command Compatibility

## Summary
✅ **Compatible** with Unix `uniq`

## Test Coverage
- **Tests:** 17 functions
- **Coverage:** 93.8%
- **Status:** ✅ All passing

## Key Behaviors

```bash
# Remove adjacent duplicates
$ echo -e "a\na\nb\nb\nc" | uniq
a
b
c

# Count occurrences
$ uniq -c
   2 a
   3 b
   1 c

# Show only duplicates
$ uniq -d
a
b

# Show only unique
$ uniq -u
c
```

Core filtering functionality matches Unix `uniq`.

