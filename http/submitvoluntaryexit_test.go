// Copyright © 2020, 2021 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http_test

import (
	"context"
	"os"
	"testing"

	client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/stretchr/testify/require"
)

func TestSubmitVoluntaryExit(t *testing.T) {
	tests := []struct {
		name string
		exit *phase0.SignedVoluntaryExit
	}{
		{
			name: "InvalidSignature",
			exit: &phase0.SignedVoluntaryExit{
				Message: &phase0.VoluntaryExit{
					ValidatorIndex: 12345,
					Epoch:          2,
				},
				Signature: phase0.BLSSignature{},
			},
		},
	}

	service, err := http.New(context.Background(),
		http.WithTimeout(timeout),
		http.WithAddress(os.Getenv("HTTP_ADDRESS")),
	)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := service.(client.VoluntaryExitSubmitter).SubmitVoluntaryExit(context.Background(), test.exit)
			require.Contains(t, err.Error(), "400")
		})
	}
}
