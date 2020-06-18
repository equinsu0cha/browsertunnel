package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDomain(t *testing.T) {
	tests := []struct {
		topDomain string
		domain    string
		fails     bool
		output    msgFragment
	}{
		{
			topDomain: "tunnel.example.com.",
			domain:    "2jkhm3.592.0.jf2ca2ltebqxiidxn5zgwidfozsxe6lxnbsxezjmebthk3tdoruw63tjnztsa43.nn5xxi2dmpeqgc5baoruw2zltfqqgc5ban52gqzlseb2gs3lfomqgs3ramzuxi4.zamfxgiidtorqxe5dtfyqes5bamjzgkylunbsxglbanf2ca2dfmf2hglbanf2ca.zlborzs4icjoqqhg2djorzsaylomq.tunnel.example.com.",
			output: msgFragment{
				id:        "2jkhm3",
				totalSize: 592,
				offset:    0,
				data:      "jf2ca2ltebqxiidxn5zgwidfozsxe6lxnbsxezjmebthk3tdoruw63tjnztsa43nn5xxi2dmpeqgc5baoruw2zltfqqgc5ban52gqzlseb2gs3lfomqgs3ramzuxi4zamfxgiidtorqxe5dtfyqes5bamjzgkylunbsxglbanf2ca2dfmf2hglbanf2cazlborzs4icjoqqhg2djorzsaylomq",
			},
			fails: false,
		},
		{
			topDomain: "tunnel.example.com.",
			domain:    "example.com.",
			fails:     true,
		},
		{
			topDomain: "tunnel.example.com.",
			domain:    "2jkhm3.FAIL.0.jf2ca2ltebqxiidxn5zgwidfozsxe6lxnbsxezjmebthk3tdoruw63tjnztsa43.nn5xxi2dmpeqgc5baoruw2zltfqqgc5ban52gqzlseb2gs3lfomqgs3ramzuxi4.zamfxgiidtorqxe5dtfyqes5bamjzgkylunbsxglbanf2ca2dfmf2hglbanf2ca.zlborzs4icjoqqhg2djorzsaylomq.tunnel.example.com.",
			fails:     true,
		},
		{
			topDomain: "tunnel.example.com.",
			domain:    "2jkhm3.592.FAIL.jf2ca2ltebqxiidxn5zgwidfozsxe6lxnbsxezjmebthk3tdoruw63tjnztsa43.nn5xxi2dmpeqgc5baoruw2zltfqqgc5ban52gqzlseb2gs3lfomqgs3ramzuxi4.zamfxgiidtorqxe5dtfyqes5bamjzgkylunbsxglbanf2ca2dfmf2hglbanf2ca.zlborzs4icjoqqhg2djorzsaylomq.tunnel.example.com.",
			fails:     true,
		},
		{
			topDomain: "tunnel.example.com.",
			domain:    "2jkhm3.592.0.tunnel.example.com.",
			fails:     true,
		},
	}
	for _, test := range tests {
		got, err := parseDomain(test.topDomain, test.domain)
		if test.fails {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
		}
		require.Equal(t, test.output, got)
	}
}

