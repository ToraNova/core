/*
 * Copyright (c) 2012-2020 MIRACL UK Ltd.
 *
 * This file is part of MIRACL Core
 * (see https://github.com/miracl/core).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* Boneh-Lynn-Shacham signature API Functions */

/* Loosely (for now) following https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-bls-signature-00 */

// Minimal-signature-size variant

package org.miracl.core.XXX;

import org.miracl.core.RAND;
import org.miracl.core.HMAC;

public class BLS192 {
    public static final int BFS = CONFIG_BIG.MODBYTES;
    public static final int BGS = CONFIG_BIG.MODBYTES;
    public static final int BLS_OK = 0;
    public static final int BLS_FAIL = -1;

    public static FP8[] G2_TAB;

    static int ceil(int a,int b) {
        return (((a)-1)/(b)+1);
    }

    static FP[] hash_to_field(int hash,int hlen,byte[] DST,byte[] M,int ctr) {
        BIG q = new BIG(ROM.Modulus);
        int L = ceil(q.nbits()+CONFIG_CURVE.AESKEY*8,8);
        FP [] u = new FP[ctr];
        byte[] fd=new byte[L];

        byte[] OKM=HMAC.XMD_Expand(hash,hlen,L*ctr,DST,M);
        for (int i=0;i<ctr;i++)
        {
            for (int j=0;j<L;j++)
                fd[j]=OKM[i*L+j];
            u[i]=new FP(DBIG.fromBytes(fd).mod(q));
        }
    
        return u;
    }    

/* output u \in F_p 
    static BIG hash_to_base(int hash,int hlen,byte[] DST,byte[] M,int ctr) {
        BIG q = new BIG(ROM.Modulus);
        int L = ceil(q.nbits()+CONFIG_CURVE.AESKEY*8,8);

        String tag= new String("H2C");
        byte[] TAG=tag.getBytes();
        byte[] INFO=new byte[TAG.length+1];
        for (int i=0;i<TAG.length;i++)
            INFO[i]=TAG[i];
        INFO[TAG.length]=(byte)ctr;

        byte[] PRK=HMAC.HKDF_Extract(hash,hlen,DST,M);
        byte[] OKM=HMAC.HKDF_Expand(hash,hlen,L,PRK,INFO);

        DBIG dx=DBIG.fromBytes(OKM);
        BIG u=dx.mod(q);
        return u;
    } */

    /* hash a message to an ECP point, using SHA2, random oracle method */
    public static ECP bls_hash_to_point(byte[] M) {
        String dst= new String("BLS_SIG_ZZZG1_XMD:SHA384-SSWU-RO-_NUL_");
        FP[] u=hash_to_field(HMAC.MC_SHA2,CONFIG_CURVE.HASH_TYPE,dst.getBytes(),M,2);

        ECP P=ECP.map2point(u[0]);
        ECP P1=ECP.map2point(u[1]);
        P.add(P1);
        P.cfp();
        P.affine();
        return P;
    }

    public static int init() {
        ECP4 G = ECP4.generator();
        if (G.is_infinity()) return BLS_FAIL;
        G2_TAB = PAIR192.precomp(G);
        return BLS_OK;
    }

    /* generate key pair, private key S, public key W */
    public static int KeyPairGenerate(byte[] IKM, byte[] S, byte[] W) {
        BIG r = new BIG(ROM.CURVE_Order);     
        int L = ceil(3*ceil(r.nbits(),8),2);
        ECP4 G = ECP4.generator();
        String salt=new String("BLS-SIG-KEYGEN-SALT-");
        String info=new String("");
        byte[] PRK=HMAC.HKDF_Extract(HMAC.MC_SHA2,CONFIG_CURVE.HASH_TYPE,salt.getBytes(),IKM);
        byte[] OKM=HMAC.HKDF_Expand(HMAC.MC_SHA2,CONFIG_CURVE.HASH_TYPE,L,PRK,info.getBytes());

        DBIG dx=DBIG.fromBytes(OKM);
        BIG s=dx.mod(r);

        s.toBytes(S);
        G = PAIR192.G2mul(G, s);
        G.toBytes(W,true);
        return BLS_OK;
    }

    /* Sign message M using private key S to produce signature SIG */
    public static int core_sign(byte[] SIG, byte[] M, byte[] S) {
        ECP D = bls_hash_to_point(M);
        BIG s = BIG.fromBytes(S);
        D = PAIR192.G1mul(D, s);
        D.toBytes(SIG, true);
        return BLS_OK;
    }

    /* Verify signature given message m, the signature SIG, and the public key W */

    public static int core_verify(byte[] SIG, byte[] M, byte[] W) {
        ECP HM = bls_hash_to_point(M);

        ECP D = ECP.fromBytes(SIG);
        if (!PAIR192.G1member(D)) return BLS_FAIL;
        D.neg();

        ECP4 PK = ECP4.fromBytes(W);

// Use new multi-pairing mechanism
        FP24[] r = PAIR192.initmp();
        PAIR192.another_pc(r, G2_TAB, D);
        PAIR192.another(r, PK, HM);
        FP24 v = PAIR192.miller(r);

//.. or alternatively
//		ECP4 G=ECP4.generator();
//		if (G.is_infinity()) return BLS_FAIL;
//		FP24 v=PAIR192.ate2(G,D,PK,HM);

        v = PAIR192.fexp(v);
        if (v.isunity())
            return BLS_OK;
        return BLS_FAIL;
    }
}
