package TCPserver

import "net"

// Server is simple tcp server
type Server struct {
	address  string
	working  bool
	stopChan chan struct{}
	listner  net.Listener
	handler  func(net.Conn)
}
// NewServer crates new tcp server, requred listning address in "ip:port" format and handler function. 
func NewServer(address string, handler func(net.Conn)) *Server {
	return &Server{
		address:  address,
		working:  false,
		stopChan: make(chan struct{}),
		handler:  handler,
	}
}

// IsWork is function for checking server status true if server works
func (s *Server) IsWork() bool {
	return s.working
}

// IsStop is function for checking server status true if server stoped
func (s *Server) IsStop() bool {
	return !s.working
}

// Start is function to start server
func (s *Server) Start() (err error) {
	if s.IsStop() {
		s.listner, err = net.Listen("tcp", s.address)
		if err != nil {
			return err
		}
		go s.acceptConn()
		s.working = true
	}
	return nil
}

// Stop is function to stop server
func (s *Server) Stop() (err error) {
	if s.IsWork() {
		err = s.listner.Close()
		if err != nil {
			return err
		}
		s.stopChan <- struct{}{}
		s.working = false
	}
	return nil

}

// Internal loop function for accepting new connections
func (s *Server) acceptConn() {
	for {
		select {
		default:
			conn, err := s.listner.Accept()
			if err != nil {
				break
			}
			go s.handler(conn)

		case <-s.stopChan:
			return
		}
	}
}