func TestAssemble(t *testing.T) {
	tests := []struct {
		input  msgFragmentList
		output string
		fails  bool
	}{
		{
			input: msgFragmentList{
				totalSize: 592,
				fragments: map[int]msgFragment{
					0: msgFragment{
						id:        "2jkhm3",
						totalSize: 592,
						offset:    0,
						data:      "jf2ca2ltebqxiidxn5zgwidfozsxe6lxnbsxezjmebthk3tdoruw63tjnztsa43nn5xxi2dmpeqgc5baoruw2zltfqqgc5ban52gqzlseb2gs3lfomqgs3ramzuxi4zamfxgiidtorqxe5dtfyqes5bamjzgkylunbsxglbanf2ca2dfmf2hglbanf2cazlborzs4icjoqqhg2djorzsaylomq",
					},
					218: msgFragment{
						id:        "2jkhm3",
						totalSize: 592,
						offset:    218,
						data:      "qgm5ldnnzs4icxnbqxiidbebwws43umfvwkidun4qgqylwmuqgk5tfoiqhgyljmqqhi2dfebuwilraiv3gk4tzo5ugk4tfebuxiidjomqg2yldnbuw4zlt4kaji4tfmfwca33omvzsyidon52caztjm52xeylunf3gkidpnzsxgoranvqwg2djnzsxgideojuxm2lom4qg65dimvzca3lbmn",
					},
					434: msgFragment{
						id:        "2jkhm3",
						totalSize: 592,
						offset:    434,
						data:      "ugs3tfomwca3lbmnugs3tfomqgezljnztsazdsnf3gk3ramj4sa33unbsxeidnmfrwq2lomvzsyidxnf2gqidbnrwca5dimuqg4zldmvzxgylspeqgg33vobwgs3thomqgc3teebrw63tomvrxi2lpnzzs4000",
					},
				},
			},
			output: "It is at work everywhere, functioning smoothly at times, at other times in fits and starts. It breathes, it heats, it eats. It shits and fucks. What a mistake to have ever said the id. Everywhere it is machines—real ones, not figurative ones: machines driving other machines, machines being driven by other machines, with all the necessary couplings and connections.",
			fails:  false,
		},
		{
			input: msgFragmentList{
				totalSize: 10,
				fragments: map[int]msgFragment{
					30: msgFragment{
						id:        "2jkhm3",
						totalSize: 10,
						offset:    30,
						data:      "jf2ca2ltebqxiidxn5zgwidfozsxe6lxnbsxezjmebthk3tdoruw63tjnztsa43nn5xxi2dmpeqgc5baoruw2zltfqqgc5ban52gqzlseb2gs3lfomqgs3ramzuxi4zamfxgiidtorqxe5dtfyqes5bamjzgkylunbsxglbanf2ca2dfmf2hglbanf2cazlborzs4icjoqqhg2djorzsaylomq",
					},
				},
			},
			fails: true,
		},
		{
			input: msgFragmentList{
				totalSize: 10,
				fragments: map[int]msgFragment{
					10: msgFragment{
						id:        "2jkhm3",
						totalSize: 10,
						offset:    0,
						data:      "not base32",
					},
				},
			},
			fails: true,
		},
	}
	for _, test := range tests {
		got, err := test.input.assemble()
		if test.fails {
			require.NotNil(t, err)
		} else {
			require.Nil(t, err)
		}
		require.Equal(t, test.output, got)
	}
}

func TestListenDomains(t *testing.T) {
	tun := &tunnel{
		Messages:  make(chan string, 256),
		topDomain: "tunnel.example.com.",
		domains:   make(chan string, 256),
		fragments: make(map[string]*msgFragmentList),
	}
	go tun.listenDomains()

	tun.domains <- "i42ftq.592.0.jf2ca2ltebqxiidxn5zgwidfozsxe6lxnbsxezjmebthk3tdoruw63tjnztsa43.nn5xxi2dmpeqgc5baoruw2zltfqqgc5ban52gqzlseb2gs3lfomqgs3ramzuxi4.zamfxgiidtorqxe5dtfyqes5bamjzgkylunbsxglbanf2ca2dfmf2hglbanf2ca.zlborzs4icjoqqhg2djorzsaylomq.tunnel.example.com."
	tun.domains <- "i42ftq.592.218.qgm5ldnnzs4icxnbqxiidbebwws43umfvwkidun4qgqylwmuqgk5tfoiqhgyljm.qqhi2dfebuwilraiv3gk4tzo5ugk4tfebuxiidjomqg2yldnbuw4zlt4kaji4tf.mfwca33omvzsyidon52caztjm52xeylunf3gkidpnzsxgoranvqwg2djnzsxgid.eojuxm2lom4qg65dimvzca3lbmn.tunnel.example.com."
	tun.domains <- "i42ftq.592.434.ugs3tfomwca3lbmnugs3tfomqgezljnztsazdsnf3gk3ramj4sa33unbsxeidnm.frwq2lomvzsyidxnf2gqidbnrwca5dimuqg4zldmvzxgylspeqgg33vobwgs3th.omqgc3teebrw63tomvrxi2lpnzzs4000.tunnel.example.com."

	got := <-tun.Messages
	expected := "It is at work everywhere, functioning smoothly at times, at other times in fits and starts. It breathes, it heats, it eats. It shits and fucks. What a mistake to have ever said the id. Everywhere it is machines—real ones, not figurative ones: machines driving other machines, machines being driven by other machines, with all the necessary couplings and connections."
	require.Equal(t, expected, got)
}
