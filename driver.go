package converter

type Driver interface {
	ConvertToPNG(input []byte) (output []byte, err error)
	ConvertToJpg(input []byte) (output []byte, err error)
	Supports() []string
}
