package serverutils

import (
	"log"
	"msgbroker/utils"
)

type Server struct {
	Exchanges map[string]*utils.Exchange
	Queues    map[string]*utils.Queue
}

func NewServer() *Server {

	s := Server{}

	s.Exchanges = make(map[string]*utils.Exchange)
	s.Queues = make(map[string]*utils.Queue)

	return &s

}
func (s *Server) AddQueue(name string) {

	s.Queues[name] = &utils.Queue{Channel: make(chan utils.Data)}

	return
}

func (s *Server) BindQueue(exchange string, queue string) {

	if s.Exchanges[exchange] == nil {
		log.Println(exchange, " exchange not found !")
		return
	}

	if s.Queues[queue] == nil {
		log.Println(queue, " queue not found")
		return
	}

	s.Exchanges[exchange].Queues[queue] = s.Queues[queue]

	return

}
