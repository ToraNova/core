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

#![allow(non_snake_case)]
extern crate core;
use core::hmac;

fn ceil(a: usize,b: usize) -> usize {
    return (a-1)/b+1;
}

fn hash_to_field_ed25519(hash: usize,hlen: usize,dst: &[u8],msg: &[u8],ctr: usize) -> [core::ed25519::fp::FP;2] {

    use core::ed25519::big::BIG;
    use core::ed25519::dbig::DBIG;
    use core::ed25519::fp::FP;
    use core::ed25519::rom;

    let mut u: [FP; 2] = [FP::new(),FP::new()];

    let q = BIG::new_ints(&rom::MODULUS);
    let k=q.nbits();
    let r = BIG::new_ints(&rom::CURVE_ORDER);
    let m=r.nbits();
    let L=ceil(k+ceil(m,2),8);
    let mut okm: [u8;512]=[0;512];
    hmac::xmd_expand(hash,hlen,&mut okm,L*ctr,&dst,&msg);
    let mut fd: [u8;256]=[0;256];
    for i in 0..ctr {
        for j in 0..L {
            fd[j]=okm[i*L+j];
        }
		let mut dx=DBIG::frombytes(&fd[0..L]);
		let w=FP::new_big(&dx.dmod(&q));
		u[i].copy(&w);
    }

    return u;
}

fn htp_ed25519(mess: &str) {
    use core::ed25519::ecp;
    use core::ed25519::ecp::ECP;
    println!("Random Access - message= {}",mess);
    let m = mess.as_bytes();
    let mut dst = "edwards25519_XMD:SHA-256_ELL2_RO_TESTGEN".as_bytes();
    let mut u=hash_to_field_ed25519(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,2);
    println!("u[0]= {}",u[0].tostring());
    println!("u[1]= {}",u[1].tostring());
    let mut P=ECP::map2point(&u[0]);
    println!("Q[0]= {}",P.tostring());
    let P1=ECP::map2point(&u[1]);
    println!("Q[1]= {}",P1.tostring());
    P.add(&P1);
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());

    println!("\nNon-Uniform");
    dst = "edwards25519_XMD:SHA-256_ELL2_NU_TESTGEN".as_bytes();
    u=hash_to_field_ed25519(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,1);
    println!("u[0]= {}",u[0].tostring());
    P=ECP::map2point(&u[0]);
    println!("Q= {}",P.tostring());
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());
}

fn hash_to_field_nist256(hash: usize,hlen: usize,dst: &[u8],msg: &[u8],ctr: usize) -> [core::nist256::fp::FP;2] {

    use core::nist256::big::BIG;
    use core::nist256::dbig::DBIG;
    use core::nist256::fp::FP;
    use core::nist256::rom;

    let mut u: [FP; 2] = [FP::new(),FP::new()];

    let q = BIG::new_ints(&rom::MODULUS);
    let k=q.nbits();
    let r = BIG::new_ints(&rom::CURVE_ORDER);
    let m=r.nbits();
    let L=ceil(k+ceil(m,2),8);
    let mut okm: [u8;512]=[0;512];
    hmac::xmd_expand(hash,hlen,&mut okm,L*ctr,&dst,&msg);
    let mut fd: [u8;256]=[0;256];
    for i in 0..ctr {
        for j in 0..L {
            fd[j]=okm[i*L+j];
        }
		let mut dx=DBIG::frombytes(&fd[0..L]);
		let w=FP::new_big(&dx.dmod(&q));
		u[i].copy(&w);
    }

    return u;
}

fn htp_nist256(mess: &str) {
    use core::nist256::ecp;
    use core::nist256::ecp::ECP;
    println!("Random Access - message= {}",mess);
    let m = mess.as_bytes();
    let mut dst = "P256_XMD:SHA-256_SSWU_RO_TESTGEN".as_bytes();
    let mut u=hash_to_field_nist256(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,2);
    println!("u[0]= {}",u[0].tostring());
    println!("u[1]= {}",u[1].tostring());
    let mut P=ECP::map2point(&u[0]);
    println!("Q[0]= {}",P.tostring());
    let P1=ECP::map2point(&u[1]);
    println!("Q[1]= {}",P1.tostring());
    P.add(&P1);
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());

    println!("\nNon-Uniform");
    dst = "P256_XMD:SHA-256_SSWU_NU_TESTGEN".as_bytes();
    u=hash_to_field_nist256(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,1);
    println!("u[0]= {}",u[0].tostring());
    P=ECP::map2point(&u[0]);
    println!("Q= {}",P.tostring());
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());
}

