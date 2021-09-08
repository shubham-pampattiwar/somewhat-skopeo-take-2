package functions

import (
	"context"
	"github.com/containers/common/libimage"
)

func Export (imageNames []string, format string, exportPath string, options *libimage.SaveOptions) error {
	runtime, err := libimage.RuntimeFromStore(defaultStore(), &libimage.RuntimeOptions{SystemContext: options.SystemContext})
	if err != nil {
		return err
	}

	err = runtime.Save(context.Background(), imageNames, format, exportPath, options)
	if err != nil {
		return err
	}

	return nil
}
