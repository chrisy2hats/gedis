gedis
A redis-like program supporting a subset of redis commands with the same syntax.
Just a project for learning Golang, Redis and Sockets. Don't ever use this in production!

DONE
Handle multiple clients at once
get
set 
keys 

TODO
Protect against concurrent access
 
Lists
rpush
lpop
lindex
lrange

Hashes (maybe)
HMSET 
HGETALL

Tidyup:
Avoid switch statement in HandleConnection
DRY out some of the parsing code
Report instruction parsing errors back to client. Make JSON always returned so a proper error field can exist. Returning unstructed strings makes this a PIA
