# Payslip Employee

1. run query di postgres untuk mengaktifkan fitur uuid 

```CREATE EXTENSION IF NOT EXISTS "uuid-ossp";```

2. import file user_csv

3. import file employee_csv

4. login admin <b>(username : admin password:1234567890)</b>

5. login user <b>(username: sesuai pada csv, password: 1234567890)</b>

6. hit url ```/login``` untuk mendapatkan token

7. gunakan pada bearer token untuk meng hit route api yang lain nya

8. untuk contoh payload postman bisa di import ```postman_collection.json```

9. buat ```config.yaml``` silahkan copy dari config.sample

10. untuk menjalankan setelah git clone

11. run ```go mod tidy```

12. ```go run cmd/app/main.go```
    