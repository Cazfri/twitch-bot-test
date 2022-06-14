# Twitch Chat Test Server

This quick proof-of-concept will simultaneously monitor a Twitch chat and run a websocket server. The server will distribute all chat messages from Twitch in a standardized format to a client that connects via websocket to `/messages`.

## How to run

First, modify `config/config.go` to specify the Twitch chat that the server should pull from, as well as the port to run the message server on, and a list of allowed commands (only allowed commands will be passed through to clients).

Once that's all done, pull dependencies with

```
go get ./...
```

Then run the server with

```
go run main.go
```

You can also start up a test client that will connect to the server and consume messages, make sure the server is running first, then run

```
go run testclient/ws_client.go
```

## Caveats and TODOs

There are plenty of things I would want to add before calling this done, most of which are marked with `TODO` comments.

Probably the biggest one is that all connected clients currently share the same queue of messages. This is not ideal, as this means the server can only really serve one client completely at a time. See `main.go` for a (someone discombobulated) description of how I might solve this in the future.