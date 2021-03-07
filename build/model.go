package build

type Item struct {
	ID string
}

type Build struct {
	Name string
	Items []*Item
}