# Tabungan API Go

## Docker
Clone this repository and run:
```
docker-compose up
```

Following endpoints:

| Method | Route                | Body                                               |
| ------ | -------------------- |----------------------------------------------------|
| POST   | /daftar              | `{"nama": "Dian", "nik": "1234", "no_hp": "0821" }`|
| POST   | /tabung              | `{"no_rekening": "00009", "nominal": 50000`        |
| POST   | /tarik               | `{"no_rekening": "00009", "nominal": 50000`        |
| GET    | /saldo/:no_rekening  |                                                    |
| GET    | /mutasi/:no_rekening |                                                    |

