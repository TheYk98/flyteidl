// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package admin

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "endpoint"), defaultConfig.Endpoint.String(), "For admin types,  specify where the uri of the service is located.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "insecure"), defaultConfig.UseInsecureConnection, "Use insecure connection.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "maxBackoffDelay"), defaultConfig.MaxBackoffDelay.String(), "Max delay for grpc backoff")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "perRetryTimeout"), defaultConfig.PerRetryTimeout.String(), "gRPC per retry timeout")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "maxRetries"), defaultConfig.MaxRetries, "Max number of gRPC retries")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "useAuth"), defaultConfig.UseAuth, "Whether or not to try to authenticate with options below")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "clientId"), defaultConfig.ClientId, "Client ID")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "clientSecretLocation"), defaultConfig.ClientSecretLocation, "File containing the client secret")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "scopes"), []string{}, "List of scopes to request")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "authorizationServerUrl"), defaultConfig.AuthorizationServerURL, "This is the URL to your IDP's authorization server'")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "tokenUrl"), defaultConfig.TokenURL, "Your IDPs token endpoint")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "authorizationHeader"), defaultConfig.AuthorizationHeader, "Custom metadata header to pass JWT")
	return cmdFlags
}
