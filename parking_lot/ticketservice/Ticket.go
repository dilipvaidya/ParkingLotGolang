package ticketservice

// Status holds the validity of the ticket
type Status int

const (
	// Valid ticket
	Valid Status = iota
	// OverTime - Valid but extra time utilised
	OverTime
	// Invalid - vehicle with this ticket exited
	Invalid
)

// Ticket structure
type Ticket struct {
	tickerNumber string
	checkInTime  int64
	checkInGate  int
	paymentTime  int64
	amountPaid   float32
	ticketStatus Status
}

// GetNewTicket gets new tikcket for default hour of parking
func GetNewTicket(licensePlate string, checkInTime int64, checkInGate int, paymentTime int64, amountPaid float32) *Ticket {
	return &Ticket{
		tickerNumber: licensePlate,
		checkInTime:  checkInTime,
		checkInGate:  checkInGate,
		paymentTime:  paymentTime,
		amountPaid:   amountPaid,
		ticketStatus: Valid,
	}
}

// SetStatus sets the ticket status
func (t *Ticket) SetStatus(s Status) {
	t.ticketStatus = s
}

// GetStatus returns ticket status
func (t *Ticket) GetStatus() Status {
	return t.ticketStatus
}
