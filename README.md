# KaiJi API Server

[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![Build Status](https://travis-ci.org/boennemann/badges.svg?branch=master)](https://travis-ci.org/boennemann/badges)    

KaiJi REST API service.

## Build

```
$ docker-compose build
```

## Docker Image

```
$ docker pull allensyk/kaiji-data-provider
```

## API
### Get Game

PATH: ` GET /kaiji/v1/games`

#### Query Parameter

| Parameter | Type | Mandatory | Description |
| :---: | :---: | :---: | :--- |
| game-type | string | Y | One of NBA, MLB, NPB. |
| begin | string | Y | Begin date(include), ex: 20200131. |
| end | string | Y | End date(include), ex: 20200228. |

### Example

#### Request

```
GET /kaiji/v1/games?game-type=NBA&begin=20191004&end=20191005
```

#### Response

```
HTTP/1.1 200 OK
Content-Type: application/json
...

[
    {
        "_id": "5f3fee94a48cf4de2d416516",
        "game_time": "2019-10-04T01:00:00Z",
        "gamble_id": "313",
        "game_type": "NBA",
        "guest": {
            "name": "HOU",
            "score": 109
        },
        "host": {
            "name": "LAC",
            "score": 96
        },
        "gamble_info": {
            "total_point": {
                "threshold": 224.5,
                "response": {
                    "under": 1.75,
                    "over": 1.75
                },
                "judgement": "under",
                "prediction": {
                    "under": {
                        "percentage": 30,
                        "population": 147
                    },
                    "over": {
                        "percentage": 70,
                        "population": 335
                    }
                }
            },
            "spread_point": {
                "guest": -4.5,
                "host": 4.5,
                "response": {
                    "guest": 1.7,
                    "host": 1.8
                },
                "judgement": "guest",
                "prediction": {
                    "guest": {
                        "percentage": 53,
                        "population": 322
                    },
                    "host": {
                        "percentage": 47,
                        "population": 283
                    },
                    "major": true
                }
            },
            "original": {
                "response": {},
                "judgement": "guest",
                "prediction": {
                    "guest": {},
                    "host": {}
                }
            }
        }
    },
    {
        "_id": "5f3fee94a48cf4de2d416517",
        "game_time": "2019-10-04T09:30:00Z",
        "gamble_id": "323",
        "game_type": "NBA",
        "guest": {
            "name": "IND",
            "score": 132
        },
        "host": {
            "name": "SAC",
            "score": 131
        },
        "gamble_info": {
            "total_point": {
                "threshold": 222.5,
                "response": {
                    "under": 1.75,
                    "over": 1.75
                },
                "judgement": "over",
                "prediction": {
                    "under": {
                        "percentage": 43,
                        "population": 162
                    },
                    "over": {
                        "percentage": 57,
                        "population": 213
                    },
                    "major": true
                }
            },
            "spread_point": {
                "guest": 6.5,
                "host": -6.5,
                "response": {
                    "guest": 1.75,
                    "host": 1.75
                },
                "judgement": "guest",
                "prediction": {
                    "guest": {
                        "percentage": 43,
                        "population": 254
                    },
                    "host": {
                        "percentage": 57,
                        "population": 334
                    }
                }
            },
            "original": {
                "response": {},
                "judgement": "guest",
                "prediction": {
                    "guest": {},
                    "host": {}
                }
            }
        }
    },
    {
        "_id": "5f3fee9aa48cf4de2d416518",
        "game_time": "2019-10-05T09:30:00Z",
        "gamble_id": "362",
        "game_type": "NBA",
        "guest": {
            "name": "SAC",
            "score": 106
        },
        "host": {
            "name": "IND",
            "score": 130
        },
        "gamble_info": {
            "total_point": {
                "threshold": 229.5,
                "response": {
                    "under": 1.75,
                    "over": 1.75
                },
                "judgement": "over",
                "prediction": {
                    "under": {
                        "percentage": 72,
                        "population": 66
                    },
                    "over": {
                        "percentage": 28,
                        "population": 26
                    }
                }
            },
            "spread_point": {
                "guest": -3.5,
                "host": 3.5,
                "response": {
                    "guest": 1.75,
                    "host": 1.75
                },
                "judgement": "host",
                "prediction": {
                    "guest": {
                        "percentage": 58,
                        "population": 63
                    },
                    "host": {
                        "percentage": 42,
                        "population": 46
                    }
                }
            },
            "original": {
                "response": {},
                "judgement": "host",
                "prediction": {
                    "guest": {},
                    "host": {}
                }
            }
        }
    }
]
```