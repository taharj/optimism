// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100725760003560e01c8063e03110e111610050578063e03110e114610106578063e15926111461012e578063fef2b4ed1461014357600080fd5b806361238bde146100775780638542cf50146100b5578063c0c220c9146100f3575b600080fd5b6100a26100853660046104df565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6100e36100c33660046104df565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100ac565b6100a2610101366004610501565b610163565b6101196101143660046104df565b610238565b604080519283526020830191909152016100ac565b61014161013c36600461053c565b610329565b005b6100a26101513660046105b8565b60006020819052908152604090205481565b600061016f8686610432565b905061017c836008610600565b8211806101895750602083115b156101c0576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558484528752808320948352938652838220558181529384905292205592915050565b6000828152600260209081526040808320848452909152812054819060ff166102c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b50600083815260208181526040909120546102dd816008610600565b6102e8856020610600565b1061030657836102f9826008610600565b6103039190610618565b91505b506000938452600160209081526040808620948652939052919092205492909150565b604435600080600883018611156103485763fe2549876000526004601cfd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316176104d8818360408051600093845233602052918152606090922091527effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b9392505050565b600080604083850312156104f257600080fd5b50508035926020909101359150565b600080600080600060a0868803121561051957600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b60008060006040848603121561055157600080fd5b83359250602084013567ffffffffffffffff8082111561057057600080fd5b818601915086601f83011261058457600080fd5b81358181111561059357600080fd5b8760208285010111156105a557600080fd5b6020830194508093505050509250925092565b6000602082840312156105ca57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115610613576106136105d1565b500190565b60008282101561062a5761062a6105d1565b50039056fea164736f6c634300080f000a"

var PreimageOracleDeployedSourceMap = "306:3865:130:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;537:68;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;413:25:286;;;401:2;386:18;537:68:130;;;;;;;;680:66;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;614:14:286;;607:22;589:41;;577:2;562:18;680:66:130;449:187:286;1367:1165:130;;;;;;:::i;:::-;;:::i;789:536::-;;;;;;:::i;:::-;;:::i;:::-;;;;1205:25:286;;;1261:2;1246:18;;1239:34;;;;1178:18;789:536:130;1031:248:286;2574:1595:130;;;;;;:::i;:::-;;:::i;:::-;;419:50;;;;;;:::i;:::-;;;;;;;;;;;;;;;1367:1165;1529:12;1634:36;1663:6;1634:28;:36::i;:::-;1627:43;-1:-1:-1;1764:9:130;:5;1772:1;1764:9;:::i;:::-;1750:11;:23;:37;;;;1785:2;1777:5;:10;1750:37;1746:90;;;1810:15;;;;;;;;;;;;;;1746:90;1905:12;2005:4;1998:18;;;2106:3;2102:15;;;2089:29;;2138:4;2131:19;;;;2240:18;;2330:20;;;:14;:20;;;;;;:33;;;;;;;;:40;;;;2366:4;2330:40;;;;;;2380:19;;;;;;;;:32;;;;;;;;;:39;2496:21;;;;;;;;;:29;2345:4;1367:1165;-1:-1:-1;1367:1165:130:o;789:536::-;865:12;914:20;;;:14;:20;;;;;;;;:29;;;;;;;;;865:12;;914:29;;906:62;;;;;;;2839:2:286;906:62:130;;;2821:21:286;2878:2;2858:18;;;2851:30;2917:22;2897:18;;;2890:50;2957:18;;906:62:130;;;;;;;;-1:-1:-1;1099:14:130;1116:21;;;1087:2;1116:21;;;;;;;;1167:10;1116:21;1176:1;1167:10;:::i;:::-;1151:12;:7;1161:2;1151:12;:::i;:::-;:26;1147:87;;1216:7;1203:10;:6;1212:1;1203:10;:::i;:::-;:20;;;;:::i;:::-;1193:30;;1147:87;-1:-1:-1;1290:19:130;;;;:13;:19;;;;;;;;:28;;;;;;;;;;;;789:536;;-1:-1:-1;789:536:130:o;2574:1595::-;2870:4;2857:18;2675:12;;2999:1;2989:12;;2973:29;;2970:210;;;3074:10;3071:1;3064:21;3164:1;3158:4;3151:15;2970:210;3423:3;3419:14;;;3323:4;3407:27;3454:11;3428:4;3573:16;3454:11;3555:41;3786:29;;;3790:11;3786:29;3780:36;3838:20;;;;3985:19;3978:27;4007:11;3975:44;4038:19;;;;4016:1;4038:19;;;;;;;;:32;;;;;;;;:39;;;;4073:4;4038:39;;;;;;4087:18;;;;;;;;:31;;;;;;;;;:38;;;;4135:20;;;;;;;;;;;:27;;;;-1:-1:-1;;;;2574:1595:130:o;492:353:129:-;752:11;777:19;765:32;;749:49;824:14;749:49;1277:21;1426:15;;;1467:8;1461:4;1454:22;1595:4;1582:18;;1602:19;1578:44;1624:11;1575:61;;1222:430;824:14;817:21;492:353;-1:-1:-1;;492:353:129:o;14:248:286:-;82:6;90;143:2;131:9;122:7;118:23;114:32;111:52;;;159:1;156;149:12;111:52;-1:-1:-1;;182:23:286;;;252:2;237:18;;;224:32;;-1:-1:-1;14:248:286:o;641:385::-;727:6;735;743;751;804:3;792:9;783:7;779:23;775:33;772:53;;;821:1;818;811:12;772:53;-1:-1:-1;;844:23:286;;;914:2;899:18;;886:32;;-1:-1:-1;965:2:286;950:18;;937:32;;1016:2;1001:18;988:32;;-1:-1:-1;641:385:286;-1:-1:-1;641:385:286:o;1284:659::-;1363:6;1371;1379;1432:2;1420:9;1411:7;1407:23;1403:32;1400:52;;;1448:1;1445;1438:12;1400:52;1484:9;1471:23;1461:33;;1545:2;1534:9;1530:18;1517:32;1568:18;1609:2;1601:6;1598:14;1595:34;;;1625:1;1622;1615:12;1595:34;1663:6;1652:9;1648:22;1638:32;;1708:7;1701:4;1697:2;1693:13;1689:27;1679:55;;1730:1;1727;1720:12;1679:55;1770:2;1757:16;1796:2;1788:6;1785:14;1782:34;;;1812:1;1809;1802:12;1782:34;1857:7;1852:2;1843:6;1839:2;1835:15;1831:24;1828:37;1825:57;;;1878:1;1875;1868:12;1825:57;1909:2;1905;1901:11;1891:21;;1931:6;1921:16;;;;;1284:659;;;;;:::o;1948:180::-;2007:6;2060:2;2048:9;2039:7;2035:23;2031:32;2028:52;;;2076:1;2073;2066:12;2028:52;-1:-1:-1;2099:23:286;;1948:180;-1:-1:-1;1948:180:286:o;2315:184::-;2367:77;2364:1;2357:88;2464:4;2461:1;2454:15;2488:4;2485:1;2478:15;2504:128;2544:3;2575:1;2571:6;2568:1;2565:13;2562:39;;;2581:18;;:::i;:::-;-1:-1:-1;2617:9:286;;2504:128::o;2986:125::-;3026:4;3054:1;3051;3048:8;3045:34;;;3059:18;;:::i;:::-;-1:-1:-1;3096:9:286;;2986:125::o"

func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
}
