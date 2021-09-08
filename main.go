package main

import (
	"fmt"
	"github.com/containers/common/libimage"
	"github.com/containers/storage/pkg/reexec"
	"somewhat-skopeo-take-2/main/functions"
)

func main() {
	if reexec.Init() {
		return
	}
	// initializing storage
	functions.InitDefaultStoreOptions()
	functions.Show()
	//functions.ImagePull("docker://alpine:latest")
	//functions.InitDefaultStoreOptions()
	//functions.ClearStuff()

	imageNames := pullImage()

	saveImage(imageNames)

	functions.ClearStuff()

	loadedImageNames := loadImage("/home/shubham/take-2-img-skopeo")

	lookupImage(loadedImageNames[0])

	//importedImageName := importImage("/home/shubham/take-2-img-skopeo")
	//fmt.Printf("\n Imported image name is %v \n", importedImageName)
	//lookupImage(importedImageName)

	functions.Show()
}

func pullImage() (imageNames []string) {
	imageID, imageNames, err := functions.Pull("docker://alpine:latest", &libimage.PullOptions{})
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("Pulled image ID: %s ", imageID)
	return imageNames
}

func saveImage(imageNames []string) {
	// defaulting to docker-archive format because we are taking alpine docker image for the sake of POC
	err := functions.Save(imageNames, "docker-archive", "/home/shubham/take-2-img-skopeo", &libimage.SaveOptions{})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func importImage(importImagePath string) string {
	name, err := functions.Import(importImagePath, &libimage.ImportOptions{})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return name
}

func lookupImage(imageName string) {
	err := functions.Lookup(imageName, &libimage.ImportOptions{})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func loadImage(path string) []string {
	loadedImageNames, err := functions.Load(path, &libimage.LoadOptions{})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return loadedImageNames
}
