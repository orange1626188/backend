package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEmployee - ฟังก์ชันสำหรับดึงข้อมูลพนักงาน
func GetEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Employee",
	})
}

// GetEmployee - ฟังก์ชันสำหรับดึงข้อมูลพนักงาน
func GetEmployeebyID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

// PostEmployee - ฟังก์ชันสำหรับเพิ่มข้อมูลพนักงาน
func PostEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Post Employee",
	})
}

// PutEmployee - ฟังก์ชันสำหรับแก้ไขข้อมูลพนักงาน
func PutEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Put Employee",
	})
}

// DeleteEmployee - ฟังก์ชันสำหรับลบข้อมูลพนักงาน
func DeleteEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Employee",
	})
}
func GetEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}
func PostEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}
func PutEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}
func DELETEEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}