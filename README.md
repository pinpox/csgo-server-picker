# csgo-server-picker

A blocker for specified IPs

This is a project stub, nothing is implemented yet!

The executable has three sub-commands

#### ./exec update

Retrieves the latest server list from https://github.com/SteamDatabase/SteamTracking/blob/master/Random/NetworkDatagramConfig.json

```
./exec update
```

#### ./exec block
Lets you block a server or a rang by name

```
./exec block "amsterdam"
```

#### ./exec cleanup
Removes all created blocking rules. This is only needed if the program crashes or is killed by the user.

```
./exec cleanup
```
