package collection

import "fmt"

type Gambler struct {
	Name     string `json:"name,omitempty"`
	Capital  int    `json:"capital"`
	Strategy struct {
		Bet struct {
			Name       string                 `json:"name"`
			Parameters map[string]interface{} `json:"parameters,omitempty"`
		} `json:"bet"`
		Put struct {
			Name       string                 `json:"name"`
			Parameters map[string]interface{} `json:"parameters,omitempty"`
		} `json:"put"`
	}
}

func (g *Gambler) String() string {
	return fmt.Sprintf("gambler name: %s, bet strategy: %s, put strategy: %s", g.Name, g.Strategy.Bet.Name, g.Strategy.Put.Name)
}
