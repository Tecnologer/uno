package classic

//rawCard is a struct to load cards from json configuration
type rawCard struct {
	Value  string   `json:"value"`
	Count  int      `json:"count"`
	Colors []string `json:"colors"`
}

type card struct {
	Value string `json:"value"`
	Color string `json:"count"`
}

func (c *card) GetValue() string {
	return c.Value
}

func (c *card) GetColor() string {
	return c.Color
}
