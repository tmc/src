# ctx-exec

ctx-exec is a command-line tool that executes shell commands and wraps their output in XML-like tags.

## Usage

```
ctx-exec "your shell command here"
```

Example:
```
ctx-exec "echo Hello, World\!"
```

Output:
```
<exec-output cmd="echo Hello, World\!">Hello, World\!
</exec-output>
```

