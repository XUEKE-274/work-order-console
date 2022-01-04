package response

import "work-order-console/domain/entity"

type TicketCountResponse struct {
	Account  string   `json:"account"`
	Count    int64	  `json:"count"`
}

type TicketDetailResponse struct {
	TicketNumber    string                    `json:"ticketNumber"`
	Id              string                    `json:"id"`
	CustomFieldList *[]entity.CustomFieldList `json:"customFieldList"`
}