# goconfig
[![CircleCI](https://circleci.com/gh/steadyequipment/goconfig/tree/master.svg?style=svg)](https://circleci.com/gh/steadyequipment/goconfig/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/steadyequipment/goconfig)](https://goreportcard.com/report/github.com/steadyequipment/goconfig)

Poursteady's [golang](https://golang.org) config file and command line parsing.  

* Provides are verbose, simple, configurable way of pulling configuration information from both a config file and from the command line, with neither required
* Values specified in the command line override values specified in a config file
* Pretty prints usage ala [75lb/command-line-usage](https://github.com/75lb/command-line-usage)

Thanks to [spf13](https://www.github.com/spf13)'s wonderful [pflag](https://github.com/spf13/pflag) and [viper](https://github.com/spf13/viper) of which this library wraps.
