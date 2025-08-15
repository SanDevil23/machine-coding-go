package main

// import "fmt"

type TicketMaster struct {
	Parkings *ParkingSpaces
	Tickets  map[string]*Ticket
}

func NewTicketMaster() *TicketMaster {
	return &TicketMaster{
		Tickets: make(map[string]*Ticket),
	}
}

func (tm *TicketMaster) AddTicket(ticket *Ticket) {
	tm.Tickets[ticket.TicketId] = ticket
}

func (tm *TicketMaster) GetTicket(ticketId string) *Ticket {
	return tm.Tickets[ticketId]
}
