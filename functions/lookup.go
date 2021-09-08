package functions

import (
	"context"
	"fmt"
	"github.com/containers/common/libimage"
)

// check for lookup options vs import options in params
func Lookup(imageName string, options *libimage.ImportOptions) error {
	runtime, err := libimage.RuntimeFromStore(defaultStore(), &libimage.RuntimeOptions{SystemContext: options.SystemContext})
	if err != nil {
		return err
	}

	importedImage, _ , err := runtime.LookupImage(imageName, nil)
	manifest, mimeType, _  := importedImage.Manifest(context.Background())
	fmt.Printf("\n Lookedup image manifest: %v \n", string(manifest))
	fmt.Printf("\n Lookedup image mime type: %v \n", mimeType)
	return nil
}
