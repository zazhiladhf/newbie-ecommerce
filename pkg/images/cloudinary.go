package images

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Object cloudinanry
type Cloudinary struct {
	// client cloudinary
	client    *cloudinary.Cloudinary
	cloud     string
	apiKey    string
	apiSecret string
}

// construct object cloudinary
func NewCloudinary(cloud, apiKey, apiSecret string) (resp Cloudinary, err error) {
	// try to initiate cloudinary client
	client, err := cloudinary.NewFromParams(cloud, apiKey, apiSecret)
	if err != nil {
		return
	}

	return Cloudinary{
		client:    client,
		cloud:     cloud,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}, nil
}

// implement from contract
func (c Cloudinary) Upload(ctx context.Context, file interface{}, pathDestination string, quality string) (uri string, err error) {
	// try to change filename with uuid
	filename := time.Now().UnixNano()

	// proses cloudinary nge-upload image
	res, err := c.client.Upload.Upload(ctx, file, uploader.UploadParams{
		// publicID => adalah lokasi dan nama file nya.
		// jadi nanti nya akan ada di folder "NBID_Training"
		PublicID: "Ecommerce/" + pathDestination + "/" + fmt.Sprintf("%v", filename),

		// for handle transformation
		Eager: quality, // berfungsi untuk mengirim qualitas yang akan digunakan
	})

	if err != nil {
		return "", err
	}

	// check if there are any eager in response
	if len(res.Eager) > 0 {
		// will return secure url with transformation
		return res.Eager[0].SecureURL, nil
	}

	// if no, will use secure url (without transformation)
	url := res.SecureURL

	return url, nil
}

// func (c Cloudinary) Remove(ctx context.Context, path string) error {
// 	res, err := c.Cloud.Upload.Destroy(ctx, uploader.DestroyParams{
// 		//public id must not contains format
// 		//e.g: NBID-Training/IniDariPublicId
// 		//format: <file>/<public_id>
// 		PublicID: path,
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if strings.Contains(res.Result, "not found") {
// 		return errors.New("image not found")
// 	}

// 	fmt.Printf("%+v\n", res)

// 	return err
// }
