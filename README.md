# scopeX
This Go script is designed to assist with the management of scope in a bug bounty program. It accepts a list of subdomains as input via stdin and an exclude JSON object. The script processes the input and filters out any subdomains specified in the exclude JSON object. The resulting output is a list of subdomains that are not excluded and can be used as the updated scope for the bug bounty program. This script can be particularly useful in ensuring that out-of-scope domains are not included in the scope by mistake, allowing bounty hunters to focus their efforts on domains that are within the scope of the program.

## Install
```bash
go install github.com/xhzeem/scopeX@latest
```

## Examples

### Mode 1
Mode 1: filter by subdmoains, input [sub]

> all-xhzeem.me.txt
```
www.xhzeem.me
beta.xhzeem.me
test.beta.xhzeem.me
local.xhzeem.me
...
```

```bash
echo all-xhzeem.me.txt | scopeX -x '{"sub":["*beta.xhzeem.me"],"dns":["8.254.3.254", "cnm.github.com"]}' -m 1

# all-xhzeem.me.txt will have a list of subdomains, will filter and remove anything maches ^*.beta.xhzeem.me$
```

### Mode 2
Mode 1: filter by DNS record, input [sub ip,cname,txt]

> dns-xhzeem.me.txt
```
www.xhzeem.me 10.11.22.42,xhzeem.github.io
beta.xhzeem.me 101.44.11.88,90.144.112.90
test.beta.xhzeem.me 8.254.3.254,test.xhzeem.me
local.xhzeem.me 127.0.0.1
...
```

```bash
echo dns-xhzeem.me.txt | scopeX -x '{"sub":["*.beta.xhzeem.me"],"dns":["8.254.3.254", "cnm.github.com"]}' -m 2

# dns-xhzeem.me.txt will have a list of subdomains, then a space, then a comma separated DNS records, will filter and remove any in the list ["8.254.3.254", "cnm.github.com"]
```
