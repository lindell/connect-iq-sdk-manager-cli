<h1 align="center">
  connect-iq-sdk-manager
</h1>

connect-iq-sdk-manager handles and downloads SDKs and other resources connected to ConnectIQ development. It aims to have parity with the official GUI SDK manager, but have some additional features such as only downloading devices from a specified manifest file.

## How to use

### Use on your local machine
```bash
connect-iq-sdk-manager agreement view # View and read through the agreement
connect-iq-sdk-manager agreement accept # Accept the license agreement
connect-iq-sdk-manager login
connect-iq-sdk-manager sdk set 6.2.2 # Downloads as sets the sdk as the current one
export PATH=`connect-iq-sdk-manager sdk current-path --bin`:$PATH # Make the SDK binaries can be callable
connect-iq-sdk-manager device download --manifest=your-app/manifest.xml # Download the devices used in your project

# You can now use the SDK to build your project. Here we build an IQ file.
monkeyc -f your-app/monkey.jungle -w -e -o output.iq
```

### Use in CI/CD

First, view and read through the agreement on your local machine:
```bash
connect-iq-sdk-manager agreement view
```

Take the acceptance hash to the next step
```bash
# Accept the license agreement and login
# To ensure we accept the latest agreement, use the `--acceptance-hash` flag
connect-iq-sdk-manager agreement accept --acceptance-hash=<THE HASH FROM THE PREVIOUS STEP>
export GARMIN_USERNAME="<YOUR GARMIN USERNAME>"
export GARMIN_PASSWORD="<YOUR GARMIN PASSWORD>"
connect-iq-sdk-manager login

connect-iq-sdk-manager sdk set ^6.2.2 # Downloads as sets the sdk as the current one
export PATH=`connect-iq-sdk-manager sdk current-path --bin`:$PATH # Make the SDK binaries can be callable
connect-iq-sdk-manager device download --manifest=your-app/manifest.xml # Download the devices used in your project

# You can now use the SDK to build your project. Here we build an IQ file.
monkeyc -f your-app/monkey.jungle -w -e -o output.iq
```

## Install

