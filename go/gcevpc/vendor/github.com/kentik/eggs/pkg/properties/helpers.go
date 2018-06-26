package properties

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func loadPropertiesFromFilesystem(root string) (map[string]string, error) {
	props := make(map[string]string)

	visit := func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("loadPropertiesFromFilesystem: could not walk path '%s': %+v", path, err)
		}
		if path != root && len(path) > len(root)+1 {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("loadPropertiesFromFilesystem: could not read '%s': %+v", path, err)
			}
			propName := path[len(root)+1:]
			props[propName] = string(content)
		}
		return nil
	}

	if err := filepath.Walk(root, visit); err != nil {
		return nil, fmt.Errorf("loadPropertiesFromFilesystem: could not walk root '%s': %+v", root, err)
	}
	return props, nil
}
