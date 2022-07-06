package set

import (
	"fmt"
	"reflect"
	"strings"
)

type Set[T comparable] map[T]struct{}

func New[T comparable](v ...T) Set[T] {
	set := make(Set[T], len(v))
	set.AddAll(v...)
	return set
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) HasAll(vals ...T) bool {
	for _, v := range vals {
		if !s.Has(v) {
			return false
		}
	}

	return true
}

func (s Set[T]) HasAny(vals ...T) bool {
	for _, v := range vals {
		if s.Has(v) {
			return true
		}
	}

	return false
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) AddAll(vals ...T) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Values() []T {
	t := make([]T, 0, len(s))
	for v := range s {
		t = append(t, v)
	}
	return t
}

func (s Set[T]) GoString() string {
	var buf strings.Builder
	buf.WriteString(`set.New[`)
	buf.WriteString(reflect.TypeOf(*new(T)).String())
	buf.WriteString(`](`)

	first := true
	for v := range s {
		if !first {
			buf.WriteString(`, `)
		}
		first = false
		buf.WriteString(fmt.Sprintf(`%#v`, v))
	}
	buf.WriteString(`)`)

	return buf.String()
}

func Intersect[T comparable](a, b Set[T]) Set[T] {
	lenA, lenB := len(a), len(b)
	if lenA > lenB {
		lenA = lenB
	}
	out := make(Set[T], lenA)
	for v := range a {
		if _, ok := b[v]; ok {
			out[v] = struct{}{}
		}
	}
	return out
}

func Union[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T], len(a)+len(b))
	for v := range a {
		out[v] = struct{}{}
	}
	for v := range b {
		out[v] = struct{}{}
	}
	return out
}
