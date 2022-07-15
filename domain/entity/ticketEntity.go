package entity

type TicketEntity struct {
	BaseEntity
	TicketNumber    string             `json:"ticketNumber"`
	//CustomFieldList *[]CustomFieldList `json:"customFieldList"`
}

type CustomFieldList struct {
	BaseEntity
	FieldType int32       `bson:"field_type" json:"fieldType"`
	IsFixed   bool        `bson:"is_fixed" json:"isFixed"`
	Name      string      `bson:"name" json:"name"`
	Value     string      `bson:"value" json:"value"`
	TicketId  string      `json:"ticket_id"`

	//Optional  *[]Optional `bson:"optional" json:"optional"`
}

type Optional struct {
	BaseEntity
	Value string `bson:"value" json:"value"`
	FieldId string `json:"field_id"`
}
