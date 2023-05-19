# Get Data Cart

Melihat data produk yang ada didalam keranjang sesuai user

**URL** : `/carts`

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
            "cart_id": 24,
            "user_id": 2,
            "product_id": 1,
            "qty": 7
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
