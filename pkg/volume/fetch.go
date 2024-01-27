package volume

import (
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Fetcher struct {
	Logger *log.Logger
}

func (f *Fetcher) Fetch(dir string, pvc string, metadataDir string) error {
	f.Logger.Printf("Fetching %q...", pvc)

	projectMetadataFile, err := os.Create(path.Join(metadataDir, "project-metadata.toml"))
	if err != nil {
		return errors.Wrapf(err, "invalid metadata destination '%s/project-metadata.toml' for volume: %s", metadataDir, pvc)
	}
	defer projectMetadataFile.Close()

	projectMd := Project{
		Source: Source{
			Type: "volume",
			Metadata: Metadata{
				PersistentVolumeClaim: pvc,
			},
		},
	}
	if err := toml.NewEncoder(projectMetadataFile).Encode(projectMd); err != nil {
		return errors.Wrapf(err, "invalid metadata destination '%s/project-metadata.toml' for volume: %s", metadataDir, pvc)
	}

	f.Logger.Printf("Successfully fetched %q in path %q", pvc, dir)

	return nil
}

type Project struct {
	Source Source `toml:"source"`
}

type Source struct {
	Type     string   `toml:"type"`
	Metadata Metadata `toml:"metadata"`
}

type Metadata struct {
	PersistentVolumeClaim string `toml:"persistentVolumeClaim"`
}
