package fileutils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipFolder(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, _ := zip.FileInfoHeader(info)
		header.Name, _ = filepath.Rel(source, path)
		if info.IsDir() {
			header.Name += "/"
		}
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	return nil
}
