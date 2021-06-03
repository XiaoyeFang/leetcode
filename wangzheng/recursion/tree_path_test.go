package recursion

import "testing"

func TestFile_TreePath(t *testing.T) {
	var file101 =&File{Name: "011_file",Path: "00",Files: []*File{}}
	var file10 =&File{Name: "010_file",Path: "00",Files: []*File{file101}}
	var file23 =&File{Name: "023_file",Path: "00",Files: []*File{}}
	var file1 =&File{Name: "01_file",Path: "00",Files: []*File{file10}}
	var file2 =&File{Name: "02_file",Path: "00",Files: []*File{file23}}
	var file =&File{Name: "00_file",Path: "00",Files: []*File{file1,file2}}
	file.TreePath()
}
