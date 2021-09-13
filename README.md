# somewhat-skopeo-take-2
Another jab at copying container images for the sake of backup and restore ops, lets gooo...

Workflows and Findings:
- Case 1:
  - Intialize Storage
  - Pull imge
  - save image
  - load image without wiping out the local storage
  - lookup image data and compare
  - Results: Image Data inspection reveals that there is change in `Config`, `Size`, `VirtualSize` , `GraphDriver` and `RootFS`. Also, an additional `RepoDigest` is present in the loaded image from tar.
- Case 2:
  - Intialize Storage
  - Pull imge
  - save image
  - load image after wiping out the local storage
  - lookup image data and compare
  - Results: Along with the case 1 results, the `digest SHA` is also different


Reference links:
- Import image steps: https://github.com/containers/common/blob/main/libimage/import.go
- Export Image steps: https://github.com/containers/common/blob/main/libimage/save.go
- blog for container/storage and containers/image: https://iximiuz.com/en/posts/working-with-container-images-in-go/
- Docker cmd info: https://appfleet.com/blog/how-to-transfer-move-a-docker-image-to-another-system/
- Podman export code: https://github.com/containers/podman/blob/23f9565547ae2a6b0154e6913abf7f1232f0ece0/pkg/domain/infra/tunnel/containers.go#L288
- Most important lib, combines storage and image: https://github.com/containers/common/tree/main/libimage
- Podman save: https://github.com/containers/podman/blob/536f23c0b78dd8feafee4e40b743988dbb03bfa2/vendor/github.com/containers/common/libimage/save.go#L35