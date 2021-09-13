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

	image, _ , err := runtime.LookupImage(imageName, nil)
	if err != nil {
		return err
	}
	imageData, err := image.Inspect(context.Background(), true)
	if err != nil {
		return err
	}
	manifest, mimeType, _  := image.Manifest(context.Background())
	fmt.Printf("\n Lookedup image manifest: %v \n", string(manifest))
	fmt.Printf("\n Lookedup image mime type: %v \n", mimeType)
	fmt.Printf("\n Looked up image data: %+v", imageData)
	return nil
}