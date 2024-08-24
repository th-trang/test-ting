// file: Logic/GetKolLogic.go
package Logic

import (
	"errors"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

// GetKolLogic retrieves KOLs from the database based on the pageIndex and pageSize.
func GetKolLogic(pageIndex int, pageSize int) ([]*DTO.KolDTO, error) {
	var kols []Models.Kol
	offset := (pageIndex - 1) * pageSize

	// Fetch KOLs from the database with pagination
	result := Initializers.DB.Offset(offset).Limit(pageSize).Find(&kols)
	if result.Error != nil {
		return nil, result.Error
	}

	// Check if no records found
	if len(kols) == 0 {
		return nil, errors.New("no KOLs found")
	}

	// Convert Models.Kol to DTO.KolDTO
	var kolDTOs []*DTO.KolDTO
	for _, kol := range kols {
		kolDTO := &DTO.KolDTO{
			KolID:                kol.KolID,
			UserProfileID:        kol.UserProfileID,
			Language:             kol.Language,
			Education:            kol.Education,
			ExpectedSalary:       kol.ExpectedSalary,
			ExpectedSalaryEnable: kol.ExpectedSalaryEnable,
			ChannelSettingTypeID: kol.ChannelSettingTypeID,
			IDFrontURL:           kol.IDFrontURL,
			IDBackURL:            kol.IDBackURL,
			PortraitURL:          kol.PortraitURL,
			RewardID:             kol.RewardID,
			PaymentMethodID:      kol.PaymentMethodID,
			TestimonialsID:       kol.TestimonialsID,
			VerificationStatus:   kol.VerificationStatus,
			Enabled:              kol.Enabled,
			ActiveDate:           kol.ActiveDate,
			Active:               kol.Active,
			CreatedBy:            kol.CreatedBy,
			CreatedDate:          kol.CreatedDate,
			ModifiedBy:           kol.ModifiedBy,
			ModifiedDate:         kol.ModifiedDate,
			IsRemove:             kol.IsRemove,
			IsOnBoarding:         kol.IsOnBoarding,
			Code:                 kol.Code,
			PortraitRightURL:     kol.PortraitRightURL,
			PortraitLeftURL:      kol.PortraitLeftURL,
			LivenessStatus:       kol.LivenessStatus,
		}
		kolDTOs = append(kolDTOs, kolDTO)
	}

	return kolDTOs, nil
}
