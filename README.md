# UniStore
## Anggota 1
Nama	= Ibrohim Mubarok <br />
NPM		= 1214081 <br />
Kelas	= 3C <br />
## Anggota 2
Nama	= Fitrah Ali Akbar Setiawan <br />
NPM		= 1214085 <br />
Kelas	= 3C <br />

## Registrasi Akun

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/UniStoreMembuatUser
```

Body

```
{
    "username": "input username di sini",
    "password": "input password di sini",
	"role": "input role di sini"
}
```

Response

```
{"status":true,"message":"Berhasil Input data"}
```

## Login Akun

### Membuat Token

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/UniStoreMembuatTokenUser
```

Body

```
{
    "username": "input username di sini",
    "password": "input password di sini"
}
```

Response bila berhasil

```
{"status":true,"token":"token yang didapat","message":"Selamat Datang"}
```

Response bila gagal

```
{"status":false,"message":"Password Salah"}
```

### Menyimpan Token

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/UniStoreMenyimpanTokenUser

```

Header

```
Login : masukkan token di sini
```

Response bila berhasil

```
{
    "status": true,
    "message": "data User berhasil diambil",
    "data": [
        {
            "username": "data",
            "password": "data",
            "role": "role"
        },
        {
            "username": "data",
            "password": "data",
            "role": "role"
        }
    ]
}
```

Response bila gagal

```
{"status":false,"message":"Data Username tidak ada di database"}
```

## Delete Akun

Link API-nya

```
https://asia-southeast2-gis3-401509.cloudfunctions.net/UniStoreHapusUser

```

Body

```
{
    "username": "input username di sini"
}
```

Response bila berhasil

```
{"status":false,"message":"Berhasil Delete data"}
```

Response bila gagal

```
{"status":false,"message":"error parsing application/json: EOF"}
```