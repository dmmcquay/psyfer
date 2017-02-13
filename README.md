# Psyfer

[![GoDoc](https://godoc.org/github.com/dmmcquay/psyfer?status.svg)](https://godoc.org/github.com/dmmcquay/psyfer)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmmcquay/psyfer)](https://goreportcard.com/report/github.com/dmmcquay/psyfer)
 
# Overview

Psyfer is a app which lets you perform various different ciphers to input 
strings.  It also provides the ability to guess certain ciphers. 


# Getting Started

If you have a working Go enviroment on your computer, you can download pull 
down this repo and build this code. 

Once you have psyfer, you can run commands like:

    > psyfer vig -k xyz "this is my string"

which will perform the vigenere cipher with a key of "xyz" on the given input

# Usage

## psyfer trans

Running `psyfer trans [mode] [flags] [input]` will perform various different 
transposition ciphers to your input string.  the possible options are `random`, 
`railfence`, and `split`.

Here is an example:

```
psyfer trans railfence "my input string"
```

Likewise, you can also decrypt by adding the `-d` flag:

```
psyfer trans railfence "m nu tigyiptsrn" -d
```

## psyfer sub
Running `psyfer sub [flags] [input]` will perform a substitution cipher according
to a key file.  An example key file is found in the repo called key.json.

Here is an example:

```
psyfer sub -k key.json "my input string"
```

Likewise, you can also decrypt by adding the `-d` flag:

```
psyfer sub -k key.json "ax nmpir krqnmb" -d
```

## psyfer vig

Running `psyfer vig [flags] [input]` will perform the vigenere cipher with 
the provided key.  

Here is an example:

```
psyfer vig -k mykey "my input string"
```

Likewise, you can also decrypt by adding the `-d` flag:

```
psyfer vig -k mykey "decrypt this string" -d
```

## psyfer aes

Running `psyfer aes [flags] [input]` will perform the AES cipher with keysizes 
128, 192 or 256 bit in Electronic Codebook (ECB). The key is "baked" into the
program (example keys found in FIPS publication 197.

The default behavior is to return the hex representation of the string. the
`-a` flag will display in ascii.

Here is an example:

```
psyfer aes -k 128 "my input string"
```

Likewise, you can also decrypt by adding the `-d` flag:

```
psyfer aes -k 128 "my input string" -d
```

## psyfer guess

Running `psyfer guess [flags] [input]` will return the top five most likely
solutions to the caesar cipher provided.

Here is an example:

```
psyfer guess "wklv lv d vhfuhw phvvdjh"
``` 

which will return "THISISASECRETMESSAGE" as the most likely answer
