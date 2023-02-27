package controller
 
import (
    "encoding/json"
    "fmt"
    "log"
	"os/exec"
    "net/http"
	"time"
	"io/ioutil"
    "github.com/model"
    "github.com/config"
)
 
// GetAll = Select Users API
func GetAll(w http.ResponseWriter, r *http.Request) {
    var user model.User
    var response model.Response
    var arrayUser []model.User
 
    db := config.Connect()
    defer db.Close()
 
    rows, err := db.Query("SELECT id, user_name, city, cpf, email, phone, created_at, updated_at FROM user")
 
    if err != nil {
        log.Print(err)
    }
 
    for rows.Next() {
        err = rows.Scan(&user.Id, &user.Name, &user.City, &user.Cpf, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            log.Fatal(err.Error())
        } else {
            arrayUser = append(arrayUser, user)
        }
    }
 
    response.Status = 200
    response.Message = "Success"
    response.Data = arrayUser
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}
 
// Insert = Insert user API
func Insert(w http.ResponseWriter, r *http.Request) {
    var response model.Response
	user := &model.User{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err) //TODO: erro generico
	}
	errUnmarshal := json.Unmarshal(reqBody, &user)
	if errUnmarshal != nil{
		panic(err) //TODO: erro generico
	}
    db := config.Connect()
    defer db.Close()

    // err := r.ParseMultipartForm(4096)
    // if err != nil {
    //     panic(err)
    // }
    user_name := user.Name
    city := user.City
	cpf := user.Cpf
    email := user.Email
	phone := user.Phone
	id ,_:= exec.Command("uuidgen").Output()
	created_at := time.Now()
	updated_at := created_at
    

    _, err = db.Exec("INSERT INTO user(id, user_name, city, cpf, email, phone, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", id, user_name, city, cpf, email, phone, created_at, updated_at)
 
    if err != nil {
        log.Print(err)
        return
    }
    response.Status = 200
    response.Message = "Insert data successfully"
    fmt.Print("Insert data to database")
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}

// Update = Update user API
func Update(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	// err := r.ParseMultipartForm(4096)

	// if err != nil {
	// 	panic(err)
	// }

	user := &model.User{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err) //TODO: erro generico
	}
	errUnmarshal := json.Unmarshal(reqBody, &user)
	if errUnmarshal != nil{
		panic(err) //TODO: erro generico
	}

	_, errDb := db.Exec("UPDATE user SET user_name=?, city=?, cpf=?, email=?, phone=?, updated_at=? WHERE id=?", user.Name, user.City, user.Cpf, user.Email, user.Phone, time.Now(), user.Id)
    if errDb != nil {
        log.Print(err)
        return
    }	

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Update data successfully"
	fmt.Print("Update data successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Delete = Delete user API
func Delete(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	// err := r.ParseMultipartForm(4096)

	// if err != nil {
	// 	panic(err)
	// }

	id := r.URL.Query().Get("id")

	_, err := db.Exec("DELETE FROM user WHERE id=?", id)

	if err != nil {
		log.Print(err)
		return
	}

	response.Status = 200
	response.Message = "Delete data successfully"
	fmt.Print("Delete data successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


// MainPage - just to test
func MainPage(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	response.Status = 200
	response.Message = "Project is Working"
	fmt.Print("log: Project is Working")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}