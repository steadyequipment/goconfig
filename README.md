# goconfig
[![CircleCI](https://circleci.com/gh/steadyequipment/goconfig/tree/master.svg?style=svg)](https://circleci.com/gh/steadyequipment/goconfig/tree/master)
[![Travis CI](https://travis-ci.org/steadyequipment/goconfig.svg?branch=master)](https://travis-ci.org/steadyequipment/goconfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/steadyequipment/goconfig)](https://goreportcard.com/report/github.com/steadyequipment/goconfig)
[![codebeat badge](https://codebeat.co/badges/7d0b5194-cdbf-42a7-a4a0-d036359f308a)](https://codebeat.co/projects/github-com-steadyequipment-goconfig)
[![Coverage Status](https://coveralls.io/repos/github/steadyequipment/goconfig/badge.svg?branch=master)](https://coveralls.io/github/steadyequipment/goconfig?branch=master)

Poursteady's [golang](https://golang.org) config file and command line parsing.  

* Provides are verbose, simple, configurable way of pulling configuration information from both a config file and from the command line, with neither required
* Values specified in the command line override values specified in a config file
* Pretty prints usage ala [75lb/command-line-usage](https://github.com/75lb/command-line-usage)

Thanks to [spf13](https://www.github.com/spf13)'s wonderful [pflag](https://github.com/spf13/pflag) and [viper](https://github.com/spf13/viper) of which this library wraps.
