package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jtblin/go-ldap-client"
	"github.com/sirupsen/logrus"
	rldap "gopkg.in/ldap.v2"

	"github.com/dexidp/dex/connector"
	ldapConnector "github.com/dexidp/dex/connector/ldap"
)

var (
	host    = "10.0.129.127"
	port    = 30389
	tlsPort = 30636
	base    = "dc=example,dc=org"
	bindDn  = "cn=admin,dc=example,dc=org"
	bindPwd = "admin"
)

var CADATA string = `-----BEGIN CERTIFICATE-----
MIICoDCCAYgCCQDQlOqjhKVjPjANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDDAdr
dWJlLWNhMB4XDTE5MDgwOTAyMjQ1NloXDTQxMDcwNDAyMjQ1NlowEjEQMA4GA1UE
AwwHa3ViZS1jYTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAM4aKmKt
KC+aRHvF1JsmPbfcWRI3ft7gKu+LLy5K2TIitzAXQ/bWtuHoFhuNLUXcOhvfys99
mmBTbVHcF9LQMtRDoSFQjOAEnSTM/uci/gAnFYbV5M0Ea3WbUolMNESELWeUefa3
GSMKs1Pk3y2AsAkw4/aOCGn35lfZxgGy/oCRT5CAtYi8gm/XX7uaeUxXXW2SK64l
pt+Jc/9iXfyGVJhrnZNvYHsF/Ug+whtqkzIHNvz6xY40U7JrOJlb4vCSsCojguxp
QuZfENoCrKncQhodzXzDpXk7GmbfTWYFH7Rx0M5zcxy8PQC6z3OJb+RbL6A2t/RF
dhg6iMvziXSZwosCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAcrO9Y4IBHDZWSHcJ
eyn/PgqH5BXAhzLHgfV396dkEHyk5lu1s5JOHB0k2Y8cvuotqDb0eRHw+FhP0/j4
H9333nL6j3W0in7GrA8GXpuhEUm0tsgiZ0B+g+FQjKnmWzelxJvCoUPLl2AWQjyF
95gflRGN2ipmY5i2pLVllaih5z8xvaZDHwZLjGMvyk2hOGJ0A7GNgvySyrJRnSyK
26Z4mm1Ez5Luexa/mD7sP+B/kVSb3FhG4L5mNi1YYM82qtBE6/oBMTEdOvoQrBnj
1gQE0/oHY0YAtdL2Alh5J4gAiDjLHH5FUZXdpDD+eVMFJczCi4LE3zP8i8JYdT8h
Iim+rg==
-----END CERTIFICATE-----`

func main() {
	secure6361 := Secure636VerfiyServerCrt()
	verify(secure6361)

	secure6362 := Secure636NoVerifyServerCrt()
	verify(secure6362)

	secure389 := Secure389()
	verify(secure389)

	insecure389 := Insecure389()
	verify(insecure389)
}

func verify(c *ldapConnector.Config) {
	/*
		InsecureNoSSL 用来控制使用使用tls，使用时，必须为false
		InsecureSkipVerify 用来控制客户端是否校验服务端证书，为false时，RootCA或者RootCAData其一必须存在
		StartTLS 用于使用389端口，但是启用tls
	*/
	fmt.Printf(`Verify ldap with 
		Host: %v
		InsecureNoSSL: %v
		StartTLS: %v
		InsecureSkipVerify: %v
		ClientCert: %v
		ClientKey: %v
		RootCA: %v
		RootCAData: %v`, c.Host, c.InsecureNoSSL, c.StartTLS, c.InsecureSkipVerify, c.ClientCert, c.ClientKey, c.RootCA, c.RootCAData)
	c.UserSearch.BaseDN = base
	c.UserSearch.NameAttr = "cn"
	c.UserSearch.EmailAttr = "mail"
	c.UserSearch.IDAttr = "DN"
	c.UserSearch.Username = "cn"
	c.BindDN = "cn=admin,dc=example,dc=org"
	c.BindPW = "admin"

	l := &logrus.Logger{Out: ioutil.Discard, Formatter: &logrus.TextFormatter{}}
	conn, err := c.OpenConnector(l)
	if err != nil {
		fmt.Printf("open connector: %v", err)
	}

	s := connector.Scopes{OfflineAccess: true, Groups: true}

	ident, validPW, err := conn.Login(context.Background(), s, "name", "password")
	if err != nil {
		fmt.Printf("Failed to validate password: %v", err)
		return
	}

	fmt.Printf("Ident: %v, validPW: %v", ident, validPW)
}

