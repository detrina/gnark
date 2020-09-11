// Copyright 2020 ConsenSys AG
//
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

// Code generated by gnark/internal/generators DO NOT EDIT

package r1cs

import (
	bw761backend "github.com/consensys/gnark/internal/backend/bw761"

	"github.com/consensys/gurvy/bw761/fr"
)

func (r1cs *UntypedR1CS) toBW761() *bw761backend.R1CS {

	toReturn := bw761backend.R1CS{
		NbWires:         r1cs.NbWires,
		NbPublicWires:   r1cs.NbPublicWires,
		NbSecretWires:   r1cs.NbSecretWires,
		SecretWires:     r1cs.SecretWires,
		PublicWires:     r1cs.PublicWires,
		WireTags:        r1cs.WireTags,
		NbConstraints:   r1cs.NbConstraints,
		NbCOConstraints: r1cs.NbCOConstraints,
		Constraints:     r1cs.Constraints,
		Coefficients:    make([]fr.Element, len(r1cs.Coefficients)),
		Logs:            r1cs.Logs,
		DebugInfo:       r1cs.DebugInfo,
	}

	for i := 0; i < len(r1cs.Coefficients); i++ {
		toReturn.Coefficients[i].SetBigInt(&r1cs.Coefficients[i])
	}

	return &toReturn
}
