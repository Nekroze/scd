#! /usr/bin/env nix-shell
#! nix-shell -i bats -p bats
go build
scd="./scd -dryrun"

@test "github https url goes into home git tree" {
  run $scd https://github.com/Nekroze/scd.git
  [ "$status" -eq 0 ]
  [[ "$output" = "$HOME/git/github.com/Nekroze/scd" ]]
}

@test "github ssh url goes into home git tree, same as https output" {
  run $scd git@github.com:Nekroze/scd.git
  [ "$status" -eq 0 ]
  [[ "$output" = "$HOME/git/github.com/Nekroze/scd" ]]
}
