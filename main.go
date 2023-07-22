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


var dataBlog = []Data{} //variable kosong


func main() {
	e := echo.New()

	e.Static("/Aset","Aset")

	e.GET("/Home",Home)
	e.GET("/MyProject",MyProject)
	e.GET("/AddProject",Project)
	e.GET("/DetailProject/:id",DetailProject)

	e.POST("/submitProject",AdProject)
	e.POST("/submitEdit",SubmitEEDIT)
	e.GET("/EditProject/:id",EditProject)
	e.POST("/Delete",deleteBlog)


	e.Logger.Fatal(e.Start("localhost:5000"))
}
func SubmitEEDIT(c echo.Context) error {
    // Parse form data dari permintaan
    err := c.Request().ParseForm()
    if err != nil {
        return c.String(http.StatusInternalServerError, "Gagal memproses data form")
    }

    // Dapatkan id dari parameter URL
    nama := c.FormValue("Namasbl")

    // Cari proyek dalam array dataBlog dengan nama yang cocok
    var projectToUpdate *Data
    for i, project := range dataBlog {
        if project.Nama == nama {
            projectToUpdate = &dataBlog[i]
            break
        }
    }

    // Jika proyek dengan nama yang diberikan tidak ditemukan, kembalikan pesan error
    if projectToUpdate == nil {
        return c.String(http.StatusNotFound, "Proyek tidak ditemukan")
    }

    // Update data proyek dengan nilai-nilai baru dari form
    projectToUpdate.Nama = c.FormValue("Nama")
    projectToUpdate.PostDate = c.FormValue("Mulai")
    projectToUpdate.EndDate = c.FormValue("Akhir")
    projectToUpdate.Durasi = durasiTanggal(projectToUpdate.PostDate ,projectToUpdate.EndDate)
    projectToUpdate.Deskripsi = c.FormValue("Deskripsi")
    projectToUpdate.ReactJs = c.FormValue("ReactJs")
    projectToUpdate.NodeJs = c.FormValue("NodeJs")
    projectToUpdate.JavaScript = c.FormValue("JavaScript")
    projectToUpdate.Golang = c.FormValue("Golang")

    // Redirect pengguna ke halaman detail proyek setelah berhasil mengedit
    return c.Redirect(http.StatusSeeOther, "/Home")
}

 
func Home(c echo.Context) error {
	tmpl,err := template.ParseFiles("Views/index.html")
if err != nil{
	return c.JSON(http.StatusInternalServerError,err.Error())
}
	data := map[string]interface{}{
		"Projek" : dataBlog,
		
	}
	fmt.Println("dari home ",dataBlog)
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
		Nama: Nama, // STrruct , variabel dari form
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
	
	
fmt.Println("dari form",dataBlog)

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

	for index, item := range dataBlog{
		if index == stc {
			detail = Data{
				Nama: item.Nama,
				PostDate: item.PostDate,
				EndDate: item.EndDate,
				Deskripsi: item.Deskripsi,
				ReactJs: item.ReactJs,
				NodeJs: item.NodeJs,
				JavaScript: item.JavaScript,
				Golang: item.Golang,

			}
		}
	}

	data := map [string]interface{}{
		
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
	id := c.FormValue("index")
	idToInt, _ := strconv.Atoi(id)

	dataBlog = append(dataBlog[:idToInt], dataBlog[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/Home")
}