package valueobject

type Name struct {
	First string
	Last  string
}

func (n *Name) Full() string {
	if n.Last == "" {
		return n.First
	}
	return n.First + " " + n.Last
}

func NewName(first, name string) *Name {
	return &Name{
		First: first,
		Last:  name,
	}
}
