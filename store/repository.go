package store

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "mongodb://admin:MacCollector2018@ds151892.mlab.com:51892/heroku_bw74ps3m"

// DBNAME the name of the DB instance
const DBNAME = "heroku_bw74ps3m"

// COLLECTION is the name of the collection in DB
const COLLECTION = "sigfox"

var macAddressID = 10
var scanID = 10

// GetMacAddresses returns the list of MacAddresses
func (r Repository) GetMacAddresses() MacAddresses {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := MacAddresses{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// GetMacAddressById returns a unique MacAddress
func (r Repository) GetMacAddressById(id int) MacAddress {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	var result MacAddress

	fmt.Println("ID in GetMacAddressById", id)

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// GetMacAddressesByString takes a search string as input and returns macAddresses
func (r Repository) GetMacAddressesByString(query string) MacAddresses {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	result := MacAddresses{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"title": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// AddMacAddress adds a MacAddress in the DB
func (r Repository) AddMacAddress(macAddress MacAddress) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	macAddressID++
	macAddress.ID = macAddressID
	session.DB(DBNAME).C(COLLECTION).Insert(macAddress)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New MacAddress ID- ", macAddress.ID)

	return true
}

// UpdateMacAddress updates a MacAddress in the DB
func (r Repository) UpdateMacAddress(macAddress MacAddress) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	err = session.DB(DBNAME).C(COLLECTION).UpdateId(macAddress.ID, macAddress)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated MacAddress ID - ", macAddress.ID)

	return true
}

// DeleteMacAddress deletes an MacAddress
func (r Repository) DeleteMacAddress(id int) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Remove macAddress
	if err = session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted MacAddress ID - ", id)
	// Write status
	return "OK"
}

// GetScans returns the list of Scans
func (r Repository) GetScans() Scans {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := Scans{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// GetScanById returns a unique Scan
func (r Repository) GetScanById(id int) Scan {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	var result Scan

	fmt.Println("ID in GetScanById", id)

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// GetScansByString takes a search string as input and returns scan
func (r Repository) GetScansByString(query string) Scans {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	result := Scans{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"title": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// AddScan adds a Scan in the DB
func (r Repository) AddScan(scan Scan) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	scanID++
	scan.ID = scanID
	session.DB(DBNAME).C(COLLECTION).Insert(scan)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New Scan ID- ", scan.ID)

	return true
}

// UpdateScan updates a Scan in the DB
func (r Repository) UpdateScan(scan Scan) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	err = session.DB(DBNAME).C(COLLECTION).UpdateId(scan.ID, scan)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated Scan ID - ", scan.ID)

	return true
}

// DeleteScan deletes an Scan
func (r Repository) DeleteScan(id int) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Remove macScan
	if err = session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted MacScan ID - ", id)
	// Write status
	return "OK"
}
