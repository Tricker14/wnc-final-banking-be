## Makefile
### Swagger
```bash
$ make swagger
```
### Swagger UI located at
```link
http://localhost:PORT/swagger/index.html
```
---
### Wire
```bash
$ make wire
```

## Create migration files
```bash
$ make migration name=create_names_table
```

### Apply migration files
```bash
$ make migrate-up
```