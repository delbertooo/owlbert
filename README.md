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



## Example installation on a RaspberryPI

Place owlbert at `/home/pi` and create `/home/pi/webhooks` with your config files.

Install a service for owlbert `/etc/systemd/system/owlbert.service`:


```
[Unit]
Description=Owlbert
After=network.target

[Service]
Type=simple
WorkingDirectory=/home/pi
ExecStart=/home/pi/owlbert 8080
User=pi
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=owlbert

[Install]
WantedBy=multi-user.target
```

(Optional) Forward the local owlbert via ssh. Use autossh to keep the connection alive `/etc/systemd/system/owlbert-forward.service`:

```
[Unit]
Description=Owlbert Forward
After=network.target

[Service]
Type=simple
WorkingDirectory=/home/pi
ExecStart=/usr/bin/autossh -M 20000 -N user@your-server -R 9000:localhost:8080 -C
User=pi
Group=pi

[Install]
WantedBy=multi-user.target
```

This example redirects the owlbert from localhost:8080 to your-server:9000. The port 20000 is used for autossh's alive checks.


Now just enable and start the systemd services as usual.
