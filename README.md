# scopeX
Golang script to exclude out-of-scope from a list of subdomains

### Modes Available:
```js
Mode 1: filter by subdmoains, input [sub]
Mode 2: filter by IP address, input [sub ip,ip]
Mode 3: filter by cname, input: [sub cname,cname]
```

### Examples:
```bash
echo all-xhzeem.me.txt | scopeX -x '{"sub":["*.beta.xhzeem.me"],"ip":["8.254.3.254", "23.65.242.19"]}' -m 1

# all-xhzeem.me.txt will have a list of subdomains, will filter and remove anything maches ^*.beta.xhzeem.me$
```

```bash
echo all-xhzeem.me.txt | scopeX -x '{"sub":["*.beta.xhzeem.me"],"ip":["8.254.3.254", "23.65.242.19"]}' -m 2

# all-xhzeem.me.txt will have a list of subdomains, then a space, then a comma separated IPs, will filter and remove any subdomain with the IP in the list ["8.254.3.254", "23.65.242.19"]
```