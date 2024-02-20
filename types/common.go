package types

type ToGraphQLRequest interface {
	ToRequest() (map[string]interface{}, error)
}
