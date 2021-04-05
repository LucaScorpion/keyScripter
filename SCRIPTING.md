# Scripting Reference

## Quick Start

```
# Lines starting with a "#" are comments.
print "Call functions like this."
sleep 500
print "Done!"
```

## Functions

### print

Print a value to the output. The first argument of `print` must be a string, which is interpreted as a [Go format](https://golang.org/pkg/fmt/) string. The rest of the arguments can be any value, and are used as values for the format. To use a value in the format string, use `%v`.

Example:

```
print "A simple print."
print "%v is replaced by the passed %v." 42 "values" 
```

Output:

```
A simple print.
42 is replaced by the passed values.
```

### sleep

Wait a number of milliseconds before resuming the script.

Example:

```
# Sleep for 1 second:
sleep 1000
```
