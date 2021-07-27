package utils

import ()

type Queue struct {
	Channel chan Data
}

type Exchange struct {
	Call   func(map[string]*Queue, Data)
	Queues map[string]*Queue
}

func NewExchange(c func(map[string]*Queue, Data)) *Exchange {

	e := Exchange{}
	e.Queues = make(map[string]*Queue)
	e.Call = c
	return &e
}

type Data struct {
	Head string `json:Head`
	Body string `json:Body`
}
