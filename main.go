package main
import (
	"fmt"
	"log"
	"os"
	"bytes"
	// "net/http"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"

)

var AWS_S3_REGION string
var AWS_S3_BUCKET string
var AWS_ACCESS_KEY_ID string
var AWS_ACCESS_KEY_SECRET string

func main() {



	AWS_S3_REGION = os.Getenv("AWS_S3_REGION")
	AWS_S3_BUCKET = os.Getenv("AWS_S3_BUCKET")

	// Create a single AWS session (we can re use this if we're uploading many files)

	sess, err := session.NewSession(&aws.Config{
		 Region: aws.String(AWS_S3_REGION),
	})

	if err != nil {
		log.Fatal(err)
	}


	// checking if required credentials are present ref : https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	}
 
	filepath := os.Args[1]
	fmt.Println("filepath is ",filepath)

	// Upload
	err = AddFileToS3(sess, filepath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

}

func AddFileToS3(s *session.Session, fileDir string) error {

    // Open the file for use
    file, err := os.Open(fileDir)
    if err != nil {
        return err
    }
    defer file.Close()

    // Get file size and read the file content into a buffer
    fileInfo, _ := file.Stat()
    var size int64 = fileInfo.Size()
    buffer := make([]byte, size)
    file.Read(buffer)
	// fmt.Println("ContentType is ", http.DetectContentType(buffer))


    // Config settings: this is where you choose the bucket, filename, content-type etc. of the file you're uploading.
    _, err = s3.New(s).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String(AWS_S3_BUCKET),
        Key:                  aws.String("postgres_backups_1"),
        Body:                 bytes.NewReader(buffer),
        // ContentType:          aws.String(http.DetectContentType(buffer)),
        // ContentDisposition:   aws.String("attachment"),
    })

    return err
}