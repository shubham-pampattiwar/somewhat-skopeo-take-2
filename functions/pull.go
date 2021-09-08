package functions

import (
	"context"
	"fmt"
	"github.com/containers/common/libimage"
	"github.com/containers/common/pkg/config"
	store "github.com/containers/storage"
	"github.com/sirupsen/logrus"
)

var _defaultStore store.Store

func defaultStore() store.Store {
	options, err := store.DefaultStoreOptions(false, 0)
	fmt.Printf("\n filesystem path is %v \n", options.GraphRoot)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create default image store options")
	}
	//options.RunRoot = "/run/containers/storage"
	//options.GraphRoot = "/var/lib/containers/storage"
	//options.GraphDriverName = "overlay"

	if _defaultStore == nil {
		gotStorage, err := store.GetStore(options)
		if err != nil {
			logrus.WithError(err).Fatal("Could not create image store")
		}
		_defaultStore = gotStorage
	}

	return _defaultStore
}

func InitDefaultStoreOptions() {
	options, err := store.DefaultStoreOptions(false, 0)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create default image store options")
	}
	//options.RunRoot = "/run/containers/storage"
	//options.GraphRoot = "/var/lib/containers/storage"
	//options.GraphDriverName = "overlay"
	//_storeOptions = options

	if _defaultStore == nil {
		gotStorage, err := store.GetStore(options)
		if err != nil {
			logrus.WithError(err).Fatal("Could not create image store")
		}
		_defaultStore = gotStorage
	}

}


func Show() {
	imagesNew, _ := _defaultStore.Images()
	//spew.Dump(imagesNew)
	for _, img := range imagesNew {
		fmt.Printf("\n" + img.NamesHistory[0] + "\n")
	}

	if len(imagesNew) == 0 {
		fmt.Printf("\n No images available to show \n")
	}

	ctrs, _ := _defaultStore.Containers()
	for _, c := range ctrs {
		fmt.Printf("\n" + c.ID + "\n")
	}
	if len(ctrs) == 0 {
		fmt.Printf("\n No containers available to show \n")
	}
}

func ClearStuff() {
	_ = _defaultStore.Wipe()
}

func Pull(imageName string, options *libimage.PullOptions) (imageID string, err error) {
	libimageOptions := &libimage.PullOptions{}
	//libimageOptions.SignaturePolicyPath = options.SignaturePolicyPath
	//libimageOptions.RemoveSignatures = options.RemoveSignatures
	//libimageOptions.OciDecryptConfig = options.OciDecryptConfig
	//libimageOptions.AllTags = options.AllTags
	//libimageOptions.RetryDelay = options.RetryDelay
	//
	//if *options.MaxRetries > 0 {
	//	retries := *options.MaxRetries
	//	libimageOptions.MaxRetries = &retries
	//}

	runtime, err := libimage.RuntimeFromStore(defaultStore(), &libimage.RuntimeOptions{SystemContext: options.SystemContext})
	if err != nil {
		return "", err
	}

	// keeping pull policy as always for now, lets just roll with this
	pulledImages, err := runtime.Pull(context.Background(), imageName, config.PullPolicyAlways, libimageOptions)
	if err != nil {
		return "", err
	}

	if len(pulledImages) == 0 {
		return "", fmt.Errorf("some error occurred, could not pull image %s: ", imageName)
	}

	return pulledImages[0].ID(), nil

}