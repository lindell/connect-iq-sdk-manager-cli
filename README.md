<h1 align="center">
  connect-iq-sdk-manager
</h1>

connect-iq-sdk-manager description.

## Install

### Homebrew
If you are using Mac or Linux, [Homebrew](https://brew.sh/) is an easy way of installing connect-iq-sdk-manager.
```bash
brew install lindell/connect-iq-sdk-manager-cli/connect-iq-sdk-manager
```

### Manual binary install
Find the binary for your operating system from the [release page](https://github.com/lindell/connect-iq-sdk-manager-cli/releases) and download it.

### From source
You can also install from source with `go install`, this is not recommended for most cases.
```bash
go install github.com/lindell/connect-iq-sdk-manager-cli@latest
```

## Config file

All configuration can be done through command line flags, configuration files or a mix of both. If you want to use a configuration file, simply use the `--config=./path/to/config.yaml`. The file `~/.connect-iq-sdk-manager/config` be used for configuration. The priority of configs are first flags, then defined config file and lastly the static config file.




## Usage

* [version](#-usage-of-version) Get the version of connect-iq-sdk-manager.


### Usage of `version`
Get the version of connect-iq-sdk-manager.
```
Usage:
  connect-iq-sdk-manager version
```


