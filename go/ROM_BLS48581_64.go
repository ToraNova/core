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

/* Fixed Data in ROM - Field and Curve parameters */

package BLS48581

// Base Bits= 60
var Modulus= [...]Chunk {0xEDC154E6565912B,0x8FDF721A4A48AC3,0x7A5513170EE0A57,0x394F4736DAF6836,0xAF6E082ACD9CD30,0xF3975444A48AE43,0x22131BB3BE6C0F1,0x12A0056E84F8D1,0x76F313824E31D47,0x1280F73FF34}
var ROI= [...]Chunk {0xEDC154E6565912A,0x8FDF721A4A48AC3,0x7A5513170EE0A57,0x394F4736DAF6836,0xAF6E082ACD9CD30,0xF3975444A48AE43,0x22131BB3BE6C0F1,0x12A0056E84F8D1,0x76F313824E31D47,0x1280F73FF34}
var R2modp= [...]Chunk {0x79868479F1B5833,0xFB6EBA8FCB82D07,0x9CC8A7F1FD84C7F,0x402C51CF5CC3CBB,0x3F3114F078502C,0xFC90829BDC8336E,0xC7BE91DE9CA8EED,0xD4D273BB17BFADB,0x6EC7C9A81E792CA,0x1DC317A6E4}
const MConst Chunk=0x148B81FC39D5A7D
var Fra= [...]Chunk {0x62EB6CFE42AEB25,0xDB41942760AD3F9,0xA7DF2570715ECE4,0x90377B51208AC0F,0x6848493E1C8C418,0xF496307E298187E,0x58740E3CAFD6B62,0xF6067D047983E78,0x49FA75CD7E73E55,0xFD30DB501}
var Frb= [...]Chunk {0x62EB6CFE42AEB25,0xDB41942760AD3F9,0xA7DF2570715ECE4,0x90377B51208AC0F,0x6848493E1C8C418,0xF496307E298187E,0x58740E3CAFD6B62,0xF6067D047983E78,0x49FA75CD7E73E55,0xFD30DB501}

//*** rom curve parameters *****
// Base Bits= 60
// Ate Bits= 33
// G2 Table size= 36

const CURVE_A int= 0
const CURVE_Cof_I int= 0
var CURVE_Cof= [...]Chunk {0x140000382,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}
const CURVE_B_I int= 1
var CURVE_B= [...]Chunk {0x1,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}
var CURVE_Order= [...]Chunk {0x8A5FE6FCD671C01,0xBE599467C24DA11,0xC7CD0562303C4CC,0x9D34C4C92016A85,0xBC972C2E6E74196,0x3F0B3CBE003FAD6,0x615C0D6C635387A,0xE2885E233A9CCC1,0x2386F8A925,0x0}
var CURVE_Gx= [...]Chunk {0xBCE8732315AF640,0x74DA5D3A1E6D8C3,0x57DB368B11786CB,0x665D859236EBDBC,0x46A9DF6F9645847,0xEDFFB9F75445505,0xE86868CF61ABDBA,0x93F860DE3F257E0,0x40F2BAF2B73DF1E,0x2AF59B7AC3}
var CURVE_Gy= [...]Chunk {0xDBB5DE3E2587A70,0xF37AEF7B926B576,0xF77C2876D1B2E35,0x78584C3EF22F487,0xFFB98AEE53E80F6,0xD41B720EF7BB7BE,0xFEB8A52E991279D,0xB398A488A553C9E,0x31F91F86B3A2D1F,0xCEFDA44F65}

