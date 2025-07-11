package controllers

import (
	"net/http"
	"strconv"
	"time"

	"destinara-backend/database"
	"destinara-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateBookingRequest struct {
	TravelPackageID uint      `json:"travel_package_id" binding:"required"`
	StartDate       time.Time `json:"start_date" binding:"required"`
	EndDate         time.Time `json:"end_date" binding:"required"`
}

type UpdateBookingRequest struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
}

func GetUserBookings(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var bookings []models.Booking
	if err := database.DB.Preload("TravelPackage").Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar reservas"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func GetBookingByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var booking models.Booking
	if err := database.DB.Preload("TravelPackage").Preload("User").Where("id = ? AND user_id = ?", id, userID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reserva não encontrada"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

func CreateBooking(c *gin.Context) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Verificar se o pacote existe e está disponível
	var travelPackage models.TravelPackage
	if err := database.DB.First(&travelPackage, req.TravelPackageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pacote de viagem não encontrado"})
		return
	}

	if !travelPackage.Available {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pacote não está disponível"})
		return
	}

	// Verificar se as datas são válidas
	if req.StartDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data de início deve ser futura"})
		return
	}

	if req.EndDate.Before(req.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data de fim deve ser posterior à data de início"})
		return
	}

	// Calcular preço total baseado na duração
	duration := req.EndDate.Sub(req.StartDate).Hours() / 24
	totalPrice := travelPackage.Price * float64(int(duration))

	booking := models.Booking{
		UserID:          userID.(uint),
		TravelPackageID: req.TravelPackageID,
		StartDate:       req.StartDate,
		EndDate:         req.EndDate,
		TotalPrice:      totalPrice,
		Status:          "pending",
	}

	if err := database.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar reserva"})
		return
	}

	// Carregar relacionamentos para a resposta
	database.DB.Preload("TravelPackage").Preload("User").First(&booking, booking.ID)

	c.JSON(http.StatusCreated, booking)
}

func UpdateBooking(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var req UpdateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	var booking models.Booking
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reserva não encontrada"})
		return
	}

	// Atualizar apenas os campos fornecidos
	if !req.StartDate.IsZero() {
		booking.StartDate = req.StartDate
	}
	if !req.EndDate.IsZero() {
		booking.EndDate = req.EndDate
	}
	if req.Status != "" {
		booking.Status = req.Status
	}

	// Recalcular preço se as datas foram alteradas
	if !req.StartDate.IsZero() || !req.EndDate.IsZero() {
		var travelPackage models.TravelPackage
		database.DB.First(&travelPackage, booking.TravelPackageID)
		
		duration := booking.EndDate.Sub(booking.StartDate).Hours() / 24
		booking.TotalPrice = travelPackage.Price * float64(int(duration))
	}

	if err := database.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar reserva"})
		return
	}

	// Carregar relacionamentos para a resposta
	database.DB.Preload("TravelPackage").Preload("User").First(&booking, booking.ID)

	c.JSON(http.StatusOK, booking)
}

func CancelBooking(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var booking models.Booking
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reserva não encontrada"})
		return
	}

	booking.Status = "cancelled"

	if err := database.DB.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cancelar reserva"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reserva cancelada com sucesso"})
} 