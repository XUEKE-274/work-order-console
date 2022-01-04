package entity

type TicketEntity struct {
	BaseEntity
	Id              string             `bson:"_id"`
	TicketNumber    string             `bson:"ticketNumber"`
	CustomFieldList *[]CustomFieldList `bson:"customFieldList"`
}

type CustomFieldList struct {
	Id        string      `bson:"_id" json:"id"`
	FieldType int32       `bson:"fieldType" json:"fieldType"`
	IsFixed   bool        `bson:"isFixed" json:"isFixed"`
	Name      string      `bson:"name" json:"name"`
	Value     string      `bson:"value" json:"value"`
	Optional  *[]Optional `bson:"optional" json:"optional"`
}

type Optional struct {
	Id    string `bson:"_id" json:"id"`
	Value string `bson:"value" json:"value"`
}
