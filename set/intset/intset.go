// set provides a very implementation of sets for a number of basic types.
package set

// IntSet represents a set of ints
type IntSet struct {
	s map[int]bool
}

// NewIntSet creates a new IntSet
func NewIntSet(xs ...int) *IntSet {
	var i IntSet
	i.s = make(map[int]bool)
	for _, x := range xs {
		i.s[x] = true
	}
	return &i
}

// Clone returns a new copy of an IntMap
func (s *IntSet) Clone() *IntSet {
	return NewIntSet(s.ToList()...)
}

// Equal tests for set equality
func (s *IntSet) Equals(t *IntSet) bool {
	if len(s.s) != len(t.s) {
		return false
	}
	for k := range s.s {
		if _, ok := t.s[k]; !ok {
			return false
		}
	}
	return true
}

// Add inserts an arbitrary number of integers to an IntSet
func (s *IntSet) Add(xs ...int) *IntSet {
	for _, x := range xs {
		s.s[x] = true
	}
	return s
}

// Contains returns true if i is in the set and false otherwise
func (s *IntSet) Contains(i int) bool {
	_, ok := s.s[i]
	return ok
}

// ToList converts an IntSet into an int slice
func (s *IntSet) ToList() []int {
	l := make([]int, len(s.s))
	i := 0
	for k := range s.s {
		l[i] = k
		i++
	}
	return l
}

// Union forms the union of IntSet s and t
func (s *IntSet) Union(t *IntSet) *IntSet {
	for k := range t.s {
		s.s[k] = true
	}
	return s
}

// Intersect forms the intersection of IntSet s and t
func (s *IntSet) Intersect(t *IntSet) *IntSet {
	for k := range s.s {
		if _, ok := t.s[k]; !ok {
			delete(s.s, k)
		}
	}
	return s
}

// Complement subtracts IntSet t from s
func (s *IntSet) Complement(t *IntSet) *IntSet {
	for k := range s.s {
		if _, ok := t.s[k]; ok {
			delete(s.s, k)
		}
	}
	return s
}
