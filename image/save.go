package image

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Duckduckgot/memego/cli"
	"github.com/Duckduckgot/memego/image/stream"
	"github.com/Duckduckgot/memego/output"
)

// Save the passed image to disk.
func Save(opt cli.Options, st stream.Stream) string {
	var name string

	if opt.OutName != "" {
		name = opt.OutName
	} else {
		name = tempName(st.FileExt())
	}

	file, err := os.Create(name)
	output.OnError(err, "Could not create image file")
	defer file.Close()

	_, err = io.Copy(file, &st)
	output.OnError(err, "Could not save image stream to file")

	return name
}

// Generate a temporary file name.
func tempName(ext string) string {
	dir := os.TempDir()
	return filepath.Join(dir, fmt.Sprintf("meme.%s", ext))
}
