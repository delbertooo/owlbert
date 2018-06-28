# Owlbert

Owlbert is a small guy - obviously an owl - who emits his calling on gitlab webhooks.

## How to build

```
go build owlbert.go
```

## How to start

For running on port 8080:
```
./owlbert 8080
```

## How to use

1. Place a file in `webhooks/` like the provided `my-project.json`.
2. Set up a custom, random secret token for keeping your owlbert private.
3. Modify the calling for each gitlab object kind you want to handle. Every value is a duration in milliseconds (negative = LOW = no calling, positive = HIGH = calling). E.g. `[1000, -500, 2000]` would emit the calling for a second, wait half a second doing nothing and then emit another two beautiful seconds of the calling.
4. Use the owlbert url as gitlab webhook, including your port and your filename without the `.json` ending, e.g. `http://localhost:8080/webhooks/my-project`.