package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Counters struct {
	Id int `json: "id"`
	Count int `json: "count"`
	Time time.Time `json: "time"`
}

type Users struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Nama string `bson:"nama" json:"nama"`
	Kota string `bson:"kota" json:"kota"`
	Keranjang []Keranjang `bson:"keranjang" json:"keranjang"`
}

type Keranjang struct {
	IdBarang bson.ObjectId `bson:"_id,omitempty" json:"id"`
	NamaBarang string `bson:"namaBarang" json:"namaBarang"`
	HargaBarang string `bson:"hargaBarang" json:"hargaBarang"`
}

var db *mgo.Database

// Buat session database
func init() {
	session, err := mgo.Dial("database")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("godbapp")
}

// Khusus untuk users
func collectionUsers() *mgo.Collection {
	return db.C("users")
}

// Khusus untuk counter
func collectionCounter() *mgo.Collection {
	return db.C("counter")
}

// Mengambil semua user
func GetAllUsers() ([]Users, error) {
	res := []Users{}

	if err := collectionUsers().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// Mengambil satu user
func GetOne(id string) (*Users, error) {
	res := Users{}

	if err := collectionUsers().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Tambah user
func Insert(users Users) error {
	users.Id = bson.NewObjectId()
	return collectionUsers().Insert(users)
}

// Hapus user
func Remove(id string) error {
	res := Users{}

	if err := collectionUsers().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return err
	}

	return collectionUsers().Remove(&res)
}

// Tambah keranjang
func TambahKeranjang(iduser string, keranjang Keranjang) error {
	keranjang.IdBarang = bson.NewObjectId()
	who := bson.M{"_id": bson.ObjectIdHex(iduser)}
	pushto := bson.M{"$push": bson.M{"keranjang": keranjang}}
	return collectionUsers().Update(who, pushto)
}

// Hapus keranjang
func HapusKeranjang(iduser string, idbarang string) error {
	who := bson.M{"_id": bson.ObjectIdHex(iduser), "keranjang._id": bson.ObjectIdHex(idbarang)}
	pullto := bson.M{"$pull": bson.M{"keranjang": bson.M{"_id": bson.ObjectIdHex(idbarang)}}}
	return collectionUsers().Update(who, pullto)
}

// Proses counter
func CounterProc() int {
	var num = 0
	cnow := Counters{}
	err := collectionCounter().Find(bson.M{"id": 1}).One(&cnow)
	if err != nil {
		collectionCounter().Insert(&Counters{Id: 1, Count: num, Time: time.Now()})
	}else{
		num = cnow.Count
		num++
		// Update
		col := bson.M{"id": 1}
		chg := bson.M{"$set": bson.M{"count": num, "time": time.Now()}}
		collectionCounter().Update(col, chg)
	}
	return num
}
