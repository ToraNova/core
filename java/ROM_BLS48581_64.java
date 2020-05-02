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


package org.miracl.core.BLS48581;

public class ROM
{

// Base Bits= 60
public static final long[] Modulus= {0xEDC154E6565912BL,0x8FDF721A4A48AC3L,0x7A5513170EE0A57L,0x394F4736DAF6836L,0xAF6E082ACD9CD30L,0xF3975444A48AE43L,0x22131BB3BE6C0F1L,0x12A0056E84F8D1L,0x76F313824E31D47L,0x1280F73FF34L};
public static final long[] ROI= {0xEDC154E6565912AL,0x8FDF721A4A48AC3L,0x7A5513170EE0A57L,0x394F4736DAF6836L,0xAF6E082ACD9CD30L,0xF3975444A48AE43L,0x22131BB3BE6C0F1L,0x12A0056E84F8D1L,0x76F313824E31D47L,0x1280F73FF34L};
public static final long[] R2modp= {0x79868479F1B5833L,0xFB6EBA8FCB82D07L,0x9CC8A7F1FD84C7FL,0x402C51CF5CC3CBBL,0x3F3114F078502CL,0xFC90829BDC8336EL,0xC7BE91DE9CA8EEDL,0xD4D273BB17BFADBL,0x6EC7C9A81E792CAL,0x1DC317A6E4L};
public static final long MConst= 0x148B81FC39D5A7DL;
public static final long[] Fra= {0x62EB6CFE42AEB25L,0xDB41942760AD3F9L,0xA7DF2570715ECE4L,0x90377B51208AC0FL,0x6848493E1C8C418L,0xF496307E298187EL,0x58740E3CAFD6B62L,0xF6067D047983E78L,0x49FA75CD7E73E55L,0xFD30DB501L};
public static final long[] Frb= {0x62EB6CFE42AEB25L,0xDB41942760AD3F9L,0xA7DF2570715ECE4L,0x90377B51208AC0FL,0x6848493E1C8C418L,0xF496307E298187EL,0x58740E3CAFD6B62L,0xF6067D047983E78L,0x49FA75CD7E73E55L,0xFD30DB501L};

//*** rom curve parameters *****
// Base Bits= 60
// Ate Bits= 33
// G2 Table size= 36

public static final int CURVE_A= 0;
public static final int CURVE_Cof_I= 0;
public static final long[] CURVE_Cof= {0x140000382L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L};
public static final int CURVE_B_I= 1;
public static final long[] CURVE_B= {0x1L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L};
public static final long[] CURVE_Order= {0x8A5FE6FCD671C01L,0xBE599467C24DA11L,0xC7CD0562303C4CCL,0x9D34C4C92016A85L,0xBC972C2E6E74196L,0x3F0B3CBE003FAD6L,0x615C0D6C635387AL,0xE2885E233A9CCC1L,0x2386F8A925L,0x0L};
public static final long[] CURVE_Gx= {0xBCE8732315AF640L,0x74DA5D3A1E6D8C3L,0x57DB368B11786CBL,0x665D859236EBDBCL,0x46A9DF6F9645847L,0xEDFFB9F75445505L,0xE86868CF61ABDBAL,0x93F860DE3F257E0L,0x40F2BAF2B73DF1EL,0x2AF59B7AC3L};
public static final long[] CURVE_Gy= {0xDBB5DE3E2587A70L,0xF37AEF7B926B576L,0xF77C2876D1B2E35L,0x78584C3EF22F487L,0xFFB98AEE53E80F6L,0xD41B720EF7BB7BEL,0xFEB8A52E991279DL,0xB398A488A553C9EL,0x31F91F86B3A2D1FL,0xCEFDA44F65L};

public static final long[] CURVE_Bnx= {0x140000381L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L};
public static final long[] CURVE_Cru= {0x4DE9AC5E1C79B90L,0x5CD8E3F88E5DE82L,0xAB21F74F7421A20L,0x6694B9B60DB5D62L,0x73422B5FB82F431L,0xFF46A846B5FA6AAL,0x83D66C1E5FCBED6L,0x2096384F2AFA565L,0x8B75055DD5D1F4EL,0x2C6L};
public static final long[] CURVE_Pxaaa= {0x34FD0B4ACE8BFABL,0xB79766322154DECL,0x4D80491F510317L,0x3CA0612F4005030L,0xBAAD1A8C42281A6L,0x3A2EF156C46FF79L,0x344DBCCB7DE64DBL,0x2775DEBABBEFC70L,0x71E4A38237FA45AL,0x5D615D9A78L};
public static final long[] CURVE_Pxaab= {0x669B36676B47C57L,0x5556A01AFA143F1L,0x7630D979630FFD7L,0x6AFFA62504F0C3CL,0xABFEDF16214A7L,0x12307F4E1C3943AL,0xE1623E9526F6DAL,0xBC07E8B22BB6D98L,0x258512069B0E86AL,0x7C4973ECE2L};
public static final long[] CURVE_Pxaba= {0x488156CA55A3E6AL,0xEF4CDED6B3F0B46L,0xCBDFBB879D5FEA8L,0x66F0D2A6D55F028L,0xC1DBD19242FFAE7L,0xCCBAB5AB6860161L,0xAE237CA7A6D6957L,0xAD83BC73A8A6CA9L,0xF1334E1B2EA1853L,0x1FCCC70198L};
public static final long[] CURVE_Pxabb= {0x9A7033CBB7FEAFEL,0x10B8CB4E80BC3F0L,0x1C5257C200CA523L,0x43B1B279B9468C3L,0x5F63E1C776E6EC1L,0x393F8BE0CC218A9L,0x62F3E5821B7B92AL,0x54D4BFE8F5985ACL,0xEB6185C78D80129L,0xBE2218C25CL};
public static final long[] CURVE_Pxbaa= {0x39C3A1C53F8CCE5L,0x5B5F746C9D4CBB7L,0xD55FC1889AA80C6L,0xEF492AE589274FAL,0x9E48199D5AC10B2L,0xC5805386699981FL,0xB1642B5675FF0E7L,0xA9DD63007C675D0L,0x35913A3C598E4CAL,0x38B91C600BL};
public static final long[] CURVE_Pxbab= {0x2004D914A3C093AL,0x7960910FCE3370FL,0xA9F177612F097FCL,0x40B9C0B15DD7595L,0x3835D28997EB57BL,0x7BB037418181DF6L,0xEF0977A3D1A5867L,0xCDA088F7B8F35DCL,0x738603F1311E4EL,0xC96C7797EBL};
public static final long[] CURVE_Pxbba= {0x41607E60750E057L,0x4B5B0E205C3354EL,0xCBE4324C22D6333L,0xAA5EFCF3432AAD1L,0xF293B13CED0FD0CL,0xA2C0B7A449CEF11L,0x9D13852B6DB908BL,0x8AEE660DEA41B3L,0x61EE3F0197A4989L,0xB9B7951C60L};
public static final long[] CURVE_Pxbbb= {0xE19DA00FBC6AE34L,0x6AF2FC9E97C3F84L,0x9BD6AEBF9FC44E5L,0x90B7E2B0D458547L,0xA93F29CFF364A71L,0x719728A7F9F8CFCL,0xFAF47B5211CF741L,0x4AAA2B1E5D7A9DEL,0x2BDEC5282624C4FL,0x827D5C22FBL};
public static final long[] CURVE_Pyaaa= {0x3EDD3FE4D2D7971L,0x45012AB12C0FF32L,0x9ABF77EEA6D6590L,0x336D8AE5163C159L,0x35AFA27748D90F7L,0xBFC435FAAB09062L,0x59A577E6F3B39EL,0x2F3024B918B4238L,0x75B5DFA49721645L,0xEB53356C3L};
public static final long[] CURVE_Pyaab= {0x1471DB936CD5665L,0x8B423525FFC7B11L,0x2FA097D760E2E58L,0xD1892AB24E1DD21L,0x6B243B1F192C5C3L,0x64732FCBF3AFB09L,0xA325E6FBA01D729L,0x5FCADC2B75A422BL,0xE0FF144DA653181L,0x284DC75979L};
public static final long[] CURVE_Pyaba= {0x8332A526A2A8474L,0xBC7C46FC3B8FDE6L,0x1D35D51A652269CL,0x36CA3295E5E2F0CL,0xC99D0E904115155L,0xD370514475F7D5L,0x216D5B119D3A48L,0x67669EF2C2FC503L,0x8523E421EFB703L,0xB36A201DD0L};
public static final long[] CURVE_Pyabb= {0x6213DA92841589DL,0xB3D8B8A1E533731L,0x7BDA503EE5E578FL,0x817742770BA10D6L,0x224333FA40DCED2L,0x10E122D2742C89BL,0x60DCEE23DD8B0E7L,0x78762B1C2CDED33L,0xEDC0688223FBBD4L,0xAEC25A4621L};
public static final long[] CURVE_Pybaa= {0x47831F982E50137L,0x857FDDDFCF7A43FL,0x30135945D137B08L,0xCA4E512B64F59F4L,0x7FA238CDCE8A1E2L,0x5F1129857ED85C7L,0xB43DD93B5A95980L,0x88325A2554DC541L,0xA9C46916503FA5AL,0xD209D5A223L};
public static final long[] CURVE_Pybab= {0x4EEDC58CF90BEE4L,0xA59ED8226CF3A59L,0xFC198CAA72B679DL,0xF47C180D139E3AAL,0xE8C270841F6824L,0x55AB7504FA8342L,0xB16722B589D82E2L,0xD537B90421AD66EL,0x36B7A513D339D5AL,0x7D0D037457L};
public static final long[] CURVE_Pybba= {0xD41FAEAFEB23986L,0xE884017D9AA62B3L,0x40FA639F53DCCC9L,0xAB8C74B2618B5BBL,0x5AE3A2864F22C1FL,0xE4C819A6DF98F42L,0xC0841B064155F14L,0xD17AF8A006F364FL,0xE65EA25C2D05DFDL,0x896767811BL};
public static final long[] CURVE_Pybbb= {0x667FFCB732718B6L,0x5AC66E84069C55DL,0xD8C4AB33F748EL,0x333EC7192054173L,0x8E69C31E97E1AD0L,0xEF8ECA9A9533A3FL,0x6BE8E50C87549B6L,0x4F981B5E068F140L,0x9029D393A5C07E8L,0x35E2524FF8L};
public static final long[][] CURVE_W= {{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L}};
public static final long[][][] CURVE_SB= {{{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L}},{{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L}}};
public static final long[][] CURVE_WB= {{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L}};
public static final long[][][] CURVE_BB= {{{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L}},{{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L}},{{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L}},{{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L},{0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L,0x0L}}};

}

