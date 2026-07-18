package image

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"restApi/config"
)

type Repository interface {
	Save(context.Context, string, io.Reader) error
}

type LocalStorage struct {
	conf *config.ImageConfig
}

func NewLocalStorage(conf *config.ImageConfig) *LocalStorage {
	if err := os.MkdirAll(conf.LocalStoragePath, 0755); err != nil {
		panic(err)
	}
	return &LocalStorage{
		conf: conf,
	}
}

func (l *LocalStorage) Save(_ context.Context, name string, r io.Reader) error {
	dst, err := os.Create(filepath.Join(l.conf.LocalStoragePath, name))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, r)
	return err
}
