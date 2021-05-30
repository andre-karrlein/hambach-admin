# hambach-admin
Hambach-admin is an installable Progressive web app (PWA) that is the administration tool for the spvgg hambach website.

## How to install

### Chrome

1. Go to [admin.spvgg-hambach.de](https://admin.spvgg-hambach.de)
2. Click on the `+` on the right, inside the search bar

### IOS

1. Go to [admin.spvgg-hambach.de](https://admin.spvgg-hambach.de) with Safari
2. Click on the `Share` button
3. Click on the `Add to homescreen` button

## Availablitiy

- Works on all browser that support web assembly but optimized for Chrome on a desktop

## How it is built

- User interface is built with [go-app](https://github.com/maxence-charriere/go-app)
- CSS framework are from [bulma.io](https://bulma.io)

## Suggestions

For reporting issues: open a [GitHub issue](https://github.com/andre-karrlein/hambach-admin/issues).

## Development

Instructions for local development.

### Requirements
**go-app** requirements:

- [Go 1.14](https://golang.org/doc/go1.14) or newer
- [Go module](https://github.com/golang/go/wiki/Modules)

```sh
go mod download
```
To set it up properly you need a GCP auth file as a json file.

### Commands

Build the http server binary.
```sh
make build
```

Build the webassembly binary.
```sh
make wasm
```

Run the webserver to serve the static files and the webassembly binary. The run command executes the build and wasm command beforehand.
```sh
make run
```