fn hash_to_field_goldilocks(hash: usize,hlen: usize,dst: &[u8],msg: &[u8],ctr: usize) -> [core::goldilocks::fp::FP;2] {

    use core::goldilocks::big::BIG;
    use core::goldilocks::dbig::DBIG;
    use core::goldilocks::fp::FP;
    use core::goldilocks::rom;

    let mut u: [FP; 2] = [FP::new(),FP::new()];

    let q = BIG::new_ints(&rom::MODULUS);
    let k=q.nbits();
    let r = BIG::new_ints(&rom::CURVE_ORDER);
    let m=r.nbits();
    let L=ceil(k+ceil(m,2),8);
    let mut okm: [u8;512]=[0;512];
    hmac::xmd_expand(hash,hlen,&mut okm,L*ctr,&dst,&msg);
    let mut fd: [u8;256]=[0;256];
    for i in 0..ctr {
        for j in 0..L {
            fd[j]=okm[i*L+j];
        }
		let mut dx=DBIG::frombytes(&fd[0..L]);
		let w=FP::new_big(&dx.dmod(&q));
		u[i].copy(&w);
    }

    return u;
}

fn htp_goldilocks(mess: &str) {
    use core::goldilocks::ecp;
    use core::goldilocks::ecp::ECP;
    println!("Random Access - message= {}",mess);
    let m = mess.as_bytes();
    let mut dst = "edwards448_XMD:SHA-512_ELL2_RO_TESTGEN".as_bytes();
    let mut u=hash_to_field_goldilocks(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,2);
    println!("u[0]= {}",u[0].tostring());
    println!("u[1]= {}",u[1].tostring());
    let mut P=ECP::map2point(&u[0]);
    println!("Q[0]= {}",P.tostring());
    let P1=ECP::map2point(&u[1]);
    println!("Q[1]= {}",P1.tostring());
    P.add(&P1);
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());

    println!("\nNon-Uniform");
    dst = "edwards448_XMD:SHA-512_ELL2_NU_TESTGEN".as_bytes();
    u=hash_to_field_goldilocks(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,1);
    println!("u[0]= {}",u[0].tostring());
    P=ECP::map2point(&u[0]);
    println!("Q= {}",P.tostring());
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());
}

fn hash_to_field_secp256k1(hash: usize,hlen: usize,dst: &[u8],msg: &[u8],ctr: usize) -> [core::secp256k1::fp::FP;2] {

    use core::secp256k1::big::BIG;
    use core::secp256k1::dbig::DBIG;
    use core::secp256k1::fp::FP;
    use core::secp256k1::rom;

    let mut u: [FP; 2] = [FP::new(),FP::new()];

    let q = BIG::new_ints(&rom::MODULUS);
    let k=q.nbits();
    let r = BIG::new_ints(&rom::CURVE_ORDER);
    let m=r.nbits();
    let L=ceil(k+ceil(m,2),8);
    let mut okm: [u8;512]=[0;512];
    hmac::xmd_expand(hash,hlen,&mut okm,L*ctr,&dst,&msg);
    let mut fd: [u8;256]=[0;256];
    for i in 0..ctr {
        for j in 0..L {
            fd[j]=okm[i*L+j];
        }
		let mut dx=DBIG::frombytes(&fd[0..L]);
		let w=FP::new_big(&dx.dmod(&q));
		u[i].copy(&w);
    }

    return u;
}

fn htp_secp256k1(mess: &str) {
    use core::secp256k1::ecp;
    use core::secp256k1::ecp::ECP;
    println!("Random Access - message= {}",mess);
    let m = mess.as_bytes();
    let mut dst = "secp256k1_XMD:SHA-256_SVDW_RO_TESTGEN".as_bytes();
    let mut u=hash_to_field_secp256k1(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,2);
    println!("u[0]= {}",u[0].tostring());
    println!("u[1]= {}",u[1].tostring());
    let mut P=ECP::map2point(&u[0]);
    println!("Q[0]= {}",P.tostring());
    let P1=ECP::map2point(&u[1]);
    println!("Q[1]= {}",P1.tostring());
    P.add(&P1);
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());

    println!("\nNon-Uniform");
    dst = "secp256k1_XMD:SHA-256_SVDW_NU_TESTGEN".as_bytes();
    u=hash_to_field_secp256k1(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,1);
    println!("u[0]= {}",u[0].tostring());
    P=ECP::map2point(&u[0]);
    println!("Q= {}",P.tostring());
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());
}

fn hash_to_field_bls12381(hash: usize,hlen: usize,dst: &[u8],msg: &[u8],ctr: usize) -> [core::bls12381::fp::FP;2] {

    use core::bls12381::big::BIG;
    use core::bls12381::dbig::DBIG;
    use core::bls12381::fp::FP;
    use core::bls12381::rom;

    let mut u: [FP; 2] = [FP::new(),FP::new()];

    let q = BIG::new_ints(&rom::MODULUS);
    let k=q.nbits();
    let r = BIG::new_ints(&rom::CURVE_ORDER);
    let m=r.nbits();
    let L=ceil(k+ceil(m,2),8);
    let mut okm: [u8;512]=[0;512];
    hmac::xmd_expand(hash,hlen,&mut okm,L*ctr,&dst,&msg);
    let mut fd: [u8;256]=[0;256];
    for i in 0..ctr {
        for j in 0..L {
            fd[j]=okm[i*L+j];
        }
		let mut dx=DBIG::frombytes(&fd[0..L]);
		let w=FP::new_big(&dx.dmod(&q));
		u[i].copy(&w);
    }

    return u;
}

