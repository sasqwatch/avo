// Downloaded from https://github.com/dgryski/go-stadtx/raw/3c3d9b328c24a9b5ecd370654cd6e9d60a85752d/stadtx_test.go

package stadtx

import (
	"testing"
)

/*

#include "stadtx_hash.h"
#include <stdio.h>

int main()
{

    U64 v;

    U64 seed_ch[2] = { 0x0001020304050607ULL, 0x08090A0b0C0D0E0F };

    U8 data[128];

    int i;

    printf("var tests = []uint64{\n");

    for (i = 0; i < 128; i++) {
	data[i] = i;
	v = stadtx_hash((const U8 *) seed_ch, data, i);
	printf("0x%016lx,", v);
        if ((i+1) % 8 == 0) {
            printf("\n");
        }
    }

    printf("}\n");
}

*/

func TestHash(t *testing.T) {

	var tests = []uint64{
		0x0db6608af30c8cc1, 0xe3ffb1e1a2273589, 0xc34df40f05c3d43a, 0xf79284a93cd817cf, 0x178aed7315bcd4b9, 0x48a29a42948da116, 0xa78613b72330a064, 0xa23cd23f7b6eaef7,
		0x7e4c118aca1d7886, 0x37059645f8797936, 0xcb7197e3253670cb, 0x5dadeebe5a5bc19f, 0x537057981e3f75ad, 0x41d57f7870fbcbf8, 0x7fc3cd6b1bb018c2, 0xbf98dbc2a88c5050,
		0x4a3b3c5468927570, 0xeb6370986f0a28b4, 0x391b356f563a07a2, 0xdb64e2d15432f18f, 0xc8b1bb6c1c375241, 0xf7f86b2d54859f88, 0x862d58654d669d36, 0x77ffae02015534af,
		0x8f992c4b37053081, 0x7739d76bb99573e8, 0x91115ce66437eeb0, 0x4fe15afafcdb97cb, 0x1f9b53a41ef68d60, 0x87a23023439dbeac, 0xa91d254b6e408270, 0x92d296f71bd786de,
		0xe08da31cfd5179ca, 0xcc43972f2a93adcb, 0x33b0b5a0aaaa5eea, 0x572ceb8cbdbab15c, 0xefe041ad84a2e8d4, 0xcc196969c505dace, 0xcdf7e4d013499719, 0xfac29209a057fda0,
		0x77fd1bf64ee30cf2, 0x0f366d248913d0d0, 0x2821028cc5505b41, 0xa571d04d77ef675e, 0xb795018736480554, 0x0d0ff98c4b3f5f4c, 0x8513745208da7fd2, 0xa564e25dff77eb13,
		0x4375b88d3d7939b6, 0xa90a25148887f828, 0x60616c6d42f7c235, 0xc4fe22c82a03ff4f, 0x4f5a34c6dbf5b521, 0x2681b8c0acfd9b9f, 0xd9d18ed48f9ed27a, 0x420f9d6d60ef38df,
		0x4a4b79af1f66d7c8, 0xc4e9504dd72ef0cb, 0x99573e84016395f8, 0x1c3121e8c827dbcb, 0x9d24802464cc1872, 0xf8f6a1ce9fcacd83, 0x3253bd62273249db, 0x5f7f8f09fb38d49d,
		0x4d8687dd8d08cab1, 0x620daae84bfddb5b, 0x75a617b617f9f0a1, 0xd4da1af4f12745e0, 0x6b36894b4d819ed4, 0xaf630af151c8e494, 0xab5f39b817c3694d, 0x95d9d3368bb48ba1,
		0x255334a30d6340bb, 0x8fac630af7ca8789, 0xb5ed6726f8453ddc, 0x46901504ec0983d2, 0x28b2ce6c3ca0c1ea, 0x1dd3519df8229928, 0x26e3a6a8c6358f66, 0x6158403905bed2ea,
		0x4305ff8110eae857, 0x242b67463b95498c, 0x068a9ca66b87878f, 0x7ae1ecb47838c849, 0x11008d3f1dcaf002, 0x0913b479903c3ef3, 0x9e94d283f9ac4e75, 0xb0abae509c13b987,
		0x66a4f4534050428b, 0x84b82a6f343a8fcd, 0xa9c418edf1f127e5, 0x2b990abfcedb6724, 0x6b47931cb943f52f, 0x4c04efa93ea2e780, 0xf21e88b212e8c44b, 0x086f4fe25e8a9eda,
		0x707bc1c224794651, 0x88d1a9194b06112f, 0x09935bcf7618afad, 0x750fe3517d46e8f5, 0xeac04dfd871b8393, 0x0861101436a7d359, 0x1e971720b0c00991, 0x8f7eb2608c40a30c,
		0xb2e142a7da8288e9, 0x320949b842c373dc, 0xa2edee4716e95cbc, 0x1fd86a7a946d7e01, 0xd4b9d24e8936b228, 0x5ab2d7cab3834f7a, 0x9ae42d7a8b241d00, 0xf9dc368e8df9bb11,
		0x3e538bcf908bc657, 0x1f621c025155d15a, 0x2878d0e561531226, 0x9fb44daf75b7698e, 0xf19125829320cd64, 0x72965e1625d10cb9, 0x493679ecc71d278c, 0x563ef209b56abc24,
		0xb6d39808be0528cd, 0xdb1d268d396264ee, 0xd68ce668d90aeed7, 0x1fcb217fa263821f, 0x9015c4b806730fd9, 0xf870f30533d0109d, 0x1688ab70df2abf39, 0x3a891995e76992bd,
	}

	seed := []uint64{0x0001020304050607, 0x08090A0B0C0D0E0F}

	state := SeedState(seed)

	input := make([]byte, 128)

	for i := range input {
		input[i] = byte(i)

		if h := Hash(&state, input[:i]); h != tests[i] {
			t.Errorf("Hash(..., input[:%d])=%016x, want %016x\n", i, h, tests[i])
		}
	}
}
