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

//
//  rom.swift
//
//  Created by Michael Scott on 12/06/2015.
//  Copyright (c) 2015 Michael Scott. All rights reserved.
//

public struct ROM{
 
#if D32


// Base Bits= 28
//  bls461 Curve Modulus


static let Modulus:[Chunk] = [0xAAAAAAB,0xAC0000A,0x54AAAAA,0x5555,0x400020,0x91557F0,0xF26AA,0xFA5C1CC,0xB42A8DF,0x7B14848,0x8BACCA4,0x6F1E32D,0x4935FBD,0x55D6941,0xD5A555A,0x5545554,0x1555]
static let ROI:[Chunk] = [0xAAAAAAA,0xAC0000A,0x54AAAAA,0x5555,0x400020,0x91557F0,0xF26AA,0xFA5C1CC,0xB42A8DF,0x7B14848,0x8BACCA4,0x6F1E32D,0x4935FBD,0x55D6941,0xD5A555A,0x5545554,0x1555]
static let R2modp:[Chunk] = [0xC9B6A33,0x2ECD087,0x3CCB2B1,0xCD461FE,0x8CB5AB2,0xC5B9635,0x5312E92,0xB659F64,0x3B596FA,0x8679006,0xA92E2B3,0x3CE05E3,0x363550F,0x7C07A8E,0x382C083,0x6347FEA,0xBD]
static let MConst:Chunk = 0xFFFFFFD
static public let Fra:[Chunk] = [0xB812A3A,0x7117BF9,0x99C400F,0xC6308A5,0x5BF8A1,0x510E075,0x45FA5A6,0xCE4858D,0x770B31A,0xBC2CB04,0xE2FC61E,0xD073588,0x4366190,0x4DFEFA8,0x69E55E2,0x504B7F,0x12E4]
static public let Frb:[Chunk] = [0xF298071,0x3AE8410,0xBAE6A9B,0x39D4CAF,0xFE4077E,0x404777A,0xBAF8104,0x2C13C3E,0x3D1F5C5,0xBEE7D44,0xA8B0685,0x9EAADA4,0x5CFE2C,0x7D7999,0x6BBFF78,0x50409D5,0x271]

// bls461 Curve

static let CURVE_Cof_I:Int = 0
static let CURVE_A:Int = 0
static let CURVE_B_I:Int = 9
static public let CURVE_B:[Chunk] = [0x9,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]
static public let CURVE_Order:[Chunk] = [0x1,0x0,0xFFFFC00,0x7FEFFFE,0x110000,0x7FFC800,0x801FC01,0x5FD000E,0x17FE0,0xFFFC018,0xFFFFFF7,0x0,0x0,0x0,0x0,0x0,0x0]
static public let CURVE_Gx:[Chunk] = [0xADEE93D,0x4D026A8,0x74B7411,0xD9C00EE,0x31AC7F2,0xC3981B5,0x9218229,0xD3564DC,0xA096650,0x6F7C292,0x9743616,0xBE922B1,0x12CF668,0xC81327,0x463B73A,0xE74E99B,0xAD0]
static public let CURVE_Gy:[Chunk] = [0xAD1D465,0xF763157,0xC4FF470,0x17884C8,0xB8D215D,0xA819E66,0xF4959D0,0xE5C3245,0xB84910A,0xB8BFA40,0xBE96EEC,0x8BF9F8C,0xF277ACC,0x5F1C3F2,0x5F68C9,0xCDB14B3,0x77B]

static public let CURVE_Bnx:[Chunk] = [0x0,0xFBFFFE0,0x1FFFFF,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]

static public let CURVE_Cof:[Chunk] = [0x1,0xFBFFFE0,0x1FFFFF,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]
//static public let CURVE_Cof:[Chunk] = [0xAAAAAAB,0xA7FFFEA,0x1556AA,0xD55AAAB,0x554FFFF,0x1555,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]
static public let CURVE_Cru:[Chunk] = [0xFFFFFFE,0x40001F,0xFE00000,0xFFE7FFF,0xF0FFF6F,0x7200C47,0x7BCC604,0x15796DB,0xCF47771,0x9875433,0x613F0E8,0x5000502,0xEBFFF60,0x1FFFFF,0x0,0x0,0x0]
static public let CURVE_Pxa:[Chunk] = [0x6D0A37C,0x5B50318,0x75DCC46,0xC2E492E,0xD6878A9,0xE01F919,0xF92F564,0x86DB74F,0x66803F0,0x46D581A,0x7ED78D,0x2F97C29,0xC270C89,0xF679453,0x6A50A9A,0x54138A0,0x10CC]
static public let CURVE_Pxb:[Chunk] = [0x2C1C0AD,0xF85CA8C,0x25CADE9,0x6CD66C4,0xA289609,0xC612951,0xEE2401A,0x529ABEB,0xF65B17D,0xBA09D33,0xD4C5AF5,0x4D4371E,0x46A672E,0xA279D22,0xACEA37C,0x1FB4FE5,0x95C]
static public let CURVE_Pya:[Chunk] = [0x2FB006,0xCCD0C1B,0xA12A337,0x3D194A4,0xC92C895,0x4960CFC,0x39FC68B,0x3A9B00F,0xED1BA0F,0xA7DBBC5,0xA9CDFD8,0x27CC2F7,0x4E73ED2,0x6070F4F,0xEBA7E67,0xAC848E7,0x226]
static public let CURVE_Pyb:[Chunk] = [0xDF1457C,0xA506ADF,0x4C20A8,0xD6A31DC,0x36E3FB4,0xEA9A8F1,0x92F5668,0x3C3BE44,0x67A1297,0x74BEABA,0x56A20BE,0x4C42E38,0x45157F0,0x2AB1D00,0xBB402EA,0x101B4FA,0xE38]
static let CURVE_W:[[Chunk]] = [[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]]
static let CURVE_SB:[[[Chunk]]] = [[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]],[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]]] 
static let CURVE_WB:[[Chunk]] = [[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]]
static let CURVE_BB:[[[Chunk]]] = [[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]],[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]],[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]],[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]]]

