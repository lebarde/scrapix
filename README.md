Scrapix is a command-line tool that periodically checks the changes to web pages.

![url du logo]

[Presentation](#presentation) | [Quick install guide](#quick-install Get started easily) | [Code](http://github.com/lebarde/scrapix) | [Contribute](#contribute)

## Overview

Scrapix downloads all web pages you want to watch/monitor and tells you if there were changes.

As a musician, I was bored of repeatedly using my F5 key to check if this or that opera had a job proposal to offer. So I created Scrapix.

Practically, you write a configuration file and Scrapix saves all changes in a database. Under Linux, configuration file is supposed to be in `$HOME/.config/scrapix/config.yaml`, and the cache database is in `$HOME/.cache/scrapix/cache.db`.

Scrapix is free and open-source, and always will. Contributions are welcome, and don't hesitate to open issues if you find a bug!

Technically, the project is written in Go language and uses SQLite for its cache database.

## Supported architectures
Currently, Scrapix can run under Linux. It might work under *BSD systems too, but they have not been tested.

Windows, Mac OS X and Android are not supported, but they are on the [roadmap](TODO link) and contributions are welcome.

## Quick start guide

### Binary install
Scrapix works well as a lone binary. Just dowload it (here)(TODO link), make it executable and add it in the $PATH. Run the following in a terminal:

```
$ mkdir -p $HOME/bin
$ wget TODO-URL
$ chmod +x scrapix
$ echo 'PATH=$PATH;$HOME/bin'
```

Install done! You can now run scrapix as this:

```
$ scrapix
```

### Install from source

The following packages need to be installed on your machine: git, go and eventually build-essential if you are on a debian machine. The Golang environment needs to be set (GOROOT and so on). See [here](TODO link) for details.

After having set your Go environment, run the following:

```
$ go install -v github.com/scrapix
```

Here it is!

### Play with Scrapix

Create a config file inside `$HOME/.config/scrapix` and put the urls you want to watch:

```
params: "whatever (TODO)"

urls:
  "http://www.topologix.fr/exercices/":
    refresh: 1d
    watch: ".wrapper"
  "http://www.topologix.fr/sujets/":
```

Then run `~$ scrapix`.


## Usage and functionalities

## Contribute
