# SCD

Smart Change Directory or `scd` seeks to provide a parallel experience to your average `cd` command with subtle improvements such as detecting git repositories and automatically cloning them as neccessary.

The goal is to provide a simple means for moving about in the terminal while also keeping your `$HOME` tidy.

## Installation

Installation can be done directly via the outstanding [Go][1] programming languages tools:

```bash
[nekroze ~]$ go get github.com/Nekroze/scd
```

If your `$GOPATH` and `$PATH` environment variables are properly configured the `scd` command will now be available in your terminal.

## Usage

There are many ways you can let `scd` help simplify your life.

### Replacing `cd`

You can use `scd` the same way you would use `cd` normally:

```bash
[nekroze ~]$ mkdir things
[nekroze ~]$ scd things
[nekroze ~/scd]$
```

### Managing Repsitories

When `scd` is given vcs repository URI it will clone it if it has not already been cloned then change directory into the repository:

```bash
[nekroze ~]$ scd https://github.com/Nekroze/scd
... runs a git clone ...
[nekroze ~/git/github.com/Nekroze/scd]$
[nekroze ~]$ scd git@github.com/Nekroze/scd
[nekroze ~/git/github.com/Nekroze/scd]$
```

This will allow `scd` to attempt to decide a smart location to put things and reduce sprawl while still making sense to a human.

### Changing Directory With Fuzzy Matching

To better facilitate rapid switching between directories `scd` can allow fuzzy matching of directories:

```bash
[nekroze ~]$ scd https://github.com/Nekroze/scd
[nekroze ~/git/github.com/Nekroze/scd]$ cd
[nekroze ~]$ scd scd
[nekroze ~/git/github.com/Nekroze/scd]$ cd
[nekroze ~]$ scd cd
[nekroze ~/git/github.com/Nekroze/scd]$
```

By default `scd` records weights for directories `scd` has changed you too.
