// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/minio/minio-go/v7"
// 	"github.com/minio/minio-go/v7/pkg/credentials"
// )

// func main() {
// 	endpoint := "127.0.0.1:9000"
// 	accessKeyID := "minioadmin"
// 	secretAccessKey := "minioadmin"
// 	useSSL := false

// 	// Initialize minio client object.
// 	minioClient, err := minio.New(endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
// 		Secure: useSSL,
// 	})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())

// 	defer cancel()

// 	objectCh := minioClient.ListObjects(ctx, "testbucket", minio.ListObjectsOptions{
// 		//Prefix:    "myprefix",
// 		Recursive: true,
// 	})
// 	for object := range objectCh {
// 		if object.Err != nil {
// 			fmt.Println(object.Err)
// 			return
// 		}
// 		//fmt.Printf("ETag: %v | Name: %v |Size: %v\n", object.ETag, object.Key, object.Size)

// 		downloadObject(*minioClient, "testbucket", object.Key, fmt.Sprintf("/tmp/miniodownloads/%s", object.Key))
// 	}

// 	uploadObject(*minioClient, "testbucket", "uploadfile.txt")

// }

// func downloadObject(minioClient minio.Client, bucket string, object string, downloadpath string) error {

// 	err := minioClient.FGetObject(context.Background(), bucket, object, downloadpath, minio.GetObjectOptions{})
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	fmt.Println("downloaded " + object + " sucessfully")
// 	return nil
// }

// func uploadObject(minioClient minio.Client, bucket string, filename string) error {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	defer file.Close()

// 	fileStat, err := file.Stat()
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	uploadInfo, err := minioClient.PutObject(context.Background(), bucket, fmt.Sprintf("/user/%s", filename), file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
// 	return nil
// }
package examples
