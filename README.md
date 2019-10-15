# maildrop.cli
an unofficial cli client for maildrop.cc


### install

install `maildrop.cli` with `go get github.com/say4n/maildrop.cli`

### usage

```bash
▶ maildrop.cli
NAME:
   maildrop - an unofficial cli client to maildrop.cc

USAGE:
   maildrop.cli [global options] command [command options] [arguments...]

VERSION:
   v1 (build 2019-10-15T15:38:03+0530)

COMMANDS:
   inbox, i   show emails in inbox
   view, v    list emails from inbox
   delete, d  delete email from inbox
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --logging, -l  enable logging
   --help, -h     show help
   --version, -v  print the version

```

#### inbox

```bash
▶ maildrop.cli inbox
NAME:
   maildrop.cli inbox - show emails in inbox

USAGE:
   maildrop.cli inbox [command options] [arguments...]

OPTIONS:
   --address value, -a value  mailbox address

```

#### view

```bash
▶ maildrop.cli view
NAME:
   maildrop.cli view - read email from inbox

USAGE:
   maildrop.cli view [command options] [arguments...]

OPTIONS:
   --address value, -a value  mailbox address
   --uid value, -u value      unique ID of the email

```

#### delete

```bash
▶ maildrop.cli delete
NAME:
   maildrop.cli delete - delete email from inbox

USAGE:
   maildrop.cli delete [command options] [arguments...]

OPTIONS:
   --address value, -a value  mailbox address
   --uid value, -u value      unique ID of the email

```


## author

© 2019, Sayan Goswami