#endif

#if D64

// Base Bits= 60
// bls461 Modulus

static let Modulus:[Chunk] = [0xAAC0000AAAAAAAB,0x20000555554AAAA,0x6AA91557F004000,0xA8DFFA5C1CC00F2,0xACCA47B14848B42,0x935FBD6F1E32D8B,0xD5A555A55D69414,0x15555545554]
static let ROI:[Chunk] = [0xAAC0000AAAAAAAA,0x20000555554AAAA,0x6AA91557F004000,0xA8DFFA5C1CC00F2,0xACCA47B14848B42,0x935FBD6F1E32D8B,0xD5A555A55D69414,0x15555545554]
static let R2modp:[Chunk] = [0x96D08774614DDA8,0xCD45F539225D5BD,0xD712EB760C95AB1,0xB3B687155F30B55,0xC4E62A05C3F5B81,0xBA1151676CA3CD0,0x7EDD8A958F442BE,0x12B89DD3F91]
static let MConst:Chunk = 0xC0005FFFFFFFD
static public let Fra:[Chunk] = [0xF7117BF9B812A3A,0xA1C6308A599C400,0x5A6510E07505BF8,0xB31ACE4858D45FA,0xFC61EBC2CB04770,0x366190D073588E2,0x69E55E24DFEFA84,0x12E40504B7F]
static public let Frb:[Chunk] = [0xB3AE8410F298071,0x7E39D4CAFBAE6A9,0x104404777AFE407,0xF5C52C13C3EBAF8,0xB0685BEE7D443D1,0x5CFE2C9EAADA4A8,0x6BBFF7807D79990,0x27150409D5]

//  bls461 Curve

static let CURVE_Cof_I:Int = 0
static let CURVE_A:Int = 0
static let CURVE_B_I:Int = 9
static public let CURVE_B:[Chunk] = [0x9,0x0,0x0,0x0,0x0,0x0,0x0,0x0]
static public let CURVE_Order:[Chunk] = [0x1,0x7FEFFFEFFFFC0,0xC017FFC80001100,0x7FE05FD000E801F,0xFFFF7FFFC018001,0xFF,0x0,0x0]
static public let CURVE_Gx:[Chunk] = [0x14D026A8ADEE93D,0xF2D9C00EE74B741,0x229C3981B531AC7,0x6650D3564DC9218,0x436166F7C292A09,0x2CF668BE922B197,0x463B73A0C813271,0xAD0E74E99B]
static public let CURVE_Gy:[Chunk] = [0xF763157AD1D465,0x5D17884C8C4FF47,0x9D0A819E66B8D21,0x910AE5C3245F495,0x96EECB8BFA40B84,0x277ACC8BF9F8CBE,0x5F68C95F1C3F2F,0x77BCDB14B3]

static public let CURVE_Bnx:[Chunk] = [0xFFBFFFE00000000,0x1FFFF,0x0,0x0,0x0,0x0,0x0,0x0]
static public let CURVE_Cof:[Chunk] = [0xFFBFFFE00000001,0x1FFFF,0x0,0x0,0x0,0x0,0x0,0x0]
//static public let CURVE_Cof:[Chunk] = [0xAA7FFFEAAAAAAAB,0xFFD55AAAB01556A,0x1555554FF,0x0,0x0,0x0,0x0,0x0]
static public let CURVE_Cru:[Chunk] = [0x40001FFFFFFFE,0x6FFFE7FFFFE0000,0x6047200C47F0FFF,0x777115796DB7BCC,0x3F0E89875433CF4,0xBFFF60500050261,0x1FFFFFE,0x0] 
static public let CURVE_Pxa:[Chunk] = [0x65B503186D0A37C,0xA9C2E492E75DCC4,0x564E01F919D6878,0x3F086DB74FF92F,0xED78D46D581A668,0x270C892F97C2907,0x6A50A9AF679453C,0x10CC54138A0]
static public let CURVE_Pxb:[Chunk] = [0x9F85CA8C2C1C0AD,0x96CD66C425CADE,0x1AC612951A2896,0xB17D529ABEBEE24,0xC5AF5BA09D33F65,0x6A672E4D4371ED4,0xACEA37CA279D224,0x95C1FB4FE5]
static public let CURVE_Pya:[Chunk] = [0x7CCD0C1B02FB006,0x953D194A4A12A33,0x68B4960CFCC92C8,0xBA0F3A9B00F39FC,0xCDFD8A7DBBC5ED1,0xE73ED227CC2F7A9,0xEBA7E676070F4F4,0x226AC848E7]
static public let CURVE_Pyb:[Chunk] = [0x8A506ADFDF1457C,0xB4D6A31DC04C20A,0x668EA9A8F136E3F,0x12973C3BE4492F5,0xA20BE74BEABA67A,0x5157F04C42E3856,0xBB402EA2AB1D004,0xE38101B4FA]
static let CURVE_W:[[Chunk]] = [[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]]
static let CURVE_SB:[[[Chunk]]] = [[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]],[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]]]
static let CURVE_WB:[[Chunk]] = [[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]]
static let CURVE_BB:[[[Chunk]]] = [[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]],[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]],[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]],[[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0],[0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0]]]

#endif

}

