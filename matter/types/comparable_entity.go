package types

type ComparableEntity interface {
	Entity
	Equals(Entity) bool
}
