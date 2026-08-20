#!/usr/bin/env python

""" Build one Go source into an executable, reproducing the Makefile's
`go build -o <output> <input>`. Invoked by the generator as
go_build.py <input.go> <output.elf>. """

import os
import subprocess
import sys


def main():
    """ main entry point """
    source, output = sys.argv[1], sys.argv[2]
    os.makedirs(os.path.dirname(output), exist_ok=True)
    sys.exit(subprocess.call(["go", "build", "-o", output, source]))


if __name__ == "__main__":
    main()
