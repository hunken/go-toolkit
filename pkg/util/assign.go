package util

import (
	"github.com/tidwall/gjson"
	"time"
)

const TIME_OFFSET_VALUE = -7

const STT_PENDING = 1                          // Đã tạo đơn (chưa tiếp nhận)
const STT_APPROVED = 2                         // Đã tạo đơn (đã tiếp nhận)
const STT_PICKED = 3                           // Đã lấy hàng
const STT_SHIPPING = 4                         // Đang giao hàng
const STT_NOT_PICKED = 7                       // Không lấy được hàng
const STT_DELAY_PICKING = 8                    // Delay lấy hàng
const STT_DELAY_SHIPPING = 10                  // Delay giao hàng
const STT_PICKING = 12                         // Đang lấy hàng
const STT_REFUND = 13                          // Bồi hoàn
const STT_CHATTING = 30                        // Đang chat
const STT_CONFIMED = 31                        // Đang chốt
const STT_DRAFT = 32                           // Nháp
const STT_SHIPPED_ALL_NOT_AUDITED = 50         // Đã giao hàng toàn bộ/ chưa đối soát
const STT_SHIPPED_PART_NOT_AUDITED = 51        // Đã giao hàng 1 phần/ chưa đối soát
const STT_SHIPPED_ALL_AUDITED = 52             // Giao hàng toàn bộ/ đã đối soát
const STT_SHIPPED_PART_AUDITED = 53            // Giao hàng 1 phần/ đã đối soát
const STT_NOT_DELIVERED_NOT_AUDITED = 90       // Không giao được hàng/ chưa đối soát
const STT_NOT_DELIVERED_AUDITED = 91           // Không giao được hàng/ đã đối soát
const STT_STORED_RETURN_ALL_NOT_AUDITED = 180  // Lưu kho trả toàn bộ/ chưa đối soát
const STT_STORED_RETURN_PART_NOT_AUDITED = 181 // Lưu kho trả 1 phần/ chưa đối soát
const STT_STORED_RETURN_ALL_AUDITED = 182      // Lưu kho trả toàn bộ/ đã đối soát
const STT_STORED_RETURN_PART_AUDITED = 183     // Lưu kho trả 1 phần/ đã đối soát
const STT_WAIT_RETURN_ALL_NOT_AUDITED = 190    // Chờ trả hàng toàn bộ/ chưa đối soát
const STT_WAIT_RETURN_PART_NOT_AUDITED = 191   // Chờ trả hàng 1 phần/ chưa đối soát
const STT_WAIT_RETURN_ALL_AUDITED = 192        // Chờ trả hàng toàn bộ/ đã đối soát
const STT_WAIT_RETURN_PART_AUDITED = 193       // Chờ trả hàng 1 phần/ đã đối soát
const STT_RETURNING_ALL_NOT_AUDITED = 200      // Đang trả hàng toàn bộ/ chưa đối soát
const STT_RETURNING_PART_NOT_AUDITED = 201     // Đang trả hàng 1 phần/ chưa đối soát
const STT_RETURNING_ALL_AUDITED = 202          // Đang trả hàng toàn bộ/ đã đối soát
const STT_RETURNING_PART_AUDITED = 203         // Đang trả hàng 1 phần / đã đối soát
const STT_RETURNED_ALL_NOT_AUDITED = 210       // Đã trả hàng toàn bộ/ chưa đối soát
const STT_RETURNED_PART_NOT_AUDITED = 211      // Đã trả hàng 1 phần/ chưa đối soát
const STT_RETURNED_ALL_AUDITED = 212           // Đã trả hàng toàn bộ/ đã đối soát
const STT_RETURNED_PART_AUDITED = 213          // Đã trả hàng 1 phần/ đã đối soát
const STT_NOT_RETURNED_ALL_NOT_AUDITED = 220   // Không trả được hàng toàn bộ/chưa đối soát
const STT_NOT_RETURNED_PART_NOT_AUDITED = 221  // Không trả được hàng 1 phần /chưa đối soát)
const STT_NOT_RETURNED_ALL_AUDITED = 222       // Không trả được hàng toàn bộ /đã đối soát
const STT_NOT_RETURNED_PART_AUDITED = 223      // Không trả được hàng 1 phần/đã đối soát

func AssignIntValue(jsonValue gjson.Result) (pValue *int) {
	if IsNull(jsonValue) || jsonValue.String() == "" {
		pValue = nil
	} else {
		value := int(jsonValue.Int())
		pValue = &value
	}

	return
}

func AssignInt64Value(jsonValue gjson.Result) (pValue *int64) {
	if IsNull(jsonValue) || jsonValue.String() == "" {
		pValue = nil
	} else {
		value := jsonValue.Int()
		pValue = &value
	}

	return
}

func AssignStringValue(jsonValue gjson.Result) (pValue *string) {
	if IsNull(jsonValue) {
		pValue = nil
	} else {
		value := jsonValue.String()
		pValue = &value
	}

	return
}

func AssignTimeValue(jsonValue gjson.Result) (pValue *time.Time) {
	if jsonValue.Int() == 0 {
		pValue = nil
	} else {
		timestampValue := jsonValue.Int()
		timestampValue = AddTimestamp(timestampValue, TIME_OFFSET_VALUE)
		value := time.Unix(timestampValue/1000, 0)
		pValue = &value
	}

	return
}

func AssignNotNilValue(value *int) (statusValue int) {
	statusValue = 0

	if value != nil {
		statusValue = *value
	}

	return
}

