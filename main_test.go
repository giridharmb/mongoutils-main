package main

import (
	"fmt"
	"github.com/giridharmb/mongoutils"
	"log"
	"os"
	"testing"
)

/*
--------UNIT_TEST---------

# go test -v

=== RUN   TestInsertRecord
2021/01/05 13:26:15 TestInsertRecord...
2021-01-05T13:26:15.929-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:89	Connected to MongoDB!
2021/01/05 13:26:16 InsertGenericIfNotExists() : Matched 1 documents and updated 1 documents.
2021-01-05T13:26:16.121-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:85	Connection to MongoDB closed.
--- PASS: TestInsertRecord (0.73s)
=== RUN   TestDeleteRecord
2021/01/05 13:26:16 TestDeleteRecord...
2021-01-05T13:26:16.784-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:89	Connected to MongoDB!
2021/01/05 13:26:16 InsertGenericIfNotExists() : Matched 1 documents and updated 0 documents.
2021-01-05T13:26:16.969-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:85	Connection to MongoDB closed.
2021-01-05T13:26:17.636-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:168	Connected to MongoDB!
2021/01/05 13:26:17 Deleted (1) # of documents
2021/01/05 13:26:17 deleteGenericIfExists() : Deleted Document with key (test1)
2021-01-05T13:26:17.834-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:164	Connection to MongoDB closed.
--- PASS: TestDeleteRecord (1.71s)
=== RUN   TestFindRecord
2021/01/05 13:26:17 TestFindRecord...
2021-01-05T13:26:18.492-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:89	Connected to MongoDB!
2021/01/05 13:26:18 InsertGenericIfNotExists() : Matched 0 documents and updated 0 documents.
2021-01-05T13:26:18.676-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:85	Connection to MongoDB closed.
2021-01-05T13:26:19.172-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:479	Connected to MongoDB!
2021/01/05 13:26:19 FindGenericRecord() : Found a single document: {Key:test1 Value:[{Key:animal Value:cow} {Key:sound Value:mooooo}]}
2021-01-05T13:26:19.658-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:475
Connection to MongoDB closed.
--- PASS: TestFindRecord (1.82s)
PASS
ok  	galaxy_tfe/mongotest	4.899s
*/

/*
--------BENCHMARK---------

# go test -bench=.

2021/01/05 13:27:18 TestInsertRecord...
2021-01-05T13:27:19.700-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:89	Connected to MongoDB!
2021/01/05 13:27:19 InsertGenericIfNotExists() : Matched 1 documents and updated 0 documents.
2021-01-05T13:27:19.885-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:85	Connection to MongoDB closed.
2021/01/05 13:27:19 TestDeleteRecord...
2021-01-05T13:27:20.550-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:89	Connected to MongoDB!
2021/01/05 13:27:20 InsertGenericIfNotExists() : Matched 1 documents and updated 0 documents.
2021-01-05T13:27:20.759-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:85	Connection to MongoDB closed.
2021-01-05T13:27:21.414-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:168	Connected to MongoDB!
2021/01/05 13:27:21 Deleted (1) # of documents
2021/01/05 13:27:21 deleteGenericIfExists() : Deleted Document with key (test1)
2021-01-05T13:27:21.619-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:164	Connection to MongoDB closed.
2021/01/05 13:27:21 TestFindRecord...
2021-01-05T13:27:22.177-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:89	Connected to MongoDB!
2021/01/05 13:27:22 InsertGenericIfNotExists() : Matched 0 documents and updated 0 documents.
2021-01-05T13:27:22.554-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:85	Connection to MongoDB closed.
2021-01-05T13:27:23.135-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:479	Connected to MongoDB!
2021/01/05 13:27:23 FindGenericRecord() : Found a single document: {Key:test1 Value:[{Key:animal Value:cow} {Key:sound Value:mooooo}]}
2021-01-05T13:27:23.505-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:475
Connection to MongoDB closed.
2021-01-05T13:27:23.993-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:89	Connected to MongoDB!
2021/01/05 13:27:24 InsertGenericIfNotExists() : Matched 1 documents and updated 0 documents.
2021-01-05T13:27:24.174-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:85	Connection to MongoDB closed.
2021-01-05T13:27:25.142-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:479	Connected to MongoDB!
2021/01/05 13:27:25 FindGenericRecord() : Found a single document: {Key:test1 Value:[{Key:animal Value:cow} {Key:sound Value:mooooo}]}
2021-01-05T13:27:25.361-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:475
Connection to MongoDB closed.
2021-01-05T13:27:27.024-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:168	Connected to MongoDB!
2021/01/05 13:27:27 Deleted (1) # of documents
2021/01/05 13:27:27 deleteGenericIfExists() : Deleted Document with key (test1)
2021-01-05T13:27:27.550-0800	DEBUG	mongoutils@v0.0.0-20201230214651-1d8101ab7b41/mongoutils.go:164	Connection to MongoDB closed.
goos: darwin
goarch: amd64
pkg: galaxy_tfe/mongotest
BenchmarkGrpc-16    	       1	4042274480 ns/op
PASS
ok  	galaxy_tfe/mongotest	9.613s
*/

