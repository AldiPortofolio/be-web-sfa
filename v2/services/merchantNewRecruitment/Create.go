package merchantNewRecruitment

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"

	"github.com/astaxie/beego/logs"
)

// Create ..
func (svc *ServiceMerchantNewRecruitment) Create(token string, req dbmodels.MerchantNewRecruitments, res *models.Response) {
	fmt.Println(">>> Create - ServiceMerchantNewRecruitment <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	//validation
	if req.CustomerCode == "" { //&& req.PhoneNumber == "" {
		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.create.failed")
		return
	}

	//if req.CustomerCode != "" && req.PhoneNumber != "" {
	//	data, _ := svc.Database.ValidationCheckCustomerCode(req.CustomerCode)
	//	if data.CustomerCode == req.CustomerCode {
	//		data, _ := svc.Database.ValidationCheckPhoneNumber(req.PhoneNumber)
	//		if data.PhoneNumber == req.PhoneNumber {
	//			res.Meta = utils.GetMetaResponse("merchant.new.recruitment.create.phone.number.exist")
	//			return
	//		}else{
	//			data, _ := svc.Database.ValidationCheckCustomerCodeStatusPending(req.CustomerCode)
	//			if data.CustomerCode == req.CustomerCode {
	//				fmt.Println("Update by Customer Code (Phone Number Not Exist)")
	//				err = svc.Database.UpdateByCustomerCode(req)
	//				if err != nil {
	//					res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.failed")
	//					return
	//				}
	//				res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.success")
	//				return
	//			}else{
	//				res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.failed")
	//				return
	//			}
	//		}
	//	}else{
	//		data, _ := svc.Database.ValidationCheckPhoneNumberStatusPending(req.PhoneNumber)
	//		if data.PhoneNumber == req.PhoneNumber {
	//			fmt.Println("Update by Phone Number (Customer Code Not Exist)")
	//			err = svc.Database.UpdateByPhoneNumber(req)
	//			if err != nil {
	//				res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.failed")
	//				return
	//			}
	//			res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.success")
	//			return
	//		}
	//	}
	//}else

	//get salestype
	salesType, err := svc.Database.GetSalesTypeIDBySubArea(req.SubAreaChannelID)
	if err != nil {
		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.create.failed")
		return
	}
	req.SalesTypeId = salesType.SalesAreaChannelID

	if req.CustomerCode != "" { //&& req.PhoneNumber == "" {
		data, _ := svc.Database.ValidationCheckCustomerCodeStatusPending(req.CustomerCode)
		if data.CustomerCode == req.CustomerCode {
			fmt.Println("Update by Customer Code (Phone Number Null)")
			// res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.success")
			res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
			err = svc.Database.UpdateByCustomerCode(req)
			if err != nil {
				res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.failed")
				return
			}
			res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.success")
			return
		}
	}
	//else if req.CustomerCode == "" && req.PhoneNumber != "" {
	//	data, _ := svc.Database.ValidationCheckPhoneNumberStatusPending(req.PhoneNumber)
	//	if data.PhoneNumber == req.PhoneNumber {
	//		fmt.Println("Update by Phone Number (Customer Code Null)")
	//		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.success")
	//		err := svc.Database.UpdateByPhoneNumber(req)
	//		if err != nil {
	//			res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.failed")
	//			return
	//		}
	//		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.update.success")
	//		return
	//	}
	//}

	_, err = svc.Database.MerchantNewRecruitmentCreate(req)
	if err != nil {
		logs.Error("error recruitment ", err.Error())
		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.create.failed")
		return
	}

	// res.Meta = utils.GetMetaResponse("merchant.new.recruitment.create.success")
	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	//res.Data = data

	return
}
