{ pkgs ? import ../../../nix { } }:
let bfhd = (pkgs.callPackage ../../../. { });
in
bfhd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-bfhd.patch
  ];
})
