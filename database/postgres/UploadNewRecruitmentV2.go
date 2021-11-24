package postgres

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

// MerchantNewRecruitmentUploadV2 ..
func (database *DbPostgres) MerchantNewRecruitmentUploadV2(fileBytes []byte) (models.BulkLinkError, error) {
	fmt.Println(">>> MerchantNewRecruitmentV2 - Upload - Postgres <<<")
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
		errorMessages := ValidateParamRecruitEmptyValues(params)
		if errorMessages != "" {
			rowError := SetRowError(i, errorMessages, params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		village, vErr := GetVillage(params.VillageID)
		if vErr != nil {
			rowError := SetRowError(i, "KeluruhanID tidak terdaftar", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		district := GetDistrict(village.DistrictId)
		city := GetCity(uint(district.CityID))

		var subArea dbmodels.SubArea
		sErr := Dbcon.Where("name = ?", params.SubAreaChannel).Find(&subArea).Error
		if sErr != nil {
			rowError := SetRowError(i, "Sub Area Channel tidak terdaftar", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		var salesType dbmodels.SalesType
		stErr := Dbcon.Where("name = ?", params.SalesType).Find(&salesType).Error
		if stErr != nil {
			rowError := SetRowError(i, "Sub Area Channel tidak terdaftar", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		var salesTypeBySAC dbmodels.SalesType
		saErr := Dbcon.Where("name = ?", params.SalesType).Find(&salesTypeBySAC).Error
		if saErr != nil {
			rowError := SetRowError(i, "Sales Type tidak match dengan Sub Area Channel", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

		var merchant dbmodels.MerchantNewRecruitments
		err := Dbcon.Where("customer_code = ? ", params.CustomerCode).Find(&merchant).Error
		if err == nil && merchant.Status != "Pending" {
			rowError := SetRowError(i, "merchant telah terdaftar disistem", params)
			dataErrors = append(dataErrors, rowError)
			continue
		}

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

// ValidateParamRecruitEmptyValues ..
func ValidateParamRecruitEmptyValues(params models.BulkNewRecruitment) string {
	var errorMessages []string
	if params.CustomerCode == "" {
		errorMessages = append(errorMessages, "ID pelanggan tidak boleh kosong")
	}

	if params.StoreName == "" || params.Institution == "" || params.Address == "" || params.VillageID == "" || params.SubAreaChannel == "" || params.SalesType == "" {
		errorMessages = append(errorMessages, "Nama Toko, Institusi, Alamat, KelurahanID, Sales Type, atau SubAreaChannelCode tidak boleh kosong")
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
