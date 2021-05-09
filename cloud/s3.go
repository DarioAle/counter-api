package cloud

import (
	"bytes"
	"context"

	"github.com/DarioAle/counter-api"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// credenciales en la instancia de aws
func (s *storage) PutImage(ctx context.Context, countImage counter.Image) (counter.Image, error) {
	reader := bytes.NewReader(countImage.Binary)

	input := &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &countImage.KeyID,
		Body:   reader,
	}

	_, err := s.client.PutObject(ctx, input)
	if err != nil {
		return counter.Image{}, err
	}

	return countImage, nil
}
