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
https://asia-southeast2-unistore-403306.cloudfunctions.net/MembuatUser
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

Link API-nya

```
https://asia-southeast2-unistore-403306.cloudfunctions.net/MembuatTokenUser
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

## Delete Akun

Link API-nya

```
https://asia-southeast2-unistore-403306.cloudfunctions.net/HapusUser

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
