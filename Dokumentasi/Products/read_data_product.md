# Read Data Product

Melihat data produk secara keseluruhan

**URL** : `/products`

**Method** : `GET`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content**

```json
{
    "status": 200,
    "message": "Success",
    "data": [
        {
            "product_id": 1,
            "product_name": "One Pice",
            "price": 50000,
            "stock": 20,
            "genre": {
                "genre_id": 1,
                "genre_name": "Action"
            }
        },
        {
            "product_id": 2,
            "product_name": "High School DxD",
            "price": 55000,
            "stock": 20,
            "genre": {
                "genre_id": 5,
                "genre_name": "Hentai"
            }
        },
        {
            "product_id": 3,
            "product_name": "Kimetsu no Yaiba",
            "price": 57000,
            "stock": 15,
            "genre": {
                "genre_id": 1,
                "genre_name": "Action"
            }
        },
        {
            "product_id": 4,
            "product_name": "No Game No Life",
            "price": 57000,
            "stock": 15,
            "genre": {
                "genre_id": 6,
                "genre_name": "Game"
            }
        }
    ]
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

Kembali ke halaman [awal](../README.md)
