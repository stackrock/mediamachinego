package mediamachine

import "encoding/json"

type Creds interface {
	isCreds()
}

type CredsNamed string

// CredsAWS - re-usable aws credentials that can be attached to a BlobStore
// Allows callers to mix-and-match blob stores. Callers are encouraged to
// provide the smallest surface-area credentials.
type CredsAWS struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

// CredsAzure - re-usable azure credentials that can be attached to a BlobStore
// Allows callers to mix-and-match blob stores. Callers are encouraged to
// provide the smallest surface-area credentials.
type CredsAzure struct {
	AccountName string
	AccountKey  string
}

// CredsGCP - re-usable gcp bucket credentials that can be used as InputCreds/OutputCreds.
// For Google Cloud credentials, you need to provide the contents of the json credentials file.
// See https://cloud.google.com/iam/docs/creating-managing-service-account-keys#iam-service-account-keys-create-console
type CredsGCP json.RawMessage

func (CredsNamed) isCreds() {}
func (CredsAWS) isCreds()   {}
func (CredsAzure) isCreds() {}
func (CredsGCP) isCreds()   {}
