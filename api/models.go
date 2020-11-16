package api

type Todo struct {
	Id          string   `json:"_id" bson:"_id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
