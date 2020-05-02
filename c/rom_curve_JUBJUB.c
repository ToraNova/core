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
#include "ecp_JUBJUB.h"

/* JUBJUB Curve  */

#if CHUNK==16

#error Not supported

#endif

#if CHUNK==32

const int CURVE_A_JUBJUB= -1;
const int CURVE_Cof_I_JUBJUB= 8;
const BIG_256_29 CURVE_Cof_JUBJUB= {0x8,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0};
const int CURVE_B_I_JUBJUB= 0;
const BIG_256_29 CURVE_B_JUBJUB= {0x16343EB1,0x832FEB6,0x15E74980,0x1AFEDA6E,0x17FD4292,0x903F35E,0xD23D7F6,0x1CE97F45,0x2A9318};
const BIG_256_29 CURVE_Order_JUBJUB= {0x16F72CB7,0x4B872F6,0x120420B4,0x10412799,0x3B00A66,0x1D80809A,0x1EA4199C,0x1D4CA675,0xE7DB4};
const BIG_256_29 CURVE_Gx_JUBJUB= {0x14F976C4,0x1A7678D3,0x1CE7B79F,0x18A8D7E5,0x4882000,0x1A6F801C,0xE329892,0x55F1DFE,0x518397};
const BIG_256_29 CURVE_Gy_JUBJUB= {0x1349702E,0x35D6B84,0x9A0CEC5,0x190E0FFF,0x1C308096,0x62D5ECE,0x10B27A3F,0x8E5945F,0x3B43F8};
#endif

#if CHUNK==64

const int CURVE_A_JUBJUB= -1;
const int CURVE_Cof_I_JUBJUB= 8;
const BIG_256_56 CURVE_Cof_JUBJUB= {0x8L,0x0L,0x0L,0x0L,0x0L};
const int CURVE_B_I_JUBJUB= 0;
const BIG_256_56 CURVE_B_JUBJUB= {0x65FD6D6343EB1L,0x7F6D37579D2601L,0x7E6BD7FD4292DL,0x4BFA2B48F5FD92L,0x2A9318E7L};
const BIG_256_56 CURVE_Order_JUBJUB= {0x970E5ED6F72CB7L,0x2093CCC81082D0L,0x101343B00A668L,0x6533AFA906673BL,0xE7DB4EAL};
const BIG_256_56 CURVE_Gx_JUBJUB= {0x4ECF1A74F976C4L,0x546BF2F39EDE7FL,0xDF00384882000CL,0xF8EFF38CA624B4L,0x5183972AL};
const BIG_256_56 CURVE_Gy_JUBJUB= {0x6BAD709349702EL,0x8707FFA6833B14L,0x5ABD9DC308096CL,0x2CA2FC2C9E8FCCL,0x3B43F847L};

#endif

