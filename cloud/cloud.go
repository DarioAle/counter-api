package cloud

// instanciar paquete de S3

import (
	"github.com/DarioAle/counter-api"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// repository struct holds the S3 client y bucket name
type storage struct {
	bucket string
	client *s3.Client
}

// New returns a storage instance
func New(
	b string,
	c *s3.Client,
) counter.Storage {
	return &storage{
		bucket: b,
		client: c,
	}
}
