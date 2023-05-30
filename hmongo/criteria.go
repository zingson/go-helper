package hmongo

import "go.mongodb.org/mongo-driver/bson"

// 2023-05-28 开发中

type Criteria struct {
	filter bson.D
	k      string
}

func NewFilter() *Criteria {
	return &Criteria{filter: bson.D{}}
}

func (criteria *Criteria) And(key string) *Criteria {
	return &Criteria{
		k:      key,
		filter: criteria.filter,
	}
}
func (criteria *Criteria) Eq(val any) *Criteria {
	criteria.filter = append(criteria.filter, bson.E{criteria.k, val})
	return criteria
}
func (criteria *Criteria) Ne(val any) *Criteria {
	criteria.filter = append(criteria.filter, bson.E{criteria.k, bson.M{"$ne": val}})
	return criteria
}
func (criteria *Criteria) Gt(val any) *Criteria {

	return criteria
}
func (criteria *Criteria) Gte(val any) *Criteria {

	return criteria
}
func (criteria *Criteria) Lt(val any) *Criteria {

	return criteria
}
func (criteria *Criteria) Lte(val any) *Criteria {

	return criteria
}
func (criteria *Criteria) In(val any) *Criteria {

	return criteria
}
func (criteria *Criteria) Nin(val any) *Criteria {

	return criteria
}
func (criteria *Criteria) Regex(val any) *Criteria {

	return criteria
}
