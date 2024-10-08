package config

import (
	"fmt"

	"example.com/project-sa-g03/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.Locks{},
		&entity.Gender{},
		&entity.Users{},
		&entity.Categories{},
		&entity.Shop{},
		&entity.Reserve{},
		&entity.ReserveDetails{},
		&entity.Payment{},
	)

	// สร้างข้อมูลล็อค
	locks := []entity.Locks{
		{Id: "A00", Status: "ว่าง", Price: 200, Size: "2x2"},
		{Id: "A01", Status: "ว่าง", Price: 200, Size: "2x2"},
		{Id: "A02", Status: "ไม่ว่าง", Price: 200, Size: "2x2"},
		{Id: "A03", Status: "ว่าง", Price: 200, Size: "2x2"},
		{Id: "A04", Status: "ว่าง", Price: 200, Size: "2x2"},
		{Id: "A05", Status: "ไม่ว่าง", Price: 200, Size: "2x2"},
		{Id: "A06", Status: "ว่าง", Price: 200, Size: "2x2"},
		{Id: "A07", Status: "ว่าง", Price: 200, Size: "2x2"},
		{Id: "A08", Status: "ว่าง", Price: 200, Size: "2x2"},
		{Id: "A09", Status: "ไม่ว่าง", Price: 200, Size: "2x2"},
		{Id: "B00", Status: "ไม่ว่าง", Price: 250, Size: "2x2"},
		{Id: "B01", Status: "ว่าง", Price: 250, Size: "2x2"},
		{Id: "B02", Status: "ว่าง", Price: 250, Size: "2x2"},
		{Id: "B03", Status: "ว่าง", Price: 250, Size: "2x2"},
		{Id: "B04", Status: "ว่าง", Price: 250, Size: "2x2"},
		{Id: "B05", Status: "ว่าง", Price: 250, Size: "2x2"},
		{Id: "B06", Status: "ไม่ว่าง", Price: 250, Size: "2x2"},
		{Id: "B07", Status: "ว่าง", Price: 250, Size: "2x2"},
		{Id: "B08", Status: "ว่าง", Price: 250, Size: "2x2"},
		{Id: "B09", Status: "ว่าง", Price: 250, Size: "2x2"},
		{Id: "C00", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "C01", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "C02", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "C03", Status: "ไม่ว่าง", Price: 300, Size: "2x2"},
		{Id: "C04", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "C05", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "C06", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "C07", Status: "ไม่ว่าง", Price: 300, Size: "2x2"},
		{Id: "C08", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "C09", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "D00", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "D01", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "D02", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "D03", Status: "ไม่ว่าง", Price: 300, Size: "2x2"},
		{Id: "D04", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "D05", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "D06", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "D07", Status: "ไม่ว่าง", Price: 300, Size: "2x2"},
		{Id: "D08", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "D09", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "E00", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "E01", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "E02", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "E03", Status: "ไม่ว่าง", Price: 300, Size: "2x2"},
		{Id: "E04", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "E05", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "E06", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "E07", Status: "ไม่ว่าง", Price: 300, Size: "2x2"},
		{Id: "E08", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "E09", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "F00", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "F01", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "F02", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "F03", Status: "ไม่ว่าง", Price: 300, Size: "2x2"},
		{Id: "F04", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "F05", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "F06", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "F07", Status: "ไม่ว่าง", Price: 300, Size: "2x2"},
		{Id: "F08", Status: "ว่าง", Price: 300, Size: "2x2"},
		{Id: "F09", Status: "ว่าง", Price: 300, Size: "2x2"},

	}

	for _, lock := range locks {
		db.FirstOrCreate(&lock, entity.Locks{Id: lock.Id}) // ตรวจสอบเงื่อนไขก่อนสร้างข้อมูล
	}

	// สร้างข้อมูลเพศ
	genders := []entity.Gender{
		{GenderName: "Male"},
		{GenderName: "Female"},
		{GenderName: "Ohter.."},
		{GenderName: "Not specified"},
	}

	for _, gender := range genders {
		db.FirstOrCreate(&gender, entity.Gender{GenderName: gender.GenderName})
	}

	// สร้างข้อมูลผู้ใช้
	hashedPassword, _ := HashPassword("4")

	users := []entity.Users{
		{
			FirstName: "Four",
			LastName:  "Nuay",
			Email:     "Nuay@gmail.com",
			Password:  hashedPassword,
			Tel:       "000000000",
			Profile:   "https://th.bing.com/th/id/OIP.arl_DZPqNL6ZTs9u_OAdYwAAAA?rs=1&pid=ImgDetMain",
			GenderID:  1,
		},
		{
			FirstName: "โฟร์",
			LastName:  "นวย",
			Email:     "Nuay1@gmail.com",
			Password:  hashedPassword,
			Tel:       "000000000",
			Profile:   "https://th.bing.com/th/id/OIP.arl_DZPqNL6ZTs9u_OAdYwAAAA?rs=1&pid=ImgDetMain",
			GenderID:  2,
		},
	}

	for _, user := range users {
		db.FirstOrCreate(&user, entity.Users{Email: user.Email})
	}

	// สร้างข้อมูลหมวดหมู่
	categories := []entity.Categories{
		{CategoryName: "นวย1"},
		{CategoryName: "นวย2"},
		{CategoryName: "นวย3"},
		{CategoryName: "นวย4"},
	}

	for _, category := range categories {
		db.FirstOrCreate(&category, entity.Categories{CategoryName: category.CategoryName})
	}

	// สร้างข้อมูลร้านค้า
	shops := []entity.Shop{
		{
			NationalID: "123456789", CategoryID: 1, ShopName: "nuay",
			Description: "nuayqqqq", ShopImg: "https://th.bing.com/th/id/OIP.arl_DZPqNL6ZTs9u_OAdYwAAAA?rs=1&pid=ImgDetMain", UserID: 1,
		},
		{
			NationalID: "12345678922", CategoryID: 1, ShopName: "nuayeiei",
			Description: "nuayqqqq", ShopImg: "https://th.bing.com/th/id/OIP.arl_DZPqNL6ZTs9u_OAdYwAAAA?rs=1&pid=ImgDetMain", UserID: 2,
		},
	}

	for _, shop := range shops {
		db.FirstOrCreate(&shop, entity.Shop{NationalID: shop.NationalID})
	}
}

