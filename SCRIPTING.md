# Scripting Reference

## Quick Start

```
# Lines starting with a "#" are comments.
print "Call functions like this."
sleep 500
print "Done!"
```

## Functions

### `print`

Print zero or more a value to the output. All arguments are concatenated together with a space.

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

### `sleep`

Wait a number of milliseconds before resuming the script.

```
# Sleep for 1 second:
sleep 1000
```

### `vKeyPress`, `vKeyDown`, `vKeyUp`

Press, down, or up a virtual key. The list of virtual key codes can be found on: https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes

```
vKeyPress 65 # Press a 
vKeyDown  16 # Down shift
vKeyPress 65 # Press a
vKeyUp    16 # Up shift
```
