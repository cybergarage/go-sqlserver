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
	"regexp"

	"github.com/cybergarage/go-authenticator/auth"
	"github.com/cybergarage/go-authenticator/auth/tls"
)

type manager struct {
	auth.Manager
	credStore        map[string]auth.Credential
	commonNameRegexp []*regexp.Regexp
}

// NewManager returns a new auth manager.
func NewManager() Manager {
	return &manager{
		Manager:          auth.NewManager(),
		credStore:        map[string]auth.Credential{},
		commonNameRegexp: []*regexp.Regexp{},
	}
}

// SetCommonNameRegexps sets common name regular expressions.
func (mgr *manager) SetCommonNameRegexps(regexps ...string) error {
	for _, re := range regexps {
		r, err := regexp.Compile(re)
		if err != nil {
			return err
		}
		mgr.commonNameRegexp = append(mgr.commonNameRegexp, r)
	}
	return nil
}

// SetCredential sets a credential.
func (mgr *manager) SetCredentials(creds ...auth.Credential) error {
	for _, cred := range creds {
		mgr.credStore[cred.Username()] = cred
	}
	return nil
}

// LookupCredential looks up a credential.
func (mgr *manager) LookupCredential(q auth.Query) (auth.Credential, bool, error) {
	user := q.Username()
	cred, ok := mgr.credStore[user]
	return cred, ok, nil
}

// VerifyCertificate verifies the client certificate.
func (mgr *manager) VerifyCertificate(conn tls.Conn) (bool, error) {
	if len(mgr.commonNameRegexp) == 0 {
		return true, nil
	}
	for _, cert := range conn.ConnectionState().PeerCertificates {
		for _, re := range mgr.commonNameRegexp {
			if re.MatchString(cert.Subject.CommonName) {
				return true, nil
			}
		}
	}
	return false, nil
}
