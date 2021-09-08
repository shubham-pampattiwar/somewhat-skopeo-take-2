package functions

import (
	"context"
	"fmt"
	"github.com/containers/common/libimage"
)

func Load(path string, options *libimage.LoadOptions) ([]string, error) {
	runtime, err := libimage.RuntimeFromStore(defaultStore(), &libimage.RuntimeOptions{SystemContext: options.SystemContext})
	if err != nil {
		return nil,err
	}

	loadedImageNames, err := runtime.Load(context.Background(), path, options)
	fmt.Printf("Loaded Image Name: %v", loadedImageNames)
	return loadedImageNames, nil
}
