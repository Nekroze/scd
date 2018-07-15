#! /usr/bin/env nix-shell
#! nix-shell -i bats -p bats
go build
scd="./scd -dryrun"

@test "no parameters gives \$HOME" {
  run $scd
  [ "$status" -eq 0 ]
  [[ "$output" = "$HOME" ]]
}

@test "single dir gives same output" {
  run $scd foobar
  [ "$status" -eq 0 ]
  [[ "$output" = "foobar" ]]
}

@test "multiple dirs gives same output" {
  run $scd foo/bar
  [ "$status" -eq 0 ]
  [[ "$output" = "foo/bar" ]]
}

@test "tilde gives \$HOME" {
  run $scd ~
  [ "$status" -eq 0 ]
  [[ "$output" = "$HOME" ]]
}

@test "tilde and multiple dirs gives expanded output" {
  run $scd ~/foo/bar
  [ "$status" -eq 0 ]
  [[ "$output" = "$HOME/foo/bar" ]]
}
