package main

import (
	utl "autilities"
	"embed"
)

var header = utl.Header{}

/*
embed directives accept paths relative to the directory containing the Go source file. This directive embeds the contents of the
file into the string variable immediately following it.
*/
//go:embed folder/single_file.txt
var fileString string

/*
Or embed the contents of the file into a []byte.
*/
//go:embed folder/single_file.txt
var fileByte []byte

/*
We can also embed multiple files or even folders with wildcards. This uses a variable of the embed.FS type, which implements a
simple virtual file system.
*/
//go:embed folder/single_file.txt
//go:embed folder/file1.hash
//go:embed folder/file2.hash
//go:embed folder/*.hash
var folder embed.FS

/*
//go:embed is a compiler directive that allows programs to include arbitrary files and folders in the Go binary at build time.
Import the embed package; if you don’t use any exported identifiers from this package, you can do a blank import with _ "embed".

******************** Create these files and folders in the project directory: ********************
mkdir -p folder
echo "This is Single File." > folder/single_file.txt
echo "123 from file1.hash" > folder/file1.hash
echo "456 from file2.hash" > folder/file2.hash
*/
func main() {
	/*
		Import the embed package; if you don’t use any exported identifiers from this package, you can do a blank import with
		_ "embed".
	*/
	header.DisplayHeader("Showing Embed Directive")

	print("File Content: ", fileString)
	print("File Content string(fileByte): ", string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print("\n", string(content2))

	// Read all files in the folder
	utl.PLine("\n\nReading all files in the folder")
	files, _ := folder.ReadDir("folder")
	for _, file := range files {
		content, _ := folder.ReadFile("folder/" + file.Name())
		print("\n", string(content))
	}
}
