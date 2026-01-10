# loocup
Find C function declarations by signature.

Essentially [Hoogle](https://hoogle.haskell.org/) for C. Inspired by [Coogle](https://www.youtube.com/playlist?list=PLpM-Dvs8t0VYhYLxY-i7OcvBbDsG4izam)

## Getting started
```sh
$ git clone https://github.com/lauchimoon/loocup.git
$ cd loocup/
$ go build .
```

## Usage
```sh
loocup <signature> <file>
  <signature> looks like type(arg1, arg2, ..., argN)
  <file> is a .c file or .h file
```
