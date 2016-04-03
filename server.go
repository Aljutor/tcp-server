package TCPserver

import "net"

// Server is basic game server
type Server struct {
	address  string
	working  bool
	stopChan chan struct{}
	listner  net.Listener
	handler  func(net.Conn)
}

//NewServer is method to create new Game Server
func NewServer(address string, handler func(net.Conn)) *Server {
	return &Server{
		address:  address,
		working:  false,
		stopChan: make(chan struct{}),
		handler:  handler,
	}
}

// IsWork is function for checking server status
func (S *Server) IsWork() bool {
	return S.working
}

// IsStop is function for checking server status
func (S *Server) IsStop() bool {
	return !S.working
}

// Start is function to start server
func (S *Server) Start() (err error) {
	if S.IsStop() {
		S.listner, err = net.Listen("tcp", S.address)
		if err != nil {
			return err
		}
		go S.acceptConn()
		S.working = true
	}
	return nil
}

// Stop is function to stop server
func (S *Server) Stop() (err error) {
	if S.IsWork() {
		err = S.listner.Close()
		if err != nil {
			return err
		}
		S.stopChan <- struct{}{}
		S.working = false
	}
	return nil

}

func (S *Server) acceptConn() {
	for {
		select {
		default:
			conn, err := S.listner.Accept()
			if err != nil {
				break
			}
			go S.handler(conn)

		case <-S.stopChan:
			return
		}
	}
}
