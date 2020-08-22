package collection

import "time"

type SportsData struct {
	Id       string    `json:"_id" bson:"_id"`
	GameTime time.Time `json:"game_time" bson:"game_time"`
	GambleId string    `json:"gamble_id,omitempty" bson:"gamble_id,omitempty"`
	GameType string    `json:"game_type,omitempty" bson:"game_type,omitempty"`
	Guest    struct {
		Name  string `json:"name" bson:"name"`
		Score int    `json:"score,omitempty" bson:"score,omitempty"`
	} `json:"guest" bson:"guest"`
	Host struct {
		Name  string `json:"name" bson:"name"`
		Score int    `json:"score,omitempty" bson:"score,omitempty"`
	} `json:"host" bson:"host"`
	GambleInfo struct {
		TotalPoint struct {
			Threshold float64 `json:"threshold,omitempty" bson:"threshold,omitempty"`
			Response  struct {
				Under float64 `json:"under,omitempty" bson:"under,omitempty"`
				Over  float64 `json:"over,omitempty" bson:"over,omitempty"`
			} `json:"response,omitempty" bson:"response,omitempty"`
			Judgement  string `json:"judgement,omitempty" bson:"judgement,omitempty"`
			Prediction struct {
				Under struct {
					Percentage float64 `json:"percentage,omitempty" bson:"percentage,omitempty"`
					Population int     `json:"population,omitempty" bson:"population,omitempty"`
				} `json:"under,omitempty" bson:"under,omitempty"`
				Over struct {
					Percentage float64 `json:"percentage,omitempty" bson:"percentage,omitempty"`
					Population int     `json:"population,omitempty" bson:"population,omitempty"`
				} `json:"over,omitempty" bson:"over,omitempty"`
				Major bool `json:"major,omitempty" bson:"major,omitempty"`
			} `json:"prediction,omitempty" bson:"prediction,omitempty"`
		} `json:"total_point,omitempty" bson:"total_point,omitempty"`
		SpreadPoint struct {
			Guest    float64 `json:"guest,omitempty" bson:"guest,omitempty"`
			Host     float64 `json:"host,omitempty" bson:"host,omitempty"`
			Response struct {
				Guest float64 `json:"guest,omitempty" bson:"guest,omitempty"`
				Host  float64 `json:"host,omitempty" bson:"host,omitempty"`
			} `json:"response,omitempty" bson:"response,omitempty"`
			Judgement  string `json:"judgement,omitempty" bson:"judgement,omitempty"`
			Prediction struct {
				Guest struct {
					Percentage float64 `json:"percentage,omitempty" bson:"percentage,omitempty"`
					Population int     `json:"population,omitempty" bson:"population,omitempty"`
				} `json:"guest,omitempty" bson:"guest,omitempty"`
				Host struct {
					Percentage float64 `json:"percentage,omitempty" bson:"percentage,omitempty"`
					Population int     `json:"population,omitempty" bson:"population,omitempty"`
				} `json:"host,omitempty" bson:"host,omitempty"`
				Major bool `json:"major,omitempty" bson:"major,omitempty"`
			} `json:"prediction,omitempty" bson:"prediction,omitempty"`
		} `json:"spread_point,omitempty" bson:"spread_point,omitempty"`
		Original struct {
			Response struct {
				Guest interface{} `json:"guest,omitempty" bson:"guest,omitempty"`
				Host  interface{} `json:"host,omitempty" bson:"host,omitempty"`
			} `json:"response,omitempty" bson:"response,omitempty"`
			Judgement  string `json:"judgement,omitempty" bson:"judgement,omitempty"`
			Prediction struct {
				Guest struct {
					Percentage float64 `json:"percentage,omitempty" bson:"percentage,omitempty"`
					Population int     `json:"population,omitempty" bson:"population,omitempty"`
				} `json:"guest,omitempty" bson:"guest,omitempty"`
				Host struct {
					Percentage float64 `json:"percentage,omitempty" bson:"percentage,omitempty"`
					Population int     `json:"population,omitempty" bson:"population,omitempty"`
				} `json:"host,omitempty" bson:"host,omitempty"`
				Major bool `json:"major,omitempty" bson:"major,omitempty"`
			} `json:"prediction,omitempty" bson:"prediction,omitempty"`
		} `json:"original,omitempty" bson:"original,omitempty"`
	} `json:"gamble_info,omitempty" bson:"gamble_info,omitempty"`
}