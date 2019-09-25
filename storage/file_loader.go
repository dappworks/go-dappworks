// Copyright (C) 2018 go-dappley authors
//
// This file is part of the go-dappley library.
//
// the go-dappley library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dappley library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dappley library.  If not, see <http://www.gnu.org/licenses/>.
//

package storage

import (
	"bytes"
	"io/ioutil"
	"os"
        "path/filepath"
	logger "github.com/sirupsen/logrus"
)

type FileLoader struct {
	filePath string
}

func NewFileLoader(filePath string) *FileLoader {
	return &FileLoader{
		filePath: filePath,
	}
}

func (fl *FileLoader) ReadFromFile() ([]byte, error) {

	if _, err := os.Stat(fl.filePath); os.IsNotExist(err) {
		return nil, err
	} else if err != nil {
		logger.Panic(err)
	}

	fileContent, err := ioutil.ReadFile(fl.filePath)
	if err != nil {
		logger.Panic(err)
	}
	return fileContent, nil

}

func (fl *FileLoader) SaveToFile(buffer bytes.Buffer) {
	dir := filepath.Dir(fl.filePath)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir,os.ModePerm)
			if err != nil {
				logger.Errorf("Create folder error: %v", err.Error())
			}
		}
	}
	err = ioutil.WriteFile(fl.filePath, buffer.Bytes(), 0644)
	if err != nil {
		logger.Panic(err)
	}
}
