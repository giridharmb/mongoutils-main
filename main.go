package main

import (
	"encoding/json"
	"fmt"
	"github.com/giridharmb/mongoutils"
	"go.uber.org/zap"
	"log"
	"os"
)

/*
Logger ...
*/
var Logger *zap.Logger

/*
DoPrint ...
Does Only Printf
*/
func DoPrint(data interface{}) {
	Logger.Debug(fmt.Sprintf("%v", data))
}

/*
GetStringFromInterface ...
*/
func GetStringFromInterface(data interface{}) string {
	jsonStr := ""
	marshalledData, err := json.Marshal(data)
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		log.Println(errorMessage)
	} else {
		jsonStr = string(marshalledData)
	}
	return jsonStr
}

/*
GetInterfaceFromString ...
*/
func GetInterfaceFromString(jsonStr string) interface{} {
	var data interface{}
	_ = json.Unmarshal([]byte(jsonStr), &data)
	return data
}

/*
GetMapStringStringFromString ...
*/
func GetMapStringStringFromString(jsonStr string) map[string]string {
	returnData := make(map[string]string)
	myMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(jsonStr), &myMap)
	for key, value := range myMap {
		myKey := fmt.Sprintf("%v", key)
		_, ok := value.(string)
		if ok {
			myValue := fmt.Sprintf("%v", value.(string))
			returnData[myKey] = myValue
		} else {
			continue
		}

	}
	return returnData
}

/*
GetMapStringStringFromString ...
*/
func GetMapStringStringFromInterface(data interface{}) map[string]string {
	dataMap := data.(map[string]interface{})
	returnData := make(map[string]string)
	for key, value := range dataMap {
		myKey := fmt.Sprintf("%v", key)
		_, ok := value.(string)
		if ok {
			myValue := fmt.Sprintf("%v", value.(string))
			returnData[myKey] = myValue
		} else {
			continue
		}
	}
	return returnData
}

func main() {

	// set (export) the mongoDB hostname "MONGOHOST"
	// in your environment variable
	mongoHOST := os.Getenv("MONGOHOST")
	if mongoHOST == "" {
		log.Fatalf("environment variable MONGOHOST is not set !")
	}

	mongoURL := fmt.Sprintf("mongodb://%v", mongoHOST)
	mongoutils.Initialize(mongoURL)

	Logger, _ = zap.NewDevelopment()

	Logger.Info("Starting Main Program")

	errorMessage := ""

	fmt.Println(errorMessage)

	testData1 := make(map[string]string)
	testData1["animal"] = "cow"
	testData1["sound"] = "mooooo"

	testData2 := make(map[string]string)
	testData2["animal"] = "dog"
	testData2["sound"] = "bowwww"

	testData3 := make(map[string]interface{})
	testData3["employee"] = "jack"
	testData3["employeeDetails"] = map[string]interface{}{
		"department":   "testing",
		"jobID":        "3463",
		"hoursPerWeek": 40,
	}

	testData4 := make(map[string]interface{})
	testData4["employee"] = "peter"
	testData4["employeeDetails"] = map[string]interface{}{
		"department":   "development",
		"jobID":        "6836",
		"hoursPerWeek": 40,
	}

	myData := mongoutils.GenericData{Key: "test1", Value: testData1}

	testData1String := GetStringFromInterface(testData1)
	testData1Map := GetMapStringStringFromString(testData1String)
	DoPrint(testData1Map)

	///////////// Insert Record /////////////

	DoPrint("--- Inserting A New Record ---")

	result, err := mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(result)

	///////////// Delete a record even though it does not exist /////////////

	DoPrint("--- Deleting A Record That Does Not Exist ---")

	result, err = mongoutils.DeleteGenericIfExists("test123", "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(result)

	///////////// Delete a record that *does* exist /////////////

	DoPrint("--- Deleting A Record That Does Exist ---")

	result, err = mongoutils.DeleteGenericIfExists("test1", "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(result)

	///////////// Fetching an entire collection /////////////

	DoPrint("--- Fetching An Entire Collection ---")

	collection, err := mongoutils.FetchCollection("db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(collection)

	///////////// Insert New Data (Multiple Items) /////////////

	DoPrint("--- Insert New Data (Multiple Items) ---")

	myData = mongoutils.GenericData{Key: "test1", Value: testData1}
	result, err = mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(result)

	myData = mongoutils.GenericData{Key: "test2", Value: testData2}
	result, err = mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(result)

	myData = mongoutils.GenericData{Key: "test3", Value: testData3}
	result, err = mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(result)

	myData = mongoutils.GenericData{Key: "test4", Value: testData4}
	result, err = mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(result)

	///////////// Fetching an entire collection (2) /////////////

	DoPrint("--- Fetching An Entire Collection (2) ---")

	collection, err = mongoutils.FetchCollection("db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(collection)

	///////////// Deleting Multiple Items /////////////

	DoPrint("--- Deleting multiple items ---")

	itemsToDelete := []string{"test1", "test2", "test100", "test200"}

	deletionResults, err := mongoutils.DeleteManyIfExists(itemsToDelete, "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(deletionResults)

	///////////// Inserting Multiple Items In One Shot /////////////

	DoPrint("--- Inserting multiple items in one shot ---")

	testData5 := make(map[string]interface{})
	testData5["employee"] = "joseph"
	testData5["employeeDetails"] = map[string]interface{}{
		"department":   "human-resources",
		"jobID":        "5336",
		"hoursPerWeek": 40,
	}

	testData6 := make(map[string]interface{})
	testData6["employee"] = "joseph"
	testData6["employeeDetails"] = map[string]interface{}{
		"department":   "finance",
		"jobID":        "3532",
		"hoursPerWeek": 40,
	}

	myData5 := mongoutils.GenericData{Key: "data5", Value: testData5}
	myData6 := mongoutils.GenericData{Key: "data6", Value: testData6}

	listOfItems := []mongoutils.GenericData{myData5, myData6}

	insertionResults, err := mongoutils.InsertManyIfNotExists(listOfItems, "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(insertionResults)

	///////////// Finding A Specific Record Which Does Not Exist /////////////

	DoPrint("--- Finding a specific record which does not exist ---")

	record, err := mongoutils.FindGenericRecord("data100", "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(record)

	///////////// Finding A Specific Record Which Does Exist /////////////

	DoPrint("--- Finding a specific record which does exist ---")

	record, err = mongoutils.FindGenericRecord("data5", "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(record)

	///////////// Finding Multiple Records /////////////

	DoPrint("--- Finding multiple records ---")

	recordsToFind := []string{"data5", "data6", "data100"}

	allRecords, err := mongoutils.FindMultipleGenericRecord(recordsToFind, "db1", "collection1")
	if err != nil {
		errorMessage := fmt.Sprintf("OOPS: %v", err.Error())
		DoPrint(errorMessage)
		return
	}
	DoPrint(allRecords)

}
