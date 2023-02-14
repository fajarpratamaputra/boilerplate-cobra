package mongo

type (
	InteractionScoreSum struct {
		ID         int     `bson:"_id"`
		Count      int     `bson:"count"`
		TotalScore float64 `bson:"total_score"`
	}
)
