package handler

import (
	"C:\\Users\\Rodion\\stepic\\inernal\\models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetHandler(c echo.Context) error {
	var messages []models.Message
	if err := DB.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "Error",
			Message: "Could not find the messages",
		})
	}
	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message models.Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add to message",
		})
	}

	if err := DB.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not create the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message successfully added",
	})
}

func PathcHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}
	var updatedMessage Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Invalid input",
		})
	}

	if err := DB.Model(&Message{}).Where("id=?", id).Update("text", updatedMessage.Text).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Succes",
		Message: "Message was updated",
	})

}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	if err := DB.Delete(&Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not delete the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Succes",
		Message: "Message was delete",
	})

}
