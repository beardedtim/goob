package engine

/*
An Engine is our core business logic that takes in requests
and outputs any response from that request along with any
requests that need to be made as a result
*/
type Engine struct{}

func New() Engine {
	engine := Engine{}

	return engine
}
