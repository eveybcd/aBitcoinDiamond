aBcd
====

aBcd (advanced bitcoin diamond) is an alternative full node bitcoin diamond implementation written in Go (golang).

Bitcoin Diamond (BCD) is a fork of Bitcoin that occurs at the predetermined height of block 495866 
and therewith a new chain will be generated as the BCD. Bitcoin Diamond miners will begin creating 
blocks with a new proof-of-work algorithm, and will consecutively develop and enhance the protection 
for account transfer and privacy based on original features of BTC. This will cause a bifurcation of 
the Bitcoin blockchain. The original Bitcoin blockchain will continue on unaltered, but a new branch
of the blockchain will split off from the original chain. It shares the same transaction history 
with Bitcoin until it starts branching and coming into a unique block from which it diverges. 
As a result of this process, a new cryptocurrency was created which we call “Bitcoin Diamond”.
For more information, see https://www.bitcoindiamond.org

This project is currently under active development and forked from btcd, for more btcd information you can refer https://github.com/btcsuite/btcd.


## Requirements

[Go](http://golang.org) 1.11 or newer.

## Installation

#### Windows - MSI Available

https://github.com/eveybcd/aBitcoinDiamond/releases

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

- Run the following commands to obtain btcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/aBitcoinDiamond
$ make
$ make install

```

- aBcd (and utilities) will now be installed in ```$GOPATH/bin```.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Run the following commands to update aBcd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/aBitcoinDiamond
$ git pull
$ make
$ make install

```

## Getting Started

aBcd has several configuration options available to tweak how it runs, but all
of the basic operations described in the intro section work with zero
configuration.

#### Windows (Installed from MSI)

Launch aBcd from your Start menu.

#### Linux/BSD/POSIX/Source

```bash
$ ./aBcd
```

## Documentation

The documentation is a work-in-progress.  It is located in the [docs](https://github.com/eveybcd/aBitcoinDiamond/tree/master/docs) folder.


## License

aBcd is licensed under the [copyfree](http://copyfree.org) ISC License.
