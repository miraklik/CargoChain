package types

type ShipmentStatus string

const (
	StatusCreate    ShipmentStatus = "CREATE"
	StatusInTrans   ShipmentStatus = "IN_TRANS"
	StatusShipped   ShipmentStatus = "SHIPPED"
	StatusDelivered ShipmentStatus = "DELIVERED"
	StatusCancelled ShipmentStatus = "CANCELLED"
)

func NewShipment(creator, trackNum string, status string, location string, timestamp int64) *Shipment {
	return &Shipment{
		Creator:        creator,
		TrackingNumber: trackNum,
		Status:         status,
		Location:       location,
		Timestamp:      timestamp,
	}
}