func TestInsertRecord(t *testing.T) {

	log.Printf("TestInsertRecord...")

	// set (export) the mongoDB hostname "MONGOHOST"
	// in your environment variable
	mongoHOST := os.Getenv("MONGOHOST")
	if mongoHOST == "" {
		t.Errorf("environment variable MONGOHOST is not set !")
	}

	mongoURL := fmt.Sprintf("mongodb://%v", mongoHOST)
	mongoutils.Initialize(mongoURL)

	testData1 := make(map[string]string)
	testData1["animal"] = "cow"
	testData1["sound"] = "mooooo"

	myData := mongoutils.GenericData{Key: "test1", Value: testData1}

	//testData1String := GetStringFromInterface(testData1)
	//testData1Map := GetMapStringStringFromString(testData1String)

	// insertion ////////////////////////////////

	result, err := mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
	if err != nil {
		t.Errorf("OOPS: %v", err.Error())
	}
	_, ok := result["insertion"]
	if !ok {
		t.Errorf("key 'insertion' is missing in the response")
	}
	resultOfInsertion := result["insertion"]
	if resultOfInsertion == "failed" {
		t.Errorf("Expecting insertion to be 'successful' , instead got %v", resultOfInsertion)
	}

}

func TestDeleteRecord(t *testing.T) {
	log.Printf("TestDeleteRecord...")

	// set (export) the mongoDB hostname "MONGOHOST"
	// in your environment variable
	mongoHOST := os.Getenv("MONGOHOST")
	if mongoHOST == "" {
		t.Errorf("environment variable MONGOHOST is not set !")
	}

	mongoURL := fmt.Sprintf("mongodb://%v", mongoHOST)
	mongoutils.Initialize(mongoURL)

	testData1 := make(map[string]string)
	testData1["animal"] = "cow"
	testData1["sound"] = "mooooo"

	myData := mongoutils.GenericData{Key: "test1", Value: testData1}

	//testData1String := GetStringFromInterface(testData1)
	//testData1Map := GetMapStringStringFromString(testData1String)

	// insertion ////////////////////////////////

	result, err := mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
	if err != nil {
		t.Errorf("OOPS: %v", err.Error())
	}
	_, ok := result["insertion"]
	if !ok {
		t.Errorf("key 'insertion' is missing in the response")
	}
	resultOfInsertion := result["insertion"]
	if resultOfInsertion == "failed" {
		t.Errorf("Expecting resultOfInsertion to be 'successful' , instead got %v", resultOfInsertion)
	}

	// deletion ////////////////////////////////

	deletionResult, err := mongoutils.DeleteGenericIfExists("test1", "db1", "collection1")
	if err != nil {
		t.Errorf("OOPS: %v", err.Error())
	}
	_, ok = deletionResult["deletion"]
	if !ok {
		t.Errorf("key 'deletion' is missing in the response")
	}
	resultOfDeletion := deletionResult["deletion"]
	if resultOfInsertion == "failed" {
		t.Errorf("Expecting deletion to be 'successful' , instead got %v", resultOfDeletion)
	}
}

