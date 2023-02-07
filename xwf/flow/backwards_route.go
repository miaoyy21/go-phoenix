package flow

import (
	"github.com/sirupsen/logrus"
)

type BackwardsRoute struct {
	start Flowable

	routes map[int][]int
}

func NewBackwardsRoute(node Flowable) *BackwardsRoute {
	route := &BackwardsRoute{
		start:  node,
		routes: make(map[int][]int),
	}

	route.routes[node.Key()] = []int{node.Key()}
	return route
}

func (route BackwardsRoute) append(key, next int) {
	children, ok := route.routes[key]
	if !ok {
		logrus.Panic("unreachable")
	}

	children = append(children, next)
	route.routes[next] = children
}

func (route BackwardsRoute) Routes() map[int][]int {
	return route.routes
}
func (route BackwardsRoute) Backwards(values string) ([]Flowable, error) {
	return route.backwards(route.start, values)
}

func (route BackwardsRoute) backwards(node Flowable, values string) ([]Flowable, error) {
	nodes := make([]Flowable, 0, 3)

	// Backwards
	backwards, err := node.Backwards(values)
	if err != nil {
		return nil, err
	}

	// For Backwards
	for _, next := range backwards {
		// Next Flow
		nodes = append(nodes, next)
		route.append(node.Key(), next.Key())

		// Branch
		if _, ok := next.(BranchFlowable); ok {
			flows, err := route.backwards(next, values)
			if err != nil {
				return nil, err
			}

			nodes = append(nodes, flows...)
		}
	}

	return nodes, nil
}
