# tax-calculator

## deploy

1. create file log at "/var/log/tax/tax-api.log"
2. run file with command 
  ```bash
go run cmd/app/tax-calculator/main.go
```

3. visit url http://localhost:8081
4. generate billing:
    - url : http://localhost:8081/create
    - method : post
    - payload json:
    ```bash
    [
      {
        "name" : "Lucky Stretch",
        "tax_code" : 2,
        "price" : 1000
      },
      {
        "name" : "Big Mac",
        "tax_code" : 1,
        "price" : 1000
      },
      {
        "name" : "Movie",
        "tax_code" : 3,
        "price" : 150
      }
    ]
    ```
    - response json:
    ```bash
    {
        "items": [
            {
                "name": "Lucky Stretch",
                "tax_code": 2,
                "type": "Tobacco",
                "refundable": "NO",
                "price": 1000,
                "tax": 30,
                "amount": 1030
            },
            {
                "name": "Big Mac",
                "tax_code": 1,
                "type": "Food & Beverages",
                "refundable": "YES",
                "price": 1000,
                "tax": 100,
                "amount": 1100
            },
            {
                "name": "Movie",
                "tax_code": 3,
                "type": "Entertainment",
                "refundable": "NO",
                "price": 150,
                "tax": 0.5,
                "amount": 150.5
            }
        ],
        "price_subtotal": 2150,
        "tax_subtotal": 130.5,
        "grand_total": 2280.5
    }
    ```
5. Run mysql migration from sql/init.sql
