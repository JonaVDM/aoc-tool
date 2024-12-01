{
  description = "Aoc Tool";

  inputs.nixpkgs.url = "nixpkgs/nixos-24.11";

  outputs = {
    self,
    nixpkgs,
  }: let
    lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
    version = builtins.substring 0 8 lastModifiedDate;

    supportedSystems = ["x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin"];
    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    nixpkgsFor = forAllSystems (system: import nixpkgs {inherit system;});
  in {
    packages = forAllSystems (system: let
      pkgs = nixpkgsFor.${system};
    in {
      default = pkgs.buildGoModule {
        pname = "aoc-tool";
        inherit version;
        src = ./.;

        vendorHash = "sha256-ermZ3cA2FSOMsxkNVbXujcL2fEFhDUe3ut7imdl4Af0=";
      };
    });

    devShells = forAllSystems (system: let
      pkgs = nixpkgsFor.${system};
    in {
      default = pkgs.mkShell {
        packages = with pkgs; [
          go
        ];
      };
    });

    overlays.default = final: _prev: {
      oac-tool = self.packages.${final.system}.default;
    };
  };
}
