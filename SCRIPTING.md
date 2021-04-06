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

Print zero or more a value to the output. All arguments are concatenated together with a space.

Example:

```
print "A simple print."
print
print "Concatenate multiple" 42 "values" 123
```

Output:

```
A simple print.

Concatenate multiple 42 values 123
```

### sleep

Wait a number of milliseconds before resuming the script.

Example:

```
# Sleep for 1 second:
sleep 1000
```
