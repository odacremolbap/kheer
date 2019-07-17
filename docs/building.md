# Building

Kheer is a Go application that runs on Kubernetes using the operator pattern.

## Prerequisites

1. Go compiler version 1.12, Kheer uses Go modules.
2. Make
3. Docker

## Source

Clone the repository

```sh
git clone git@github.com:kheer/kheer.git
```

If you are contributing, please fork the project and add both repositories as remotes

```sh
$ # check the existing remote, here pointing to upstream
$ git remote -v
origin  git@github.com:kheer/kheer.git (fetch)
origin  git@github.com:kheer/kheer.git (push)

$ # rename remote as upstream
$ git remote rename origin upstream

$ # add new remote pointing to your fork
$ git remote add origin git@github.com:<YOUR-USER>/kheer.git
```

## Make targets

_User facing_ targets are

- `clean`

    Deletes `./_output` folder.

- `all` or `build`

    Is the default target. It compiles an amd64 binary for your current OS.
    Computer architecture and OS can be set using `OS` and `ARCH` environment variables. A docker Go image is used for compiling, and the result is stored at `./_output`.

    Compilation process creates a docker image (`COMPILER_IMAGE` at Makefile) that includes all dependencies for caching.

- `build-all`

    Builds the binary for all supported OSes and architectures.

- `check`

    Perfoms unit tests, race detection, vet, and format checks.

- `clean-compiler-images`

    Builds generate a Go based image that includes cached dependencies. Compiling cached images can be removed using this target. This might come handy when changing the `GO_IMAGE` tag when using a new Go version, to delete existing images that will no longer be used.

    If the image is deleted, a new one will be generated at the next build.
