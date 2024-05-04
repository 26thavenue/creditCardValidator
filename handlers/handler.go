package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "github.com/26thavenue/creditCardValidator/models"
    "github.com/26thavenue/creditCardValidator/validators"
    "gorm.io/gorm"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var user models.User
	if err := db.Preload("CreditCards").First(&user, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var users []models.User
	db.Preload("CreditCards").Find(&users)

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if err := db.Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)

	json.NewEncoder(w).Encode(user)


}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.Delete(&user)

	w.WriteHeader(http.StatusNoContent)
}


func AddCreditCardHandler(w http.ResponseWriter, r *http.Request) {
    db := r.Context().Value("db").(*gorm.DB)

    var creditCard models.CreditCard
    if err := json.NewDecoder(r.Body).Decode(&creditCard); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    
    cardNumber := &validators.CardNumber{Number: creditCard.Number}
    creditCard.IsValid = cardNumber.ValidateCard()

    userID, err := strconv.Atoi(chi.URLParam(r, "userId"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    creditCard.UserID = uint(userID)

    if err := db.Create(&creditCard).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(creditCard)
}

