package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/xuri/excelize/v2"
)

var db *sqlx.DB

func initDB() {
	var err error
	connStr := "user=postgres dbname=gin_test sslmode=disable password=1234"
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to the database!")
}

func main2() {
	
	initDB()
	
	r := gin.Default()

	r.POST("/items", createItem)
	r.GET("/items", getItems)
	r.GET("/items/:id", getItemByID)
	r.PUT("/items", updateItem)
	r.DELETE("/items/:id", deleteItem)
	r.GET("/export", exportToExcel) // excell export etmek ucin
	r.POST("/import", importFromExcel) // excell import etmek ucin

	r.Run()
}

type Item struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Price int    `db:"price" json:"price"`
}

func createItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	_, err := db.Exec(`INSERT INTO items (name, price) VALUES ($1, $2)`, item.Name, item.Price)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, "Post success")
}

func getItems(c *gin.Context) {
	items := []Item{}
	err := db.Select(&items, "SELECT * FROM items")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, items)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")
	var item Item
	err := db.Get(&item, "SELECT * FROM items WHERE id=$1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, item)
}

func updateItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	_, err :=  db.Exec(`UPDATE items SET name=$1, price=$2 WHERE id=$3`, item.Name, item.Price, item.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"message": "edit succ!"})
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM items WHERE id=$1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"message": "delete succ!"})
}

func exportToExcel(c *gin.Context) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Price")

	items := []Item{}
	err := db.Select(&items, "SELECT * FROM items order by id")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	for i, item := range items {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), item.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), item.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), item.Price)
	}

	if err := f.SaveAs("items.xlsx"); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "Excel file createdd :)"})
}

func importFromExcel(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    filePath := "./" + file.Filename

    if err := exportToExcelFun(filePath); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "Excel file imported successfully"})
}

func exportToExcelFun(filePath string) error {
    f, err := excelize.OpenFile(filePath)
    if err != nil {
        return err
    }

    rows, err := f.GetRows("Sheet1")

    rows = rows[1:] //title bolany un goyup gidyas

    for _, row := range rows {
        id, err := strconv.Atoi(row[0]) 
        if err != nil {
            return err
        }
        name := row[1]  
        price, err := strconv.Atoi(row[2]) 
        if err != nil {
            return err
        }

        _, err = db.Exec(`INSERT INTO items2 (id, name, price) VALUES ($1, $2, $3)`, id, name, price)
        if err != nil {
            return err
        }
    }

    return nil
}

