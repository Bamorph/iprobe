# iprobe

Take a list of domains and probe for ip addresses.

## Install

```
go install github.com/Bamorph/iprobe@latest
```

## Basic Usage

iprobe accepts line-delimited domains on stdin:

```
cat hosts.txt
localhost

cat hosts.txt | iprobe
127.0.0.1
```
