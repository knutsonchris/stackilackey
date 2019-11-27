# stackilackey

Programagically interact with a Stacki Frontend using Go.



### Install

`go get github.com/knutsonchris/stacki-lackey`



### Setup

To communicate with the Frontend, the lackey needs the following environment variables:

`STACKILACKEY_FRONTEND_IP: <your frontend ip>`

`STACKILACKEY_USERNAME: <the username from stacki-ws.cred>`

`STACKILACKEY_PASSWORD: <the key from stack-ws.cred>`



### Usage

Import 

`stackilackey "github.com/knutsonchris/stacki-lackey"`

Create a `StackCommand` struct

```go
stack := stackilackey.StackCommand{}
```

Use it just like you would use the command line

```go
bytes, err := stack.Add.Host.Interface("host", "channel", "interface", etc...)
```

The lackey will return the API response as a json byte slice, as well as an error if the API complains about your command.

In the case of list commands, the lackey will return a list of structs instead of the json byte slice.

``` go
ifaces, err := stack.List.Host.Interface("hostName")
if err != nil{
  // handle error
}
for _, iface := range ifaces{
  fmt.Println(iface.IP) // ip, channel, host, etc...
}
```



### Testing

#### Option 1

Build yourself a Stacki Docker image using the instruction on this page:

 https://github.com/Teradata/stacki/tree/master/docker 



#### Option 2

If you do not want to wait the ~30 minutes to build, you may download a pre-built image from Docker Hub:

`docker pull stacki/frontend-centos`

If you choose this route, you may find yourself with an older version of Stacki



Check the `docker run` command in the Makefile for the image name and tag. If the docker image you just downloaded or built has a different name or tag, replace the name in the Makefile. At the time of writing, using option 1, you should have `stacki/frontend-centos:05.03.00.00`



Run make and enjoy the show