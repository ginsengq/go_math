package geometry

import "errors"

func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("Zero length not allowed")
	}
}

//semantic versioning - universally accepted format to tack release of packages,
//modules directly to version string

//vMAJOR.MINOR.PATCH
//v1.2.3
//major - change when you make incompatible changes
//minor - when add functionality in backward minor
//patch - when make backward compatible fixes

//git tag -a v1.0.0 -m "initial release"
//git push origin master --tags
//we published module with its corresponding tag/version
