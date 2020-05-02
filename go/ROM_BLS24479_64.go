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

package BLS24479

// Base Bits= 56
var Modulus = [...]Chunk{0x44C1674A06152B, 0xFFE2E82D30DAF8, 0x6F1C5CBDB6A642, 0x3220DF068A328B, 0xE09E1F24406187, 0xBA825079733568, 0x6E803F2E77E4C1, 0x3CCC5BA839AEC, 0x555C0078}
var ROI = [...]Chunk{0x44C1674A06152A, 0xFFE2E82D30DAF8, 0x6F1C5CBDB6A642, 0x3220DF068A328B, 0xE09E1F24406187, 0xBA825079733568, 0x6E803F2E77E4C1, 0x3CCC5BA839AEC, 0x555C0078}
var R2modp = [...]Chunk{0x6A4A1FE013DF5B, 0xE8E46D4D1BDE65, 0x1F841391F45C67, 0x9148A4516FB28, 0x4398524EDF4C88, 0x41C0E241B6DCE8, 0xE42C208C19411, 0xA7FE6FD73A7B1C, 0xFCCCA76}

const MConst Chunk = 0xBD5D7D8095FE7D

var Fra = [...]Chunk{0x5CA74ABBF96F1D, 0x1FF8BD0C6FFBAD, 0x49E9E26237469C, 0x3CECA48407F8E5, 0x69D68FF59267B7, 0x5D199E33127CBD, 0xB97549184F313A, 0x4E77242DA52D8D, 0x4BBC87B9}
var Frb = [...]Chunk{0xE81A1C8E0CA60E, 0xDFEA2B20C0DF4A, 0x25327A5B7F5FA6, 0xF5343A828239A6, 0x76C78F2EADF9CF, 0x5D68B24660B8AB, 0xB50AF61628B387, 0xB555A18CDE6D5E, 0x99F78BE}

const CURVE_Cof_I int = 0
const CURVE_A int = 0
const CURVE_B_I int = 19

var CURVE_B = [...]Chunk{0x13, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Order = [...]Chunk{0x1A08FFF0000001, 0x1E7033FF551190, 0x6ADE7EE322DDAF, 0x848FC9D0CED13A, 0x50D81729CC224, 0x1F0F05B98BB44A, 0x10010010005A0, 0x0, 0x0}
var CURVE_Gx = [...]Chunk{0x6760F5EBE3CCD4, 0xEFE2DAED9F4564, 0x783F08EBA1FCC1, 0xC6F8D95AF88134, 0xDCA8D1AE2D8477, 0x9077586CEFE4BF, 0x8B7FEA5D99BC1D, 0x17CAF9486DE9E1, 0x1AB2BE34}
var CURVE_Gy = [...]Chunk{0xCBA5CAD21E5245, 0x6D6608C55DF6C4, 0xB3ED294F39746B, 0x145824920FF3C8, 0x63AA4FD63E5A64, 0x492A2BF79CE00F, 0x66A7A4529FF79A, 0x6C53E477B861CA, 0x47FCB70C}

var CURVE_Bnx = [...]Chunk{0x100020011FF80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Cof= [...]Chunk {0x100020011FF7F,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}
//var CURVE_Cof = [...]Chunk{0xC1FFBFF9F415AB, 0x5556AAB7FF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
var CURVE_Cru = [...]Chunk{0xBC27146DD794A9, 0x3A30938AF33A43, 0xB112175223DDC6, 0x125CFBB4236DFB, 0x2358E379CE607, 0xD680C6EB20806E, 0x314C200860FF77, 0x3CBC5A88268E4, 0x555C0078}
var CURVE_Pxaa = [...]Chunk{0xE2935374E24678, 0xC34342582408B, 0xF765CCDEFC69E, 0xC33AAD2888D7F9, 0x7FD2458967473A, 0x52908ED55CBAB3, 0x786671EB14AB88, 0xA3EC96077958C8, 0x959DE53}
var CURVE_Pxab = [...]Chunk{0x7F9EBAFFB099B8, 0x3775A012A47038, 0x6B5D1B1FC23856, 0x7F0A26A730F9E3, 0x1C38F85DB2A5CA, 0x76A753E17E6926, 0x2D39D1BE5AD0F9, 0x31733DFC651E4C, 0x3B0DED08}
var CURVE_Pxba = [...]Chunk{0xA1CDE711AD15D3, 0x853178DF6E16ED, 0x64BF43EA3E09A1, 0x2D8CD6DE566B2F, 0xF21C26C74FDB8B, 0x47BCC89E3F6B1E, 0x3FE2103F329F00, 0x4E507AF2AA28C3, 0x3EC27FAD}
var CURVE_Pxbb = [...]Chunk{0x7AB2875EE0F480, 0x4556E43D6C4B8C, 0xFB22DF80E1CB99, 0xF70FD0122F1FFD, 0xD5DB25698EF5EA, 0x4805CE1AF1BA3A, 0x1DA7CE2E465CB7, 0xCA0799F7E65855, 0xA5B38DB}
var CURVE_Pyaa = [...]Chunk{0x86499314781AA0, 0x609DA303B70AB1, 0xA52A6145FC44BB, 0x462E04C42A3124, 0xC383AE19AE68BB, 0xA1B34F6BE4FCAD, 0x198F901AD0BF4, 0x736C094362CED0, 0x5057F35D}
var CURVE_Pyab = [...]Chunk{0xBBEC57EEAE08FA, 0x78774BAA5F96AD, 0x64CAF099A42CA0, 0xC89FBBCCF70478, 0x6B720FEF855245, 0x97F916376F7B3E, 0x60F5587B5DF7E1, 0x61EE89637816BD, 0x2CE2B496}
var CURVE_Pyba = [...]Chunk{0x730276A5F0CC41, 0xF89325530AA1F5, 0xD9CD879AF8A147, 0xEE53E8A9FE2880, 0x420F07D3715390, 0x4C15D519B71F3A, 0x1A39DD3CB5B9B1, 0x3EE631A6BE39F8, 0x18070466}
var CURVE_Pybb = [...]Chunk{0xF1B2E6515C1CAE, 0xD40D355B0988DC, 0xC243FDC38A7772, 0x5D338136B675CA, 0x164E8A1D72FCDF, 0xBBAE5CD0961AC, 0xD6D04691771EB1, 0xD9BDEC8B792840, 0x499D14EA}
var CURVE_W = [2][9]Chunk{{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}
var CURVE_SB = [2][2][9]Chunk{{{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, {{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}}
var CURVE_WB = [4][9]Chunk{{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}
var CURVE_BB = [4][4][9]Chunk{{{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, {{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, {{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, {{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}}
