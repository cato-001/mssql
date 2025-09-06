{ pkgs }:

pkgs.buildGoApplication {
  pname = "mssql";
  version = "0.1";
  pwd = ./.;
  src = ./src;
  modules = ./gomod2nix.toml;
}
