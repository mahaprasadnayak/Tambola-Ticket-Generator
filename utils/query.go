package utils

var (
	Insert_data = "INSERT INTO tambola (SetName, TicketNumber, NumberArray,InsertedAt) VALUES ($1, $2, $3,$4)"
	Select_ticket= "SELECT setname,ticketnumber,numberarray FROM tambola"
)
