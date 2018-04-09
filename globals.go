package main

import (
	"crypto/rsa"
	"log"

	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/pkg/fabric-ca-client"
	client "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client"
	kvs "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/keyvaluestore"
)

const (
	ConfigTestFile    = "../config/config_test.yaml"
	org1Name          = "Org1"
	KeyValueStorePath = "C:\\Users\\vardan_nadkarni\\enroll_user"
	privKeyPath       = "C:\\Users\\vardan_nadkarni\\Desktop\\vardan_gopath\\src\\id-rest-manager\\keys\\id_rsa"
	pubKeyPath        = "C:\\Users\\vardan_nadkarni\\Desktop\\vardan_gopath\\src\\id-rest-manager\\keys\\id_rsa.pub"
)

var (
	Trace            *log.Logger
	Info             *log.Logger
	Warning          *log.Logger
	Error            *log.Logger
	verifyKey        *rsa.PublicKey
	signKey          *rsa.PrivateKey
	testFabricConfig apiconfig.Config
	fabricCaClient   fabricCAClient.FabricCA
	fabricClient     client.Client
	teststore        kvs.FileKeyValueStore
)
