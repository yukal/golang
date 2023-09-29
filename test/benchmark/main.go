package test

type FilterNested2 struct {
	IsActual bool `query:"isActual" json:"isActual"`
	IsUniq   bool `query:"isUniq" json:"isUniq"`
}

type FilterNested1 struct {
	Nested FilterNested2 `query:"nes2" json:"nes2"`
}

type FilterMinMax struct {
	Min uint8 `query:"min" json:"min"`
	Max uint8 `query:"max" json:"max"`
}

type FilterData struct {
	RegionId uint8         `query:"regionId" json:"regionId"`
	Phone    string        `query:"phone" json:"phone"`
	Age      FilterMinMax  `query:"age" json:"age"`
	Nested   FilterNested1 `query:"nes" json:"nes"`
}

type Filter struct {
	Data    []FilterData `query:"data" json:"data"`
	Text    string       `query:"text" json:"text"`
	OrderBy []string     `query:"orderBy" json:"orderBy"`
	Limit   int64        `query:"limit" json:"limit"`
	Offset  int64        `query:"offset" json:"offset"`
}

type Payload struct {
	Filter Filter `query:"filter" json:"filter"`
}

var result any
