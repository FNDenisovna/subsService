package main

type Msg struct {
	TraceId string `json:"trace_id"`
	Offer   `json:"offer"`
}

/*{
	"$id": "message.schema.json",
	"type": "object",
	"properties": {
	  "trace_id": {
		"type": "string"
	  },
	  "offer": {
		"$ref": "offer.schema.json"
	  }
	},
	"required": [
	  "trace_id",
	  "offer"
	]
  }
}*/