var CURVE_Bnx= [...]Chunk {0x140000381,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}
var CURVE_Cru= [...]Chunk {0x4DE9AC5E1C79B90,0x5CD8E3F88E5DE82,0xAB21F74F7421A20,0x6694B9B60DB5D62,0x73422B5FB82F431,0xFF46A846B5FA6AA,0x83D66C1E5FCBED6,0x2096384F2AFA565,0x8B75055DD5D1F4E,0x2C6}
var CURVE_Pxaaa= [...]Chunk {0x34FD0B4ACE8BFAB,0xB79766322154DEC,0x4D80491F510317,0x3CA0612F4005030,0xBAAD1A8C42281A6,0x3A2EF156C46FF79,0x344DBCCB7DE64DB,0x2775DEBABBEFC70,0x71E4A38237FA45A,0x5D615D9A78}
var CURVE_Pxaab= [...]Chunk {0x669B36676B47C57,0x5556A01AFA143F1,0x7630D979630FFD7,0x6AFFA62504F0C3C,0xABFEDF16214A7,0x12307F4E1C3943A,0xE1623E9526F6DA,0xBC07E8B22BB6D98,0x258512069B0E86A,0x7C4973ECE2}
var CURVE_Pxaba= [...]Chunk {0x488156CA55A3E6A,0xEF4CDED6B3F0B46,0xCBDFBB879D5FEA8,0x66F0D2A6D55F028,0xC1DBD19242FFAE7,0xCCBAB5AB6860161,0xAE237CA7A6D6957,0xAD83BC73A8A6CA9,0xF1334E1B2EA1853,0x1FCCC70198}
var CURVE_Pxabb= [...]Chunk {0x9A7033CBB7FEAFE,0x10B8CB4E80BC3F0,0x1C5257C200CA523,0x43B1B279B9468C3,0x5F63E1C776E6EC1,0x393F8BE0CC218A9,0x62F3E5821B7B92A,0x54D4BFE8F5985AC,0xEB6185C78D80129,0xBE2218C25C}
var CURVE_Pxbaa= [...]Chunk {0x39C3A1C53F8CCE5,0x5B5F746C9D4CBB7,0xD55FC1889AA80C6,0xEF492AE589274FA,0x9E48199D5AC10B2,0xC5805386699981F,0xB1642B5675FF0E7,0xA9DD63007C675D0,0x35913A3C598E4CA,0x38B91C600B}
var CURVE_Pxbab= [...]Chunk {0x2004D914A3C093A,0x7960910FCE3370F,0xA9F177612F097FC,0x40B9C0B15DD7595,0x3835D28997EB57B,0x7BB037418181DF6,0xEF0977A3D1A5867,0xCDA088F7B8F35DC,0x738603F1311E4E,0xC96C7797EB}
var CURVE_Pxbba= [...]Chunk {0x41607E60750E057,0x4B5B0E205C3354E,0xCBE4324C22D6333,0xAA5EFCF3432AAD1,0xF293B13CED0FD0C,0xA2C0B7A449CEF11,0x9D13852B6DB908B,0x8AEE660DEA41B3,0x61EE3F0197A4989,0xB9B7951C60}
var CURVE_Pxbbb= [...]Chunk {0xE19DA00FBC6AE34,0x6AF2FC9E97C3F84,0x9BD6AEBF9FC44E5,0x90B7E2B0D458547,0xA93F29CFF364A71,0x719728A7F9F8CFC,0xFAF47B5211CF741,0x4AAA2B1E5D7A9DE,0x2BDEC5282624C4F,0x827D5C22FB}
var CURVE_Pyaaa= [...]Chunk {0x3EDD3FE4D2D7971,0x45012AB12C0FF32,0x9ABF77EEA6D6590,0x336D8AE5163C159,0x35AFA27748D90F7,0xBFC435FAAB09062,0x59A577E6F3B39E,0x2F3024B918B4238,0x75B5DFA49721645,0xEB53356C3}
var CURVE_Pyaab= [...]Chunk {0x1471DB936CD5665,0x8B423525FFC7B11,0x2FA097D760E2E58,0xD1892AB24E1DD21,0x6B243B1F192C5C3,0x64732FCBF3AFB09,0xA325E6FBA01D729,0x5FCADC2B75A422B,0xE0FF144DA653181,0x284DC75979}
var CURVE_Pyaba= [...]Chunk {0x8332A526A2A8474,0xBC7C46FC3B8FDE6,0x1D35D51A652269C,0x36CA3295E5E2F0C,0xC99D0E904115155,0xD370514475F7D5,0x216D5B119D3A48,0x67669EF2C2FC503,0x8523E421EFB703,0xB36A201DD0}
var CURVE_Pyabb= [...]Chunk {0x6213DA92841589D,0xB3D8B8A1E533731,0x7BDA503EE5E578F,0x817742770BA10D6,0x224333FA40DCED2,0x10E122D2742C89B,0x60DCEE23DD8B0E7,0x78762B1C2CDED33,0xEDC0688223FBBD4,0xAEC25A4621}
var CURVE_Pybaa= [...]Chunk {0x47831F982E50137,0x857FDDDFCF7A43F,0x30135945D137B08,0xCA4E512B64F59F4,0x7FA238CDCE8A1E2,0x5F1129857ED85C7,0xB43DD93B5A95980,0x88325A2554DC541,0xA9C46916503FA5A,0xD209D5A223}
var CURVE_Pybab= [...]Chunk {0x4EEDC58CF90BEE4,0xA59ED8226CF3A59,0xFC198CAA72B679D,0xF47C180D139E3AA,0xE8C270841F6824,0x55AB7504FA8342,0xB16722B589D82E2,0xD537B90421AD66E,0x36B7A513D339D5A,0x7D0D037457}
var CURVE_Pybba= [...]Chunk {0xD41FAEAFEB23986,0xE884017D9AA62B3,0x40FA639F53DCCC9,0xAB8C74B2618B5BB,0x5AE3A2864F22C1F,0xE4C819A6DF98F42,0xC0841B064155F14,0xD17AF8A006F364F,0xE65EA25C2D05DFD,0x896767811B}
var CURVE_Pybbb= [...]Chunk {0x667FFCB732718B6,0x5AC66E84069C55D,0xD8C4AB33F748E,0x333EC7192054173,0x8E69C31E97E1AD0,0xEF8ECA9A9533A3F,0x6BE8E50C87549B6,0x4F981B5E068F140,0x9029D393A5C07E8,0x35E2524FF8}
var CURVE_W=[2][10]Chunk {{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}}
var CURVE_SB=[2][2][10]Chunk {{{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}},{{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}}}
var CURVE_WB=[4][10]Chunk {{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}}
var CURVE_BB=[4][4][10]Chunk {{{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}},{{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}},{{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}},{{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0},{0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}}}