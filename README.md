API 


POST /signup

Request

{
    "Username":"sunny.dua",
    "Email" : "sd.ud@abc.com",
    "Password":"please@12"
}

POST /login

Request

{
	"Username":"sunny.dua",
    "Password":"please@12"
}
Response

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6IjEiLCJjcmVhdGVkIjoxNTU1Nzg2Mzk1fQ.8-SQ9qFrlx6uMrwypwlnPXvmNxs_oqzNsTlekAvAgso"
}

Offer Creation - POST /offer

Request
{
	"bid_price":2000,
	"go_live":"2019-04-20T10:00:00Z",
	"life_time":100,
	"photo_url":"",
	"title":"TV",
	"sold":false
}
Response

{
    "id": "1",
    "bid_price": 2000,
    "go_live": "2019-04-20T10:00:00Z",
    "lifetime": 100,
    "photo_url": "",
    "title": "TV",
    "created_by": "sunny.dua",
    "sold": false
}
Fetching Offers from Db - GET /offer or /offer?size=10&page=0&sortKey=go_live

Bid Creation - POST /bids

Request

{
	"bid_price" : 2005,
	"offer_id" :1 
}

Respone

{
    "Id": 5,
    "bid_price": 2005,
    "offer_id": 1,
    "Client": {
        "Id": 1,
        "username": "sunny.dua",
        "email": "sd.ud@abc.com",
        "password": ""
    },
    "time_stamp": "2019-04-21T00:23:38.582723+05:30",
    "accepted": false,
    "client_id": 1
}

Accept Bid for an offer - PUT /vi/bid/:bidid

RESPONSE

{
    "Id": 3,
    "bid_price": 2001,
    "offer_id": 1,
    "Client": {
        "Id": 1,
        "username": "sunny.dua",
        "email": "sd.ud@abc.com",
        "password": ""
    },
    "time_stamp": "2019-04-20T23:51:24.914943+05:30",
    "accepted": true,
    "client_id": 1
}

Fetch Sold offers GET - /sold

[
    {
        "Id": 1,
        "bid_price": 2000,
        "go_live": "2019-04-20T17:37:29Z",
        "life_time": 100,
        "photo_url": "",
        "title": "TV",
        "sold": true,
        "created_by": "sunny.dua",
        "bid_id": 3,
        "Bid": {
            "Id": 3,
            "bid_price": 2001,
            "offer_id": 1,
            "Client": {
                "Id": 0,
                "username": "",
                "email": "",
                "password": ""
            },
            "time_stamp": "0001-01-01T00:00:00Z",
            "accepted": true,
            "client_id": 0
        }
    }
]