fn htp_bls12381(mess: &str) {
    use core::bls12381::ecp;
    use core::bls12381::ecp::ECP;
    println!("Random Access - message= {}",mess);
    let m = mess.as_bytes();
    let mut dst = "BLS12381G1_XMD:SHA-256_SVDW_RO_TESTGEN".as_bytes();
    let mut u=hash_to_field_bls12381(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,2);
    println!("u[0]= {}",u[0].tostring());
    println!("u[1]= {}",u[1].tostring());
    let mut P=ECP::map2point(&u[0]);
    println!("Q[0]= {}",P.tostring());
    let P1=ECP::map2point(&u[1]);
    println!("Q[1]= {}",P1.tostring());
    P.add(&P1);
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());

    println!("\nNon-Uniform");
    dst = "BLS12381G1_XMD:SHA-256_SVDW_NU_TESTGEN".as_bytes();
    u=hash_to_field_bls12381(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,1);
    println!("u[0]= {}",u[0].tostring());
    P=ECP::map2point(&u[0]);
    println!("Q= {}",P.tostring());
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());
}

fn hash_to_field2_bls12381(hash: usize,hlen: usize,dst: &[u8],msg: &[u8],ctr: usize) -> [core::bls12381::fp2::FP2;2] {

    use core::bls12381::big::BIG;
    use core::bls12381::dbig::DBIG;
    use core::bls12381::fp::FP;
    use core::bls12381::fp2::FP2;
    use core::bls12381::rom;

    let mut u: [FP2; 2] = [FP2::new(),FP2::new()];

    let q = BIG::new_ints(&rom::MODULUS);
    let k=q.nbits();
    let r = BIG::new_ints(&rom::CURVE_ORDER);
    let m=r.nbits();
    let L=ceil(k+ceil(m,2),8);
    let mut okm: [u8;512]=[0;512];
    hmac::xmd_expand(hash,hlen,&mut okm,2*L*ctr,&dst,&msg);
    let mut fd: [u8;256]=[0;256];
    for i in 0..ctr {
        for j in 0..L {
            fd[j]=okm[2*i*L+j];
        }
		let mut dx=DBIG::frombytes(&fd[0..L]);
		let w1=FP::new_big(&dx.dmod(&q));

        for j in 0..L {
            fd[j]=okm[(2*i+1)*L+j];
        }
		dx=DBIG::frombytes(&fd[0..L]);
		let w2=FP::new_big(&dx.dmod(&q));
		u[i].copy(&FP2::new_fps(&w1,&w2));
    }

    return u;
}

fn htp2_bls12381(mess: &str) {
    use core::bls12381::ecp;
    use core::bls12381::ecp2::ECP2;
    println!("Random Access - message= {}",mess);
    let m = mess.as_bytes();
    let mut dst = "BLS12381G2_XMD:SHA-256_SVDW_RO_TESTGEN".as_bytes();
    let mut u=hash_to_field2_bls12381(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,2);
    println!("u[0]= {}",u[0].tostring());
    println!("u[1]= {}",u[1].tostring());
    let mut P=ECP2::map2point(&u[0]);
    println!("Q[0]= {}",P.tostring());
    let P1=ECP2::map2point(&u[1]);
    println!("Q[1]= {}",P1.tostring());
    P.add(&P1);
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());

    println!("\nNon-Uniform");
    dst = "BLS12381G2_XMD:SHA-256_SVDW_NU_TESTGEN".as_bytes();
    u=hash_to_field2_bls12381(hmac::MC_SHA2,ecp::HASH_TYPE,dst,m,1);
    println!("u[0]= {}",u[0].tostring());
    P=ECP2::map2point(&u[0]);
    println!("Q= {}",P.tostring());
    P.cfp();
    P.affine();
    println!("P= {}",P.tostring());
}

fn main() {
    println!("\nTesting HTP for curve ed25519\n");
    htp_ed25519("");
    htp_ed25519("abc");
    htp_ed25519("abcdef0123456789");
    htp_ed25519("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");

    println!("\nTesting HTP for curve nist256\n");
    htp_nist256("");
    htp_nist256("abc");
    htp_nist256("abcdef0123456789");
    htp_nist256("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");

    println!("\nTesting HTP for curve goldilocks\n");
    htp_goldilocks("");
    htp_goldilocks("abc");
    htp_goldilocks("abcdef0123456789");
    htp_goldilocks("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");

    println!("\nTesting HTP for curve secp256k1\n");
    htp_secp256k1("");
    htp_secp256k1("abc");
    htp_secp256k1("abcdef0123456789");
    htp_secp256k1("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");

    println!("\nTesting HTP for curve bls12381_G1\n");
    htp_bls12381("");
    htp_bls12381("abc");
    htp_bls12381("abcdef0123456789");
    htp_bls12381("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");

    println!("\nTesting HTP for curve bls12381_G2\n");
    htp2_bls12381("");
    htp2_bls12381("abc");
    htp2_bls12381("abcdef0123456789");
    htp2_bls12381("a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");

}