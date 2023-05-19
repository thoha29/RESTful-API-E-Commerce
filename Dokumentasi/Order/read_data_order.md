# Get All Order

Melihat data produk yang sudah diorder

**URL** : `/orders`

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
            "order_id": 1,
            "user_id": 2,
            "product_id": 4,
            "qty": 5,
            "status": "pending"
        },
        {
            "order_id": 5,
            "user_id": 2,
            "product_id": 2,
            "qty": 3,
            "status": "berhasil"
        },
        {
            "order_id": 6,
            "user_id": 2,
            "product_id": 2,
            "qty": 3,
            "status": "pending"
        },
        {
            "order_id": 8,
            "user_id": 2,
            "product_id": 2,
            "qty": 6,
            "status": "pending"
        },
        {
            "order_id": 9,
            "user_id": 2,
            "product_id": 2,
            "qty": 3,
            "status": "pending"
        },
        {
            "order_id": 10,
            "user_id": 2,
            "product_id": 1,
            "qty": 4,
            "status": "pending"
        },
        {
            "order_id": 11,
            "user_id": 2,
            "product_id": 2,
            "qty": 5,
            "status": "berhasil"
        },
        {
            "order_id": 12,
            "user_id": 2,
            "product_id": 2,
            "qty": 5,
            "status": "antar"
        },
        {
            "order_id": 13,
            "user_id": 2,
            "product_id": 1,
            "qty": 11,
            "status": "pending"
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
