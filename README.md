# Arquitetura Hexagonal
FullCycle 3.0 | Arquitetura Hexagonal

### Como subir a aplicação

```bash
docker compose up -d
docker exec -it appproduct bash
```

### Execução dos testes

```bash
go test ./...
```

### Execução do main.go CLI

```bash
go run main.go cli -a=create -n="Product CLI" -p=25.0
go run main.go cli -a=get -i="4d7479ab-a1c7-446f-8282-25c3a0344c12"
```

### Execução do main.go HTTP

#### getProduct
```bash
go run main.go http
```

Em outro terminal rodar:
```bash
curl http://localhost:9000/product/4d7479ab-a1c7-446f-8282-25c3a0344c12
```

#### createProduct
```makefile
TYPE:
POST

URL:
http://localhost:9000/product

BODY:
{
    "name": "Test product",
    "price": 13.33
}
```

#### enable / disable

```bash
curl http://localhost:9000/product/4d7479ab-a1c7-446f-8282-25c3a0344c12/enable

curl http://localhost:9000/product/4d7479ab-a1c7-446f-8282-25c3a0344c12/disable
```