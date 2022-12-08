package main

import (
	"fmt"
	"strings"
)

// l and r are left and right compartment strings
// assumes all runes are ascii chars
// also scores
func recurring(l, r string) int {
	// use go map to count both occurences
	setl := make(map[rune]int, 0)
	setr := make(map[rune]int)
	for _, c := range l {
		setl[c] = 1
	}

	for _, c := range r {
		setr[c] += 1
	}

	sum := 0
	// look for both occurences
	for k := range setl {
		if _, ok := setr[k]; ok {
			fmt.Println(string(k), k)
			prio := int(k) - 64 + 26
			if int(k) >= 96 {
				prio = int(k) - 96
			}
			fmt.Println(int(prio))
			sum += int(prio)
		}
	}

	return sum
}

// part 2, groups of three elves
// also scores
func groupBy(b1, b2, b3 string) int {
	fill := func(s string) map[rune]int {
		m := make(map[rune]int, 0)
		for _, c := range s {
			m[c] += 1
		}
		return m
	}
	m1, m2, m3 := fill(b1), fill(b2), fill(b3)
	fmt.Println(b1, b2, b3)

	sum := 0
	for k := range m1 {
		if _, ok := m2[k]; ok {
			if _, ok = m3[k]; ok {
				fmt.Println(string(k), k)
				prio := int(k) - 64 + 26
				if int(k) >= 96 {
					prio = int(k) - 96
				}
				fmt.Println(int(prio))
				sum += int(prio)
			}
		}

	}
	return sum
}

func main() {
	lines := strings.Split(bags, "\n")
	sum := 0
	for i, l := range lines {
		leftbag, rightbag := l[0:len(l)/2], l[len(l)/2:]
		fmt.Println(i, l)
		fmt.Println("left side is: ", leftbag)
		fmt.Println("right side is: ", rightbag)
		sum += recurring(leftbag, rightbag)
	}
	fmt.Println(sum)

	sum2 := 0
	for i := 0; i < len(lines); i = i + 3 {
		sum2 += groupBy(lines[i], lines[i+1], lines[i+2])
	}
	fmt.Println(sum2)
}