func Secure636VerfiyServerCrt() *ldapConnector.Config {
	c := &ldapConnector.Config{}
	c.Host = "192.168.1.179:30636"

	// If InsecureSkipVerify is set to true, then c.RootCA or c.RootCAData is not required
	c.InsecureSkipVerify = false

	// Using tls, then InsecureNoSSL must be false, otherwise error(unable to read LDAP response packet) occurs
	c.InsecureNoSSL = false

	c.ClientCert = "/Users/alauda/Projects/ldap/ssl/cert.pem"
	c.ClientKey = "/Users/alauda/Projects/ldap/ssl/key.pem"
	c.RootCA = "/Users/alauda/Projects/ldap/ssl/ca-client.pem"

	return c
}

func Secure636NoVerifyServerCrt() *ldapConnector.Config {
	c := &ldapConnector.Config{}
	c.Host = "192.168.1.179:30636"

	// If InsecureSkipVerify is set to true, then c.RootCA or c.RootCAData is not required
	c.InsecureSkipVerify = true

	// Using tls, then InsecureNoSSL must be false, otherwise error(unable to read LDAP response packet) occurs
	c.InsecureNoSSL = false

	c.ClientCert = "/Users/alauda/Projects/ldap/ssl/cert.pem"
	c.ClientKey = "/Users/alauda/Projects/ldap/ssl/key.pem"

	return c
}

func Secure389() *ldapConnector.Config {
	c := &ldapConnector.Config{}
	c.Host = "192.168.1.179:30389"

	// If InsecureSkipVerify is set to true, then c.RootCA or c.RootCAData is not required
	c.InsecureSkipVerify = true

	// Using tls, then InsecureNoSSL must be false, otherwise error(unable to read LDAP response packet) occurs
	c.InsecureNoSSL = false

	c.StartTLS = true

	c.ClientCert = "/Users/alauda/Projects/ldap/ssl/cert.pem"
	c.ClientKey = "/Users/alauda/Projects/ldap/ssl/key.pem"

	return c

}

func Insecure389() *ldapConnector.Config {
	c := &ldapConnector.Config{}
	c.Host = "192.168.1.179:30389"

	// If InsecureSkipVerify is set to true, then c.RootCA or c.RootCAData is not required
	// c.InsecureSkipVerify = true

	// InsecureNoSSL must be true when using 389 port
	c.InsecureNoSSL = true

	return c
}

func rawLdap() {
	client := &ldap.LDAPClient{
		Base:               base,
		Host:               host,
		Port:               port,
		UseSSL:             false,
		SkipTLS:            true,
		InsecureSkipVerify: true,
		BindDN:             bindDn,
		BindPassword:       bindPwd,
	}

	tlsClient := &ldap.LDAPClient{
		Base:               base,
		Host:               host,
		Port:               tlsPort,
		UseSSL:             true,
		SkipTLS:            false,
		InsecureSkipVerify: true,
		BindDN:             bindDn,
		BindPassword:       bindPwd,
		ServerName:         "ldap.com",
	}

	tlsClient.ClientCertificates = nil

	// It is the responsibility of the caller to close the connection
	defer client.Close()
	client.Connect()
	client.Conn.Bind(client.BindDN, client.BindPassword)

	searchRequest := rldap.NewSearchRequest(
		client.Base,
		rldap.ScopeWholeSubtree, rldap.NeverDerefAliases, 0, 0, false,
		"(objectclass=inetOrgPerson)",
		[]string{"cn"}, // can it be something else than "cn"?
		nil,
	)

	sr, err := client.Conn.Search(searchRequest)
	if err != nil {
		log.Fatalf("Seaching user with err: %+v", err)
	}
	for _, entry := range sr.Entries {
		log.Println(entry.GetAttributeValue("cn"))
	}

	defer tlsClient.Close()
	err = tlsClient.Connect()
	if err != nil {
		log.Fatalf("Conn to ldap with err with tls: %+v", err)
		return
	}
	tlsClient.Conn.Bind(tlsClient.BindDN, tlsClient.BindPassword)
	sr, err = tlsClient.Conn.Search(searchRequest)
	if err != nil {
		log.Fatalf("Seaching user with err with tls: %+v", err)
	}
	for _, entry := range sr.Entries {
		log.Println(entry.GetAttributeValue("cn"))
	}

}
