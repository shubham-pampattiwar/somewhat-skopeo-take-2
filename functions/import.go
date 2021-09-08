package functions

import (
	"context"
	"github.com/containers/common/libimage"
)

func Import(importPath string, options *libimage.ImportOptions) (string, error) {
	runtime, err := libimage.RuntimeFromStore(defaultStore(), &libimage.RuntimeOptions{SystemContext: options.SystemContext})
	if err != nil {
		return "", err
	}
	options.CopyOptions.ManifestMIMEType = "application/vnd.docker.distribution.manifest.v2+json"

	importedImageName, err := runtime.Import(context.Background(), importPath, options)
	if err != nil {
		return "", err
	}

	return importedImageName, nil
}
