package walk_directory

import (
	"io/fs"
	"os"
)

// WalkDirectory walks the directory and returns the total size of all files in the directory and the total number of files in the directory.
func WalkDirectory(dir string) (uint64, uint64, error) {
	var totalSize uint64
	var totalFiles uint64

	err := fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		// Increment totalFiles
		totalFiles++
		// Increment totalSize by the size of the file
		info, err := d.Info()
		if err != nil {
			return err
		}
		totalSize += uint64(info.Size())

		return nil
	})

	if err != nil {
		return 0, 0, err
	}
	return totalSize, totalFiles, nil
}
