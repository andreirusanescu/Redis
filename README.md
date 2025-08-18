## Redis
Redis clone that allows you to use strings and hashes.

Written a parser to understand **RESP**, which allows the server to receive commands and respond with responses.

Used *GO Routines* to handle multiple connections simultaneously.

Written data to disk using the *Append Only File* (AOF), which is one of the methods Redis uses for persistence.
This way, if the server crashes or restarts, we can restore the data.
