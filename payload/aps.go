package payload

// Apple-defined keys
type Aps struct {
	alert            interface{}
	badge            int
	sound            interface{}
	threadId         string
	category         string
	contentAvailable int
	mutableContent   int
	targetContentId  string
}
