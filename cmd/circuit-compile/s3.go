package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/vocdoni/davinci-dkg/log"
)

// S3Config holds the configuration for DigitalOcean Spaces / S3 uploads.
type S3Config struct {
	Enabled   bool
	HostBase  string
	AccessKey string
	SecretKey string
	Space     string // bucket name (e.g. "circuits")
	Bucket    string // folder / release channel (e.g. "dev")
}

// NewDefaultS3Config returns an S3Config pre-populated with the same CDN
// endpoint used by davinci-node.
func NewDefaultS3Config() *S3Config {
	return &S3Config{
		Enabled:  false,
		HostBase: "ams3.digitaloceanspaces.com",
		Space:    "circuits",
		Bucket:   "dev",
	}
}

// S3Uploader handles artifact uploads to DigitalOcean Spaces via the S3 API.
type S3Uploader struct {
	client *s3.Client
	config *S3Config
}

// NewS3Uploader creates a new S3Uploader from the given config.
func NewS3Uploader(cfg *S3Config) (*S3Uploader, error) {
	if !cfg.Enabled {
		return nil, fmt.Errorf("s3 upload not enabled")
	}
	if cfg.AccessKey == "" || cfg.SecretKey == "" {
		return nil, fmt.Errorf("s3 access key and secret key are required")
	}

	sdkCfg, err := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.AccessKey, cfg.SecretKey, "",
		)),
		awsconfig.WithRegion("us-east-1"), // required by SDK; ignored by DO Spaces
	)
	if err != nil {
		return nil, fmt.Errorf("load aws config: %w", err)
	}

	client := s3.NewFromConfig(sdkCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s", cfg.HostBase))
		o.UsePathStyle = true
	})
	return &S3Uploader{client: client, config: cfg}, nil
}

// UploadFile uploads localPath to S3, using the file's base name as the object key
// within the configured bucket folder.  Returns the full object key.
func (u *S3Uploader) UploadFile(ctx context.Context, localPath string) (string, error) {
	return u.UploadFileAs(ctx, localPath, filepath.Base(localPath))
}

// UploadFileAs uploads localPath to S3 under the given remoteName (without the
// bucket prefix).  Use this when the local file has no extension but the CDN
// key should (e.g. "<hash>.ccs").  Returns the full object key.
func (u *S3Uploader) UploadFileAs(ctx context.Context, localPath, remoteName string) (string, error) {
	f, err := os.Open(localPath)
	if err != nil {
		return "", fmt.Errorf("open %s: %w", localPath, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Warnw("close file", "error", err)
		}
	}()

	objectKey := fmt.Sprintf("%s/%s", u.config.Bucket, remoteName)
	log.Infow("uploading artifact", "remote", objectKey, "space", u.config.Space)
	_, err = u.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(u.config.Space),
		Key:    aws.String(objectKey),
		Body:   f,
	})
	if err != nil {
		return "", fmt.Errorf("put object %s: %w", objectKey, err)
	}
	return objectKey, nil
}

// SetPublicACL grants public-read access to the listed object keys.
func (u *S3Uploader) SetPublicACL(ctx context.Context, objectKeys []string) error {
	for _, key := range objectKeys {
		log.Infow("setting public ACL", "object", key)
		_, err := u.client.PutObjectAcl(ctx, &s3.PutObjectAclInput{
			Bucket: aws.String(u.config.Space),
			Key:    aws.String(key),
			ACL:    types.ObjectCannedACLPublicRead,
		})
		if err != nil {
			return fmt.Errorf("set ACL %s: %w", key, err)
		}
	}
	return nil
}

// TestS3Connection verifies that the credentials and endpoint are reachable.
func TestS3Connection(ctx context.Context, cfg *S3Config) error {
	if !cfg.Enabled {
		return nil
	}
	uploader, err := NewS3Uploader(cfg)
	if err != nil {
		return err
	}
	_, err = uploader.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  aws.String(cfg.Space),
		MaxKeys: aws.Int32(1),
	})
	if err != nil {
		return fmt.Errorf("S3 connection test: %w", err)
	}
	log.Infow("S3 connection OK", "host", cfg.HostBase, "space", cfg.Space, "bucket", cfg.Bucket)
	return nil
}
