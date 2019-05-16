# snowboard

[![Build Status](https://travis-ci.org/bukalapak/snowboard.svg?branch=master)](https://travis-ci.org/bukalapak/snowboard)
[![GoDoc](https://godoc.org/github.com/kppk/snowboard?status.svg)](https://godoc.org/github.com/kppk/snowboard)
[![Docker Repository on Quay](https://quay.io/repository/bukalapak/snowboard/status)](https://quay.io/repository/bukalapak/snowboard)
[![GitHub release](https://img.shields.io/github/release/bukalapak/snowboard.svg)](https://github.com/kppk/snowboard)

API blueprint toolkit.

## Installation

The latest executables for supported platforms are available from the [release page](https://github.com/kppk/snowboard/releases).

Just extract and start using it:

```
$ wget https://github.com/kppk/snowboard/releases/download/${version}/snowboard-${version}.${os}-${arch}.tar.gz
$ tar -zxvf snowboard-${version}.${os}-${arch}.tar.gz
$ ./snowboard -h
```

Alternatively, you can also use options below:

### Homebrew

```sh
$ brew tap bukalapak/packages
$ brew install snowboard
```

> Note: If you want build from master branch you can use `brew install --HEAD snowboard`

### Arch Linux

Snowboard available as [AUR package](https://aur.archlinux.org/packages/snowboard/).

```sh
$ pacaur -S snowboard
```

### Docker

You can also use automated build docker image on `quay.io/bukalapak/snowboard`:

```
$ docker pull quay.io/bukalapak/snowboard
$ docker run -it --rm quay.io/bukalapak/snowboard help
```

To run snowboard with the current directory mounted to `/doc`:

```
$ docker run -it --rm -v $PWD:/doc quay.io/bukalapak/snowboard html -o output.html API.apib
```

### Manual

```sh
$ git clone https://github.com/kppk/snowboard.git
$ cd snowboard
$ make install
```

> Note: ensure you have installed [Go](https://golang.org/doc/install#tarball) and configured your `GOPATH` and `PATH`.

## Usage

Let's say we have API Blueprint document called `API.apib`, like:

```apib
# API
## GET /message
+ Response 200 (text/plain)

        Hello World!
```

There are some scenarios we can perform:


### Generate HTML Documentation

To generate HTML documentation we can do:

```
$ snowboard html -o output.html API.apib
```

Above command will generate `ouput.html` using `snowboard` default template (called `alpha`).

### Using Custom Template

If you want to use custom template, you can use flag `-t` for that:

```
$ snowboard html -o output.html -t awesome-template.html API.apib
```

To see how the template looks like, you can see `snowboard` default template located in [templates/alpha.html](templates/alpha.html).

### Serve HTML Documentation

If you want to access HTML documentation via HTTP, especially on local development, you can pass `-s` flag:

```
$ snowboard html -o output.html -t awesome-template.html -s API.apib
```

With this flag, You can access HTML documentation on `localhost:8088`.

If you need to customize binding address, you can use flag `-b`.

#### Auto-regeneration

To enable auto-regeneration on both input and template file updates, you can add global flag `--watch`

```
$ snowboard --watch html -o output.html -t awesome-template.html -s API.apib
```

Optionally, you can also use `--watch-interval` to enable polling interval.

```
$ snowboard --watch --watch-interval 100ms html -o output.html -t awesome-template.html -s API.apib
```

#### Serve HTML from Docker container

If you want to serve HTML documentation from Docker container, don't forget to bind address and port in the contaier plus bind ports of host and container by `-p` option of Docker command.

```
$ docker run -it --rm -v $(pwd):/doc -p 8088:8088 bukalapak/snowboard html -o output.html -b 0.0.0.0:8088 -s API.apib
```

### Generate formatted API blueprint

When you have documentation splitted across files, you can customize flags `-o` to allow `snowboard` to produce single formatted API blueprint.

```
$ snowboard apib -o API.apib project/splitted.apib
```

### Validate API blueprint

Besides render to HTML, snowboard also support validates API blueprint document. You can use `lint` subcommand.

```
$ snowboard lint API.apib
```

### Mock server from API blueprint

Another snowboard useful feature is having mock server. You can use `mock` subcommand for that.

```
$ snowboard mock API.apib
```

Then you can use `localhost:8087` for accessing mock server. You can customize the address by passing flag `-b`.

For multiple responses, you can set `X-Status-Code` or `Prefer` header to select specific response:

```
X-Status-Code: 200

// or

Prefer: status=200
```

## External Files

You can split your API blueprint document to several files and use `partial` helper to includes it to your main document.

```
{{partial "some-resource.apib"}}
```

Alternatively, you can also use HTML comment syntax to include those files:

```html
<!-- partial(some-resource.apib) -->
```

or

```html
<!-- include(some-resource.apib) -->
```

## Seed Files

As your API blueprint document become large, you might move some value to separate file for easier organization and modification. Snowboard supports this as well.

Just place your values into a json file, say, `seed.json`:

```json
{
  "official": {
    "username": "olaf"
  }
}
```

Then on your API blueprint document you can use `seed` comment helper:

```apib
# API

<!-- seed(seed.json) -->

Our friendly username is {{.official.username}}.
```

Multiple seeds are also supported.

## API Element JSON

In case you need to get API element JSON output for further processing, you can use:

```
$ snowboard json API.apib
```

## Help

As usual, you can also see all supported flags by passing `-h`:

```
$ snowboard help
NAME:
   snowboard - API blueprint toolkit

USAGE:
   snowboard [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]

COMMANDS:
     lint     Validate API blueprint
     html     Render HTML documentation
     apib     Render API blueprint
     json     Render API element json
     mock     Run Mock server
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## FAQ

- I am using Vim and snowboard file watcher doesn't trigger auto-regeneration, any idea?

  It is known issue due Vim backup scheme. You can set on your `.vimrc`:
	
    ```
    set nobackup
    set nowritebackup
    ```

## Examples

You can see examples of `snowboard` default template outputs, in [examples/alpha](examples/alpha) directory. They looks like:

- [Named Resource and Actions](https://htmlpreview.github.io/?https://github.com/kppk/snowboard/blob/master/examples/alpha/03.%20Named%20Resource%20and%20Actions.html)
- [Real World API](https://htmlpreview.github.io/?https://github.com/kppk/snowboard/blob/master/examples/alpha/Real%20World%20API.html)
- And many more...

All of the examples are generated from official [API Blueprint examples](https://github.com/apiaryio/api-blueprint/tree/master/examples)
