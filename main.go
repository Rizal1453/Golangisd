package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Data struct{
	Nama string
	PostDate string
	EndDate string
	Durasi string
	Deskripsi string
	ReactJs string
	NodeJs string
	JavaScript string
	Golang string
}

var dataBlog = []Data{}

func main() {
	e := echo.New()

	e.Static("/Aset","Aset")

	e.GET("/Home",Home)
	e.GET("/MyProject",MyProject)
	e.GET("/AddProject",Project)
	e.GET("/DetailProject/:id",DetailProject)

	e.POST("/AdProject",AdProject)
	e.GET("/EditProject/:id",EditProject)
	e.GET("/Detail/:id",EditProject)
	e.GET("/Delete/:id",deleteBlog)


	e.Logger.Fatal(e.Start("localhost:5000"))
}

func Home(c echo.Context) error {
	tmpl,err := template.ParseFiles("Views/index.html")
if err != nil{
	return c.JSON(http.StatusInternalServerError,err.Error())
}
	data := map[string]interface{}{
		"Projek" : dataBlog,
		
	}
	fmt.Println(dataBlog)
return tmpl.Execute(c.Response(),data)
}

func MyProject(c echo.Context) error {
	tmpl,err := template.ParseFiles("Views/myproject.html")
if err != nil{
	return c.JSON(http.StatusInternalServerError,err.Error())
}
	
return tmpl.Execute(c.Response(),nil)
}

func Project(c echo.Context) error {
	tmpl,err := template.ParseFiles("Views/AddProject.html")
if err != nil{
	return c.JSON(http.StatusInternalServerError,err.Error())
}
	
return tmpl.Execute(c.Response(),nil)
}

func DetailProject(c echo.Context) error {

	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/DetailProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage": err.Error()})
	}
	stc, _ := strconv.Atoi(id)
	details := Data{}
	for index, data := range dataBlog {
		if index == stc {
			details = Data{
				Nama:       data.Nama,
				PostDate:   data.PostDate,
				EndDate:    data.EndDate,
				Deskripsi:  data.Deskripsi,
				NodeJs:     data.NodeJs,
				ReactJs: data.ReactJs,
				JavaScript: data.JavaScript,
				Golang: data.Golang,
			
			}
		}
	}

	data := map[string]interface{}{
		"Id":     id,
		"ProjekDetail": details,
	}
	return tmpl.Execute(c.Response(), data)
}

func AdProject(c echo.Context)error{
	Nama := c.FormValue("Nama")
	Deskripsi := c.FormValue("Deskripsi")
	Start := c.FormValue("Mulai")
	End := c.FormValue("Akhir")
	React := c.FormValue("React")
	NodeJs := c.FormValue("NodeJs")
	JavaScript := c.FormValue("JavaScript")
	Goolang := c.FormValue("Golang")
	Durasi := durasiTanggal(Start ,End)

	BlogData := Data{
		Nama: Nama,
		PostDate: Start,
		EndDate: End,
		Deskripsi: Deskripsi,
		ReactJs: React,
		NodeJs: NodeJs,
		JavaScript: JavaScript,
		Golang: Goolang,
		Durasi: Durasi,
		

	}


	dataBlog = append(dataBlog, BlogData)
	
	


	return c.Redirect(http.StatusMovedPermanently,"/Home")
	
}


func EditProject(c echo.Context)error  {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("Views/EditProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"massage" : err.Error(),
		})
	}

	stc,_ := strconv.Atoi(id)
	detail := Data{}
	for index, data := range dataBlog{
		if index == stc {
			detail = Data{
				Nama: data.Nama,
				PostDate: data.PostDate,
				EndDate: data.EndDate,
				Deskripsi: data.Deskripsi,
				ReactJs: data.ReactJs,
				NodeJs: data.NodeJs,
				JavaScript: data.JavaScript,
				Golang: data.Golang,

			}
		}
	}

	data := map [string]interface{}{
		"id" : id,
		"ProjekDetail" : detail,
	}

return tmpl.Execute(c.Response(),data)

	
}

func durasiTanggal(start string, endDate string) string {
	awalMulai, _ := time.Parse("2006-01-02", start)

	akhirMulai, _ := time.Parse("2006-01-02", endDate)
	// untuk mengurangi dua waktu dan menghasilkan selisih waktu di antara keduanya.
	durasi := akhirMulai.Sub(awalMulai)

	years := durasi.Hours() / 24 / 365
	yearsInt := int(years)

	months := (durasi.Hours() / 24) / 30
	monthsInt := int(months)

	days := durasi.Hours() / 24
	daysInt := int(days)

	durasii := fmt.Sprintf("Durasi: %d tahun, %d bulan, %d hari", yearsInt, monthsInt, daysInt)

	return durasii

}

func deleteBlog(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	dataBlog = append(dataBlog[:idToInt], dataBlog[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/Home")
}