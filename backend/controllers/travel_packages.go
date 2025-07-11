package controllers

import (
	"net/http"
	"strconv"

	"destinara-backend/database"
	"destinara-backend/models"

	"github.com/gin-gonic/gin"
)

type CreatePackageRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Destination string  `json:"destination" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Duration    int     `json:"duration" binding:"required,gt=0"`
	ImageURL    string  `json:"image_url"`
}

type UpdatePackageRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Destination string  `json:"destination"`
	Price       float64 `json:"price"`
	Duration    int     `json:"duration"`
	ImageURL    string  `json:"image_url"`
	Available   *bool   `json:"available"`
}

func GetAllPackages(c *gin.Context) {
	var packages []models.TravelPackage
	
	query := database.DB.Model(&models.TravelPackage{})
	
	// Filtrar por disponibilidade se especificado
	if available := c.Query("available"); available != "" {
		if available == "true" {
			query = query.Where("available = ?", true)
		} else if available == "false" {
			query = query.Where("available = ?", false)
		}
	}
	
	// Filtrar por destino se especificado
	if destination := c.Query("destination"); destination != "" {
		query = query.Where("destination ILIKE ?", "%"+destination+"%")
	}
	
	if err := query.Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pacotes"})
		return
	}

	c.JSON(http.StatusOK, packages)
}

func GetPackageByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var package models.TravelPackage
	if err := database.DB.First(&package, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pacote não encontrado"})
		return
	}

	c.JSON(http.StatusOK, package)
}

func CreatePackage(c *gin.Context) {
	var req CreatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	package := models.TravelPackage{
		Name:        req.Name,
		Description: req.Description,
		Destination: req.Destination,
		Price:       req.Price,
		Duration:    req.Duration,
		ImageURL:    req.ImageURL,
		Available:   true,
	}

	if err := database.DB.Create(&package).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pacote"})
		return
	}

	c.JSON(http.StatusCreated, package)
}

func UpdatePackage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdatePackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	var package models.TravelPackage
	if err := database.DB.First(&package, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pacote não encontrado"})
		return
	}

	// Atualizar apenas os campos fornecidos
	if req.Name != "" {
		package.Name = req.Name
	}
	if req.Description != "" {
		package.Description = req.Description
	}
	if req.Destination != "" {
		package.Destination = req.Destination
	}
	if req.Price > 0 {
		package.Price = req.Price
	}
	if req.Duration > 0 {
		package.Duration = req.Duration
	}
	if req.ImageURL != "" {
		package.ImageURL = req.ImageURL
	}
	if req.Available != nil {
		package.Available = *req.Available
	}

	if err := database.DB.Save(&package).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar pacote"})
		return
	}

	c.JSON(http.StatusOK, package)
}

func DeletePackage(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var package models.TravelPackage
	if err := database.DB.First(&package, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pacote não encontrado"})
		return
	}

	if err := database.DB.Delete(&package).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar pacote"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pacote deletado com sucesso"})
} 