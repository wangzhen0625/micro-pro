package account

import ()

type node struct{}

func CreateNode() *node {
	return &node{}
}

type role struct{}

func CreateRole() *role {
	return &role{}
}

type app struct{}

func CreateApp() *app {
	return &app{}
}
