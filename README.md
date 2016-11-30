# jimmify-server
A REST api to talk to jimmy

##Getting Started

In order to set up the repo:
* install go (I'm on 1.7)
* run ```go get -u```
* run ```go install```

##Static Site

This module builds the jimmify-web static site as a dependency, and is able to automatically build and serve it. Running ```go get -u``` will automatically pull the most recent jimmify-web.  

You will need [node](https://nodejs.org/en/) installed.

You will also need to install grunt: ```npm install -g grunt-cli```

##Usage

This server is built in Go and uses SQLite. In order to build and run use:

```bash
go install
jimmify-server
```

The server has two command line options:
* -resetdb - clears and sets up the SQL database.
* -log - turns on file logging

I have also built a CLI in python for testing the endpoints. It can be run using

```bash
python3 cli.py
```

##Documentation
* The API is fully documented in the wiki.
* Go documentation in Wiki.