func TestFindRecord(t *testing.T) {

	log.Printf("TestFindRecord...")

	// set (export) the mongoDB hostname "MONGOHOST"
	// in your environment variable
	mongoHOST := os.Getenv("MONGOHOST")
	if mongoHOST == "" {
		t.Errorf("environment variable MONGOHOST is not set !")
	}

	mongoURL := fmt.Sprintf("mongodb://%v", mongoHOST)
	mongoutils.Initialize(mongoURL)

	testData1 := make(map[string]string)
	testData1["animal"] = "cow"
	testData1["sound"] = "mooooo"

	myData := mongoutils.GenericData{Key: "test1", Value: testData1}

	//testData1String := GetStringFromInterface(testData1)
	//testData1Map := GetMapStringStringFromString(testData1String)

	// insertion ////////////////////////////////

	result, err := mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
	if err != nil {
		t.Errorf("OOPS: %v", err.Error())
	}
	_, ok := result["insertion"]
	if !ok {
		t.Errorf("key 'insertion' is missing in the response")
	}
	resultOfInsertion := result["insertion"]
	if resultOfInsertion == "failed" {
		t.Errorf("Expecting resultOfInsertion to be 'successful' , instead got %v", resultOfInsertion)
	}

	fetchedData, err := mongoutils.FindGenericRecord("test1", "db1", "collection1")
	if err != nil {
		t.Errorf("OOPS: %v", err.Error())
	}
	_, ok = fetchedData["search_result"]
	if !ok {
		t.Errorf("key 'search_result' is missing in the response")
	}
	searchResult, ok := fetchedData["search_result"].(string)
	if !ok {
		t.Errorf("search_result value is not a string")
	}
	if searchResult == "not_found" {
		t.Errorf("search_result value is %v , expecting value to be found", searchResult)
	}
}

func BenchmarkGrpc(b *testing.B) {
	mongoHOST := os.Getenv("MONGOHOST")
	if mongoHOST == "" {
		b.Errorf("environment variable MONGOHOST is not set !")
	}

	mongoURL := fmt.Sprintf("mongodb://%v", mongoHOST)
	mongoutils.Initialize(mongoURL)

	testData1 := make(map[string]string)
	testData1["animal"] = "cow"
	testData1["sound"] = "mooooo"

	myData := mongoutils.GenericData{Key: "test1", Value: testData1}

	for i := 0; i < b.N; i++ {
		/////////// insert //////////////
		result, err := mongoutils.InsertGenericIfNotExists(myData, "db1", "collection1")
		if err != nil {
			b.Errorf("OOPS: %v", err.Error())
		}
		_, ok := result["insertion"]
		if !ok {
			b.Errorf("key 'insertion' is missing in the response")
		}
		resultOfInsertion := result["insertion"]
		if resultOfInsertion == "failed" {
			b.Errorf("Expecting resultOfInsertion to be 'successful' , instead got %v", resultOfInsertion)
		}

		/////////// find //////////////

		fetchedData, err := mongoutils.FindGenericRecord("test1", "db1", "collection1")
		if err != nil {
			b.Errorf("OOPS: %v", err.Error())
		}
		_, ok = fetchedData["search_result"]
		if !ok {
			b.Errorf("key 'search_result' is missing in the response")
		}
		searchResult, ok := fetchedData["search_result"].(string)
		if !ok {
			b.Errorf("search_result value is not a string")
		}
		if searchResult == "not_found" {
			b.Errorf("search_result value is %v , expecting value to be found", searchResult)
		}

		/////////// delete //////////////

		deletionResult, err := mongoutils.DeleteGenericIfExists("test1", "db1", "collection1")
		if err != nil {
			b.Errorf("OOPS: %v", err.Error())
		}
		_, ok = deletionResult["deletion"]
		if !ok {
			b.Errorf("key 'deletion' is missing in the response")
		}
		resultOfDeletion := deletionResult["deletion"]
		if resultOfInsertion == "failed" {
			b.Errorf("Expecting deletion to be 'successful' , instead got %v", resultOfDeletion)
		}
	}
}
