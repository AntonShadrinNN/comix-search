# Stem
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

Stem CLI reduces inflected words to their word stem. Currently [snowball module](https://github.com/kljensen/snowball.git) is used.

## Install
```sh
git clone git@github.com:AntonShadrinNN/comix-search.git
cd stem
make build && sudo make install
```

## Uninstall
```sh
cd stem
sudo make uninstall && make clean
```

## Usage
After installation you can simply run `stem -s="String to stem"` and see the result.

> [!IMPORTANT]
> Make sure to provide flag `-s`. It is mandatory
