package fs

import (
	"io"
	"io/fs"
	"os"
	"path"
)

// FS wraps a directory and an embed FS that are expected to have the same contents.
// This was copied from the Buffalo project and is useful in development mode to avoid
// recompilation for Templates.
type FS struct {
	embed fs.FS
	dir   fs.FS
}

// NewFS returns a new FS that wraps the given directory and embedded FS.
// the embed.FS is expected to embed the same files as the directory FS.
func NewFS(embed fs.ReadDirFS, dir string) FS {
	return FS{
		embed: embed,
		dir:   os.DirFS(dir),
	}
}

// Open implements the FS interface.
func (f FS) Open(name string) (fs.File, error) {
	file, err := f.getFile(name)
	if path.Ext(name) == ".go" {
		return nil, os.ErrNotExist
	}

	if name == "." {
		return file, err
	}

	return file, err
}

func (f FS) ReadFile(name string) ([]byte, error) {
	file, err := f.Open(name)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return b, err
}

func (f FS) getFile(name string) (fs.File, error) {
	file, err := f.dir.Open(name)
	if err == nil {
		return file, nil
	}

	return f.embed.Open(name)
}
