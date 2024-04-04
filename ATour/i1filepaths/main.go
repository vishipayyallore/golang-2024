package main

import (
	utl "autilities"
	"path/filepath"
	"strings"
)

var header = utl.Header{}

/*
The filepath package provides functions to parse and construct file paths in a way that is portable between operating systems; dir/file on Linux vs. dir\file on Windows, for example.
*/
func main() {
	/*
	 */
	header.DisplayHeader("Showing File Paths")

	// Join should be used to construct paths in a portable way. It takes any number of arguments and constructs a hierarchical path from them.
	p := filepath.Join("dir1", "dir2", "filename")
	utl.PLine("P: ", p)

	// You should always use Join instead of concatenating /s or \s manually. In addition to providing portability,
	// Join will also normalize paths by removing superfluous separators and directory changes.
	utl.PLine("Removing separators: ", filepath.Join("dir1//", "filename"))
	utl.PLine("Removing directory changes: ", filepath.Join("dir1/..", "dir2", "filename"))
	utl.PLine("Removing: ", filepath.Join("dir1/../dir1", "filename"))

	// Dir and Base can be used to split a path to the directory and the file. Alternatively, Split will return both in the same call.
	utl.PLine("Dir(p):", filepath.Dir(p))
	utl.PLine("Base(p):", filepath.Base(p))

	// We can check whether a path is absolute.
	utl.PLine("IsAbs: ", filepath.IsAbs("dir/file"))
	utl.PLine("IsAbs: ", filepath.IsAbs("/dir/file"))

	filename := "config.json"
	// Some file names have extensions following a dot. We can split the extension out of such names with Ext.
	ext := filepath.Ext(filename)
	utl.PLine("Ext: ", ext)

	// To find the file’s name with the extension removed, use strings.TrimSuffix.
	utl.PLine("TrimSuffix: ", strings.TrimSuffix(filename, ext))

	// Rel finds a relative path between a base and a target. It returns an error if the target cannot be made relative to base.
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	utl.PLine("Rel: ", rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	utl.PLine("Rel: ", rel)
}
