/*
   Copyright (C) 2019 MIRACL UK Ltd.

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.


    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

     https://www.gnu.org/licenses/agpl-3.0.en.html

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   You can be released from the requirements of the license by purchasing
   a commercial license. Buying such a license is mandatory as soon as you
   develop commercial activities involving the MIRACL Core Crypto SDK
   without disclosing the source code of your own applications, or shipping
   the MIRACL Core Crypto SDK with a closed source product.
*/

/* Run through some test vectors for hash-to-curve draft standard */

package main

import "fmt"

import "github.com/miracl/core/go/core"
import "github.com/miracl/core/go/core/ED25519"
import "github.com/miracl/core/go/core/NIST256"
import "github.com/miracl/core/go/core/GOLDILOCKS"
import "github.com/miracl/core/go/core/SECP256K1"
import "github.com/miracl/core/go/core/BLS12381"

func ceil(a int,b int) int {
    return (((a)-1)/(b)+1)
}

func hash_to_field_ED25519(hash int,hlen int,DST []byte,M []byte,ctr int) []*ED25519.FP {
	var u []*ED25519.FP
	q := ED25519.NewBIGints(ED25519.Modulus)
	k := q.Nbits()
	r := ED25519.NewBIGints(ED25519.CURVE_Order)
	m := r.Nbits();
	L:= ceil(k+ceil(m,2),8)
	var fd=make([]byte,L)
	OKM:=core.XMD_Expand(hash,hlen,L*ctr,DST,M)
    for i:=0;i<ctr;i++ {
        for j:=0;j<L;j++ {
            fd[j]=OKM[i*L+j];
        }
		dx:=ED25519.DBIG_fromBytes(fd)
		w:=ED25519.NewFPbig(dx.Mod(q))
		u=append(u,w)
    }
	return u;
}

func htp_ED25519(mess []byte) {
	
	fmt.Println("Random Access - message= "+string(mess))
	DST := []byte("edwards25519_XMD:SHA-256_ELL2_RO_TESTGEN")
	u:=hash_to_field_ED25519(core.MC_SHA2,ED25519.HASH_TYPE,DST,mess,2)
	fmt.Printf("u[0]= %s\n",u[0].ToString())
	fmt.Printf("u[1]= %s\n",u[1].ToString())
	P:=ED25519.ECP_map2point(u[0])
	fmt.Printf("Q[0]= %s\n",P.ToString())
	P1:=ED25519.ECP_map2point(u[1])
	fmt.Printf("Q[0]= %s\n",P1.ToString())
	P.Add(P1)
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

	fmt.Printf("Non-Uniform\n")
    DST=[]byte("edwards25519_XMD:SHA-256_ELL2_NU_TESTGEN");
    u=hash_to_field_ED25519(core.MC_SHA2,ED25519.HASH_TYPE,DST,mess,1);
	fmt.Printf("u= %s\n",u[0].ToString())
	P=ED25519.ECP_map2point(u[0])
	fmt.Printf("Q= %s\n",P.ToString())
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

}

func hash_to_field_NIST256(hash int,hlen int,DST []byte,M []byte,ctr int) []*NIST256.FP {
	var u []*NIST256.FP
	q := NIST256.NewBIGints(NIST256.Modulus)
	k := q.Nbits()
	r := NIST256.NewBIGints(NIST256.CURVE_Order)
	m := r.Nbits();
	L:= ceil(k+ceil(m,2),8)

	var fd=make([]byte,L)
	OKM:=core.XMD_Expand(hash,hlen,L*ctr,DST,M)
    for i:=0;i<ctr;i++ {
        for j:=0;j<L;j++ {
            fd[j]=OKM[i*L+j];
        }
		dx:=NIST256.DBIG_fromBytes(fd)
		w:=NIST256.NewFPbig(dx.Mod(q))
		u=append(u,w)
    }
	return u;
}

