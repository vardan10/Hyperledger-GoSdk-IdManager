package main

import (

	/* Formatting Imports */
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	/* Router Imports */

	jwt "github.com/dgrijalva/jwt-go"

	/* GO SDK Imports */
	ca "github.com/hyperledger/fabric-sdk-go/api/apifabca"
	cryptosuite "github.com/hyperledger/fabric-sdk-go/pkg/cryptosuite/bccsp/sw"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/pkg/fabric-ca-client"
	client "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/identity"
	kvs "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/keyvaluestore"
)

/**
 * @name enrollAdmin
 * Enrolls the Admin if not already enrolled and saved in kvs store
 */
func enrollAdmin() {
	mspID, err := testFabricConfig.MspID(org1Name)
	check("GetMspId() returned error: %v", err)

	cryptoSuiteProvider, err := cryptosuite.GetSuiteByConfig(testFabricConfig)
	check("Failed getting cryptosuite from config", err)

	caClient, err := fabricCAClient.NewFabricCAClient(org1Name, testFabricConfig, cryptoSuiteProvider)
	check("NewFabricCAClient returned error", err)

	client := client.NewClient(testFabricConfig)
	client.SetCryptoSuite(cryptoSuiteProvider)
	stateStore, err := kvs.NewFileKeyValueStore(keyValueStoreOptions())
	check("CreateNewFileKeyValueStore return error", err)
	client.SetStateStore(stateStore)

	adminUser, err := client.LoadUserFromStateStore("admin")
	check("client.LoadUserFromStateStore return error", err)

	// Enroll Admin user
	if adminUser == nil {
		key, cert, err := caClient.Enroll("admin", "adminpw")
		check("Enroll return error", err)
		if key == nil {
			fmt.Printf("private key return from Enroll is nil")
		}
		if cert == nil {
			fmt.Printf("cert return from Enroll is nil")
		}

		certPem, _ := pem.Decode(cert)
		if certPem == nil {
			fmt.Printf("Fail to decode pem block")
		}

		cert509, err := x509.ParseCertificate(certPem.Bytes)
		if err != nil {
			fmt.Printf("x509 ParseCertificate return error: %v", err)
		}
		if cert509.Subject.CommonName != "admin" {
			fmt.Printf("CommonName in x509 cert is not the enrollmentID")
		}

		adminUser2 := identity.NewUser("admin", mspID)
		adminUser2.SetPrivateKey(key)
		adminUser2.SetEnrollmentCertificate(cert)

		// Set Role as Admin
		var roles []string
		roles = append(roles, "admin")
		adminUser2.SetRoles(roles)

		// Save Admin
		err = client.SaveUserToStateStore(adminUser2)
		if err != nil {
			fmt.Printf("client.SaveUserToStateStore return error: %v", err)
		}
		adminUser, err = client.LoadUserFromStateStore("admin")
		if err != nil {
			fmt.Printf("client.LoadUserFromStateStore return error: %v", err)
		}
		if adminUser == nil {
			fmt.Printf("client.LoadUserFromStateStore return nil")
		}
	} else {
		fmt.Printf("Admin User already enrolled")
	}
}

/**
 * @name RegisterAndEnrollUser
 * Registers and Enrolls User
 */
func RegisterAndEnrollUser(username string, password string) (string, error) {

	cryptoSuiteProvider, err := cryptosuite.GetSuiteByConfig(testFabricConfig)
	check("Failed getting cryptosuite from config", err)

	caClient, err := fabricCAClient.NewFabricCAClient(org1Name, testFabricConfig, cryptoSuiteProvider)
	check("NewFabricCAClient returned error", err)

	caConfig, err := testFabricConfig.CAConfig(org1Name)
	check("GetCAConfig returned error", err)

	// Load admin user form keyValueStore
	client := client.NewClient(testFabricConfig)
	client.SetCryptoSuite(cryptoSuiteProvider)
	stateStore, err := kvs.NewFileKeyValueStore(keyValueStoreOptions())
	check("CreateNewFileKeyValueStore return error", err)
	client.SetStateStore(stateStore)
	adminUser, err := client.LoadUserFromStateStore("admin")
	if adminUser == nil {
		fmt.Printf("Admin user is nil")
	}

	// Register and Enroll a random user
	userName := username
	registerRequest := ca.RegistrationRequest{
		Name:        userName,
		Type:        "user",
		Affiliation: "org1.department1",
		CAName:      caConfig.CAName,
		Secret:      password,
	}

	enrolmentSecret, err := caClient.Register(adminUser, &registerRequest)
	if err != nil {
		fmt.Printf("Error from Register: %s", err)
	}
	fmt.Printf("Registered User: %s, Secret: %s", userName, enrolmentSecret)

	// Enroll User
	return EnrollUser(username, password)

}

/**
 * @name EnrollUser
 * Enrolls User
 */
func EnrollUser(username string, password string) (string, error) {
	mspID, err := testFabricConfig.MspID(org1Name)
	check("GetMspId() returned error: %v", err)

	cryptoSuiteProvider, err := cryptosuite.GetSuiteByConfig(testFabricConfig)
	check("Failed getting cryptosuite from config", err)

	caClient, err := fabricCAClient.NewFabricCAClient(org1Name, testFabricConfig, cryptoSuiteProvider)
	check("NewFabricCAClient returned error", err)

	ekey, ecert, err := caClient.Enroll(username, password)
	if err != nil {
		return "error Enrolling User", err
	}

	enrolleduser := identity.NewUser(username, mspID)
	enrolleduser.SetEnrollmentCertificate(ecert)
	enrolleduser.SetPrivateKey(ekey)

	return createJWTToken(username), nil
}

/**
 * @name createJWTToken
 * Creates JWT Token Based on Username and returns Payload
 */
func createJWTToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": username,
		"expiry":   time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(signKey)

	return tokenString
}
