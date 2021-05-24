package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"fmt"
	"database/sql"
	 _ "github.com/go-sql-driver/mysql"
	"time"
	"strconv"
)

type Barang struct {
	Id string
	Nama string
	CreatAt string
	UpdatAt string		
}

type Detail struct {
	Nama string
	Stock string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	
}

func main() {
	mux := httprouter.New()

	mux.GET("/", indeks)
	mux.POST("/user", user)
	mux.GET("/user/:nama", userGET)
	mux.POST("/logout", logout)
	mux.GET("/signup", signup)
	mux.POST("/signup", signupPOST)
	mux.GET("/create", createBarang)
	mux.POST("/create", createBarangPost)
	mux.GET("/read", readBarang)
	mux.GET("/update", updateBarang)
	mux.GET("/delete", deleteBarang)
	mux.POST("/update", updateBarangPOST)
	mux.POST("/delete", deleteBarangPOST)
	mux.GET("/transaksi", transaksi)
	http.ListenAndServe(":8080", mux)
}

func transaksi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	var result string	
	var total int
	var err error

	db, err := sql.Open("mysql", "root:dede1234@/larestock")
	
	if err != nil {
		panic(err.Error())
	}
	
	defer db.Close()

	res, err := db.Query("SELECT barang.nama, stocks.jumlah FROM barang where barang.id = stocks.barangid")


	if err != nil {
		
		panic(err.Error())

	}
	
	result += fmt.Sprintf("LAPORAN HARIAN <br /> TANGGAL: %v <br />", time.Now())

	for res.Next() {
		var detail Detail
		err = res.Scan(&detail.Nama, &detail.Stock)

		total, err = strconv.Atoi(detail.Stock)

		if err != nil {
			panic(err.Error())
		}
		
		result += fmt.Sprintf("<br />  %s   %s  <br />", detail.Nama, detail.Stock)
	}
	
	result += fmt.Sprintf("<br /> TOTAL: %d <br />", total) 

	tpl.ExecuteTemplate(w, "transaksi.gohtml", result)
}

func signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func signupPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	r.ParseForm()
	
	db, err := sql.Open("mysql", "root:dede1234@/larestock")
	
	if err != nil {
		panic(err.Error())
	}
	
	defer db.Close()

	id := r.FormValue("id")
	username := r.FormValue("username")
	password := r.FormValue("password")
	nama := r.FormValue("nama")
	email := r.FormValue("email")
	nohp := r.FormValue("nohp")
	fmt.Println(id, nama)

	sql := "INSERT INTO users VALUES('"+id+"','"+username+"','"+password+"', '"+nama+"', '"+email+"', '"+nohp+"' , current_timestamp, current_timestamp);"
	fmt.Println(sql)
	res, err := db.Exec(sql)

	
	if err != nil {

		panic(err.Error())
	}

	fmt.Println("SUKSES INSERT users: ", res)

	http.Redirect(w, r, "/user/"+username, http.StatusSeeOther)
	
}

func createBarang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "createBarang.gohtml", nil)
}

func createBarangPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	
	db, err := sql.Open("mysql", "root:dede1234@/larestock")
	
	if err != nil {
		panic(err.Error())
	}
	
	defer db.Close()

	id := r.FormValue("id")

	nama:= r.FormValue("nama")

	fmt.Println(id, nama)

	sql := "INSERT INTO barang VALUES('"+id+"','"+nama+"', current_timestamp, current_timestamp);"
	fmt.Println(sql)
	res, err := db.Exec(sql)

	
	if err != nil {

		panic(err.Error())
	}

	fmt.Println("SUKSES INSERT BARANG: ", res)

	http.Redirect(w, r, "/create", http.StatusSeeOther)
}

func readBarang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	var result string	

	db, err := sql.Open("mysql", "root:dede1234@/larestock")
	
	if err != nil {
		panic(err.Error())
	}
	
	defer db.Close()

	res, err := db.Query("SELECT * FROM barang")


	if err != nil {
		
		panic(err.Error())

	}
	
	for res.Next() {
		var barang Barang
		err := res.Scan(&barang.Id, &barang.Nama, &barang.CreatAt, &barang.UpdatAt)
		if err != nil {
			panic(err.Error())
		}
		
		result += fmt.Sprintf("<br /> %s   %s   %s   %s <br />", barang.Id, barang.Nama, barang.CreatAt, barang.UpdatAt)
	}

	tpl.ExecuteTemplate(w, "readBarang.gohtml", result)
}

func updateBarang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "update.gohtml", nil) 
}

func deleteBarang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "delete.gohtml", nil) 
}

func updateBarangPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	
	db, err := sql.Open("mysql", "root:dede1234@/larestock")
	
	if err != nil {
		panic(err.Error())
	}
	
	defer db.Close()

	id := r.FormValue("id")

	nama:= r.FormValue("nama")

	fmt.Println(id, nama)

	sql := "UPDATE barang set nama='"+nama+"', updatedat=current_timestamp where id="+id+";"
	fmt.Println(sql)
	res, err := db.Exec(sql)

	
	if err != nil {

		panic(err.Error())
	}

	fmt.Println("SUKSES update BARANG: ", res)

	http.Redirect(w, r, "/update", http.StatusSeeOther)
}

func deleteBarangPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	
	db, err := sql.Open("mysql", "root:dede1234@/larestock")
	
	if err != nil {
		panic(err.Error())
	}
	
	defer db.Close()

	id := r.FormValue("id")

	fmt.Println(id)

	sql := "DELETE FROM barang WHERE id="+id+";"
	fmt.Println(sql)
	res, err := db.Exec(sql)

	
	if err != nil {

		panic(err.Error())
	}

	fmt.Println("SUKSES DELETE BARANG: ", res)

	http.Redirect(w, r, "/delete", http.StatusSeeOther)
}

func indeks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "indeks.gohtml", nil)
}

func user(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	nama := r.FormValue("username")
	tpl.ExecuteTemplate(w, "user.gohtml", nama)
}

func userGET(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//http.Error(w, "FORBIDDEN", 403)
	tpl.ExecuteTemplate(w, "user.gohtml", p.ByName("nama"))
}

func logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	
}
