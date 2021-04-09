package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"

	"github.com/alazyer/gocodes/pkg/ldap"
	"github.com/sirupsen/logrus"
)

func newLogger() (logrus.FieldLogger, error) {
	return &logrus.Logger{
		Level: logrus.DebugLevel,
	}, nil
}

func mai12n() {
	configFile := "./pkg/ldap/config.yaml"
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Print("failed to read config file %s: %v", configFile, err)
	}

	var c ldap.DEXConfig
	if err := yaml.Unmarshal(configData, &c); err != nil {
		fmt.Print("error parse config file %s: %v", configFile, err)
	}

	ctx := context.TODO()
	fmt.Println(ctx)

	for _, connector := range c.Connectors {
		fmt.Print("Name: ", connector.Name, " ")
		config := connector.Config
		data, _ := yaml.Marshal(config)
		data, _ = json.Marshal(config)
		fmt.Println(base64.StdEncoding.EncodeToString(data))
		ldapConnector, err := config.Open()
		if err != nil {
			fmt.Print("err when connect to ldap: %v", connector)
		}
		fmt.Println("ldapConnector", ldapConnector)

	}
}

/*
eyJob3N0IjoiMTI5LjI4LjE4Mi4xOTc6NjA2MSIsImluc2VjdXJlTm9TU0wiOnRydWUsImluc2VjdXJlU2tpcFZlcmlmeSI6dHJ1ZSwic3RhcnRUTFMiOmZhbHNlLCJyb290Q0EiOiIiLCJyb290Q0FEYXRhIjpudWxsLCJiaW5kRE4iOiJjbj1hZG1pbixkYz1leGFtcGxlLGRjPW9yZyIsImJpbmRQVyI6ImFkbWluIiwidXNlcm5hbWVQcm9tcHQiOiIiLCJ1c2VyU2VhcmNoIjp7ImJhc2VETiI6ImRjPWV4YW1wbGUsZGM9b3JnIiwiZmlsdGVyIjoiKG9iamVjdENsYXNzPWluZXRPcmdQZXJzb24pIiwidXNlcm5hbWUiOiJtYWlsIiwic2NvcGUiOiIiLCJpZEF0dHIiOiJ1aWQiLCJlbWFpbEF0dHIiOiJtYWlsIiwibmFtZUF0dHIiOiJjbiJ9LCJncm91cFNlYXJjaCI6eyJiYXNlRE4iOiJkYz1leGFtcGxlLGRjPW9yZyIsImZpbHRlciI6IihvYmplY3RDbGFzcz1wb3NpeEdyb3VwKSIsInNjb3BlIjoiIiwidXNlckF0dHIiOiJ1aWQiLCJncm91cEF0dHIiOiJtZW1iZXJ1aWQiLCJuYW1lQXR0ciI6ImNuIn19

eyJob3N0IjoiMTI5LjI4LjE4Mi4xOTc6NjA2MSIsImluc2VjdXJlTm9TU0wiOnRydWUsImluc2VjdXJlU2tpcFZlcmlmeSI6dHJ1ZSwic3RhcnRUTFMiOmZhbHNlLCJyb290Q0EiOiIiLCJyb290Q0FEYXRhIjpudWxsLCJiaW5kRE4iOiJjbj1hZG1pbixkYz1leGFtcGxlLGRjPW9yZyIsImJpbmRQVyI6ImFkbWluIiwidXNlcm5hbWVQcm9tcHQiOiLnlKjmiLflkI0iLCJ1c2VyU2VhcmNoIjp7ImJhc2VETiI6ImRjPWV4YW1wbGUsZGM9b3JnIiwiZmlsdGVyIjoiKG9iamVjdENsYXNzPWluZXRPcmdQZXJzb24pIiwidXNlcm5hbWUiOiJtYWlsIiwic2NvcGUiOiIiLCJpZEF0dHIiOiJ1aWQiLCJlbWFpbEF0dHIiOiJtYWlsIiwibmFtZUF0dHIiOiJjbiJ9LCJncm91cFNlYXJjaCI6eyJiYXNlRE4iOiJkYz1leGFtcGxlLGRjPW9yZyIsImZpbHRlciI6IihvYmplY3RDbGFzcz1wb3NpeEdyb3VwKSIsInNjb3BlIjoiIiwidXNlckF0dHIiOiJ1aWQiLCJncm91cEF0dHIiOiJtZW1iZXJ1aWQiLCJuYW1lQXR0ciI6ImNuIn19

eyJob3N0IjoiMTI5LjI4LjE4Mi4xOTc6NjA2MSIsImluc2VjdXJlTm9TU0wiOnRydWUsImluc2VjdXJlU2tpcFZlcmlmeSI6dHJ1ZSwic3RhcnRUTFMiOmZhbHNlLCJyb290Q0EiOiIiLCJyb290Q0FEYXRhIjpudWxsLCJiaW5kRE4iOiJjbj1hZG1pbixkYz1leGFtcGxlLGRjPW9yZyIsImJpbmRQVyI6ImFkbWluIiwidXNlcm5hbWVQcm9tcHQiOiJteXVzZXJuYW1lIiwidXNlclNlYXJjaCI6eyJiYXNlRE4iOiJkYz1leGFtcGxlLGRjPW9yZyIsImZpbHRlciI6IihvYmplY3RDbGFzcz1pbmV0T3JnUGVyc29uKSIsInVzZXJuYW1lIjoibWFpbCIsInNjb3BlIjoiIiwiaWRBdHRyIjoidWlkIiwiZW1haWxBdHRyIjoibWFpbCIsIm5hbWVBdHRyIjoiY24ifSwiZ3JvdXBTZWFyY2giOnsiYmFzZUROIjoiZGM9ZXhhbXBsZSxkYz1vcmciLCJmaWx0ZXIiOiIob2JqZWN0Q2xhc3M9cG9zaXhHcm91cCkiLCJzY29wZSI6IiIsInVzZXJBdHRyIjoidWlkIiwiZ3JvdXBBdHRyIjoibWVtYmVydWlkIiwibmFtZUF0dHIiOiJjbiJ9fQ==
*/
