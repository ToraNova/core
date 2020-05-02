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

/* Fixed Data in ROM - Field and Curve parameters */


package org.miracl.core.C25519;

public class ROM
{
// Base Bits= 29
public static final int[] Modulus= {0x1FFFFFED,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x7FFFFF};
public static final int[] ROI= {0xA0EA0B0,0x770D93A,0xBF91E31,0x6300D5A,0x1D7A72F4,0x4C9EFD,0x1C2CAD34,0x1009F83B,0x2B8324};
public static final int[] R2modp= {0x169000,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0};
public static final int MConst= 0x13;

public static final int CURVE_Cof_I= 8;
public static final int[] CURVE_Cof= {0x8,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0};
public static final int CURVE_A= 486662;
public static final int CURVE_B_I= 0;
public static final int[] CURVE_B= {0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0};
public static final int[] CURVE_Order= {0x1CF5D3ED,0x9318D2,0x1DE73596,0x1DF3BD45,0x14D,0x0,0x0,0x0,0x100000};
public static final int[] CURVE_Gx= {0x9,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0};
public static final int[] CURVE_Gy= {0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0};

}
