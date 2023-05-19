# Create Data Product

Menambahkan produk baru

**URL** : `/products`

**Method** : `POST`

**Auth required** : YES

**Data constraints**

```json
{
    "product_name": "[name in plain text]",
    "price": "[price in int]",
    "stock": "[stock in int]",
    "genre_id": "[genre_id in int]"
}
      
```

**Data example**

```json
{
    "product_name": "One Pice",
    "price": 50000,
    "stock": 20,
    "genre_id": 1
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
**Condition** : Jika tidak terauterisasi

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
