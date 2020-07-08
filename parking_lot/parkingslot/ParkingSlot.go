package parkingslot

// SlotType is enum type for different types of slot
type SlotType int

const (
	handicapped SlotType = iota
	motorcycle
	compact
	large
)

// SlotStatus is enum type for different Slot Status
type SlotStatus int

const (
	Occupied SlotStatus = iota
	Free
	NotInUsed
)

// Slot is a single parking slot
type Slot struct {
	slotNumber int
	//slotType   int
	slotStatus SlotStatus
}

// GetNewSlot return new slot for given type with initial given status
//func GetNewSlot(slotNumber int, slotType int, SlotStatus int) *Slot {
func GetNewSlot(slotNumber int, SlotStatus SlotStatus) *Slot {
	return &Slot{
		slotNumber: slotNumber,
		//slotType:   slotType,
		slotStatus: SlotStatus,
	}
}

// GetSlotStatus returns slot status
func (s *Slot) GetSlotStatus() SlotStatus {
	return s.slotStatus
}

// SetSlotStatus returns slot status
func (s *Slot) SetSlotStatus(slotStatus SlotStatus) {
	s.slotStatus = slotStatus
}
