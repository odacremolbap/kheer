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

- `clean`

    Deletes `./_output` folder.

- `build`

    Is the default taget. It compiles an amd64 binary for your current OS.
    Computer architecture and OS can be set using `OS` and `ARCH` environment variables. A docker Go image is used for compiling, and the result is stored at `./_output`.

- `build-all`

    Builds the binary for all supported OSes and architectures.

- `check`

    Perfoms unit tests, race detection, vet, and format checks.
