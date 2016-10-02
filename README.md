# File-Presence-Covert-Channel
Implementation of a simple covert channel.  Contains a sender and a receiver.  The two processes communicate on the same system via the presence of certain files on the system.  This covert channel is a timing channel.  Once the transmit file has been set, the receiver will start listening for bits until it is removed.  The receiver must be started before the sender.

##Compilation instructions
Build the following files

```
$ go build receiver.go
$ go build sender.go
```
