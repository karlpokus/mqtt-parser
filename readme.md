# mqtt-parser
Parse mqtt messages into nicely formatted lines like `<from>:<to> <ControlPacketType`. See http://docs.oasis-open.org/mqtt/mqtt/v3.1.1/os/mqtt-v3.1.1-os.html#_Toc398718027 for details.

Only works for unencrypted traffic. Requires `ngrep`. Useful in debugging connection reconnect attempts.

# usage
````bash
# assuming you have a broker listening on port 1883 on localhost
$ sudo ngrep -d lo -x port 1883 | go run .
````

# test
````bash
$ go test
````

# todo
- [x] tests. see https://github.com/karlpokus/logspam/blob/master/logspam_test.go
- [ ] decrypt tls
- [ ] parse client_id from CONNECT and use in output
- [ ] replace ngrep w tcpdump
- [ ] consider encoding bytes from `net.Conn` to hex instead in https://github.com/karlpokus/prxy
