# Regist

Registrasi ini digunakan untuk membuat user baru

**URL** : `/users`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "name": "[password in plain text]"
    "pass": "[password in plain text]"
    "email": "[password in plain text]",
    "alamat": "[alamat in plain text]"
}
```

**Data example**

```json
{
    "name": "Thoha"
    "pass": "pass123"
    "email": "thoha@email",
    "alamat": "Konoha"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
   "status": 200,
    "message": "Success",
    "data": {
        "lastInsertedId": 22
    }
}
```

## Error Response

**Condition** : Jika ada yang dimasukan ada yang kosong

**Code** : `400 Bad Request`

**Content** :

```json
{
    "status": 400,
    "message": "Data can't be empty",
    "data": null
}
```

Kembali ke halaman [awal](../README.md)
