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

// GetAllInvoices returns all invoices
// @Summary List all invoices
// @Description Get all invoices stored in the database
// @Tags Invoices
// @Produce json
// @Router /invoices [get]
// @Success 200 {object} models.Invoice
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func GetAllInvoices(w http.ResponseWriter, r *http.Request){
	s, e, err := paginate(w, r)
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
	result, err := q.GetAllInvoices(ctx)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	if *e < len(result) && len(result[*s:*e]) == pageSize {
		result = result[*s:*e]
	} else if *e >= len(result) && *s < len(result) {
		result = result[*s:]
	} else if *e >= len(result) && *s >= len(result) && result != nil {
		*s = 0
		*e = pageSize
		result = result[*s:*e]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result,
	})
}

// CreateInvoice writes an invoice 
// @Summary Creates an invoice
// @Description Creates an invoice in the database
// @Tags Invoices
// @Produce json
// @Router /invoices [post]
// @Success 200 {object} models.Invoice
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
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

// RetrieveInvoice renders the invoice with the given id 
// @Summary Get invoice by id
// @Description RetrieveInvoice returns a single invoice by id
// @Tags Invoice
// @Produce json
// @Param id path string true "invoice id"
// @Router /invoices/{id} [get]
// @Success 200 {object} models.Invoice
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
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

// UpdateInvoice modifies the invoice with the given id 
// @Summary Modify invoice by id
// @Description UpdateInvoice modifies a single invoice by id
// @Tags Invoice
// @Produce json
// @Param id path string true "invoice id"
// @Router /invoices/{id} [patch]
// @Success 200 {object} models.Invoice
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
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


	var payload types.UpdateInvoicePayload
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
		TotalAmount: payload.TotalAmount,
		Tax: payload.Tax,
		Discount: payload.Discount,
		GrandTotal: payload.GrandTotal,
		ID: invoiceID,
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
	str := fmt.Sprintf("Successfully updated invoice with id %s", invoiceID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
}
