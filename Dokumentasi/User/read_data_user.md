# Get All User

Dugunakan untuk mengambil semua data user yang ada didalam database

**URL** : `/users`

**Method** : `GET`

**Auth required** : YES

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "status": 200,
    "message": "Success",
    "data": [
        {
            "user_id": 2,
            "user_name": "Thoha",
            "user_pass": "$2a$10$UhP.pEL7Lp5wrVo2mNmau.wwoAlHh2gFO8bGTPPyRjF5Cs4l3k8Fm",
            "user_email": "thoha@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 3,
            "user_name": "Mellisa",
            "user_pass": "$2a$10$ZVVZMzen7NP9aYwkt4Dp5.a.KYoUO6IRl10ODRfz1XQ3WpiK65qwO",
            "user_email": "mellisa@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 4,
            "user_name": "Himawari",
            "user_pass": "$2a$10$vvywvoBcj.TO9oq4njrpKuC9NpRrAls0F9ig7pIfhAbJ87Ceum8f.",
            "user_email": "himawari@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 5,
            "user_name": "Sakura",
            "user_pass": "$2a$10$YXPS3zBVb27/dMqmnnR3ZO6qDx.0D0X8KiMIC5.qXRtIcZa9t/SYq",
            "user_email": "sakura@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 8,
            "user_name": "Temarin",
            "user_pass": "$2a$10$LF53FqUSyVfv8olWfFK3GesKUpYxGEQcCfHfSJ4Hi32DjYtUaqJ6a",
            "user_email": "temarin@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 9,
            "user_name": "Tenten",
            "user_pass": "$2a$10$b96.i.mcfpujFpESX/ZGh.DgHdDk5RFSI/FKPJrHmJoJ7AdM9QVKe",
            "user_email": "tenten@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 11,
            "user_name": "Hinata",
            "user_pass": "$2a$10$Jo2SxEY89Pp1AqG5.M8MA.xIh7PgTfTULWcffC.FoyLqvTLBDYKYe",
            "user_email": "hinata@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 12,
            "user_name": "Naruto",
            "user_pass": "$2a$10$7mOEi1VoCw0JeRYzvMLKJ.bLcZ36G.acT8YkXcFrliQahVbEPdXg2",
            "user_email": "naruto@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 13,
            "user_name": "Ino",
            "user_pass": "$2a$10$kaGr0eQdk9tpFMD81zt8FuMj.Ac1vukN0Q7f.bTY4DdEBCYWMai/q",
            "user_email": "ino@email.com",
            "alamat": "konoha"
        },
        {
            "user_id": 22,
            "user_name": "ganteng",
            "user_pass": "$2a$10$JhuV7hKvWkcFRJPlpejohuDJ8uWAKrID3FA34iszuQvQ.8Iv5p4Cy",
            "user_email": "naruto@email.com",
            "alamat": "konoha"
        }
    ]
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

Kembali ke halaman [awal](../README.md)
