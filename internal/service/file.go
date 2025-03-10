package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"jeanfo_mix/config"
	"jeanfo_mix/internal/model"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

type FileService struct {
	DB *gorm.DB
}

func (s *FileService) UploadFile(file io.ReadSeeker, fileName string, userID uint, saveMeta bool) (metaID string, relativePath string, err error) {
	metaID, relativePath, err = "", "", nil

	hash := md5.New()
	file.Seek(0, io.SeekStart)
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}
	hashSum := hex.EncodeToString(hash.Sum(nil))
	// 补充后缀名
	filePostfix := filepath.Ext(fileName)
	hashFilename := hashSum + filePostfix

	baseDir := config.GetConfig().Web.UploadDir
	firstLevelDir := hashSum[:2]
	secondLevelDir := hashSum[2:4]
	storageDir := filepath.Join(baseDir, firstLevelDir, secondLevelDir)
	err = os.MkdirAll(storageDir, 0755)
	if err != nil {
		err = errors.New("fail create storage dir: " + err.Error())
		return
	}

	relativePath = filepath.Join(firstLevelDir, secondLevelDir, hashFilename)
	storagePath := filepath.Join(baseDir, relativePath)

	outfile, err := os.Create(storagePath)
	if err != nil {
		err = errors.New("create file on saving fail: " + err.Error())
		return
	}
	defer outfile.Close()

	file.Seek(0, io.SeekStart)
	_, err = io.Copy(outfile, file)
	if err != nil {
		err = errors.New("copy file on saving fail: " + err.Error())
		return
	}

	if saveMeta {
		fileModel := &model.File{
			UserID: userID, FileName: fileName, RelativePath: relativePath,
		}
		err = fileModel.Create(s.DB)
		if err != nil {
			err = errors.New("save file meta to DB fail: " + err.Error())
			return
		}
		metaID = fileModel.MetaID
	}

	return
}
