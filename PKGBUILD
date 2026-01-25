# Maintainer: Warren Wu <warrenweiwu04@gmail.com>
pkgname=bananatype
pkgver=0.0.1
pkgrel=1
pkgdesc="monkeytype in your terminal"
arch=('x86_64')
url="https://github.com/hungrymonkeystudio/bananas"
license=('MIT')
depends=()
makedepends=('git' 'go')
source=("git+${url}.git#tag=v${pkgver}")
sha256sums=('SKIP')
options=('!debug' '!strip')

build() {
    cd "$srcdir/$pkgname"
    go build -ldflags="-X main.Build=prod" -o main
}

package() {
    cd "$srcdir/$pkgname"
    
    # install binary to /usr/bin
    # -D flag automatically creates directory
    install -Dm755 main "$pkgdir/usr/bin/$pkgname"

    # move resources to /usr/share/wallflower
    install -d "$pkgdir/usr/share/$pkgname"
    cp -r resources/common-words.txt "$pkgdir/usr/share/$pkgname/"

    # REQUIRED: install License file
    install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
