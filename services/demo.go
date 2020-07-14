package services

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var d struct {
	a int
	db *sql.DB
}

func Connect() {
	db,_:=sql.Open("mysql","root:u4i3o2p1@(localhost)/lalala")
	//defer db.Close()
	err := db.Ping()
	if err != nil{
		fmt.Println("111")
	}
	d.a = 1
	d.db = db
//	fmt.Println(db)
}

type Weather struct {
	Id int
	Temperaturer int
	Other1 string
	Other2 string
	Place string
	Update_time string
}

func GetWeather(place string) Weather {
	if d.a != 1 {
		Connect()
	}
	//fmt.Println(d)
	//db,_:=sql.Open("mysql","root:u4i3o2p1@(localhost)/lalala")
	//defer db.Close()
	var w Weather
	sql := "select * from weather where place =\"" + place + "\" order by id desc limit 1"
	fmt.Println(sql)
	//sql = "select * from weather order by id desc limit 1"
	rows:=d.db.QueryRow(sql)
	var id, l1 int
	var l2, l3, l4 string
	var l5 string

	rows.Scan(&id,&l1,&l2,&l3,&l4,&l5)
	w = Weather{id,l1,l2,l3,l4,l5}
	fmt.Println(w)
	return w
}
func GetView(name string) int{
	if d.a !=1 {
		Connect()
	}
	sql := "SELECT view FROM Himawari WHERE title = \"" + name + "\" ORDER BY id DESC limit 1"
	fmt.Println(sql)
	var aview int
	d.db.QueryRow(sql).Scan(&aview)
	fmt.Println(aview)
	return aview
}

func GetAllMsg() []string {
	if d.a != 1 {
		Connect()
	}
	sql := "select message_board from message order by id desc limit 5"
	rows,_ := d.db.Query(sql)
	arr := make([]string, 0)
	for rows.Next(){
		var msg string
		rows.Scan(&msg)
		arr = append(arr, msg)
	}
	return arr
}

func WriteMsg(msg string) {
	if d.a != 1{
		Connect()
	}

	sql := "insert into message(message_board) values(?)"
	d.db.Exec(sql, msg)
}



