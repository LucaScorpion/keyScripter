# Scripting Reference

## Quick Start

```
# Lines starting with a "#" are comments.
print "Call functions like this."

# Store values in variables for easy reusability.
timeToSleep = 500
sleep timeToSleep

# Create custom functions to reuse blocks of logic.
myFunc = (first second) {
    print first
    print second
}
myFunc "This is a" "custom function"

# Use timestamp blocks for easy specific timings.
timestamps {
    0    print "This runs immediately"
    1000 print "This runs after one second"
    2000 print "It's like using sleep, but easier"
}

print "That's all folks!"
```

## Values

There are 2 kinds of values: strings and integers. Strings are values wrapped in double quotes (`"`). Integers are numbers, which can be either in decimal or hexadecimal notation.

```
print "I am a string"
vKeyPress 65   # Decimal
vKeyPress 0x41 # Hexadecimal
```

## Variables

Values can be stored in a variable. These variables can be passed to functions just like other values. Variable names can contain letters, numbers, and underscores. Note that they cannot begin with a number.

```
a = 0x41
vKeyPress a
```

## Custom Functions

```
myFunc = () {
    print "Inside myFunc!"
}
myFunc
```

A function can accept arguments by putting the names in the brackets. The types of the arguments are inferred.

```
modPress = (mod key) {
    vKeyDown mod
    vKeyPress key
    vKeyUp mod
}

shift = 0x10
a = 0x41
modPress shift a
```

Variables defined in a function are only available within that function. The following example will error, because the "hoist" variable is not defined.

```
noHoisting = () {
    hoist = "nope"
}
noHoisting
print hoist # This will error.
```

## Timestamps

Timestamp blocks allow you to call a function on specific times. Each line of a timestamp block starts with a number, which is the amount of milliseconds from the start of the block at which the function should run.

```
timestamps {
    0    print "Start"
    300  print "Wait a bit"
    2000 print "Done"
}

# Is functionally the same as:

print "Start"
sleep 300
print "Wait a bit"
sleep 1700
print "Done"
```

The millisecond values should always go up. This example will error:

```
timestamps {
    300 print "300"
    100 print "Not allowed"
}
```

## Builtin Functions

### `pause`

Pause script execution. This waits for the user to press enter before resuming.

```
print "Before"
pause
print "After"
```

Output:

```
Before
Press enter to continue...
After
```

### `print`

Print zero or more a values to the output. All arguments are concatenated together with a space.

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

### `scKeyDown`, `scKeyUp`

Down or up a key based on a scan code. Note that scancodes for different keys can differ per application.

```
scKeyDown 28
scKeyUp   28
```

### `sleep`

Wait a number of milliseconds before resuming the script.

```
# Sleep for 1 second
sleep 1000
```

### `vKeyPress`, `vKeyDown`, `vKeyUp`

Press, down, or up a virtual key. The list of virtual key codes can be found on: https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes

```
# Type "aA"
vKeyPress 0x41 # Press a 
vKeyDown  0x10 # Down shift
vKeyPress 0x41 # Press a
vKeyUp    0x10 # Up shift
```
