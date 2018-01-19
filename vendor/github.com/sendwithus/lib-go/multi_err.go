package swu

// MultiError is a type that can concatenate together multiple errors into a
// friendly error message.
//
// Usage:
// ```
// func MyFunc() error {
//  	// do things that give two errors
// 	var err1 error
// 	var err2 error
// 	m := MultiError{err1}
//      return m.Append(err2)
// }
//
// ```
type MultiError []error

// Retrieve the error. This concatenates all errors using ','
func (m MultiError) Error() string {
	if len(m) == 1 {
		return m[0].Error()
	}
	s := ""
	for idx, e := range m {
		s += e.Error()
		if idx != len(m)-1 {
			s += ", "
		}
	}
	return s
}

// Append an error to the list. This will return the MultiError instance for chaining.
func (m *MultiError) Append(err error) MultiError {
	*m = append(*m, err)
	return *m
}
