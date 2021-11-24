package postgres

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

// MerchantNewRecruitmentUpload ..
func (database *DbPostgres) MerchantNewRecruitmentUpload(fileBytes []byte) (models.BulkLinkError, error) {
	fmt.Println(">>> MerchantNewRecruitment - Upload - Postgres <<<")
	// sugarLogger := database.General.OttoZaplog
	var recruitCSV []models.BulkNewRecruitment
	recruitErrResponse := models.BulkLinkError{
		ErrorFile: "",
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use pipe as delimiter
	})

	err := gocsv.UnmarshalBytes(fileBytes, &recruitCSV)
	if err != nil {
		log.Println("Failed to extract file : ", err)
		return recruitErrResponse, err
	}

	var dataErrors []models.DataNewRecruitmentErrorByRow
	var successCount int
	for i, params := range recruitCSV {
		fmt.Println("csv new recruitment ------>>", params)
		errorMessages := ValidateRecruitEmptyValues(params)
		if errorMessages != "" {
			rowError := SetRowError(i, errorMessages, params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		village, vErr := GetVillage(params.VillageID)
		if vErr != nil {
			rowError := SetRowError(i, "KeluruhanID tidak terfdaftar", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		district := GetDistrict(village.DistrictId)
		city := GetCity(uint(district.CityID))

		var subArea dbmodels.SubArea
		sErr := Dbcon.Where("name = ?", params.SubAreaChannel).Find(&subArea).Error
		if sErr != nil {
			rowError := SetRowError(i, "Sub Area Channel tidak terfdaftar", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}
		var merchant dbmodels.MerchantNewRecruitments
		// var merchant, merchantByPhone dbmodels.MerchantNewRecruitments
		err := Dbcon.Where("customer_code = ? ", params.CustomerCode).Find(&merchant).Error
		// Dbcon.Where("phone_number = ? ", params.MerchantPhone).Find(&merchantByPhone)
		if err == nil && merchant.Status != "Pending" {
			rowError := SetRowError(i, "merchant telah terdaftar disistem", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		// if merchantByPhone.PhoneNumber != "" && merchantByPhone.Id != merchant.Id {
		// 	rowError := SetRowError(i, "Nomor HP telah terdaftar", params)
		// 	dataErrors = append(dataErrors, rowError)
		// 	continue
		// }

		if params.MerchantPhone != "" && params.MerchantPhone[0:1] != "0" {
			rowError := SetRowError(i, "Pastikan nomor HP yang di masukan sesuai", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		merchant.Name = params.StoreName
		merchant.PhoneNumber = params.MerchantPhone
		merchant.CustomerCode = params.CustomerCode
		merchant.InstitutionCode = params.Institution
		merchant.SubAreaChannelID = int64(subArea.ID)
		merchant.SubAreaChannelName = params.SubAreaChannel
		merchant.OwnerName = params.Owner
		merchant.Address = params.Address
		merchant.Longitude = params.Longitude
		merchant.Latitude = params.Latitude
		merchant.ProvinceId = int64(city.ProvinceID)
		merchant.CityId = int64(district.CityID)
		merchant.DistrictId = int64(village.DistrictId)
		merchant.VillageId = int64(village.Id)
		merchant.Status = "Pending"
		merchant.IdCard = params.IDCard

		mErr := Dbcon.Save(&merchant).Error
		if mErr != nil {
			rowError := SetRowError(i, "Failed to save merchant new recruitment : "+mErr.Error(), params)
			dataErrors = append(dataErrors, rowError)
		} else {
			successCount++
		}

	}

	if len(dataErrors) > 0 {
		dataErrCount := strconv.Itoa(len(dataErrors))
		succCount := strconv.Itoa(successCount)
		recruitErrResponse.ErrorFile = generateURLErrorFileNewRecruitment(dataErrors, succCount, dataErrCount)
		return recruitErrResponse, errors.New(succCount + " data berhasil diproses " + dataErrCount + " data gagal diproses.")
	}
	//else {
	return recruitErrResponse, nil
	//}
}

// generateURLErrorFileNewRecruitment ..
func generateURLErrorFileNewRecruitment(dataErrors []models.DataNewRecruitmentErrorByRow, calCount string, dataErrCount string) string {
	csvContent, err := gocsv.MarshalString(&dataErrors) // Get all clients as CSV string

	fmt.Println("Export ------ ", csvContent)
	if err != nil {
		panic(err)
	}
	encodedString := base64.StdEncoding.EncodeToString([]byte(csvContent))
	errorCSVContent := "data:text/csv;base64," + encodedString
	// errorFileURL := ""
	errorFileURL := utils.UploadFileError(errorCSVContent, "NewRecruitment")
	log.Println("=====>>> URL --->", errorFileURL)

	if errorFileURL != "" {
		bulkError := dbmodels.BulkErrorFile{
			ErrorFile: errorFileURL,
			BulkType:  "assignment",
			Message:   calCount + " data berhasil diproses. " + dataErrCount + " data gagal diproses.",
		}

		bulkErr := Dbcon.Create(&bulkError).Error

		if bulkErr != nil {
			log.Println("Failed to create file error : ", bulkErr)
		}
	}

	return errorFileURL
}

// SetRowError ..
func SetRowError(idx int, msg string, params models.BulkNewRecruitment) models.DataNewRecruitmentErrorByRow {
	var rowError models.DataNewRecruitmentErrorByRow
	err := errors.New(msg)
	data, _ := json.Marshal(params)
	json.Unmarshal(data, &rowError)
	rowError.NoRow = idx + 1
	rowError.ErrorMessages = err.Error()

	return rowError
}

// ValidateRecruitEmptyValues ..
func ValidateRecruitEmptyValues(params models.BulkNewRecruitment) string {
	var errorMessages []string
	if params.CustomerCode == "" {
		errorMessages = append(errorMessages, "ID pelanggan tidak boleh kosong")
	}

	// if params.MerchantPhone == "" {
	// 	errorMessages = append(errorMessages, "No HP tidak boleh kosong")
	// }

	if params.StoreName == "" || params.Institution == "" || params.Address == "" || params.VillageID == "" || params.SubAreaChannel == "" {
		errorMessages = append(errorMessages, "Nama Toko, Institusi, Alamat, KelurahanID, atau SubAreaChannelCode tidak boleh kosong")
	}

	if params.Longitude == "" && params.Latitude != "" {
		errorMessages = append(errorMessages, "Longitude Mohon diisi")
	}

	if params.Longitude != "" && params.Latitude == "" {
		errorMessages = append(errorMessages, "Latitude Mohon diisi")
	}

	if len(params.IDCard) > 16 {
		errorMessages = append(errorMessages, "NIK Tidak boleh lebih dari 16 digit")
	}

	st := regexp.MustCompile(`[a-z]`)
	if st.MatchString(params.IDCard) {
		errorMessages = append(errorMessages, "NIK harus angka")
	}

	return strings.Join(errorMessages, "|")
}

// GetVillage ..
func GetVillage(villageID string) (dbmodels.Village, error) {
	var village dbmodels.Village
	err := Dbcon.Where("id = ?", villageID).Find(&village).Error

	return village, err
}

// GetDistrict ..
func GetDistrict(districtID uint) dbmodels.District {
	var district dbmodels.District
	Dbcon.Where("id = ?", districtID).Find(&district)

	return district
}

// GetCity ..
func GetCity(cityID uint) dbmodels.Cities {
	var city dbmodels.Cities
	Dbcon.Where("id = ?", cityID).Find(&city)

	return city
}