func htp_NIST256(mess []byte) {
	
	fmt.Println("Random Access - message= "+string(mess))
	DST := []byte("P256_XMD:SHA-256_SSWU_RO_TESTGEN")
	u:=hash_to_field_NIST256(core.MC_SHA2,NIST256.HASH_TYPE,DST,mess,2)
	fmt.Printf("u[0]= %s\n",u[0].ToString())
	fmt.Printf("u[1]= %s\n",u[1].ToString())
	P:=NIST256.ECP_map2point(u[0])
	fmt.Printf("Q[0]= %s\n",P.ToString())
	P1:=NIST256.ECP_map2point(u[1])
	fmt.Printf("Q[0]= %s\n",P1.ToString())
	P.Add(P1)
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

	fmt.Printf("Non-Uniform\n")
    DST=[]byte("P256_XMD:SHA-256_SSWU_NU_TESTGEN");
    u=hash_to_field_NIST256(core.MC_SHA2,NIST256.HASH_TYPE,DST,mess,1);
	fmt.Printf("u= %s\n",u[0].ToString())
	P=NIST256.ECP_map2point(u[0])
	fmt.Printf("Q= %s\n",P.ToString())
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

}

func hash_to_field_GOLDILOCKS(hash int,hlen int,DST []byte,M []byte,ctr int) []*GOLDILOCKS.FP {
	var u []*GOLDILOCKS.FP
	q := GOLDILOCKS.NewBIGints(GOLDILOCKS.Modulus)
	k := q.Nbits()
	r := GOLDILOCKS.NewBIGints(GOLDILOCKS.CURVE_Order)
	m := r.Nbits();
	L:= ceil(k+ceil(m,2),8)
	var fd=make([]byte,L)
	OKM:=core.XMD_Expand(hash,hlen,L*ctr,DST,M)
    for i:=0;i<ctr;i++ {
        for j:=0;j<L;j++ {
            fd[j]=OKM[i*L+j];
        }
		dx:=GOLDILOCKS.DBIG_fromBytes(fd)
		w:=GOLDILOCKS.NewFPbig(dx.Mod(q))
		u=append(u,w)
    }
	return u;
}

func htp_GOLDILOCKS(mess []byte) {
	
	fmt.Println("Random Access - message= "+string(mess))
	DST := []byte("edwards448_XMD:SHA-512_ELL2_RO_TESTGEN")
	u:=hash_to_field_GOLDILOCKS(core.MC_SHA2,GOLDILOCKS.HASH_TYPE,DST,mess,2)
	fmt.Printf("u[0]= %s\n",u[0].ToString())
	fmt.Printf("u[1]= %s\n",u[1].ToString())
	P:=GOLDILOCKS.ECP_map2point(u[0])
	fmt.Printf("Q[0]= %s\n",P.ToString())
	P1:=GOLDILOCKS.ECP_map2point(u[1])
	fmt.Printf("Q[0]= %s\n",P1.ToString())
	P.Add(P1)
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

	fmt.Printf("Non-Uniform\n")
    DST=[]byte("edwards448_XMD:SHA-512_ELL2_NU_TESTGEN");
    u=hash_to_field_GOLDILOCKS(core.MC_SHA2,GOLDILOCKS.HASH_TYPE,DST,mess,1);
	fmt.Printf("u= %s\n",u[0].ToString())
	P=GOLDILOCKS.ECP_map2point(u[0])
	fmt.Printf("Q= %s\n",P.ToString())
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

}

func hash_to_field_SECP256K1(hash int,hlen int,DST []byte,M []byte,ctr int) []*SECP256K1.FP {
	var u []*SECP256K1.FP
	q := SECP256K1.NewBIGints(SECP256K1.Modulus)
	k := q.Nbits()
	r := SECP256K1.NewBIGints(SECP256K1.CURVE_Order)
	m := r.Nbits();
	L:= ceil(k+ceil(m,2),8)
	var fd=make([]byte,L)
	OKM:=core.XMD_Expand(hash,hlen,L*ctr,DST,M)
    for i:=0;i<ctr;i++ {
        for j:=0;j<L;j++ {
            fd[j]=OKM[i*L+j];
        }
		dx:=SECP256K1.DBIG_fromBytes(fd)
		w:=SECP256K1.NewFPbig(dx.Mod(q))
		u=append(u,w)
    }
	return u;
}