const bags = `PnJJfVPBcfVnnPnBFFcggttrtgCrjDtSjzSS
llWlLbmmdLLwLbqGdmTmbZfCQtzrMQfrjSzrtSrMqgMt
sHlZTsWZwGGlZmGTmdlZbhJNRPphVfRvVnRBsRsJJV
fsHtVbjtqstBghhwwPBw
SDQzzSzQrQMmmQlmlcNcJLZPgLrVZTdCddhgdPwwCw
JmSWSVGGlJJbRsbpWHfbRj
tJndRtwtddPvllvfrldrfPpHWDgglFDWWmMmHWmHpZlS
BBJTTjCsJWZCmSHSZD
LhqLcVzshTNjhqhcjLLTLjbTnGndfdwrfPRVRrdnwftQwJRv
wHlPJZwbbZfqbFwqFZfrrcrJrtMWSMMVtVcJht
NzzzNBjNfLzvGfDNjMhVhrrMShLchsRVLs
DDdmmgBGDNdgfgZggnZbZHln
jqNjZJqsGsRqJJqnlJJGzMzffcffTCfQcFmvcWfvTNfcvv
PdhVdrwphhVtDdSPLmFCWTLFWWTfFQQr
dSPwbbVdbpQllZMQbMjM
QQdfflqvjTvfZqLMWfNDGhwsCNGGGM
rzRRRTVTPTNhsDWDRhGC
gHSTpTnppvjQgJjcql
nzNvsFBBBFsNrnNBrvndfThwDbhVPzVVwhZZChpZPCbZ
GMQQStmcHHmlfMtPwbZVVVVhhPhbVc
GmSmMGlmWWgqQMHgQHtSvTNTTrgBBvNvFsfTfLvd
VNVBgRRBRzWRRbRTrTBgbzpZvMvMJlQJZJpJvlTMdJlp
mfdfdfdHfQplZJJh
ttnPmsLHmqsPGgWjjWRbBdbbGj
mMMzwFFpnwgMggSzhQSbSJbhWzbr
NPRlHHPbPRdPHCdRNGWWhWQhJVVGhZrCCJ
LlNbRllLLdHnpLgMTgTsDT
mWhvwmthrzJrLhvftwPrtNHlbcHgbDzNDHcGgzllHl
TCjdSCFMQCCCjMspCMQBMqpFHNGblPPnGNNbpnpccZGccclc
jVBPsqdFMFsTVMsdMmfwhwvVwLrvWhVtJR
TdhJJttgmdgctBBhBwBqBCLWqC
fjNGPblpPlRDfsfGbfcWqBqwWvnHVvBLVHss
lpRNDGjbrGfPZbmZZMFdTzTgcFdS
WPLMvWLMqLcSvrrBmDPFfRzmDPRfZDHP
dCggNCCdNCbHTHTfFnNHzM
GlJVQjdCCwjMGbGpjGdCpqrtctqcttchsrLLcthlhc
jgPHPgjGCwrqjqrvzBfFfCTTCdFpDD
QhQLtcrQQLWshQttdzzvzBnhvvBhvFTv
RJRssrQSWQSqjqqVSSGP
sPszmnmnQPQbjSRVVjbRjVVB
cdvZqMfNchMMMqCBlBZlmSmWCHBj
dFFDqJNNJvTJTqvvwmGQsGwJztgrwsts
dJGPMNStSclWJSPScMNMzqqDfzqVzwwjGfzjqfQz
hppFbrhFrrhgCpFlFhvFTgrsDfsjVzvLDfqwVzzwQDsVwL
RrgnZbphTpbgrnZgmpbhFcMlnWdMBNWJdtNJBMMSMc
HqtfLZCLmLtrSqZStzFtwpPQWlNQBpplpQWmBpNp
jMJcnRFcJjvjvcGvdJpBppBwQBpwQWPW
DcvDgbhjMMsjGvcjRRvjvgbGtsTZFHCCHSCHZHLfsSsCSrZH
wvFPFqvjlgTwjvjZbMSzHTSScMBBbN
VVmpVsmmptfDsdSZbMzZzMZBtB
LVhfGJVhQsWhVJhfmGJMwljFvqWrlPvCqrqjFFvj
cdNbtwrtwrRNLzhQhHfCzbGzQH
DMlMgFDBMSBjMTqMFWMFDQHdGZHvQdTHHvHzHChJJv
MBVWDngVqVWlqqBgBjjcsPnwnNdNsPNPpRcLws
pQqHsBmmrQHrrmsmmpwMdfdPbfWPMVbMdWbPwG
ZSvZDlvghgSDZCSzJZzgZNNPFRFMVRVfMdbFPddfMzGRWR
vWgJhSntvglSgDCHsjHBBpnTrTmjmB
hDgmmgnwDCdddRCqCPzz
sSslHbctbHsNrhpZqdqprc
WbsHSBQHQBLHTWJSlmjnVwFDvWhnvjmfgv
phwdhVLRddWJLVglLWRlWWgwhCjCrrrHZmCjZNHtjvvZrHTN
zzFzPQqFDcFsHsFPHSsGSCCNmftjjNDjCfCCCNfvrr
BSQsqFQMcPqHbcGzFMWgLlgwpLlpJLLpBwnp
GTrsHGGHfGMvsrGsvnGcHsvTwhghwhpVpJhQpwhQghhmmm
bFPFPPtSdPtSZPZWjSdZSNFdmhDJQgRQbQgmhJgJmRwVQMRh
jPdjldqlSNBdFBjtFZsfsGMqsHCMcqcfvnrf
crZHHZcZZsSrVrSCrBNTTpppBpJBCbwbBn
DzhRggfhhlRghmBNBpLJwDpDDTNp
qvwqwPRfvfzPPRvQlfPvPlRjZSHZrcScWrtVQscjjtZHMS
GRGBgsgghbvRvgsBTsgRjjznCCtSVCZtCFZtzFtwnSdj
cmqpPNnWWPrqWHplHDDcpCddFfZdFtdrwzrfdVFFdt
lPNPNPcqpcHWmPLllDDGhbRbnMgLgRbvRsbhhg
ZRBBNFWWBWPDZZprnwVvDD
bSzzhHdttMtCCzCThMTShStcvrqNNjwnrcvqvvVVpprvnnpN
CStdTMLztMmTSSMTSSSfBQNWlsFRJRPfmlNRNB
pMLptvGHTSGSHtLLwNWvcqsJmWWzmPFFms
nrQPrjrddgQlDPZjdlBcFFzFsFWhqWgshszhNW
CZDCQjjBBjrnCfQCrZrbSbHLRwGftbfPbHRMfp
sqQHBqsHbwDtmbSB
GjVGrlsMsPjbPbPmTtSwDb
lzFFGznLGGrLrnrVzMzsqvpqvWQcnNZWvCpnqpHc
nZWZjdWRwhRdBRhhdNZMtHtpPrGCCGsrTGsG
DzFFbblvbzFvVJzzzlVmvlpTpMpMrMCgGMtprJgCMrPM
zVmFzFvfzSmDzmCmzmmVQwNRqdRhcWBqShjqNndWWj
QQTdgLQlhGhQdPbwJJgwRVmtNRBV
vCSnjzFqSDDFMDjvrqvjvnVBRsNJRpJwtNJNpNrtbtbt
DfnCCjDMzcnqvHZGlcTcZThBlG
lWtmssWNcBjTjhlLpn
fwJbgfQfVfqPwmdgJjppHTjCphppBHCnhV
bgvfmvmbqdfDqJvqgbPJNDNSMscNDtccDzzcNtRt
ttBHSflTlqwGGQJBQq
jvdPdPLsDvCzvVqVQQGw
ZDdZnPrrPLPLZMZlfWSqmmTWmtTMtF
ShhwCZSwlPwZPSplwlRGGVmttWVVQcJVnJtSJVLcjc
qmzDmgNqNfgtJWWqtQLJqL
rsFgvrmFdfRhPBRvRRPZ
RQQRGJRpfJzlTQmbfbsGlsTMSHMPgcMFpgqPHHpHqqPpcP
WmLZDBBLZLNVdjWqDCFHCMccSPgSgg
thrNBVVBrWZNBWNrRGQbRtJGzlRRbmlf
NfdjBfsnDfNjfBddNBnBwsQShzQDhSDhvZSRDpFRSVFQ
lHrbTgtLmrblPzShhRpvRPZPVz
tWbqHlHmWHgWHtqrLJJsvBwnBfjswjwnNMwf
qQfdCVqZSqZmQmgVqCqCWTMcNFTwWwTccgpcWcpw
nSLrjrJnLcTwLNcz
sjRjSnJnRsBlPtnqCqCGfCbmQGHfCl
tqSmNPNcVrNNmPtHfMZbCWMCZbHwfdzW
RwgGvFwFdsssgjdb
LnLhhnFDStLJSrtw
HHGCHgfjgjgzNSTgJTDJ
WdMbDBdQLLZPWZVZZmLdzqqNbvzswSzrqqsJwSNw
QQmQBmWQZBmRPGfRjRcHccplpD
zFpjDvzDzWWvWBqDqvQcjsfbHPjHsVjcPGffGc
TttwldmlTPSCCZmhlNggwffhHsshRsGssRsfGHRhHf
twmMMmltZdNNlZLLMrPWzrrprpDr
ChCnpqhzCdndsDCnVpVJBgLTpTtgVggL
swHPGmwHbjPjvjQsGLSgVMtgSLLTVBHSrg
mRPPwvbvbvQjljvbPvPlfwWhsnhZDzFnWhDFdWcdhZWl
ndnSZSfNgmmSZlbllglVVsBCfqFBfttpPFtsqV
DJwcMzJDwLcMHMDRMRrtVWFWtCBqBWpWCzzP
wvcHhRcTJJvTRLRvRhwTQTHdQjqqGdgnbGQgSbqSbgNSdl
MHtvCFtCFvMvBDcVfjhhQf
WJZWSgZbdlwZZsbTbwDhcBVQhmTTTcfRjmjD
WwwrBwZsSJWwFttnLrMCpntn
BJgWJVSPVzNJNPZPVHBJHBfdhbjDnbQndnBnBQfndQ
MwMsNmNpmsrrwLssvsvpRpvQnCqbCCQfCfhhbDnqCfqLqn
MFvtrmGMFwRmFvcvMFmsWSVHWVHNzgZWTtWSTPSZ
wWNsHswfsWnVvNZvMLDRRpcLMBDnBlpR
QFdhqzFbzmgbhzQhFSLRrpBBBRcRBcLldpWc
SbmhzChGhqgQFFzqGJhwvtvNWCtVjVfwwZjHjT
WHGDHBRQHDBBQpGzBzQlDVjPmmVtbqbLmtPtzvmZbb
nchMhdMgCMJcwvLZvVcffjZfmq
ndhghggCwSnhCnnTNSdBpFWQFNLFlpHQpDLRGl
FZdjbqjjZTFbRHfVwgttfBZrBpBB
LJQFJWhPPnQJJQCPQpzwphczchrgtrfBrf
FFWvnlLGDGRMDHjm
jCsCNZTTfjcQMLWTtvHQvG
mgwdznllpRgBgpBDdnDDzBDmdrMHMLjtvvHJtLrWrHQWvQtv
lgjDpVSppmqsqssVPZPf
NcNwtChhScfPtwVRgmvWvWvNllVg
nZGJMJLnqqGzLzrGQvlmmgrTmWWlcVvR
LLdMzdssLZnMMbbZBJGsqLcjwjwPffpdtjFpFhPdFwjFFD
wQvZHfbLjrMZSVWMZSdZ
wwqhwtTRRpCCMmsnVlWGMpWM
qBDcqBTBtJCBDqctRDCBCzDFLfbfLwvfbQcFHNvLvvfQjQ
QDFwwvLJFjWQvWgGNnfNftCJPggP
lpqdpqqrqBBpTTBsprpsTsbrNGPfCPzftfPMMGztfzNgNflM
sqdpHrcbmTTHTdcspZNbZmDSQLQjhhSSWwvWWFHQvLLh
gjQHjfqgVSqjqMSnCRMDMvCGZDvCNCNJ
FnLswmdphJchhcZd
bslwFmLTTbWWmlmbBLwlpBHPPBjqHjBSgPSSSfrVnzqQ
JddrdBgJpChGfDLDDpcm
WnSFqRhnjjWSSlSFWqTFVSLGbDLcbLDDvDLztbftVzcv
nRWjsRjQFnQRshgwJCZwJrhw
rnbZwrcZQdpsLpHDVWBBBgDv
NNTSFCqJFMMMSJqTFFPttmNPlwgLvWvgWBffWlClLfvlVCVv
TMTTJSJFRNwmjZsnGsrhnZZR
CMZhZstZqlClJScfBrfHHHFWFFHh
RwGdpjdGVGdmNHcfczpJWWpf
dDJnQbdnwQdbggSPgPsnlqSn
hnczBfznJFmzhnzJLLMwLjmQrppgCjqC
HRRlrPRZSDtQjjPQNMPpPw
ttlZHDRGRlRtGltTtGvcbnFrbvWJhvWhndcFWF
hRsCqfRdqdqfqdtqmMmmZpStMTZZSTnp
VjDDjWbzPFFwPRPFDjzLSnTMnZSGZMnWpLLmpL
PVHFNBzDRbDwzjNNzHdvdHvlhhcHgvCglf
cRLLzhRBhLhVLLLMnnRVVhtJgNgPJmPgPvDPJgvzvvNJJP
lpFffHHWHbbMNPQgsgJJmW
lpdplSpHSHlCCTbSpSFwdnLqMwcdBRRRLdRnth
csZWwFFZbtbztcZttbnRbGlbJhqfnhlpnn
mmTjjBLdSMQdBggQHNgTmjgJJRqRhCJphRHRRClGnhlWJJ
MNMMNTmBLDmTLDSQmFDvvzwPWzsvDZFFcD
fmflRfGcCtGDDbRCRlhCmmBJQsvBBWWvtWBnBJvLLvLn
GFjdMMFGzVQvWWVWJQVJ
wjTZzjrFTHgZjTFjMwGRcfRcqlbhcccrcmmlqC
PvnvjHPnjRPPqvvnPqZNNWHTSdVDHwTLSDWDlTTTLl
MJswhgppwspJpJsGbhchbcMlLLDQdVSgVldtSlSlSTdTld
FwzwJzpfzzJRPZmPBvvmBF
ZrJhslhhclDJhrDZllcGGrZnFCCWNMCHfsFbPddwdwWwCCfP
BTjgvtBBSLqqBSLVzgTpBVMFHNPHwFbSwMPbMwbHMCPf
fRtTppTjRJmrmJDR
hJzcQcwMQhcQwZHVNmSmPCQZSm
tpfvgWFFGGfRWFBqtPrHCmSwPPVPwgVgPC
fpWBRqWqqGRpWWRpFWssMnLczJnnwcjscsJnLhbz
HggzqvvrgtlrrHlQtvlJqFdfBLFJBbnfbBBdfBBfFd
SVVmpZVsRLZGGZwRGZSTDWBbmBFBdFBnnFCnWnBfhB
wpjVppVpZLRTGZQrrjjHNzcgvjrg
GsMgMCFGMsbwbbSwwglS
mJDMqRJzTRVdPWttPffZppBBfS
dzcdqQcmQMGnHCsQ
HlPgHwlHVwsVPPFVsNwzwLnmRmCCMcJcMRRLqGtRjCjJ
dWhrWdDnmrrRMRct
BZDvBpfTDDfBddWTZhDddWDQgwsFsHVvVHgHHznFQFznlQ
FjmqFqJJGsBGGhsZdRDDRZZjTvdMdt
NnHPngrPbPpbPfVvfCrnwccTwrRwTZdcdTZLwRTD
lbbPvHWlHCNNnJzGJlQsJlsBBq
mLNRpRDcDDCmCpzTcCPcNTTDrrnrPrPvjqtqhnhhhqQQjrhF
dVMlVGgVSlGgWgsWSZlwtrHvqMjhnrhQjnBBqHqvhv
WJfsVwZfJGfdLNmcmCcCtztf
bBTtWQDMpDSjjznztCLd
JhfZfrNPPwNJrhqRcJjfCCCzLSGLbGGjGznj
RsPRZhwNbgbrgbBpBWWsQFvDMTDF
rqLzTmqqMmmwrqwrJJwPTplhbRDphhdhDvnnhnRDnP
ggfctFfSRNgZFBRVfFfWfZNSvjjvjCChpDhhDCCjdvlDDD
NQfFRFgtffcGfcgZcfBqJrGLJJqLTwwmmHJrqz
WVGJCltslWJsbbtNJbDJPMrPLBNLrpNjfqQQMMNr
nTFnTmdRnTmmmdHZZnPQpLMfMZPMPqLZDqLf
HRvRFRTvvvRmnvFcFRndcwVVJWWGbsJGwgttlCzJlGDg
rfwdfLwLwwNdwbLcgCCTtFPvFPzWtWVtzvJF
MpjqDDQQRmQjDspzTWWthVWRvRVRzn
ZTQQQqlQZsqlrlgwNlgLcSdf
QqBNgbNNJvcgnbBQQgJjQZSJWppPWsGpDpSGPpWTRDTp
rrldddlrChLFpWPSSPWvPGCR
fhFltdLhHwVfvgcZZqjzNNtQgt
BMCMQZSMvSZSQTQTWvMSvrrTccgGRNljBFglGRgBVgFGRRGc
ffmdmwdLbDJwFVNlGRRVVbRR
dwJnJLfJwLhnPQSSNMvMqhMP
RpQDDCHQwMRpwpNMCwpZhgZrPSqQdgrqPqPZqq
zbtzmzzGbVzJSTScZqrthddT
LVLJlLBnGnWNHMDRHnCCDh
wQDcLhScLLBStHhctwVwsBjNTrssTjgNsgGCgssj
fWblbRqMZWRffldFZpqWRbTCsPrsTNgnCQjgTrMGGgsr
FZRpqdlWWbZZJJJppWWdHQVmLLDmzwDcLSVcQJLw
DlllCDRlflQRsRnlCCBhzvLFVhJzJVzh
jHjNNqdHqwNgpgqNcpgZwjZhJJbvdbPVhBbbvVvBvLzbPF
cpcgSwWqgVnftSSTTf
QmLZvjwwmDWFLNLqbqfF
JJSBMBpdpgSBRBHJgdHBNVbsgqVlqfbVFVlqVFqb
pCttfdCRMDCjhPDhmQ
nDjnzdcDbDtGdcQTvssHssHbbfTq
lVgRgpWCrMRVgVglRLjJmQFHqHHHsqFvFfTQsHSfMQ
VrlpmWmgWppllLpPPWPWljPBntcnhdBNPhZBZhGdDd
fQsRRfBHvRRjjvSgTwCTphdlphqBpn
WJrDZbrNJZPWJcZmWLrrbPczpCnTddqdshqhDgTTwlwwTqhs
brcPbZPcWGWssGtVfVMjFF
GBZgqHhHGBZcLcGMjVJhrjrbStjbbb
wTTNFPQpvDQTQLQvntRSbRJjVnrSDtJV
TTwzQzPPTvmWLTTcdWfdGdZgfZHfZH
VswNdtMgdsvQDJgdVNgQrrLlhrSCLflqfLrDnzLL
RbjRbcZGTGpRbzrTqChtflhSnt
mBjPRWpGcmBZGZpmZjtRbQQVwJgsPsQMsvNJdsVFdN
TDmDTTThRDbbZtpTSnhRcFJMglldGdrJltgBMtMB
jNjsWLqqqvzWswNqdccrFNBFMBlMrdMr
LjjqwvsWCwjHwQHHTbDHbbRZmnRd
qtnncCFCqjwjcDwnFQSGGRSBGBSvFdvrhR
JPZzpgPJWNJBRhqSvQvBgB
mZmJmMPWzmWmPLbmqTTJPfDHnHjCjflHjflfDcfMwC
VfJrplTqzVVlVVlpfTMJfqDtWSLgbHgDjSHLgjHWjrjb
vhGsRdsNRPRvvRhjSZjJLPDjDJttbZ
GvGhQmhBnQNmnJGvhdFqBFVqcllqMBCMlVCc
cnmsmmDCBDmcmvVnsVBtQQWWrrltQlrBfRHB
hqphjLZFhMMjZQGQWGfftQWf
JhFqhbjSLjMpqzJFbSLdFSggDwgmgcDVsVzVfvNvDvsm
mRdspZZggmbJbpDcDJDqPhvvGvrnVqPqpzvVVV
dSjSjQtwfQlLwjQSPrrqVqvGPGPVPjnP
fMQtFQLHSfLtHFttHMflfwHWgFJsDbmJRDWdFBRJWmcbsJ
lMMftpGHMntGtMMGpJnSPQvwzvvHscsdQscQsscQss
hbfmWWLggVFjVVQdvbQvbcdzwbRR
qZWFWTgTZhrLhLmWGlGGrSnJSNfGtSll
QVGdmgVGLdllrGVgnZrdBvNvBSSjSvWtBHBqNNmv
ThRjRFwpfTNBBqSBSSvT
MFzbzfDjFFCCFJpRpFwbwJRnZndDgsZlgLrQLDgdlglZnD
PjbDQvQPjbjjvtrQcpHBWnLfqnqGcHngnLdf
VTwmMzMFnqHBwHBB
lMTTNCzCZmClVVQjstQpPlprlHSp
MtMqBtwFFsMRZMMZMjJGZR
PHbpbPqHmVVCRGGCGRjv
pdrrgdrqPbpzrHQbbqQPLHcQBsNtNnwzNhSFfhtfffFffwDs
cPjpPcRJJVtSGGtJtffN
CldrWrmDHnHnmWmTWsNzzfbbFNhtsbbNGGfb
LCDTldCmnLLVwPVgQt
NJpFrbpFZhlzhbzzqCtgMCgRbqCRMMtm
dWcnntHQWcjRHHCVMqCsRg
vLnQDvnGptrzLPwp
dwTwTHwvZHqTrTRTlWtfzt
GbFFbFBPnGQBQcQcLbPjFtCSrtgMMLrMfrrWgRgStr
PnBnQhsGPQbnJRqwdRDwvJNDwq
zzjWtnQRntzSBjZccmZmmC
bdJdbdfJdFCGflfGwBJsmhmSMgmgcsmM
qDbGDbvrFbbvGdvdffFdTTWQzrTWTLWLNnttLnCH
WMMssTqCbpGzGSlmzLbG
FVdwftRRVDZPcdZtZVRcjmvQSzzLrmlrclSGms
DVfPFZDZRfVFPFHRPFRCphhMpCTHTsNgqHNWgn
LzBGCRjHCnmnHzdzLLjHwGplflfrfPcgtflgLFLlcrcc
hqhDMSSqWZsDMVWTVWmTcglglflFVflFrpPtggcv
sshbJsTsqWMbqNWbhbzzjzGBQjCJQnzwmzBw
QnpfLpbLfGbvgjHzjgmNRqbz
SSMBMZDwMwFsqgSPPmHlHj
BhMFJhMBddDMFcTdMVddrQfvrnCqnCffpqnVCLtC
lhrTZNJZjCRjSCvRSlTSLlrvFnMHQhVDnqnmqqqmQqPDHmqF
wcpGdtwtwzcbpzggCFPMqtMQmQqmVqHP
bBfwgWbfwBdGpppGGGbcBsTRLlZTZsTRTsRNZClJsvZL
nJLgNcQDNMlQHMvCbv
zphFpmTszmwhGGFhhtppNfffVlvZvHCCVZzbfzvS
mTTsmTRGstsFhWwtWjPRdjnJdjJnLjcLNd
MfBDjllflHLTpDhhppDDbp
NZBBnGJNnNPWTcTTmVhZCh
PSzgSgwrnzrSzBGJSJrSLQqfMHQfqgRgfjHLljll
RgbNmBbqgWHWRNRqHtcMlMwJJjcDtVlD
SzpFLGPddSGnnSLQZLtJJcclDlVjDQwMhDcc
LtTZCTttRqqqvqTN`