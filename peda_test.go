package peda

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateGetData(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	datagedung := GetAllUser(mconn, "user")
	fmt.Println(datagedung)
}

// }
// func TestGCFCreateHandler(t *testing.T) {
// 	// Simulate input parameters
// 	MONGOCONNSTRINGENV := "mongodb://raulgantengbanget:0nGCVlPPoCsXNhqG@ac-oilbpwk-shard-00-00.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-01.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-02.9ofhjs3.mongodb.net:27017/test?replicaSet=atlas-13x7kp-shard-0&ssl=true&authSource=admin"
// 	dbname := "petapedia"
// 	collectionname := "user"

// 	// Create a test User
// 	datauser := User{
// 		Username: "testuser",
// 		Password: "testpassword",
// 		Role:     "user",
// 	}

// 	// Call the handler function
// 	result := GCFCreateHandler(MONGOCONNSTRINGENV, dbname, collectionname, datauser)
// 	fmt.Println(result)
// 	// You can add assertions here to validate the result, or check the database for the created user.
// }

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "unistore"
	userdata.Password = "unistore"
	userdata.Role = "admin"
	mconn := SetConnection("MONGOCONNSTRINGENV", "unistore")
	CreateNewUserRole(mconn, "users", userdata)
}

func TestDeleteUser(t *testing.T) {

	mconn := SetConnection("MONGOCONNSTRINGENV", "unistore")
	var userdata User
	userdata.Username = "unistore"
	DeleteUser(mconn, "users", userdata)
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGOCONNSTRINGENV", "unistore")
	var userdata User
	userdata.Username = "unistore"
	userdata.Password = "unistore"
	userdata.Role = "admin"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestFunciionUser(t *testing.T) {
	mconn := SetConnection("MONGOCONNSTRINGENV", "unistore")
	var userdata User
	userdata.Username = "unistore"
	userdata.Password = "unistore"
	userdata.Role = "admin"
	CreateNewUserRole(mconn, "users", userdata)
}

func TestGeneratePasswordHashh(t *testing.T) {
	password := "unistore"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestHashFunctionn(t *testing.T) {
	mconn := SetConnection("MONGOCONNSTRINGENV", "unistore")
	var userdata User
	userdata.Username = "unistore"
	userdata.Password = "unistore"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "users", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}
func TestFindUser(t *testing.T) {
	var userdata User
	userdata.Username = "unistore"
	mconn := SetConnection("MONGOCONNSTRINGENV", "unistore")
	res := FindUser(mconn, "users", userdata)
	fmt.Println(res)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "unistore"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("unistore", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOCONNSTRINGENV", "unistore")
	var userdata User
	userdata.Username = "unistore"
	userdata.Password = "unistore"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("mongodb://raulgantengbanget:0nGCVlPPoCsXNhqG@ac-oilbpwk-shard-00-00.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-01.9ofhjs3.mongodb.net:27017,ac-oilbpwk-shard-00-02.9ofhjs3.mongodb.net:27017/test?replicaSet=atlas-13x7kp-shard-0&ssl=true&authSource=admin", "petapedia")
	var userdata User
	userdata.Username = "bangsat"
	userdata.Password = "ganteng"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}
