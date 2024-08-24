// file: Controllers/GetKolsController.go
package Controllers

import (
	"net/http"
	"strconv"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetKolsController handles the HTTP request for retrieving KOLs with pagination.
func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	var guid = uuid.New().String()

	// Extract pageIndex and pageSize from query parameters
	pageIndexStr := context.DefaultQuery("pageIndex", "1")
	pageSizeStr := context.DefaultQuery("pageSize", "10")

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil || pageIndex < 1 {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = "Invalid pageIndex"
		KolsVM.Guid = guid
		context.JSON(http.StatusBadRequest, KolsVM)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = "Invalid pageSize"
		KolsVM.Guid = guid
		context.JSON(http.StatusBadRequest, KolsVM)
		return
	}

	// Fetch KOLs from the logic layer
	kols, err := Logic.GetKolLogic(pageIndex, pageSize)
	if err != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = err.Error()
		KolsVM.PageIndex = int64(pageIndex)
		KolsVM.PageSize = int64(pageSize)
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// Prepare the successful response
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = int64(pageIndex)
	KolsVM.PageSize = int64(pageSize)
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}
