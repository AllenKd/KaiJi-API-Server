package collection

import "fmt"

type Gambler struct {
	Id        string `json:"id"`
	Principle int    `json:"principle"`
	Strategy  struct {
		Name        string `json:"name"`
		PutStrategy struct {
			Name       string                 `json:"name"`
			Parameters map[string]interface{} `json:"parameters,omitempty"`
		} `json:"put_strategy"`
		Parameters map[string]interface{} `json:"parameters,omitempty"`
	}
}

func (g *Gambler) String() string {
	return fmt.Sprintf("strategy: %s, put strategy: %s", g.Strategy.Name, g.Strategy.PutStrategy.Name)
}
