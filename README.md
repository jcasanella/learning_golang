# Learning Golang

This repo is just some work to learn Golang and some notes

## How to set up Visual Studio Code

Open Visual Studio and search Golang Plugin and then install it.
Click View -> Command Pallete or type Ctrl+Shift+P and type goinstall update/tools. Check all dependencies and click OK
Create the **settings.json** as:

https://github.com/golang/vscode-go/blob/master/docs/debugging.md

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {},
            "args": []
        }
    ]
}
```

**Note**:

Maybe we need ot create the **go.mod** to be able to run Golang apps from VS Code.

https://www.programmersought.com/article/46126170856/

```
go mod init hello
```

### Guide

1. [Application Structure - Build and run](1_application_structure/README.md)

