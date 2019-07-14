# GOGIT

[![Build Status](http://192.168.0.3/job/build-go-git/badge/icon)](http://192.168.0.3/job/build-go-git/)

[![Quality Gate Status](http://192.168.0.3:9000/api/project_badges/measure?project=gogit-simple-run&metric=alert_status)](http://192.168.0.3:9000/dashboard?id=gogit-simple-run)

## Overview

GoGit is a highly efficient binary with fully compatible with GIT written in GO language. It's written to do day to day GIT command in more simple way and easy to remember commands.

## Installation

GOGIT can be downloadable from [URL](https://github.com/mohsinzaheer25/gogit/releases/download/V1.0.0/gogit) for all three major operating system or by using below wget command on linux machines.

```
wget https://github.com/mohsinzaheer25/gogit/releases/download/v1.0.0/gogit_REPLACEWITHYOUROS_X86_64
```

Copy the downloaded binary to `/usr/local/bin/` by using below command.

```
cp REPLCEWITHDOWNLOADEDPATH/gogit_REPLACEWITHYOUROS_X86_64 /usr/local/bin/gogit
```

Validate the installation is successful by running `gogit version` command.

```
$ gogit version

GOGIT v1.0.0
```

## Example

There are number of commands available by GOGIT and can be found by running `gogit help` command.

```
$ gogit help

GOGIT is command line tool to use git command in simple way.

Usage: gogit [COMMAND]

Commands:
ls          List All The Change Files
add         Add Files To The Repository
get         Get a Repository or Branch Updated Files
undo        Reset your commit to particular commit or Reset last commit
newbranch   Creates A New Branch
version     Display GOGIT Version
```

Each command usage can be found by running `gogit [command] help`.

```
$ gogit add help

GOGIT is command line tool to use git command in simple way.
	Usage: gogit add -f "[Filenames]" -c "[Comment]" -b "[Branch Name]"

Usage:
  gogit [flags]

Flags:
  -b, --branch string    Git Branch
  -c, --comment string   Comment
  -f, --files string     Single File or Multiple Files With Spaces Within Quotes
  -h, --help             help for gogit
```

To create a new branch and switch to that branch, just use `gogit newbranch [Branch Name]`

```
$ gogit newbranch testbranch

Switched to a new branch 'testbranch'
```

Likewise, you can utilize other commands to work with git in more simple way.

## Contribute

Contributions are more than welcome, if you are interested please send an email to **mohsinzaheer25@hotmail.com** until contribution guidelines get ready.
