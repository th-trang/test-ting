package Logic

import (
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(pageIndex int, pageSize int) ([]*DTO.KolDTO, error) {
	var kols []*DTO.KolDTO
	offset := (pageIndex - 1) * pageSize

	// Query the database
	result := Initializers.DB.Limit(pageSize).Offset(offset).Find(&kols).Scan(&kols)
	if result.Error != nil {
		return nil, result.Error
	}

	return kols, nil
}
