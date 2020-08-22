# Data Provider

RESTful API for query game data.

## Get Game

Path ` GET /kais/v1/games`

### Query Parameter

| Parameter | Type | Mandatory | Description |
| :---: | :---: | :---: | :---: |
| game-type | string | Y | One of NBA, MLB, NPB. |
| begin | string | Y | Begin date(include), ex: 20200131. |
| end | string | Y | End date(include), ex: 20200228. |

## Example

