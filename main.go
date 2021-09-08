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
	exportImage(imageNames)

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

func exportImage(imageNames []string) {
	// defaulting to docker-archive format because we are taking alpine docker image for the sake of POC
	err := functions.Export(imageNames,"docker-archive", "/home/shubham/take-2-img-skopeo", &libimage.SaveOptions{})
	if err != nil {
		fmt.Printf(err.Error())
	}
}