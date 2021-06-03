package recursion

import "fmt"

type File struct {
	Name  string
	Path  string
	Files []*File
}

func (f *File) TreePath() {
	fmt.Println(f.Name)

	for _, file := range f.Files {
		file.TreePath()
	}

}

func (f *File) IsDic() bool {
	if len(f.Files) == 0 {
		return false
	}
	return true
}
