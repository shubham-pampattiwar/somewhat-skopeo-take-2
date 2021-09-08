package functions

import (
	"context"
	"github.com/containers/common/libimage"
)

func Save (imageNames []string, format string, path string, options *libimage.SaveOptions) error {
	runtime, err := libimage.RuntimeFromStore(defaultStore(), &libimage.RuntimeOptions{SystemContext: options.SystemContext})
	if err != nil {
		return err
	}

	err = runtime.Save(context.Background(), imageNames, format, path, options)
	if err != nil {
		return err
	}

	return nil
}
