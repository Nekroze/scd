# SCD

Smart Change Directory or `scd` seeks to provide a parallel experience to your average `cd` command with subtle improvements such as detecting git repositories and automatically cloning them as neccessary.

The goal is to provide a simple means for moving about in the terminal while also keeping your `$HOME` tidy.

## Installation

Installation can be done directly via the outstanding [Go][1] programming languages tools:

```bash
[nekroze ~]$ go get github.com/Nekroze/scd
```

If your `$GOPATH` and `$PATH` environment variables are properly configured the `scd` command will now be available in your terminal.

Due to the way the current working directory functions in a shell you will need to utilize a wrapper around `scd` to perform the final directory change. This can all be done by sourcing the `s.sh` file in your `~/.bashrc` or equivelant.

```
source s.sh
```

Once this line is executed you should have an `s` function/command at your disposal. If you use `scd` directly it will still perform all its functionality except the directory to change to will be printed to the terminal instead of changed to.

## Usage

There are many ways you can let `scd` help simplify your life.

### Replacing `cd`

You can use `scd` the same way you would use `cd` normally:

```bash
[nekroze ~]$ mkdir things
[nekroze ~]$ s things
[nekroze ~/scd]$
```

### Managing Repsitories

When `scd` is given vcs repository URI it will clone it if it has not already been cloned then change directory into the repository:

```bash
[nekroze ~]$ s https://github.com/Nekroze/scd
... runs a git clone ...
[nekroze ~/git/github.com/Nekroze/scd]$
[nekroze ~]$ s git@github.com/Nekroze/scd
[nekroze ~/git/github.com/Nekroze/scd]$
```

This will allow `scd` to attempt to decide a smart location to put things and reduce sprawl while still making sense to a human.

### Changing Directory With Fuzzy Matching

NOTE: This feature is not yet implemented!

To better facilitate rapid switching between directories `scd` can allow fuzzy matching of directories:

```bash
[nekroze ~]$ s https://github.com/Nekroze/scd
[nekroze ~/git/github.com/Nekroze/scd]$ cd
[nekroze ~]$ s scd
[nekroze ~/git/github.com/Nekroze/scd]$ cd
[nekroze ~]$ s cd
[nekroze ~/git/github.com/Nekroze/scd]$
```

By default `scd` records weights for directories `scd` has changed you too.

NOTE: This feature is not yet implemented!
