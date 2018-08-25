[![Build Status](https://travis-ci.org/fuwalab/tools.svg?branch=master)](https://travis-ci.org/fuwalab/tools)
[![Go Report Card](https://goreportcard.com/badge/github.com/fuwalab/tools)](https://goreportcard.com/report/github.com/fuwalab/tools)
[![Maintainability](https://api.codeclimate.com/v1/badges/11b9d5fd6267ed9681e8/maintainability)](https://codeclimate.com/github/fuwalab/tools/maintainability)
[![GitHub license](https://img.shields.io/github/license/fuwalab/tools.svg)](https://github.com/fuwalab/tools/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/fuwalab/tools?status.svg)](https://godoc.org/github.com/fuwalab/tools)

# tools

## Overview
Toolkit that may need.
 
## Features
### Sub Commands
This tool has the following `sub command`.

#### Secure Account Manager
1. AddAccount
    - Able to save service name, user name and password
    - All parameters will be encrypted before saving into DB
1. ShowAccount
    - Able to show user name
    - Output user name on your terminal
1. CopyPassword
    - Able to copy password into clipboard
    
### Installation
#### git clone
Need to clone this repository on go source directory.
```bash
$ cd $GOPATH/src
$ git clone git@github.com:fuwalab/tools.git
```

#### Install dependent packages
- Require [`dep`](https://github.com/golang/dep) 
    - mac
        ```bash
        $ brew install dep
        ```

- Install dependencies
    ```bash
    $ cd tools
    $ dep ensure
    ```
    
#### Build
```bash
$ cd $GOPATH
$ go build -o bin/tools tools
```

### Usage
### How to run?
```bash
$ $GOPATH/bin/tools [subcommand] [options...]
```

#### Usage
```
AddAccount: Add a new account information.
      Run "AddAccount -h" for more detail.
ShowAccount: Show account/user name of a particular service.
      Run "ShowAccount -h" for more detail.
CopyPassword: Copy password of the particular service to clipboard.
      Run "CopyPassword -h" for more detail.
```