func htp_SECP256K1(mess []byte) {
	
	fmt.Println("Random Access - message= "+string(mess))
	DST := []byte("secp256k1_XMD:SHA-256_SVDW_RO_TESTGEN")
	u:=hash_to_field_SECP256K1(core.MC_SHA2,SECP256K1.HASH_TYPE,DST,mess,2)
	fmt.Printf("u[0]= %s\n",u[0].ToString())
	fmt.Printf("u[1]= %s\n",u[1].ToString())
	P:=SECP256K1.ECP_map2point(u[0])
	fmt.Printf("Q[0]= %s\n",P.ToString())
	P1:=SECP256K1.ECP_map2point(u[1])
	fmt.Printf("Q[0]= %s\n",P1.ToString())
	P.Add(P1)
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

	fmt.Printf("Non-Uniform\n")
    DST=[]byte("secp256k1_XMD:SHA-256_SVDW_NU_TESTGEN");
    u=hash_to_field_SECP256K1(core.MC_SHA2,SECP256K1.HASH_TYPE,DST,mess,1);
	fmt.Printf("u= %s\n",u[0].ToString())
	P=SECP256K1.ECP_map2point(u[0])
	fmt.Printf("Q= %s\n",P.ToString())
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

}

func hash_to_field_BLS12381(hash int,hlen int,DST []byte,M []byte,ctr int) []*BLS12381.FP {
	var u []*BLS12381.FP
	q := BLS12381.NewBIGints(BLS12381.Modulus)
	k := q.Nbits()
	r := BLS12381.NewBIGints(BLS12381.CURVE_Order)
	m := r.Nbits();
	L:= ceil(k+ceil(m,2),8)
	var fd=make([]byte,L)
	OKM:=core.XMD_Expand(hash,hlen,L*ctr,DST,M)
    for i:=0;i<ctr;i++ {
        for j:=0;j<L;j++ {
            fd[j]=OKM[i*L+j];
        }
		dx:=BLS12381.DBIG_fromBytes(fd)
		w:=BLS12381.NewFPbig(dx.Mod(q))
		u=append(u,w)
    }
	return u;
}

func htp_BLS12381(mess []byte) {
	
	fmt.Println("Random Access - message= "+string(mess))
	DST := []byte("BLS12381G1_XMD:SHA-256_SVDW_RO_TESTGEN")
	u:=hash_to_field_BLS12381(core.MC_SHA2,BLS12381.HASH_TYPE,DST,mess,2)
	fmt.Printf("u[0]= %s\n",u[0].ToString())
	fmt.Printf("u[1]= %s\n",u[1].ToString())
	P:=BLS12381.ECP_map2point(u[0])
	fmt.Printf("Q[0]= %s\n",P.ToString())
	P1:=BLS12381.ECP_map2point(u[1])
	fmt.Printf("Q[0]= %s\n",P1.ToString())
	P.Add(P1)
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

	fmt.Printf("Non-Uniform\n")
    DST=[]byte("BLS12381G1_XMD:SHA-256_SVDW_NU_TESTGEN");
    u=hash_to_field_BLS12381(core.MC_SHA2,BLS12381.HASH_TYPE,DST,mess,1);
	fmt.Printf("u= %s\n",u[0].ToString())
	P=BLS12381.ECP_map2point(u[0])
	fmt.Printf("Q= %s\n",P.ToString())
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

}

