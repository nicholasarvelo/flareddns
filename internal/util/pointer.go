package util

// BooleanPointer returns a pointer to a boolean value, primarily aimed at
// enhancing code readability when dealing with structs representing optional
// boolean fields in Go.
func BooleanPointer(boolean bool) *bool {
	return &boolean
}

// StringPointer returns a pointer to a string value, primarily aimed at
// enhancing code readability when dealing with structs representing optional
// string fields in Go.
func StringPointer(string string) *string {
	return &string
}
