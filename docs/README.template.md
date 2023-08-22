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
If you are using Mac or Linux, [Homebrew](https://brew.sh/) is an easy way of installing connect-iq-sdk-manager.
```bash
brew install lindell/connect-iq-sdk-manager-cli/connect-iq-sdk-manager
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
````

### From source
You can also install from source with `go install`, this is not recommended for most cases.
```bash
go install github.com/lindell/connect-iq-sdk-manager-cli@latest
```

## Config file

All configuration can be done through command line flags, configuration files or a mix of both. If you want to use a configuration file, simply use the `--config=./path/to/config.yaml`. The file `~/.connect-iq-sdk-manager/config` be used for configuration. The priority of configs are first flags, then defined config file and lastly the static config file.

{{range .Commands}}
{{if .YAMLExample}}
<details>
  <summary>All available <code>{{.Name}}</code> options</summary>

```yaml
{{ .YAMLExample }}
```
</details>
{{end}}{{end}}

## Usage
{{range .Commands}}
* [{{ .Name }}](#-usage-of-{{ .Path }}) {{ .Short }}{{end}}

{{range .Commands}}
### Usage of `{{.Name}}`
{{.Long}}
```
{{.Usage}}
```

{{end}}
