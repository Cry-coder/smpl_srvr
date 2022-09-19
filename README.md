

# Test Routes

## Ping

GET:`http://localhost:8081/ping`

## Events List

GET:`http://localhost:8007/v1/events`

## Event (by ID)

GET:`http://localhost:8007/v1/events/{id}`

## Create record 

POST:`http://localhost:8007/v1/events`

## Update 

PUT:`http://localhost:8007/v1/events`

## Delete  (by ID)

GET:`http://localhost:8007/v1/events/{id}`

# Table example
## Change password and db name in repository.go settings
`CREATE TABLE staff (
personid SERIAL,
lastname varchar(15),
firstname varchar(15),
city varchar(15)
);`

`INSERT INTO staff (lastname, firstname, city)
VALUES ('Klichko', 'Vitaliy', 'Kyiv');`

`INSERT INTO staff (lastname, firstname, city)
VALUES ('Klichko', 'Volodumir', 'Kyiv');`