func GenStatusId(pkgStatusId int, tmpPickingStatus int, tmpDeliveringStatus int, returnStatus int, returnPartPackage int) (statusId int) {

	statusId = 0

	if pkgStatusId == STT_PENDING || pkgStatusId == STT_APPROVED {
		statusId = pkgStatusId
	} else if pkgStatusId == 12 && tmpPickingStatus == 0 {
		statusId = STT_PICKING
	} else if pkgStatusId == 3 || (pkgStatusId == 12 && (tmpPickingStatus == 1 || tmpPickingStatus == 4)) {
		statusId = STT_PICKED
	} else if pkgStatusId == 8 || (pkgStatusId == 12 && (tmpPickingStatus == 3 || tmpPickingStatus == 6)) {
		statusId = STT_DELAY_PICKING
	} else if pkgStatusId == 7 || (pkgStatusId == 12 && (tmpPickingStatus == 2 || tmpPickingStatus == 5)) {
		statusId = STT_NOT_PICKED
	} else if pkgStatusId == 4 && tmpDeliveringStatus == 0 {
		statusId = STT_SHIPPING
	} else if pkgStatusId == 10 || (pkgStatusId == 4 && (tmpDeliveringStatus == 4 || tmpDeliveringStatus == 8)) {
		statusId = STT_DELAY_SHIPPING
	} else if (pkgStatusId == 5 && returnPartPackage != 1) || (pkgStatusId == 4 && (tmpDeliveringStatus == 1 || tmpDeliveringStatus == 5)) {
		statusId = STT_SHIPPED_ALL_NOT_AUDITED
	} else if (pkgStatusId == 5 && returnPartPackage == 1 && returnStatus == 0) || (pkgStatusId == 4 && (tmpDeliveringStatus == 2 || tmpDeliveringStatus == 6)) {
		statusId = STT_SHIPPED_PART_NOT_AUDITED
	} else if pkgStatusId == 6 && returnPartPackage == 1 && returnStatus == 0 {
		statusId = STT_SHIPPED_PART_AUDITED
	} else if pkgStatusId == 6 && returnPartPackage != 1 {
		statusId = STT_SHIPPED_ALL_AUDITED
	} else if (pkgStatusId == 9 && returnStatus == 0) || (pkgStatusId == 4 && (tmpDeliveringStatus == 3 || tmpDeliveringStatus == 7)) {
		statusId = STT_NOT_DELIVERED_NOT_AUDITED
	} else if pkgStatusId == 11 && returnStatus == 0 {
		statusId = STT_NOT_DELIVERED_AUDITED
	} else if pkgStatusId == 9 && returnStatus == 3 {
		statusId = STT_WAIT_RETURN_ALL_NOT_AUDITED
	} else if pkgStatusId == 5 && returnPartPackage == 1 && returnStatus == 3 {
		statusId = STT_WAIT_RETURN_PART_NOT_AUDITED
	} else if pkgStatusId == 11 && returnStatus == 3 {
		statusId = STT_WAIT_RETURN_ALL_AUDITED
	} else if pkgStatusId == 6 && returnPartPackage == 1 && returnStatus == 3 {
		statusId = STT_WAIT_RETURN_PART_AUDITED
	} else if pkgStatusId == 9 && returnStatus == 1 {
		statusId = STT_RETURNING_ALL_NOT_AUDITED
	} else if pkgStatusId == 5 && returnPartPackage == 1 && returnStatus == 1 {
		statusId = STT_RETURNING_PART_NOT_AUDITED
	} else if pkgStatusId == 11 && returnStatus == 1 {
		statusId = STT_RETURNING_ALL_AUDITED
	} else if pkgStatusId == 6 && returnPartPackage == 1 && returnStatus == 1 {
		statusId = STT_RETURNING_PART_AUDITED
	} else if pkgStatusId == 9 && returnStatus == 4 {
		statusId = STT_STORED_RETURN_ALL_NOT_AUDITED
	} else if pkgStatusId == 5 && returnPartPackage == 1 && returnStatus == 4 {
		statusId = STT_STORED_RETURN_PART_NOT_AUDITED
	} else if pkgStatusId == 11 && returnStatus == 4 {
		statusId = STT_STORED_RETURN_ALL_AUDITED
	} else if pkgStatusId == 6 && returnPartPackage == 1 && returnStatus == 4 {
		statusId = STT_STORED_RETURN_PART_AUDITED
	} else if pkgStatusId == 9 && returnStatus == 2 {
		statusId = STT_RETURNED_ALL_NOT_AUDITED
	} else if pkgStatusId == 5 && returnPartPackage == 1 && returnStatus == 2 {
		statusId = STT_RETURNED_PART_NOT_AUDITED
	} else if pkgStatusId == 11 && returnStatus == 2 {
		statusId = STT_RETURNED_ALL_AUDITED
	} else if pkgStatusId == 6 && returnPartPackage == 1 && returnStatus == 2 {
		statusId = STT_RETURNED_PART_AUDITED
	} else if pkgStatusId == 9 && returnStatus == 5 {
		statusId = STT_NOT_RETURNED_ALL_NOT_AUDITED
	} else if pkgStatusId == 5 && returnPartPackage == 1 && returnStatus == 5 {
		statusId = STT_NOT_RETURNED_PART_NOT_AUDITED
	} else if pkgStatusId == 11 && returnStatus == 5 {
		statusId = STT_NOT_RETURNED_ALL_AUDITED
	} else if pkgStatusId == 6 && returnPartPackage == 1 && returnStatus == 5 {
		statusId = STT_NOT_RETURNED_PART_AUDITED
	} else if pkgStatusId == 13 {
		statusId = STT_REFUND
	}

	return
}
