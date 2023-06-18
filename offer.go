package main

type Offer struct {
	ID             string `json:"id"`
	Price          int    `json:"price"`
	StockCount     int    `json:"stock_count"`
	PartnerContent `json:"partner_content"`
}

type PartnerContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

/*{
  "$id": "offer.schema.json",
  "type": "object",
  "properties": {
    "id": {
      "type": "string",
      "description": "Offer identifier, only numerical symbols are allowed"
    },
    "price": {
      "type": "integer",
      "description": "Offer price, value in range from 0 to 10̂9"
    },
    "stock_count": {
      "type": "integer",
      "description": "Items left on stocks, value in range from 0 to 10̂9"
    },
    "partner_content": {
      "type": "object",
      "properties": {
        "title": {
          "type": "string",
          "description": "Offer title filled in by the partner"
        },
        "description": {
          "type": "string",
          "description": "Offer description filled in by the partner"
        }
      }
    }
  },
  "required": [
    "id"
  ]
}*/
