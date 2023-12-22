package merchant

import (
	"github.com/zazhiladhf/newbie-ecommerce/pkg/helper"
)

type Merchant struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Address     string `db:"address"`
	ImageUrl    string `db:"image_url"`
	City        string `db:"city"`
	AuthId      int    `db:"auth_id"`
}

func NewMerchant() Merchant {
	return Merchant{}
}

func (m Merchant) newFromRequest(req RequestBodyCreateMerchant) (merchant Merchant, err error) {
	merchant = Merchant{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		ImageUrl:    req.ImageUrl,
		City:        req.City,
		// AuthId:      req.AuthId,
	}

	err = merchant.ValidateRequestCreateMerchant()
	return
}

func (m Merchant) ValidateRequestCreateMerchant() (err error) {
	if m.Name == "" {
		return helper.ErrEmptyNameMerchant
	}

	if m.Address == "" {
		return helper.ErrEmptyAddressMerchant
	}

	if m.PhoneNumber == "" {
		return helper.ErrEmptyPhoneNumber
	}

	if len(m.PhoneNumber) < 10 {
		return helper.ErrInvalidPhoneNumber
	}

	if m.ImageUrl == "" {
		return helper.ErrEmptyImageURLMerchant
	}

	if m.City == "" {
		return helper.ErrEmptyCity
	}

	if err != nil {
		return helper.ErrMerchantNotFound
	}

	return
}

func (m Merchant) MerchantResponse(merchant Merchant) GetMerchantResponse {
	resp := GetMerchantResponse{
		Name:        merchant.Name,
		PhoneNumber: merchant.PhoneNumber,
		Address:     merchant.Address,
		ImageUrl:    merchant.ImageUrl,
		City:        merchant.City,
	}
	return resp
}