func hash_to_field2_BLS12381(hash int,hlen int,DST []byte,M []byte,ctr int) []*BLS12381.FP2 {
	var u []*BLS12381.FP2
	q := BLS12381.NewBIGints(BLS12381.Modulus)
	k := q.Nbits()
	r := BLS12381.NewBIGints(BLS12381.CURVE_Order)
	m := r.Nbits();
	L:= ceil(k+ceil(m,2),8)
	var fd=make([]byte,L)
	OKM:=core.XMD_Expand(hash,hlen,2*L*ctr,DST,M)
    for i:=0;i<ctr;i++ {
        for j:=0;j<L;j++ {
            fd[j]=OKM[2*i*L+j];
        }
		dx:=BLS12381.DBIG_fromBytes(fd)
		w1:=BLS12381.NewFPbig(dx.Mod(q))

        for j:=0;j<L;j++ {
            fd[j]=OKM[(2*i+1)*L+j];
        }
		dx=BLS12381.DBIG_fromBytes(fd)
		w2:=BLS12381.NewFPbig(dx.Mod(q))
		u=append(u,BLS12381.NewFP2fps(w1,w2))
    }
	return u;
}

func htp2_BLS12381(mess []byte) {
	
	fmt.Println("Random Access - message= "+string(mess))
	DST := []byte("BLS12381G2_XMD:SHA-256_SVDW_RO_TESTGEN")
	u:=hash_to_field2_BLS12381(core.MC_SHA2,BLS12381.HASH_TYPE,DST,mess,2)
	fmt.Printf("u[0]= %s\n",u[0].ToString())
	fmt.Printf("u[1]= %s\n",u[1].ToString())
	P:=BLS12381.ECP2_map2point(u[0])
	fmt.Printf("Q[0]= %s\n",P.ToString())
	P1:=BLS12381.ECP2_map2point(u[1])
	fmt.Printf("Q[0]= %s\n",P1.ToString())
	P.Add(P1)
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

	fmt.Printf("Non-Uniform\n")
    DST=[]byte("BLS12381G2_XMD:SHA-256_SVDW_NU_TESTGEN");
    u=hash_to_field2_BLS12381(core.MC_SHA2,BLS12381.HASH_TYPE,DST,mess,1);
	fmt.Printf("u= %s\n",u[0].ToString())
	P=BLS12381.ECP2_map2point(u[0])
	fmt.Printf("Q= %s\n",P.ToString())
	P.Cfp()
	P.Affine();
	fmt.Printf("P= %s\n\n",P.ToString())

}


func main() {
    fmt.Printf("\nTesting HTP for curve ED25519\n");
	htp_ED25519([]byte(""))
	htp_ED25519([]byte("abc"))
	htp_ED25519([]byte("abcdef0123456789"))
	htp_ED25519([]byte("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

    fmt.Printf("\nTesting HTP for curve NIST256\n");
	htp_NIST256([]byte(""))
	htp_NIST256([]byte("abc"))
	htp_NIST256([]byte("abcdef0123456789"))
	htp_NIST256([]byte("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

    fmt.Printf("\nTesting HTP for curve GOLDILOCKS\n");
	htp_GOLDILOCKS([]byte(""))
	htp_GOLDILOCKS([]byte("abc"))
	htp_GOLDILOCKS([]byte("abcdef0123456789"))
	htp_GOLDILOCKS([]byte("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

    fmt.Printf("\nTesting HTP for curve SECP256K1\n");
	htp_SECP256K1([]byte(""))
	htp_SECP256K1([]byte("abc"))
	htp_SECP256K1([]byte("abcdef0123456789"))
	htp_SECP256K1([]byte("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

    fmt.Printf("\nTesting HTP for curve BLS12381\n");
	htp_BLS12381([]byte(""))
	htp_BLS12381([]byte("abc"))
	htp_BLS12381([]byte("abcdef0123456789"))
	htp_BLS12381([]byte("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

    fmt.Printf("\nTesting HTP for curve BLS12381_G2\n");
	htp2_BLS12381([]byte(""))
	htp2_BLS12381([]byte("abc"))
	htp2_BLS12381([]byte("abcdef0123456789"))
	htp2_BLS12381([]byte("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

}