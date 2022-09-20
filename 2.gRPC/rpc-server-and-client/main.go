package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var Db []Item

func (a *API) GetDB(s string, reply *[]Item) error {
	*reply = Db
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, v := range Db {
		if v.Title == title {
			getItem = v
		}
	}
	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	Db = append(Db, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for i, v := range Db {
		if v.Title == edit.Title {
			Db[i] = Item{edit.Title, edit.Body}
			changed = Db[i]
		}

	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for i, v := range Db {
		if v.Title == item.Title && v.Body == item.Body {
			Db = append(Db[:i], Db[i+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

	// fmt.Println("initial database ", db)

	// a := Item{"first", "a test item"}
	// b := Item{"second", "a second item"}
	// c := Item{"third", "a third item"}

	// CreateItem(a)
	// CreateItem(b)
	// CreateItem(c)

	// fmt.Println("Second db: ", db)

	// DeleteItem(b)
	// fmt.Println("third db: ", db)

	// EditItem("third", Item{"fourth", "a new item"})
	// fmt.Println("fourth db: ", db)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)

}
