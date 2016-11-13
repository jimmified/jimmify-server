# jimmify-server
A REST api to talk to jimmy

## Getting Started

In order to set up the repo:
* install go (I'm on 1.7)
* clone jimmify-server in to your go/src/ folder
* run ```go get```
* run ```go install```

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
