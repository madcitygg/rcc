rcc
===
[![Build Status](https://travis-ci.org/madcitygg/rcc.svg)](https://travis-ci.org/madcitygg/rcc)
[![Test Coverage](https://img.shields.io/codecov/c/github/madcitygg/rcc.svg)](https://codecov.io/github/madcitygg/rcc)
[![GoDoc](https://godoc.org/github.com/madcitygg/rcc?status.svg)](https://godoc.org/github.com/madcitygg/rcc)
[![License](https://img.shields.io/github/license/madcitygg/rcc.svg)](https://github.com/madcitygg/rcc/blob/master/LICENSE.md)

The RemoteConsoleConsole! Control Source servers from a convenient cross-platform command line app.

Usage
-----
A simple example:
```
$ rcc 10.10.10.10:27015
Connecting to 10.10.10.10:27015
Connection successful!
Enter password: **********
10.10.10.10:27015> status
hostname: Example CS:GO
version : 1.35.2.0/13520 262/6283 secure  [G:1:201586]
udp/ip  : 10.10.10.10:27015  (public ip: 162.243.46.7)
os      :  Linux
type    :  community dedicated
map     : de_dust2
players : 0 humans, 0 bots (20/0 max) (hibernating)

# userid name uniqueid connected ping loss state rate adr
#end
L 02/15/2016 - 18:38:26: rcon from "73.51.200.225:59394": command "status"

10.10.10.10:27015> 
```
