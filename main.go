package main

import (
	"somewhat-skopeo-take-2/main/functions"
)

func main() {
	//if reexec.Init() {
	//	return
	//}
	// initializing storage
	functions.InitDefaultStoreOptions()
	functions.Show()
	//functions.ImagePull("docker://alpine:latest")
	//functions.InitDefaultStoreOptions()
	functions.ClearStuff()
	//imageID, err := functions.Pull("docker://alpine:latest", &libimage.PullOptions{})
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//fmt.Printf("Pulled image ID: %s ", imageID)
	functions.Show()
}
