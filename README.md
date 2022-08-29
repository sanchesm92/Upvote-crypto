
# Upvote Cryptocurrency

API Restful in Golang where you can vote your favorite cryptocurrency

- Golang;
- Gin;
- Gorm;
- Postgres
- Docker;

# How to Use

- clone this repository
- install go, docker and docker-compose
- run docker-compose up -d && go run main.go 

# Endpoints

- http://localhost:3000/upvote/currencies/
- http://localhost:3000/upvote/currencies/ranking

- Examples:
- GET:
```json
[
	{
		"id": 1,
		"name": "bitcoin",
		"vote": 1,
		"created": "2022-08-29T11:47:39.31198-03:00",
		"updated": "2022-08-29T11:47:39.31198-03:00",
		"deleted": null
	}
]
```
- POST:
```json
{
	"name": "bitcoin"
}
```
- RANKING:
```json
[
	{
		"Name": "bitcoin",
		"Vote": 2
	},
	{
		"Name": "ethereum",
		"Vote": 1
	}
]
```
# Any bugs or suggestions?

Contatct me:
- https://www.linkedin.com/in/felipesanchesm/
- https://github.com/sanchesm92

# Enjoy.
