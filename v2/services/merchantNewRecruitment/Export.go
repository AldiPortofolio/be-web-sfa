package merchantNewRecruitment

import (
	"fmt"
	"io/ioutil"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

// Export ..
func (svc *ServiceMerchantNewRecruitment) Export(token string, req models.MerchantNewRecruitmentListReq, res *models.Response) {
	fmt.Println(">>> Export - ServiceMerchantNewRecruitment <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.MerchantNewRecruitmentDataExport(req)
	if err != nil {
		res.Meta = utils.GetMessageFailedError(401, err)
	}

	csvContent, err := gocsv.MarshalString(Export(data))
	if err != nil {
		panic(err)
	}

	r := strings.NewReader(csvContent)
	file, err := ioutil.ReadAll(r)
	if err != nil {
		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.export.failed")
		return
	}

	res.Meta = utils.GetMetaResponse("merchant.new.recruitment.export.success")
	res.Data = file

	return
}

// Export ..
func Export(data []models.MerchantNewRecruitmentExportRes) (res []models.MerchantNewRecruitmentExportCSVRes) {
	for i, val := range data {
		no := strconv.Itoa(i + 1)
		a := models.MerchantNewRecruitmentExportCSVRes{
			No:                 no,
			Name:               val.Name,
			OwnerName:          val.OwnerName,
			CustomerCode:       val.CustomerCode,
			PhoneNumber:        val.PhoneNumber,
			InstitutionCode:    val.InstitutionCode,
			Address:            val.Address,
			ProvinceName:       val.ProvinceName,
			CityName:           val.CityName,
			DistrictName:       val.DistrictName,
			VillageName:        val.VillageName,
			SubAreaChannelName: val.SubAreaChannelName,
			Longitude:          val.Longitude,
			Latitude:           val.Latitude,
		}
		res = append(res, a)
	}

	return res
}
