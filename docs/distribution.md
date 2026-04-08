# 📦 Distribution Process

This document outlines the build, packaging, and release process for `bananatype`.

## Arch Linux 

Distributed through the AUR.

1. Clone the existing AUR repository

```bash
git clone https://aur.archlinux.org/bananatype.git
```

2. Update the PKGBUILD file with new version and source URL

3. Push changes to AUR repository and create a new release on GitHub with the same version number.

```bash
git add PKGBUILD
git commit -m "Release vX.Y.Z"
git push origin master
```
