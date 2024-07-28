package file

import "fmt"

type FileManager struct {}

func (fm *FileManager) Read(string path) (string, error) {
    data, err := os.ReadFile(path)    
    if err != nil {
        return fmt.Errorf("Failed to read file %s with error: %v")
    }

    return string(data)
}
