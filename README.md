# awswho
Quickly output the results of the AWS sts get-caller-identity API

**awswho** does nearly the same thing as running `aws sts get-caller-identity` but is about 75% faster since it is written in Go.  It is also portable since it can be compiled
into a static, stand-alone binary executable.

## Usage
```
Usage of awswho:
  -a	output Account as well as Arn and UserId
  -n	output a newline character
  -p string
    	an aws profile listed in ~/.aws/config (default "default")
  -r string
    	aws region (default "us-east-1")
  -v	output version and then exit
```

## Installation

* macOS: `breq update; brew install jftuga/tap/awsho`
* Binaries for Linux, macOS and Windows are provided in the [releases](https://github.com/jftuga/awswho/releases) section.

## Compilation

Static compilation can be achieved by running:

`CGO_ENABLED=0 go build -ldflags="-extldflags=-static"`

## Examples

```bash
# no newline, by default -- good for scripting
jftuga@ubuntu:~$ awswho
123456789012jftuga@ubuntu:~$

# use a named profile with -p and region with -r
jftuga@ubuntu:~$ awswho -p prod -r us-east-2 -n
098765432123

# display all identity info with -a, plus add a newline with -n
$ jftuga@ubuntu:~$ awswho -a -n
123456789012 arn:aws:iam::123456789012:user/default AID123CCCCC12345AAAAA
```
