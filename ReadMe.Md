![](l0la-100.png)

l0la is a simple cli utility that combines the output of `lsof`, `ps` and `netstat` to continuously monitor cpu, 
threads, files, network connections of your process for debugging purposes.

[![Circleci Builds](https://circleci.com/gh/simonmittag/l0la.svg?style=shield)](https://circleci.com/gh/simonmittag/l0la)
[![Github Issues](https://img.shields.io/github/issues/simonmittag/l0la)](https://github.com/simonmittag/l0la/issues)
[![Github Activity](https://img.shields.io/github/commit-activity/m/simonmittag/mse6)](https://img.shields.io/github/commit-activity/m/simonmittag/l0la)  
[![Codeclimate Maintainability](https://api.codeclimate.com/v1/badges/4f80639bc17c6987eb05/maintainability)](https://codeclimate.com/github/simonmittag/l0la/maintainability)
[![Codeclimate Test Coverage](https://api.codeclimate.com/v1/badges/4f80639bc17c6987eb05/test_coverage)](https://codeclimate.com/github/simonmittag/l0la/test_coverage)
[![Go Version](https://img.shields.io/github/go-mod/go-version/simonmittag/l0la)](https://img.shields.io/github/go-mod/go-version/simonmittag/l0la)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/badge/version-0.1.3-orange)](https://github.com/simonmittag/l0la/releases/tag/v0.1.3)

## What's New
### v0.1.3
* bumped to go 1.20

## Up and running
### Homebrew
```
brew tap simonmittag/cli &&
  brew install l0la &&
  l0la 
```

## Usage
```
λ l0la -h
<l0la> v0.1.3
Usage: l0la [-v] [-h] [PID]
  -h    print help
  -v    print version
```

## Samples

Monitor a process by pid
```
l0la $pid
```

Monitor a process by name
```
pgrep j8a | xargs l0la
```

![](terminal.gif)

## Contributions
The l0la team welcomes all [contributors](https://github.com/simonmittag/mse6/blob/master/CONTRIBUTING.md). Everyone interacting with the project's codebase, issue trackers, chat rooms and mailing lists
is expected to follow the [code of conduct](https://github.com/simonmittag/mse6/blob/master/CODE_OF_CONDUCT.md)
