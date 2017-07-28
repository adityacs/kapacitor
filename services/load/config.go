package load

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const taskDir = "tasks"
const templateDir = "templates"
const handlerDir = "handlers"

type Config struct {
	Enabled bool   `toml:"enabled"`
	Dir     string `toml:"dir"`
}

func NewConfig() Config {
	return Config{
		Enabled: false,
		Dir:     "/etc/kapacitor/load", // TODO: better default
	}
}

// Validates verifies that the directory specified is an absolute path
// and that it contains the directories /tasks and /handlers. The directory
// may contain additional files, but must at least contain /tasks and /handlers.
func (c Config) Validate() error {
	if !c.Enabled {
		return nil
	}

	// Verify that the path is absolute
	if !filepath.IsAbs(c.Dir) {
		return errors.New("dir must be an absolute path")
	}

	// Verify that correct subdirectories exist
	files, err := ioutil.ReadDir(c.Dir)
	if err != nil {
		return err
	}

	dirs := map[string]bool{}

	for _, file := range files {
		if file.IsDir() {
			dirs[file.Name()] = true
		}
	}

	if !dirs[taskDir] {
		return fmt.Errorf("directory %s must be contain subdirectory %s", c.Dir, taskDir)
	}

	if !dirs[templateDir] {
		return fmt.Errorf("directory %s must be contain subdirectory %s", c.Dir, templateDir)
	}

	if !dirs[handlerDir] {
		return fmt.Errorf("directory %s must be contain subdirectory %s", c.Dir, handlerDir)
	}

	// TODO: we should probably check that we have the correct permissions to access the necessary files

	return nil
}

func (c Config) TasksDir() string {
	return filepath.Join(c.Dir, taskDir)
}

func (c Config) TemplatesDir() string {
	return filepath.Join(c.Dir, templateDir)
}

func (c Config) HandlersDir() string {
	return filepath.Join(c.Dir, handlerDir)
}
