// Copyright (C) 2025 The go-sqlserver Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"github.com/cybergarage/go-authenticator/auth"
	"github.com/cybergarage/go-authenticator/auth/tls"
)

// Conn represents a connection.
type Conn = auth.Conn

// CredentialStore represents a credential store.
type CredentialStore = auth.CredentialStore

// Query represents a query.
type Query = auth.Query

// CertificateAuthenticator represents a certificate authenticator.
type CertificateAuthenticator = auth.CertificateAuthenticator

type Manager interface {
	// SetCredentialStore sets the credential store.
	SetCredentialStore(store CredentialStore)
	// VerifyCredential verifies the client credential.
	VerifyCredential(conn Conn, q Query) (bool, error)
	// SetCertificateAuthenticator sets the certificate authenticator.
	SetCertificateAuthenticator(auth CertificateAuthenticator)
	// VerifyCertificate verifies the client certificate.
	VerifyCertificate(conn tls.Conn) (bool, error)
}
