package peda

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/whatsauth/watoken"
)

func ReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func MembuatUser(mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		response.Status = true
		hash, hashErr := HashPassword(datauser.Password)
		if hashErr != nil {
			response.Message = "Gagal Hash Password" + err.Error()
		}
		InsertUserdata(mconn, collname, datauser.Username, datauser.Role, hash)
		response.Message = "Berhasil Input data"
	}
	return ReturnStruct(response)
}

func MembuatTokenUser(privatekey, mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collname, datauser) {
			response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(privatekey))
			if err != nil {
				response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				response.Message = "Selamat Datang"
				response.Token = tokenstring
			}
		} else {
			response.Message = "Password Salah"
		}
	}

	return ReturnStruct(response)
}

func HapusUser(mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		DeleteUser(mconn, collname, datauser)
		response.Message = "Berhasil Delete data"
	}
	return ReturnStruct(response)
}

func LoginDenganRole(privatekey, mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collname, datauser) {
			response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(privatekey))
			if err != nil {
				response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				if CekRole(mconn, collname, datauser) {
					response.Message = "Selamat Datang Admin"
					response.Token = tokenstring
				} else {
					response.Message = "Selamat Datang User"
					response.Token = tokenstring
				}
			}
		} else {
			response.Message = "Password Salah"
		}
	}
	return ReturnStruct(response)
}

func LoginDenganRole2(privatekey, mongoenv, dbname, collname string, r *http.Request) string {
	var response Credential
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collname, datauser) {
			response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(privatekey))
			if err != nil {
				response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				if CekRoleAdmin(mconn, collname, datauser) {
					response.Message = "Selamat Datang Admin"
					response.Token = tokenstring
				}
				if CekRoleUser(mconn, collname, datauser) {
					response.Message = "Selamat Datang User"
					response.Token = tokenstring
				} else {
					response.Message = "Anda tidak memiliki role"
					response.Token = tokenstring
				}
			}
		} else {
			response.Message = "Password Salah"
		}
	}
	return ReturnStruct(response)
}
