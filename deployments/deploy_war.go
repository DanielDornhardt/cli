package deployments

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Scalingo/cli/config"

	"github.com/Scalingo/go-scalingo/v4"

	errgo "gopkg.in/errgo.v1"
)

type DeployWarRes struct {
	Deployment *scalingo.Deployment `json:"deployment"`
}

func DeployWar(appName, warPath, gitRef string, opts DeployOpts) error {
	var warReadStream io.ReadCloser

	var warSize int64
	var warFileName string

	c, err := config.ScalingoClient()
	if err != nil {
		return errgo.Notef(err, "fail to get Scalingo client")
	}

	if strings.HasPrefix(warPath, "http://") || strings.HasPrefix(warPath, "https://") {
		warReadStream, warSize, err = getURLInfo(warPath)
		if err != nil {
			return errgo.Mask(err, errgo.Any)
		}
		warFileName = appName + ".war"
	} else {
		warReadStream, warSize, warFileName, err = getFileInfo(warPath)
		if err != nil {
			return errgo.Mask(err, errgo.Any)
		}
	}
	defer warReadStream.Close()
	// Create the tar header
	header := &tar.Header{
		Name:       fmt.Sprintf("%s/%s", appName, warFileName),
		Typeflag:   tar.TypeReg, // Is a regular file
		Mode:       0640,
		ModTime:    time.Now(),
		AccessTime: time.Now(),
		ChangeTime: time.Now(),
	}
	if warSize != 0 {
		header.Size = warSize
	} else {
		return errgo.New("Unknown WAR size")
	}

	// Get the sources endpoints
	sources, err := c.SourcesCreate()
	if err != nil {
		return errgo.Mask(err, errgo.Any)
	}

	archiveBuffer := new(bytes.Buffer)
	gzWriter := gzip.NewWriter(archiveBuffer)
	tarWriter := tar.NewWriter(gzWriter)

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return errgo.Notef(err, "fail to create tarball")
	}

	_, err = io.Copy(tarWriter, warReadStream)
	if err != nil {
		return errgo.Notef(err, "fail to copy war content")
	}

	tarWriter.Close()
	gzWriter.Close()

	res, err := uploadArchive(sources.UploadURL, archiveBuffer, int64(archiveBuffer.Len()))
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errgo.Newf("wrong status code after upload %s", res.Status)
	}

	return Deploy(appName, sources.DownloadURL, gitRef, opts)
}

func getURLInfo(warPath string) (warReadStream io.ReadCloser, warSize int64, err error) {
	res, err := http.Get(warPath)
	if err != nil {
		err = errgo.Mask(err, errgo.Any)
		return
	}
	warReadStream = res.Body
	if res.Header.Get("Content-Length") != "" {
		i, err := strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)
		if err == nil {
			// If there is an error, we just skip this header
			warSize = i
		}
	}
	return
}

func getFileInfo(warPath string) (warReadStream io.ReadCloser, warSize int64, warFileName string, err error) {
	warFileName = filepath.Base(warPath)
	fi, err := os.Stat(warPath)
	if err == nil {
		// If there is an error, we just skip this header
		warSize = fi.Size()
	}
	warReadStream, err = os.Open(warPath)
	if err != nil {
		err = errgo.Mask(err, errgo.Any)
		return
	}
	return
}
