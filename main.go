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

	ClearFlow()
	RunFlow()
}

func ClearFlow() {
	// initializing storage
	functions.InitDefaultStoreOptions()

	functions.ClearStuff()

	functions.Show()
}

func RunFlow() {
	// initializing storage
	functions.InitDefaultStoreOptions()

	imageNames := pullImage()

	saveImage(imageNames)

	functions.ClearStuff()

	loadedImageNames := loadImage("/home/shubham/take-2-img-skopeo")

	lookupImage(loadedImageNames[0])

	functions.Show()

	//pushImage(loadedImageNames[0], "docker.io/spampatt/skopeo-take-2-test")



}

func pushImage(src string, dest string) {
	pushedManifest, err := functions.Push(src,dest,&libimage.PushOptions{
		CopyOptions: libimage.CopyOptions{
			InsecureSkipTLSVerify: 0,
		},
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("\n Pushed image manifest: %v \n", pushedManifest)


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
	// cannot use container-storage format for saving images, its not supported by the library,
	// supported format for save op are: docker-archive, oci-archive, docker-dir, oci-dir
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
