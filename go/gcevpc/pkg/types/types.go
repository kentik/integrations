package types

import (
	"encoding/binary"
	"math"
	"net"
	"strconv"
)

const (
	RETRANSMITTED_IN_PKTS  = "RETRANSMITTED_IN_PKTS"
	RETRANSMITTED_OUT_PKTS = "RETRANSMITTED_OUT_PKTS"
	OOORDER_IN_PKTS        = "OOORDER_IN_PKTS"
	OOORDER_OUT_PKTS       = "OOORDER_OUT_PKTS"
	FRAGMENTS              = "FRAGMENTS"
	CLIENT_NW_LATENCY_MS   = "CLIENT_NW_LATENCY_MS"
	SERVER_NW_LATENCY_MS   = "SERVER_NW_LATENCY_MS"
	APPL_LATENCY_MS        = "APPL_LATENCY_MS"
	KFLOW_HTTP_URL         = "KFLOW_HTTP_URL"
	KFLOW_HTTP_STATUS      = "KFLOW_HTTP_STATUS"
	KFLOW_HTTP_UA          = "KFLOW_HTTP_UA"
	KFLOW_HTTP_REFERER     = "KFLOW_HTTP_REFERER"
	KFLOW_DNS_QUERY        = "KFLOW_DNS_QUERY"
	KFLOW_DNS_QUERY_TYPE   = "KFLOW_DNS_QUERY_TYPE"
	KFLOW_DNS_RET_CODE     = "KFLOW_DNS_RET_CODE"
	KFLOW_HTTP_HOST        = "KFLOW_HTTP_HOST"
	KFLOW_DNS_RESPONSE     = "KFLOW_DNS_RESPONSE"

	KFLOW_WINDOW_SIZE       = "c_window_size"
	KFLOW_RCV_MSS           = "c_rcv_mss"
	KFLOW_SND_MSS           = "c_snd_mss"
	KFLOW_CACHE_STATUS      = "c_cache_status"
	KFLOW_SERVER_REGION     = "c_server_region"
	KFLOW_SERVER_DATACENTER = "c_server_datacenter"
	KFLOW_FASTLY_NEXTHOP    = "c_fastly_next_hop"

	EST_NUM_PKTS = 1500
)

func getUInt32(s *string) uint32 {
	n, _ := strconv.Atoi(*s)
	return uint32(n)
}

func getMSUInt32(s *string) uint32 {
	n, _ := strconv.Atoi(*s)
	nms := float64(n) / 1000
	return uint32(math.Floor(nms))
}

func getUInt64(s *string) uint64 {
	n, _ := strconv.Atoi(*s)
	return uint64(n)
}

func getInt64(s *string) int64 {
	n, _ := strconv.Atoi(*s)
	return int64(n)
}

func getPkts(bytes uint64, mss *string) uint64 {

	maxSeg, _ := strconv.Atoi(*mss)
	if maxSeg == 0 {
		maxSeg = EST_NUM_PKTS
	}

	pkts := float64(bytes) / float64(maxSeg)
	if pkts == 0 {
		return 1
	} else {
		return uint64(math.Ceil(pkts))
	}
}

func PackIP(base *string) (uint32, []byte) {
	ipr := net.ParseIP(*base)
	if ipr == nil {
		return 0, nil
	}

	if v4 := ipr.To4(); v4 != nil {
		ipv4 := binary.BigEndian.Uint32(ipr.To4())
		return ipv4, nil
	} else {
		return 0, ipr.To16()
	}
}