### Homebrew
If you are using Linux, [Homebrew](https://brew.sh/) is an easy way of installing connect-iq-sdk-manager.
```bash
brew install lindell/connect-iq-sdk-manager-cli/connect-iq-sdk-manager-cli
```

### Manual binary install
Find the binary for your operating system from the [release page](https://github.com/lindell/connect-iq-sdk-manager-cli/releases) and download it.

### Automatic binary install
To automatically install the latest version
```bash
curl -s https://raw.githubusercontent.com/lindell/connect-iq-sdk-manager-cli/master/install.sh | sh
```
Or a specific version.
```bash
curl -s https://raw.githubusercontent.com/lindell/connect-iq-sdk-manager-cli/master/install.sh | sh -s -- -d vX.X.X
```

### From source
You can also install from source with `go install`, this is not recommended for most cases.
```bash
go install github.com/lindell/connect-iq-sdk-manager-cli@latest
```

## Config file

All configuration can be done through command line flags, configuration files or a mix of both. If you want to use a configuration file, simply use the `--config=./path/to/config.yaml`. The file `~/.connect-iq-sdk-manager/config` be used for configuration. The priority of configs are first flags, then defined config file and lastly the static config file.



<details>
  <summary>All available <code>agreement accept</code> options</summary>

```yaml
# The hash of a previously read agreement.
agreement-hash:

# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info
```
</details>


<details>
  <summary>All available <code>agreement view</code> options</summary>

```yaml
# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info
```
</details>


<details>
  <summary>All available <code>device download</code> options</summary>

```yaml
# The device(s) that should be used.
device:
  - example

# Path to the manifest file.
download-all: false

# Download the fonts used for simulating the downloaded devices.
include-fonts: false

# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info

# Path to the manifest file.
manifest:
```
</details>


<details>
  <summary>All available <code>device list</code> options</summary>

```yaml
# The device(s) that should be used.
device:
  - example

# Path to the manifest file.
download-all: false

# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info

# Path to the manifest file.
manifest:
```
</details>


<details>
  <summary>All available <code>login</code> options</summary>

```yaml
# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info

# The Garmin password.
password:

# The Garmin username.
username:
```
</details>


<details>
  <summary>All available <code>sdk current-path</code> options</summary>

```yaml
# Print binary path
bin: false

# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info
```
</details>


<details>
  <summary>All available <code>sdk download</code> options</summary>

```yaml
# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info
```
</details>


<details>
  <summary>All available <code>sdk list</code> options</summary>

```yaml
# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info
```
</details>


<details>
  <summary>All available <code>sdk set</code> options</summary>

```yaml
# The file where all logs should be printed to. "-" means stdout.
log-file: "-"

# The formatting of the logs. Available values: text, json, json-pretty.
log-format: text

# The level of logging that should be made. Available values: trace, debug, info, error.
log-level: info
```
</details>


## Usage

* [agreement accept](#usage-of-agreement-accept) Accept the SDK agreement.
* [agreement view](#usage-of-agreement-view) View the SDK agreement.
* [device download](#usage-of-device-download) Download devices.
* [device list](#usage-of-device-list) List devices.
* [login](#usage-of-login) Login to be able to use some parts of the manager.
* [sdk current-path](#usage-of-sdk-current-path) Print the path to the currently active SDK
* [sdk download](#usage-of-sdk-download) Download an SDK. Without setting it as the current one.
* [sdk list](#usage-of-sdk-list) List SDKs.
* [sdk set](#usage-of-sdk-set) Set which SDK version to be used.


### Usage of `agreement accept`
Accept the the Garmin CONNECT IQ SDK License Agreement and CONNECT IQ Application Developer Agreement.
		
You can either accept the latest, just read agreement. Or accept a previously read agreement, if used in for example CI/CD.
```
Usage:
  connect-iq-sdk-manager agreement accept [flags]

Flags:
  -H, --agreement-hash string   The hash of a previously read agreement.

Global Flags:
      --config string       Path of the config file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
```


### Usage of `agreement view`
View the the Garmin CONNECT IQ SDK License Agreement and CONNECT IQ Application Developer Agreement.
```
Usage:
  connect-iq-sdk-manager agreement view [flags]

Global Flags:
      --config string       Path of the config file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
```


### Usage of `device download`
Download devices.

Either all devices can be chosen, all devices defined in a manifest file, or a list of devices.
```
Usage:
  connect-iq-sdk-manager device download [flags]

Flags:
  -F, --include-fonts   Download the fonts used for simulating the downloaded devices.

Global Flags:
      --config string       Path of the config file.
  -d, --device strings      The device(s) that should be used.
      --download-all        Path to the manifest file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
  -m, --manifest string     Path to the manifest file.
```


### Usage of `device list`
List devices.

Either all devices can be chosen, all devices defined in a manifest file, or a list of devices.
```
Usage:
  connect-iq-sdk-manager device list [flags]

Global Flags:
      --config string       Path of the config file.
  -d, --device strings      The device(s) that should be used.
      --download-all        Path to the manifest file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
  -m, --manifest string     Path to the manifest file.
```


### Usage of `login`
Login to be able to use some parts of the manager.

If used as is, you will be asked to login via the Garmin SSO OAuth flow.
Credentials can also be set via the --username and --password config,
or GARMIN_USERNAME and GARMIN_PASSWORD environment variables.
```
Usage:
  connect-iq-sdk-manager login [flags]

Flags:
      --password string   The Garmin password.
      --username string   The Garmin username.

Global Flags:
      --config string       Path of the config file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
```


### Usage of `sdk current-path`

```
Usage:
  connect-iq-sdk-manager sdk current-path [flags]

Flags:
      --bin   Print binary path

Global Flags:
      --config string       Path of the config file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
```


### Usage of `sdk download`
Download an SDK. Without setting it as the current one.

The version argument can be a specific version or a semver-range.
For example: ^6.2.0 or >=4.0.0 or 4.2.1
```
Usage:
  connect-iq-sdk-manager sdk download version [flags]

Global Flags:
      --config string       Path of the config file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
```


### Usage of `sdk list`
List SDKs.

To only list certain versions. The version argument can be used with a semver-range.
For example: ^6.2.0 or >=4.0.0
```
Usage:
  connect-iq-sdk-manager sdk list [version] [flags]

Global Flags:
      --config string       Path of the config file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
```


### Usage of `sdk set`
Set which SDK version to be used. If it does not exist, it will be downloaded.

The version argument can be a specific version or a semver-range.
For example: ^6.2.0 or >=4.0.0 or 4.2.1
```
Usage:
  connect-iq-sdk-manager sdk set version [flags]

Global Flags:
      --config string       Path of the config file.
      --log-file string     The file where all logs should be printed to. "-" means stdout. (default "-")
      --log-format string   The formatting of the logs. Available values: text, json, json-pretty. (default "text")
  -L, --log-level string    The level of logging that should be made. Available values: trace, debug, info, error. (default "info")
```


