/*
 * Subscription Manager
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package subscriptionManager

import (
	"net/http"
	"subscriptionManager/db"
	"subscriptionManager/models"

	"github.com/gin-gonic/gin"
)

// PaymentsCurrentGet - Returns an payments current page.\"
func PaymentsCurrentGet(c *gin.Context) {
	/**
	 * check authentication
	 */
	if IsApiAuthenticated(c) > 0 {
		http.Error(c.Writer, "Failed to authenticate.", http.StatusUnauthorized)
		return
	}
	user_id := GetUserId(c) //get user_id
	var UserPlans []models.UserPlans
	db.DB.Where("user_id", user_id).Where("status", "CURRENT").Find(&UserPlans)
	if len(UserPlans) > 0 {
		plan_id := UserPlans[0].PlanId
		var Plan []models.AvailablePlans
		db.DB.Where("id", plan_id).Find(&Plan)
		Aplan := []APlan{
			APlan{
				Id:          int32(Plan[0].Id),
				Name:        Plan[0].Name,
				Description: Plan[0].Description,
				Price:       Plan[0].Price,
				Recurrence:  int32(Plan[0].Recurrence),
			},
		}
		c.JSON(http.StatusOK, gin.H{"plan": Aplan})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Status 404",
		})
		return
	}
}