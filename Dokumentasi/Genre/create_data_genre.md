# Create Data Genre

Menambahkan genre baru

**URL** : `/genres`

**Method** : `POST`

**Auth required** : YES

**Data constraints**

```json
{
    "genre": "[name in plain text]",
}
      
```

**Data example**

```json
{
    "genre": "fantasy",
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
        "lastInsertedId": 6
    }
}
```

## Error Response
**Condition** : Jika tidak ter-autentikasi

**Code** : `401 Unauthorized`

**Content** :

```json
{
    "message": "invalid or expired jwt"
}
```

**Condition** : Jika yang dimasukan ada yang kosong

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
