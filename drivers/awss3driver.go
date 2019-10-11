package drivers

import (
	"fmt"
	"os"
	"io"
	"time"
	"path"
	"goftp.io/server"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"	
)

type BucketObject struct {
	server.FileInfo
	isRoot bool
	isDir bool
	name string
	size int64
	modified time.Time
	owner string
}


type AWSS3Driver struct {
	bucketName string
	session *session.Session
	client *s3.S3
	uploader *s3manager.Uploader
	downloader *s3manager.Downloader
}

type FileInfo struct {
	os.FileInfo

	mode  os.FileMode
	owner string
	group string
}

func (f *BucketObject) Mode() os.FileMode {
	return 777
}

func (f *BucketObject) Owner() string {
	return f.owner
}

func (f *BucketObject) Group() string {
	return f.owner
}

func (f *BucketObject) IsDir() bool {
	return f.isDir
}

func (f *BucketObject) Name() string {
	return f.name
}

func (f *BucketObject) Size() int64 {
	return f.size
}

func (f *BucketObject) ModTime() time.Time {
	return f.modified
}

func (f *BucketObject) Sys() interface{} {
	return nil
}

type AWSS3DriverFactory struct {
	BucketName string
	AWSSession *session.Session
}

func (factory *AWSS3DriverFactory) NewDriver() (server.Driver, error){
	fmt.Println("New AWS S3 FTP Driver")

	session, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		fmt.Println("Error creating session", err)
		return nil, err
	}

	driver := &AWSS3Driver{
		bucketName: factory.BucketName,
		session: session,
		client: s3.New(session, &aws.Config{}),
		uploader: s3manager.NewUploader(session),
		downloader: s3manager.NewDownloader(session),
	}

	return driver, nil
}

func (driver *AWSS3Driver) Init(conn *server.Conn) {
	fmt.Println("AWS S3 FTP Driver Init")
}

func (driver *AWSS3Driver) ListDir(path string, onFile func(server.FileInfo) error) error {
	fmt.Printf("AWS S3 FTP Driver ListDir %s\n", path)

	pageNum := 0
	var bucket *string
	bucket = &driver.bucketName 
	params := &s3.ListObjectsInput{Bucket: bucket}
	driver.client.ListObjectsPages(params, func(page *s3.ListObjectsOutput, lastPage bool) bool {
        pageNum++

		for _, item := range page.Contents {
			name := *item.Key
			isDir := (name[len(name) -1] == '/')
			if isDir{
				name = name[:len(name) -1]
			}

			f := &BucketObject{
				name: name,
				isDir: isDir,  
				size: *item.Size,
				modified: *item.LastModified,
				owner: *item.Owner.DisplayName,
			}

			onFile(f)
		}

        return pageNum <= 1
	})
	
	return nil
}

// Downloads an S3 Bucket Object 
func (driver *AWSS3Driver) GetFile(filename string, s int64) (int64, io.ReadCloser, error) {
	fmt.Printf("AWS S3 FTP Driver GetFile '%s'\n", filename)

	key := filename

	buffer := aws.NewWriteAtBuffer([]byte{})

	bytesRead, err := driver.downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(driver.bucketName),
		Key:    aws.String(key),
	})

	fmt.Printf("Downloaded %d bytes from %s\n", bytesRead, key)
	
	if err != nil {
		return 0, nil, fmt.Errorf("failed to download file, %v", err)
	}

	pr, pw := io.Pipe()
	
	go func() {
		bytesWritten, err := pw.Write(buffer.Bytes())
		if err != nil {
			fmt.Printf("failed to download file, %v\n", err)
			return
		}
	
		fmt.Printf("Wrote %d bytes\n", bytesWritten)
		pw.Close()
	}()

	return bytesRead, pr, nil
}

// Use S3 Upload Manager to upload file to bucket root
func (driver *AWSS3Driver) PutFile(destPath string, fileReader io.Reader, u bool) (int64, error) {

	_, key := path.Split(destPath)

	fmt.Printf("AWS S3 FTP Driver PutFile '%s' => '%s'\n", destPath, key)

	result, err := driver.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(driver.bucketName),
		Key:    aws.String(key),
		Body:   fileReader,
	})

	if err != nil {
		return 0, fmt.Errorf("failed to upload file %v", err)
	}

	fmt.Printf("file uploaded to, %s\n", result.Location)

	return 0, nil
}


func (driver *AWSS3Driver) Stat(path string) (server.FileInfo, error) {
	fmt.Printf("AWS S3 FTP Driver Stat path: %s\r\n", path)

	if path == "/" {
		return &BucketObject{ isRoot: true, name: "/", size: 0, isDir: true }, nil 
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String(driver.bucketName),
		Key:    aws.String(path),
	}

	fmt.Println(input)

	//TODO: get s3.GetObject
	return &BucketObject{}, nil
}

func (driver *AWSS3Driver) ChangeDir(string) error {
	fmt.Println("AWS S3 FTP Driver ChangeDir")
	//TODO: Not Implemented
	return nil
}

func (driver *AWSS3Driver) DeleteDir(string) error {
	fmt.Println("AWS S3 FTP Driver DeleteDir")
	//TODO: Not Implemented
	return nil
}

func (driver *AWSS3Driver) DeleteFile(string) error {
	fmt.Println("AWS S3 FTP Driver DeleteFile")
	//TODO: Not Implemented
	return nil
}

func (driver *AWSS3Driver) Rename(string, string) error {
	fmt.Println("AWS S3 FTP Driver Rename")
	//TODO: Not Implemented
	return nil
}

func (driver *AWSS3Driver) MakeDir(string) error {
	fmt.Println("AWS S3 FTP Driver MakeDir")
	//TODO: Not Implemented
	return nil
}

