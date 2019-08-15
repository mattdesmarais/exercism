//Package twofer returns strings containing stuff.
package twofer

//ShareWith returns the name with leading and trailing strings.
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."exercism.io
	}
	return "One for you, one for me."
}
