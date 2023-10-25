package server

import (
	"backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonStatusOkResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// errors

func jsonInternalServerErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func jsonBadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func jsonUnauthorizedResponse(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
}

func jsonBadRoleResponse(c *gin.Context, role string) {
	c.JSON(http.StatusForbidden, gin.H{"error": role})
}

// user
func jsonUserInfoOkResponse(c *gin.Context, user *models.User) {
	c.JSON(http.StatusOK, gin.H{"login": user.Login, "name": user.Name, "surname": user.Surname,
		"contact": user.Contact, "registration_date": user.RegisteredAt, "role": user.Role})
}

// // user

func jsonUserCreateResponse(c *gin.Context, token string) {
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func jsonUserLoginOkResponse(c *gin.Context, token string) {
	c.JSON(http.StatusOK, gin.H{"Token": token})
}

// farm
func jsonUserFarmsOkResponse(c *gin.Context, farms []models.Farm) {
	c.JSON(http.StatusOK, gin.H{"farms": farms})
}

func jsonFarmCreateResponse(c *gin.Context, farm models.Farm) {
	c.JSON(http.StatusCreated, gin.H{"name": farm.Name, "description": farm.Description,
		"address": farm.Address, "userId": farm.UserId})
}

// func jsonUserInfoOkResponse(c *gin.Context, user *models.User) {
// 	c.JSON(http.StatusOK, gin.H{"UserId": user.UserId, "Login": user.Login})
// }

// // doctor

// func jsonDoctorCreateResponse(c *gin.Context, doctor *models.Doctor, token string) {
// 	c.JSON(http.StatusCreated, gin.H{"DoctorId": doctor.DoctorId, "Login": doctor.Login,
// 		"StartTime": doctor.StartTime, "EndTime": doctor.EndTime, "Token": token})
// }

// func jsonDoctorLoginOkResponse(c *gin.Context, doctor *models.Doctor, token string) {
// 	c.JSON(http.StatusOK, gin.H{"DoctorId": doctor.DoctorId, "Login": doctor.Login, "Token": token})
// }

// func jsonDoctorInfoOkResponse(c *gin.Context, doctor *models.Doctor) {
// 	c.JSON(http.StatusOK, gin.H{"DoctorId": doctor.DoctorId, "Login": doctor.Login,
// 		"StartTime": doctor.StartTime, "EndTime": doctor.EndTime})
// }

// func jsonDoctorsOkResponse(c *gin.Context, doctors []models.Doctor) {
// 	c.JSON(http.StatusOK, gin.H{"doctors": doctors})
// }

// // records

// func jsonRecordsOkResponse(c *gin.Context, records []models.Record) {
// 	c.JSON(http.StatusOK, gin.H{"records": records})
// }

// func jsonRecordCreatedResponse(c *gin.Context, record *models.Record) {
// 	c.JSON(http.StatusCreated, gin.H{"record": record})
// }

// // pets

// func jsonPetsOkResponse(c *gin.Context, pets []models.Pet) {
// 	c.JSON(http.StatusOK, gin.H{"pets": pets})
// }

// func jsonPetCreatedResponse(c *gin.Context, pet *models.Pet) {
// 	c.JSON(http.StatusCreated, gin.H{"pet": pet})
// }

// func jsonPetOkResponse(c *gin.Context, pet *models.Pet) {
// 	c.JSON(http.StatusOK, gin.H{"pet": pet})
// }