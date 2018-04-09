package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	kvs "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/keyvaluestore"
	"github.com/hyperledger/fabric-sdk-go/test/integration"
)

/**
 * @name initKeys
 * Helper function to Initialize sign and verify keys (Public and Private keys)
 * Used for Signing and verifying JWT Token
 */
func initKeys() {
	signBytes, _ := ioutil.ReadFile(privKeyPath)
	signKey, _ = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	verifyBytes, _ := ioutil.ReadFile(pubKeyPath)
	verifyKey, _ = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
}

/**
 * @name bcNetworksetup
 * This function Suggests which Blockchain Network to connect
 */
func bcNetworksetup() {
	var err error

	testSetup := integration.BaseSetupImpl{
		ConfigFile: "config/" + ConfigTestFile,
	}

	testFabricConfig, err = testSetup.InitConfig()()
	if err != nil {
		fmt.Printf("Failed InitConfig [%s]\n", err)
		os.Exit(1)
	}

}

/**
 * @name initLogger
 * Helper functions to initialize a logger
 */
func initLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

/**
 * @name createLogger
 * Helper functions to create a logger
 */
func createLogger() {
	_ = os.Mkdir("logs", os.ModePerm)

	current := time.Now()
	Timestamp := current.Format("2006-01-02T15.04.05")
	logFileName := "logs/" + Timestamp + ".log"
	logFile, _ := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY, 0666)
	//check(err)
	initLogger(ioutil.Discard, logFile, logFile, logFile)
}

/**
 * @name check
 * Helper functions for any un-recoverable error
 */
func check(context string, err error) {
	if err != nil {
		Error.Printf(context)
		panic("Got Error : " + err.Error() + "\n" + context)
	}
}

/**
 * @name KeyValueStoreOptions
 * Helper functions to initialize key value store
 */
func keyValueStoreOptions() *kvs.FileKeyValueStoreOptions {
	return &kvs.FileKeyValueStoreOptions{
		Path: KeyValueStorePath,
		KeySerializer: func(key interface{}) (string, error) {
			keyString, ok := key.(string)
			if !ok {
				return "", errors.New("converting key to string failed")
			}
			return path.Join(KeyValueStorePath, keyString+".json"), nil
		},
	}
}
