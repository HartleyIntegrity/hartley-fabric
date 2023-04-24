package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"hartley-fabric/backend/blockchain"
	"hartley-fabric/backend/database"

	"github.com/gin-gonic/gin"
)

var db *database.Database
var bc *blockchain.Blockchain

func RegisterHandlers(router *gin.Engine, database *database.Database, blockchain *blockchain.Blockchain) {
	db = database
	bc = blockchain
	router.GET("/api/tenancy-agreements", getTenancyAgreements)
	router.POST("/api/tenancy-agreements", createTenancyAgreement)
	router.PUT("/api/tenancy-agreements/:id", updateTenancyAgreement)
	router.DELETE("/api/tenancy-agreements/:id", deleteTenancyAgreement)
	router.GET("/api/latest-hash", getLatestHash)
}

func getTenancyAgreements(c *gin.Context) {
	agreements := make([]*blockchain.TenancyAgreement, 0)
	for _, block := range bc.Blocks {
		agreement := &blockchain.TenancyAgreement{}
		json.Unmarshal(block.Data, agreement)
		agreements = append(agreements, agreement)
	}
	c.JSON(http.StatusOK, agreements)
}

func createTenancyAgreement(c *gin.Context) {
	var agreement blockchain.TenancyAgreement
	if err := c.BindJSON(&agreement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	agreement.ID = strconv.Itoa(len(bc.Blocks))
	data, _ := json.Marshal(agreement)
	bc.AddBlock(string(data))

	c.JSON(http.StatusCreated, agreement)
}

func updateTenancyAgreement(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not provided"})
		return
	}

	var updatedAgreement blockchain.TenancyAgreement
	if err := c.BindJSON(&updatedAgreement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updated := false
	for _, block := range bc.Blocks {
		agreement := &blockchain.TenancyAgreement{}
		json.Unmarshal(block.Data, agreement)
		if agreement.ID == id {
			agreement = &updatedAgreement
			data, _ := json.Marshal(agreement)
			block.Data = []byte(data)
			updated = true
			break
		}
	}

	if updated {
		c.JSON(http.StatusOK, updatedAgreement)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenancy agreement not found"})
	}
}

func getLatestHash(c *gin.Context) {
	latestHash := ""
	if len(bc.Blocks) > 0 {
		latestHash = string(bc.Blocks[len(bc.Blocks)-1].Hash)
	}
	c.JSON(http.StatusOK, gin.H{
		"latestHash": latestHash,
	})
}

func deleteTenancyAgreement(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not provided"})
		return
	}

	deleted := false
	newBlocks := make([]*blockchain.Block, 0)
	for _, block := range bc.Blocks {
		agreement := &blockchain.TenancyAgreement{}
		json.Unmarshal(block.Data, agreement)
		if agreement.ID == id {
			deleted = true
		} else {
			newBlocks = append(newBlocks, block)
		}
	}

	if deleted {
		bc.Blocks = newBlocks
		c.JSON(http.StatusOK, gin.H{"message": "Tenancy agreement deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tenancy agreement not found"})
	}
}
