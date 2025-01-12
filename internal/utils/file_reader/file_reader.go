package file_reader

type FileReader interface {
	ReadFile(filePath string) ([]byte, error)
}
