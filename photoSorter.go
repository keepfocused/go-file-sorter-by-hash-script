// photoSorter
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type photo struct {
	path string
	size int
	name string
}

func main() {

	findUniqAndMoveToFolder()
	//makeCheckWhatLeft()

}

func makeCheckWhatLeft() {

	var firstFolderFiles []photo
	var secondFolderFiles []photo

	root := "/Users/danil/Downloads/untitled_folder/mov/"

	root2 := "/Users/danil/Downloads/untitled_folder/myUniqMp4/"

	//First folder check
	firstFolderFiles = GetFilesByPath(root, ".MP4")

	//Second folder check
	secondFolderFiles = GetFilesByPath(root2, ".mp4")

	fmt.Println("first folder count", len(firstFolderFiles))
	fmt.Println("second folder count", len(secondFolderFiles))

	var difference = Difference(firstFolderFiles, secondFolderFiles)
	fmt.Println("difference count")
	fmt.Println(len(difference))

	//MoveFiles(difference)

}

func GetFilesByPath(root string, fileType string) []photo {

	var photos []photo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == fileType {

			ph := photo{path, int(info.Size()), info.Name()}

			photos = append(photos, ph)

		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return photos

}

func MoveFiles(files []photo) {

	for _, ph := range files {

		oldLocation := ph.path
		newName := []string{"/Users/danil/Downloads/untitled_folder/mapjpeg", ""}
		newName[1] = ph.name
		newLocation := strings.Join(newName, "")
		err := os.Rename(oldLocation, newLocation)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func Difference(a, b []photo) (diff []photo) {

	var exclude []photo

	dst := make([]photo, len(a))

	copy(dst, a)

	var result []photo = dst

	for _, ph := range a {

		for _, ph2 := range b {

			if ph.size == ph2.size {
				exclude = append(exclude, ph)
				break
			}
		}

	}

	fmt.Println("exclude")
	fmt.Println(len(exclude))

	for _, ph := range exclude {
		var index = FindIndex(ph, result)
		var newArray = RemoveIndex(result, index)
		result = newArray
	}

	fmt.Println("after exclude")
	fmt.Println(len(result))

	return result
}

func FindIndex(element photo, array []photo) int {
	for index, value := range array {
		if reflect.DeepEqual(element, value) {
			//fmt.Println(index)
			return index
		}
	}
	return -1
}

func RemoveIndex(array []photo, index int) []photo {
	return append(array[:index], array[index+1:]...)
}

// if you have duplicates

//find all files which has duplicates by bites
// and take only one of them from each bunch of duplicates and move it to the folder
func findUniqAndMoveToFolder() {

	var photos []photo
	var sortedPhotos []photo

	root := "/Users/danil/Downloads/untitled_folder/"

	photos = GetFilesByPath(root, ".jpeg")

	photos = append(photos, GetFilesByPath(root, ".JPG")...)

	fmt.Println("total count", len(photos))

	dict := map[int]photo{}

	for _, ph := range photos {

		if _, ok := dict[ph.size]; !ok {

			dict[ph.size] = ph
		}
	}

	fmt.Println("dict count = ", len(dict))

	//Перенос файлов

	for _, v := range dict {
		sortedPhotos = append(sortedPhotos, v)
	}

	fmt.Println("sortedphotos count = ", len(sortedPhotos))

	//MoveFiles(sortedPhotos)

}
