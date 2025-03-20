# vulnerable-repo
This is an intentionally vulnerable repo to test code scanning tools with


## Go Server

You can start the golang server by running

```
go run main.go
```

This will start a server on port `:8080`.

Access `http://localhost:8080/` and you will see it print out your client headers!


### Environment Setup

Attacker Docker Image

```
docker run -p 8090:8090 -p 4443:4443 --name attacker -d -it debian:latest
docker exec -it attacker /bin/bash
apt update
apt install ncat
nc -l 8090
```

Server

```
docker build --tag vulnerable-server .
docker run -p 8080:8080 --name server -d -it vulnerable-server
```



### Vulnerability 

You may have noticed when inspecting the code there is a bit of a vulnerability in this server...

For some reason User-Agent is being executed as if it was a command on the server.

```
    ua := r.Header.Get("User-Agent")
    out, err := exec.Command(ua).Output()
```

This means the server is vulnerable to command injection attacks!


### Exploit

You can modify the User Agent string in Chromium browers by opening the developer tools >> more tools >> network conditions.
Then deselect under User Agent "Use Browser Default" and you can now input any string you want into the user agent. 
