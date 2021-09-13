package functions

import (
	"context"
	"github.com/containers/common/libimage"
)

func Push(source string, destination string, options *libimage.PushOptions) ([]byte, error) {
	runtime, err := libimage.RuntimeFromStore(defaultStore(), &libimage.RuntimeOptions{SystemContext: options.SystemContext})
	if err != nil {
		return nil, err
	}

	manifest, err := runtime.Push(context.Background(), source, destination, options)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}