package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/amosehiguese/restaurant-api/store"
	"github.com/amosehiguese/restaurant-api/types"
	"github.com/google/uuid"
)

func GetAllInvoices(w http.ResponseWriter, r *http.Request){
	q := store.GetQuery()
	invoices, err := q.GetAllInvoices(ctx)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": invoices,
	})
}

func CreateInvoice(w http.ResponseWriter, r *http.Request){
	var payload types.InvoicePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	v := types.NewValidator()

	if err := v.Struct(payload); err != nil {
		json.NewEncoder(w).Encode(resp{
			"error": true,
			"code": http.StatusBadRequest,
			"msg":types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()

	invoice := models.CreateInvoiceParams{
		OrderID: payload.OrderID,
		InvoiceDate: time.Now(),
		TotalAmount: payload.TotalAmount,
		Tax: payload.Tax,
		Discount: payload.Discount,
		GrandTotal: payload.GrandTotal,
	}

	result, err := q.CreateInvoice(ctx, invoice)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result.ID,
	})
}
func RetrieveInvoice(w http.ResponseWriter, r *http.Request){
	id := getField(r, "id")
	invoiceID, err := uuid.Parse(id)

	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	q := store.GetQuery()
	invoice, err := q.RetrieveInvoice(ctx, invoiceID)
	if err != nil {
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusNotFound,
			"msg": fmt.Sprintf("invoice with this ID %s not found", invoiceID),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": invoice,
	})
}
func UpdateInvoice(w http.ResponseWriter, r *http.Request){
	id := getField(r, "id")
	invoiceID, err := uuid.Parse(id)

	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}


	var payload types.InvoicePayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	v := types.NewValidator()

	if err := v.Struct(payload); err != nil{
		json.NewEncoder(w).Encode(resp{
			"error": true,
			"code": http.StatusBadRequest,
			"msg":types.ValidatorErrors(err),
		})
		return
	}

	invoice := models.UpdateInvoiceParams{
		OrderID: payload.OrderID,
		TotalAmount: payload.TotalAmount,
		Tax: payload.Tax,
		Discount: payload.Discount,
		GrandTotal: payload.GrandTotal,
	}

	q := store.GetQuery()
	err = q.UpdateInvoice(ctx, invoice)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	str := fmt.Sprintf("Successfully update invoice with id %s", invoiceID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
}
