package aaguids

// metadata is a map linking unique identifier to its corresponding Entry in the Metadata.
var metadata map[string]Entry

// goPtr returns a pointer to the given value of any type.
func goPtr[T any](v T) *T {
	return &v
}

// GetEntry retrieves the metadata Entry identified by aaGuid.
// Returns the Entry and a boolean indicating if it exists in the metadata map.
func GetEntry(aaGuid string) (e Entry, exists bool) {
	e, exists = metadata[aaGuid]
	return
}
