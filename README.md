# tcp-server
Simple golang package for tcp-server

Example echo server:

``` Go

func handler (conn net.Conn){
  defer conn.Close()
  
  buf := make([]byte, 1024)
  
  len, err := conn.Read(buf)
  if err != nil{
	return
  }

  conn.Write(buf[:len])
  
}

server := TCPserver.NewServer(":8080", handler)

server.Start()

```
