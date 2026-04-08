# 💿 Installation

## Table of Contents

- [Windows Installation](#-windows-installation)
- [Mac Installation](#-mac-installation)
- [Linux Installation](#-linux-installation)
    - [Arch](#-arch)
    - [Ubuntu](#-ubuntu)
    - [Debian](#-debian)
- [Other Install Methods](#-other-install-methods)
  - [Go Install (Recommended)](#-go-install-recommended)
  - [Manual Build](#-manual-build)
- [Compatibility and Support](#-compatibility-and-support)

## Windows Installation (WIP)

This section is a work in progress. For now use other install methods.

## Mac Installation (WIP)

This section is a work in progress. For now use other install methods.

## Linux Installation

If you do not see your distro listed, use the [Other Install Methods](#-other-install-methods).

### Arch

`bananatype` is available via the [Arch AUR](https://aur.archlinux.org/packages/bananatype)!

```
yay -S bananatype
```

## Other Install Methods

These install methods are for if you do not see your OS/distro listed above. It is recommended to use the above installation methods if possible to make your life easier :).

### Go Install 

The easiest way to install is via `go install`:

```bash
go install github.com/hungrymonkeystudio/bananas@latest
```

This will download the project and install the binary into your `$GOPATH/bin` folder. Make sure that folder is in your `$PATH`.

### Manual Build

Clone the repository and build from source:

```bash
git clone https://github.com/hungrymonkeystudio/bananas
cd bananas
go build -o bananas main.go
```

You can now run `./bananas` or move the binary to your local `/usr/local/bin` directory.

## Compatibility and Support (WIP)

This section is work in progress. For now, `bananas` is designed to be cross-platform and should work on any OS with Go support. If you encounter issues, please open an issue on GitHub!

