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
	"strings"
	"testing"
	"time"

	client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/prysmaticlabs/go-bitfield"
	"github.com/stretchr/testify/require"
)

func TestSubmitAttestations(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Good",
		},
	}

	service, err := http.New(context.Background(),
		http.WithTimeout(timeout),
		http.WithAddress(os.Getenv("HTTP_ADDRESS")),
	)
	require.NoError(t, err)

	// Need to fetch current slot for attestation.
	genesis, err := service.(client.GenesisProvider).Genesis(context.Background())
	require.NoError(t, err)
	slotDuration, err := service.(client.SlotDurationProvider).SlotDuration(context.Background())
	require.NoError(t, err)

	for _, test := range tests {
		thisSlot := phase0.Slot(uint64(time.Since(genesis.GenesisTime).Seconds()) / uint64(slotDuration.Seconds()))
		t.Run(test.name, func(t *testing.T) {
			// Fetch attestation data.
			attestationData, err := service.(client.AttestationDataProvider).AttestationData(context.Background(), thisSlot, 0)
			require.NoError(t, err)
			require.NotNil(t, attestationData)

			aggregationBits := bitfield.NewBitlist(128)
			aggregationBits.SetBitAt(1, true)
			attestation := &phase0.Attestation{
				AggregationBits: aggregationBits,
				Data:            attestationData,
				Signature: phase0.BLSSignature([96]byte{
					0xb1, 0x3c, 0xa7, 0x7f, 0xda, 0xb9, 0x0f, 0xce, 0xdf, 0x0c, 0xda, 0x74, 0xe9, 0xe9, 0xda, 0x1e,
					0xdb, 0xe4, 0x32, 0x91, 0x09, 0x48, 0xca, 0xad, 0xca, 0x64, 0xbb, 0xfb, 0x93, 0x34, 0x26, 0x44,
					0xac, 0xbb, 0xd3, 0xa1, 0x02, 0x4c, 0xa3, 0x9b, 0xd3, 0x50, 0x70, 0xca, 0xb3, 0xc6, 0x90, 0xd4,
					0x07, 0x43, 0x00, 0x1b, 0x44, 0x51, 0x53, 0xff, 0x97, 0x76, 0x18, 0x3c, 0xfe, 0x94, 0xec, 0x00,
					0x33, 0x90, 0xec, 0x76, 0x08, 0x4f, 0x7e, 0x20, 0x83, 0xcf, 0x3a, 0x46, 0xe1, 0xd6, 0xca, 0x1c,
					0x72, 0xb5, 0x71, 0xab, 0x58, 0x2d, 0x3d, 0x64, 0xe2, 0x69, 0x10, 0x20, 0x80, 0x85, 0x0d, 0x82,
				}),
			}

			err = service.(client.AttestationsSubmitter).SubmitAttestations(context.Background(), []*phase0.Attestation{attestation})
			// We will get an error as the bitlist is the incorrect size (on purpose, to stop our test being broadcast).
			if err != nil {
				require.True(t, strings.Contains(err.Error(), "Aggregation bitlist size (128) does not match committee size"))
			}
		})
	}
}
