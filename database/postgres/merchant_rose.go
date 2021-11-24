package postgres

import (
	"fmt"
	"log"
	"ottosfa-api-web/models"
)

// GetMerchantDetailRose ..
func (database *DbPostgres) GetMerchantDetailRose(merchantID string) (models.MerchantDetailV2, error) {
	fmt.Println(">>> TodoList - MerchantDetailRose - DB <<<")
	var res models.MerchantDetailV2

	query := `select 
				m.id,
				m.store_name as merchant_name, 
				m.merchant_outlet_id as merchant_id, 
				m.alamat as address, 
				m.notes as note, 
				m.store_phone_number as merchant_phone, 
				m.kelurahan as village_id,
				m.sr_id as sales_type_id,
				CONCAT(o.owner_first_name, ' ', o.owner_last_name) as owner_name,
				m.patokan as address_benchmark
				from merchant m
				left join "owner" o on o.id = m.owner_id 
				where m.merchant_outlet_id = ?`

	sql := DbConRose.Raw(query, merchantID).Scan(&res).Error

	if sql != nil {
		log.Println("Failed to get admin detail: ", sql)
		return res, sql
	}
	return res, nil
}

// GetMerchantListRose ..
func (database *DbPostgres) GetMerchantListRose(keyword string) ([]models.MerchantList, error) {
	fmt.Println(">>> TodoList - MerchantListRose - DB <<<")

	var res []models.MerchantList

	keywordValue := "%" + keyword + "%"
	err := DbConRose.Raw("select m.merchant_outlet_id as merchant_id, m.store_name as name, m.store_phone_number as phone_number from merchant m where m.store_name LIKE ? OR m.store_phone_number LIKE ? limit 5", keywordValue, keywordValue).Scan(&res).Error

	if err != nil {
		log.Println("Failed to get merchant list: ", err)
		return res, err
	}
	return res, nil
}
