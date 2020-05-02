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
#include "arch.h"
#include "fp_FP512BN.h"

/* Curve FP512BN - Pairing friendly BN curve */

/* ISO fast curve */

#if CHUNK==16

#error Not supported

#endif

#if CHUNK==32
// Base Bits= 29

const BIG_512_29 Modulus_FP512BN= {0x2ADEF33,0x7594049,0x131919ED,0x14AB9CBE,0x16FE1916,0x12EF5591,0x2E39231,0x3D597D3,0x55146CF,0x88D877A,0x102EF8F0,0x1196A60F,0x1C60BA1D,0x1CF63F80,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x7FFFF};
const BIG_512_29 ROI_FP512BN= {0x2ADEF32,0x7594049,0x131919ED,0x14AB9CBE,0x16FE1916,0x12EF5591,0x2E39231,0x3D597D3,0x55146CF,0x88D877A,0x102EF8F0,0x1196A60F,0x1C60BA1D,0x1CF63F80,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x7FFFF};
const BIG_512_29 R2modp_FP512BN= {0xFD68B47,0xFCF5D2C,0x437675A,0x1BBC3FBF,0x1411E413,0x13453559,0x10B5639,0x1C34CE79,0x6D476BF,0xFD05F2B,0x15D17C28,0x6C9F76E,0x1C2375B3,0x78CCE9B,0x15F0AB33,0x1960F32E,0x1A8D44E,0x57A38};
const chunk MConst_FP512BN= 0x1CCC5C05;
const BIG_512_29 Fra_FP512BN= {0x14B73AB2,0x4B0BD8F,0xABB47D,0x2A29EC4,0x18681E17,0x104069DE,0x12EED67D,0x1553D0A5,0x398E9F8,0x7971034,0xAC9AF23,0x52DEF23,0x14EA18A5,0x1463E345,0x6DE465A,0x17F212B4,0x1AA9CF5B,0xF7B8};
const BIG_512_29 Frb_FP512BN= {0xDF6B481,0x2A882B9,0x126D6570,0x1208FDFA,0x1E95FAFF,0x2AEEBB2,0xFF4BBB4,0xE81C72D,0x1B85CD6,0xF67746,0x56549CD,0xC68B6EC,0x776A178,0x8925C3B,0x1921B9A5,0x80DED4B,0x55630A4,0x70847};



#endif

#if CHUNK==64
// Base Bits= 60

const BIG_512_60 Modulus_FP512BN= {0x4EB280922ADEF33L,0x6A55CE5F4C6467BL,0xC65DEAB236FE191L,0xCF1EACBE98B8E48L,0x3C111B0EF455146L,0xA1D8CB5307C0BBEL,0xFFFF9EC7F01C60BL,0xFFFFFFFFFFFFFFFL,0xFFFFFFFFL};
const BIG_512_60 ROI_FP512BN= {0x4EB280922ADEF32L,0x6A55CE5F4C6467BL,0xC65DEAB236FE191L,0xCF1EACBE98B8E48L,0x3C111B0EF455146L,0xA1D8CB5307C0BBEL,0xFFFF9EC7F01C60BL,0xFFFFFFFFFFFFFFFL,0xFFFFFFFFL};
const BIG_512_60 R2modp_FP512BN= {0x1FA6DCEF99812E9L,0xAB3452895A0B74EL,0xC53EA988C079E1EL,0x1E90E033BA630B9L,0xF1EA41C0714D8B0L,0xE72785387509E28L,0xD86794F834DAB00L,0x9757C2ACCD342A1L,0x44ECB079L};
const chunk MConst_FP512BN= 0x692A189FCCC5C05L;
const BIG_512_60 Fra_FP512BN= {0x49617B1F4B73AB2L,0x71514F6202AED1FL,0xF6080D3BD8681E1L,0xF8AA9E852CBBB59L,0xC8CF2E2068398E9L,0x8A5296F791AB26BL,0x196A8C7C68B4EA1L,0xCF5BBF9095A1B79L,0x1EF71AA9L};
const BIG_512_60 Frb_FP512BN= {0x5510572DF6B481L,0xF9047EFD49B595CL,0xD055DD765E95FAFL,0xD6740E396BFD2EEL,0x7341ECEE8C1B85CL,0x1786345B7615952L,0xE695124B876776AL,0x30A4406F6A5E486L,0xE108E556L};


#endif

