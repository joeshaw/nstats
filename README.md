# nstats

`nstats` is a Go program to collect simple numerical statistics based
on input from standard input.

`nstats` splits input on any whitespace.

## Example usage

    $ cat data.txt
    3 57 1 17 7 13 2 25 3 15 26 21 23 [...]

    $ cat data.txt | nstats
    N       1120
    min     1
    max     260
    sum     11625
    median  4
    mean    10.379464285714286
    stddev  19.0384524043852

