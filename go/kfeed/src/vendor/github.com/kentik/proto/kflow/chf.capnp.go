package chf

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
	"math"
)

type Custom C.Struct
type CustomValue Custom
type CustomValue_Which uint16

const (
	CUSTOMVALUE_UINT32VAL  CustomValue_Which = 0
	CUSTOMVALUE_FLOAT32VAL CustomValue_Which = 1
	CUSTOMVALUE_STRVAL     CustomValue_Which = 2
)

func NewCustom(s *C.Segment) Custom            { return Custom(s.NewStruct(16, 1)) }
func NewRootCustom(s *C.Segment) Custom        { return Custom(s.NewRootStruct(16, 1)) }
func AutoNewCustom(s *C.Segment) Custom        { return Custom(s.NewStructAR(16, 1)) }
func ReadRootCustom(s *C.Segment) Custom       { return Custom(s.Root(0).ToStruct()) }
func (s Custom) Id() uint32                    { return C.Struct(s).Get32(0) }
func (s Custom) SetId(v uint32)                { C.Struct(s).Set32(0, v) }
func (s Custom) Value() CustomValue            { return CustomValue(s) }
func (s CustomValue) Which() CustomValue_Which { return CustomValue_Which(C.Struct(s).Get16(8)) }
func (s CustomValue) Uint32Val() uint32        { return C.Struct(s).Get32(4) }
func (s CustomValue) SetUint32Val(v uint32)    { C.Struct(s).Set16(8, 0); C.Struct(s).Set32(4, v) }
func (s CustomValue) Float32Val() float32      { return math.Float32frombits(C.Struct(s).Get32(4)) }
func (s CustomValue) SetFloat32Val(v float32) {
	C.Struct(s).Set16(8, 1)
	C.Struct(s).Set32(4, math.Float32bits(v))
}
func (s CustomValue) StrVal() string      { return C.Struct(s).GetObject(0).ToText() }
func (s CustomValue) StrValBytes() []byte { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s CustomValue) SetStrVal(v string) {
	C.Struct(s).Set16(8, 2)
	C.Struct(s).SetObject(0, s.Segment.NewText(v))
}
func (s Custom) IsDimension() bool     { return C.Struct(s).Get1(80) }
func (s Custom) SetIsDimension(v bool) { C.Struct(s).Set1(80, v) }
func (s Custom) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"id\":")
	if err != nil {
		return err
	}
	{
		s := s.Id()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"value\":")
	if err != nil {
		return err
	}
	{
		s := s.Value()
		err = b.WriteByte('{')
		if err != nil {
			return err
		}
		if s.Which() == CUSTOMVALUE_UINT32VAL {
			_, err = b.WriteString("\"uint32Val\":")
			if err != nil {
				return err
			}
			{
				s := s.Uint32Val()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == CUSTOMVALUE_FLOAT32VAL {
			_, err = b.WriteString("\"float32Val\":")
			if err != nil {
				return err
			}
			{
				s := s.Float32Val()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == CUSTOMVALUE_STRVAL {
			_, err = b.WriteString("\"strVal\":")
			if err != nil {
				return err
			}
			{
				s := s.StrVal()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		err = b.WriteByte('}')
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"isDimension\":")
	if err != nil {
		return err
	}
	{
		s := s.IsDimension()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Custom) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Custom) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("id = ")
	if err != nil {
		return err
	}
	{
		s := s.Id()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("value = ")
	if err != nil {
		return err
	}
	{
		s := s.Value()
		err = b.WriteByte('(')
		if err != nil {
			return err
		}
		if s.Which() == CUSTOMVALUE_UINT32VAL {
			_, err = b.WriteString("uint32Val = ")
			if err != nil {
				return err
			}
			{
				s := s.Uint32Val()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == CUSTOMVALUE_FLOAT32VAL {
			_, err = b.WriteString("float32Val = ")
			if err != nil {
				return err
			}
			{
				s := s.Float32Val()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		if s.Which() == CUSTOMVALUE_STRVAL {
			_, err = b.WriteString("strVal = ")
			if err != nil {
				return err
			}
			{
				s := s.StrVal()
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
				if err != nil {
					return err
				}
			}
		}
		err = b.WriteByte(')')
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("isDimension = ")
	if err != nil {
		return err
	}
	{
		s := s.IsDimension()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Custom) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Custom_List C.PointerList

func NewCustomList(s *C.Segment, sz int) Custom_List {
	return Custom_List(s.NewCompositeList(16, 1, sz))
}
func (s Custom_List) Len() int        { return C.PointerList(s).Len() }
func (s Custom_List) At(i int) Custom { return Custom(C.PointerList(s).At(i).ToStruct()) }
func (s Custom_List) ToArray() []Custom {
	n := s.Len()
	a := make([]Custom, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Custom_List) Set(i int, item Custom) { C.PointerList(s).Set(i, C.Object(item)) }

type CHF C.Struct

func NewCHF(s *C.Segment) CHF                { return CHF(s.NewStruct(224, 14)) }
func NewRootCHF(s *C.Segment) CHF            { return CHF(s.NewRootStruct(224, 14)) }
func AutoNewCHF(s *C.Segment) CHF            { return CHF(s.NewStructAR(224, 14)) }
func ReadRootCHF(s *C.Segment) CHF           { return CHF(s.Root(0).ToStruct()) }
func (s CHF) TimestampNano() int64           { return int64(C.Struct(s).Get64(0)) }
func (s CHF) SetTimestampNano(v int64)       { C.Struct(s).Set64(0, uint64(v)) }
func (s CHF) DstAs() uint32                  { return C.Struct(s).Get32(8) }
func (s CHF) SetDstAs(v uint32)              { C.Struct(s).Set32(8, v) }
func (s CHF) DstGeo() uint32                 { return C.Struct(s).Get32(12) }
func (s CHF) SetDstGeo(v uint32)             { C.Struct(s).Set32(12, v) }
func (s CHF) DstMac() uint32                 { return C.Struct(s).Get32(16) }
func (s CHF) SetDstMac(v uint32)             { C.Struct(s).Set32(16, v) }
func (s CHF) HeaderLen() uint32              { return C.Struct(s).Get32(20) }
func (s CHF) SetHeaderLen(v uint32)          { C.Struct(s).Set32(20, v) }
func (s CHF) InBytes() uint64                { return C.Struct(s).Get64(24) }
func (s CHF) SetInBytes(v uint64)            { C.Struct(s).Set64(24, v) }
func (s CHF) InPkts() uint64                 { return C.Struct(s).Get64(32) }
func (s CHF) SetInPkts(v uint64)             { C.Struct(s).Set64(32, v) }
func (s CHF) InputPort() uint32              { return C.Struct(s).Get32(40) }
func (s CHF) SetInputPort(v uint32)          { C.Struct(s).Set32(40, v) }
func (s CHF) IpSize() uint32                 { return C.Struct(s).Get32(44) }
func (s CHF) SetIpSize(v uint32)             { C.Struct(s).Set32(44, v) }
func (s CHF) Ipv4DstAddr() uint32            { return C.Struct(s).Get32(48) }
func (s CHF) SetIpv4DstAddr(v uint32)        { C.Struct(s).Set32(48, v) }
func (s CHF) Ipv4SrcAddr() uint32            { return C.Struct(s).Get32(52) }
func (s CHF) SetIpv4SrcAddr(v uint32)        { C.Struct(s).Set32(52, v) }
func (s CHF) L4DstPort() uint32              { return C.Struct(s).Get32(56) }
func (s CHF) SetL4DstPort(v uint32)          { C.Struct(s).Set32(56, v) }
func (s CHF) L4SrcPort() uint32              { return C.Struct(s).Get32(60) }
func (s CHF) SetL4SrcPort(v uint32)          { C.Struct(s).Set32(60, v) }
func (s CHF) OutputPort() uint32             { return C.Struct(s).Get32(64) }
func (s CHF) SetOutputPort(v uint32)         { C.Struct(s).Set32(64, v) }
func (s CHF) Protocol() uint32               { return C.Struct(s).Get32(68) }
func (s CHF) SetProtocol(v uint32)           { C.Struct(s).Set32(68, v) }
func (s CHF) SampledPacketSize() uint32      { return C.Struct(s).Get32(72) }
func (s CHF) SetSampledPacketSize(v uint32)  { C.Struct(s).Set32(72, v) }
func (s CHF) SrcAs() uint32                  { return C.Struct(s).Get32(76) }
func (s CHF) SetSrcAs(v uint32)              { C.Struct(s).Set32(76, v) }
func (s CHF) SrcGeo() uint32                 { return C.Struct(s).Get32(80) }
func (s CHF) SetSrcGeo(v uint32)             { C.Struct(s).Set32(80, v) }
func (s CHF) SrcMac() uint32                 { return C.Struct(s).Get32(84) }
func (s CHF) SetSrcMac(v uint32)             { C.Struct(s).Set32(84, v) }
func (s CHF) TcpFlags() uint32               { return C.Struct(s).Get32(88) }
func (s CHF) SetTcpFlags(v uint32)           { C.Struct(s).Set32(88, v) }
func (s CHF) Tos() uint32                    { return C.Struct(s).Get32(92) }
func (s CHF) SetTos(v uint32)                { C.Struct(s).Set32(92, v) }
func (s CHF) VlanIn() uint32                 { return C.Struct(s).Get32(96) }
func (s CHF) SetVlanIn(v uint32)             { C.Struct(s).Set32(96, v) }
func (s CHF) VlanOut() uint32                { return C.Struct(s).Get32(100) }
func (s CHF) SetVlanOut(v uint32)            { C.Struct(s).Set32(100, v) }
func (s CHF) Ipv4NextHop() uint32            { return C.Struct(s).Get32(104) }
func (s CHF) SetIpv4NextHop(v uint32)        { C.Struct(s).Set32(104, v) }
func (s CHF) MplsType() uint32               { return C.Struct(s).Get32(108) }
func (s CHF) SetMplsType(v uint32)           { C.Struct(s).Set32(108, v) }
func (s CHF) OutBytes() uint64               { return C.Struct(s).Get64(112) }
func (s CHF) SetOutBytes(v uint64)           { C.Struct(s).Set64(112, v) }
func (s CHF) OutPkts() uint64                { return C.Struct(s).Get64(120) }
func (s CHF) SetOutPkts(v uint64)            { C.Struct(s).Set64(120, v) }
func (s CHF) TcpRetransmit() uint32          { return C.Struct(s).Get32(128) }
func (s CHF) SetTcpRetransmit(v uint32)      { C.Struct(s).Set32(128, v) }
func (s CHF) SrcFlowTags() string            { return C.Struct(s).GetObject(0).ToText() }
func (s CHF) SrcFlowTagsBytes() []byte       { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s CHF) SetSrcFlowTags(v string)        { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s CHF) DstFlowTags() string            { return C.Struct(s).GetObject(1).ToText() }
func (s CHF) DstFlowTagsBytes() []byte       { return C.Struct(s).GetObject(1).ToDataTrimLastByte() }
func (s CHF) SetDstFlowTags(v string)        { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s CHF) SampleRate() uint32             { return C.Struct(s).Get32(132) }
func (s CHF) SetSampleRate(v uint32)         { C.Struct(s).Set32(132, v) }
func (s CHF) DeviceId() uint32               { return C.Struct(s).Get32(136) }
func (s CHF) SetDeviceId(v uint32)           { C.Struct(s).Set32(136, v) }
func (s CHF) FlowTags() string               { return C.Struct(s).GetObject(2).ToText() }
func (s CHF) FlowTagsBytes() []byte          { return C.Struct(s).GetObject(2).ToDataTrimLastByte() }
func (s CHF) SetFlowTags(v string)           { C.Struct(s).SetObject(2, s.Segment.NewText(v)) }
func (s CHF) Timestamp() int64               { return int64(C.Struct(s).Get64(144)) }
func (s CHF) SetTimestamp(v int64)           { C.Struct(s).Set64(144, uint64(v)) }
func (s CHF) DstBgpAsPath() string           { return C.Struct(s).GetObject(3).ToText() }
func (s CHF) DstBgpAsPathBytes() []byte      { return C.Struct(s).GetObject(3).ToDataTrimLastByte() }
func (s CHF) SetDstBgpAsPath(v string)       { C.Struct(s).SetObject(3, s.Segment.NewText(v)) }
func (s CHF) DstBgpCommunity() string        { return C.Struct(s).GetObject(4).ToText() }
func (s CHF) DstBgpCommunityBytes() []byte   { return C.Struct(s).GetObject(4).ToDataTrimLastByte() }
func (s CHF) SetDstBgpCommunity(v string)    { C.Struct(s).SetObject(4, s.Segment.NewText(v)) }
func (s CHF) SrcBgpAsPath() string           { return C.Struct(s).GetObject(5).ToText() }
func (s CHF) SrcBgpAsPathBytes() []byte      { return C.Struct(s).GetObject(5).ToDataTrimLastByte() }
func (s CHF) SetSrcBgpAsPath(v string)       { C.Struct(s).SetObject(5, s.Segment.NewText(v)) }
func (s CHF) SrcBgpCommunity() string        { return C.Struct(s).GetObject(6).ToText() }
func (s CHF) SrcBgpCommunityBytes() []byte   { return C.Struct(s).GetObject(6).ToDataTrimLastByte() }
func (s CHF) SetSrcBgpCommunity(v string)    { C.Struct(s).SetObject(6, s.Segment.NewText(v)) }
func (s CHF) SrcNextHopAs() uint32           { return C.Struct(s).Get32(140) }
func (s CHF) SetSrcNextHopAs(v uint32)       { C.Struct(s).Set32(140, v) }
func (s CHF) DstNextHopAs() uint32           { return C.Struct(s).Get32(152) }
func (s CHF) SetDstNextHopAs(v uint32)       { C.Struct(s).Set32(152, v) }
func (s CHF) SrcGeoRegion() uint32           { return C.Struct(s).Get32(156) }
func (s CHF) SetSrcGeoRegion(v uint32)       { C.Struct(s).Set32(156, v) }
func (s CHF) DstGeoRegion() uint32           { return C.Struct(s).Get32(160) }
func (s CHF) SetDstGeoRegion(v uint32)       { C.Struct(s).Set32(160, v) }
func (s CHF) SrcGeoCity() uint32             { return C.Struct(s).Get32(164) }
func (s CHF) SetSrcGeoCity(v uint32)         { C.Struct(s).Set32(164, v) }
func (s CHF) DstGeoCity() uint32             { return C.Struct(s).Get32(168) }
func (s CHF) SetDstGeoCity(v uint32)         { C.Struct(s).Set32(168, v) }
func (s CHF) Big() bool                      { return C.Struct(s).Get1(1376) }
func (s CHF) SetBig(v bool)                  { C.Struct(s).Set1(1376, v) }
func (s CHF) SampleAdj() bool                { return C.Struct(s).Get1(1377) }
func (s CHF) SetSampleAdj(v bool)            { C.Struct(s).Set1(1377, v) }
func (s CHF) Ipv4DstNextHop() uint32         { return C.Struct(s).Get32(176) }
func (s CHF) SetIpv4DstNextHop(v uint32)     { C.Struct(s).Set32(176, v) }
func (s CHF) Ipv4SrcNextHop() uint32         { return C.Struct(s).Get32(180) }
func (s CHF) SetIpv4SrcNextHop(v uint32)     { C.Struct(s).Set32(180, v) }
func (s CHF) SrcRoutePrefix() uint32         { return C.Struct(s).Get32(184) }
func (s CHF) SetSrcRoutePrefix(v uint32)     { C.Struct(s).Set32(184, v) }
func (s CHF) DstRoutePrefix() uint32         { return C.Struct(s).Get32(188) }
func (s CHF) SetDstRoutePrefix(v uint32)     { C.Struct(s).Set32(188, v) }
func (s CHF) SrcRouteLength() uint8          { return C.Struct(s).Get8(173) }
func (s CHF) SetSrcRouteLength(v uint8)      { C.Struct(s).Set8(173, v) }
func (s CHF) DstRouteLength() uint8          { return C.Struct(s).Get8(174) }
func (s CHF) SetDstRouteLength(v uint8)      { C.Struct(s).Set8(174, v) }
func (s CHF) SrcSecondAsn() uint32           { return C.Struct(s).Get32(192) }
func (s CHF) SetSrcSecondAsn(v uint32)       { C.Struct(s).Set32(192, v) }
func (s CHF) DstSecondAsn() uint32           { return C.Struct(s).Get32(196) }
func (s CHF) SetDstSecondAsn(v uint32)       { C.Struct(s).Set32(196, v) }
func (s CHF) SrcThirdAsn() uint32            { return C.Struct(s).Get32(200) }
func (s CHF) SetSrcThirdAsn(v uint32)        { C.Struct(s).Set32(200, v) }
func (s CHF) DstThirdAsn() uint32            { return C.Struct(s).Get32(204) }
func (s CHF) SetDstThirdAsn(v uint32)        { C.Struct(s).Set32(204, v) }
func (s CHF) Ipv6DstAddr() []byte            { return C.Struct(s).GetObject(7).ToData() }
func (s CHF) SetIpv6DstAddr(v []byte)        { C.Struct(s).SetObject(7, s.Segment.NewData(v)) }
func (s CHF) Ipv6SrcAddr() []byte            { return C.Struct(s).GetObject(8).ToData() }
func (s CHF) SetIpv6SrcAddr(v []byte)        { C.Struct(s).SetObject(8, s.Segment.NewData(v)) }
func (s CHF) SrcEthMac() uint64              { return C.Struct(s).Get64(208) }
func (s CHF) SetSrcEthMac(v uint64)          { C.Struct(s).Set64(208, v) }
func (s CHF) DstEthMac() uint64              { return C.Struct(s).Get64(216) }
func (s CHF) SetDstEthMac(v uint64)          { C.Struct(s).Set64(216, v) }
func (s CHF) Custom() Custom_List            { return Custom_List(C.Struct(s).GetObject(9)) }
func (s CHF) SetCustom(v Custom_List)        { C.Struct(s).SetObject(9, C.Object(v)) }
func (s CHF) Ipv6SrcNextHop() []byte         { return C.Struct(s).GetObject(10).ToData() }
func (s CHF) SetIpv6SrcNextHop(v []byte)     { C.Struct(s).SetObject(10, s.Segment.NewData(v)) }
func (s CHF) Ipv6DstNextHop() []byte         { return C.Struct(s).GetObject(11).ToData() }
func (s CHF) SetIpv6DstNextHop(v []byte)     { C.Struct(s).SetObject(11, s.Segment.NewData(v)) }
func (s CHF) Ipv6SrcRoutePrefix() []byte     { return C.Struct(s).GetObject(12).ToData() }
func (s CHF) SetIpv6SrcRoutePrefix(v []byte) { C.Struct(s).SetObject(12, s.Segment.NewData(v)) }
func (s CHF) Ipv6DstRoutePrefix() []byte     { return C.Struct(s).GetObject(13).ToData() }
func (s CHF) SetIpv6DstRoutePrefix(v []byte) { C.Struct(s).SetObject(13, s.Segment.NewData(v)) }
func (s CHF) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"timestampNano\":")
	if err != nil {
		return err
	}
	{
		s := s.TimestampNano()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstAs\":")
	if err != nil {
		return err
	}
	{
		s := s.DstAs()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstGeo\":")
	if err != nil {
		return err
	}
	{
		s := s.DstGeo()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstMac\":")
	if err != nil {
		return err
	}
	{
		s := s.DstMac()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"headerLen\":")
	if err != nil {
		return err
	}
	{
		s := s.HeaderLen()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"inBytes\":")
	if err != nil {
		return err
	}
	{
		s := s.InBytes()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"inPkts\":")
	if err != nil {
		return err
	}
	{
		s := s.InPkts()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"inputPort\":")
	if err != nil {
		return err
	}
	{
		s := s.InputPort()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipSize\":")
	if err != nil {
		return err
	}
	{
		s := s.IpSize()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv4DstAddr\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4DstAddr()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv4SrcAddr\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4SrcAddr()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"l4DstPort\":")
	if err != nil {
		return err
	}
	{
		s := s.L4DstPort()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"l4SrcPort\":")
	if err != nil {
		return err
	}
	{
		s := s.L4SrcPort()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"outputPort\":")
	if err != nil {
		return err
	}
	{
		s := s.OutputPort()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"protocol\":")
	if err != nil {
		return err
	}
	{
		s := s.Protocol()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"sampledPacketSize\":")
	if err != nil {
		return err
	}
	{
		s := s.SampledPacketSize()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcAs\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcAs()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcGeo\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcGeo()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcMac\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcMac()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"tcpFlags\":")
	if err != nil {
		return err
	}
	{
		s := s.TcpFlags()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"tos\":")
	if err != nil {
		return err
	}
	{
		s := s.Tos()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"vlanIn\":")
	if err != nil {
		return err
	}
	{
		s := s.VlanIn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"vlanOut\":")
	if err != nil {
		return err
	}
	{
		s := s.VlanOut()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv4NextHop\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4NextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mplsType\":")
	if err != nil {
		return err
	}
	{
		s := s.MplsType()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"outBytes\":")
	if err != nil {
		return err
	}
	{
		s := s.OutBytes()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"outPkts\":")
	if err != nil {
		return err
	}
	{
		s := s.OutPkts()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"tcpRetransmit\":")
	if err != nil {
		return err
	}
	{
		s := s.TcpRetransmit()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcFlowTags\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcFlowTags()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstFlowTags\":")
	if err != nil {
		return err
	}
	{
		s := s.DstFlowTags()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"sampleRate\":")
	if err != nil {
		return err
	}
	{
		s := s.SampleRate()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"deviceId\":")
	if err != nil {
		return err
	}
	{
		s := s.DeviceId()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"flowTags\":")
	if err != nil {
		return err
	}
	{
		s := s.FlowTags()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"timestamp\":")
	if err != nil {
		return err
	}
	{
		s := s.Timestamp()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstBgpAsPath\":")
	if err != nil {
		return err
	}
	{
		s := s.DstBgpAsPath()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstBgpCommunity\":")
	if err != nil {
		return err
	}
	{
		s := s.DstBgpCommunity()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcBgpAsPath\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcBgpAsPath()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcBgpCommunity\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcBgpCommunity()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcNextHopAs\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcNextHopAs()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstNextHopAs\":")
	if err != nil {
		return err
	}
	{
		s := s.DstNextHopAs()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcGeoRegion\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcGeoRegion()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstGeoRegion\":")
	if err != nil {
		return err
	}
	{
		s := s.DstGeoRegion()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcGeoCity\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcGeoCity()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstGeoCity\":")
	if err != nil {
		return err
	}
	{
		s := s.DstGeoCity()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"big\":")
	if err != nil {
		return err
	}
	{
		s := s.Big()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"sampleAdj\":")
	if err != nil {
		return err
	}
	{
		s := s.SampleAdj()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv4DstNextHop\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4DstNextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv4SrcNextHop\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4SrcNextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcRoutePrefix\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcRoutePrefix()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstRoutePrefix\":")
	if err != nil {
		return err
	}
	{
		s := s.DstRoutePrefix()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcRouteLength\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcRouteLength()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstRouteLength\":")
	if err != nil {
		return err
	}
	{
		s := s.DstRouteLength()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcSecondAsn\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcSecondAsn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstSecondAsn\":")
	if err != nil {
		return err
	}
	{
		s := s.DstSecondAsn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcThirdAsn\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcThirdAsn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstThirdAsn\":")
	if err != nil {
		return err
	}
	{
		s := s.DstThirdAsn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv6DstAddr\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6DstAddr()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv6SrcAddr\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6SrcAddr()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"srcEthMac\":")
	if err != nil {
		return err
	}
	{
		s := s.SrcEthMac()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"dstEthMac\":")
	if err != nil {
		return err
	}
	{
		s := s.DstEthMac()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"custom\":")
	if err != nil {
		return err
	}
	{
		s := s.Custom()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteJSON(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv6SrcNextHop\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6SrcNextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv6DstNextHop\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6DstNextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv6SrcRoutePrefix\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6SrcRoutePrefix()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ipv6DstRoutePrefix\":")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6DstRoutePrefix()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s CHF) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s CHF) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("timestampNano = ")
	if err != nil {
		return err
	}
	{
		s := s.TimestampNano()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstAs = ")
	if err != nil {
		return err
	}
	{
		s := s.DstAs()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstGeo = ")
	if err != nil {
		return err
	}
	{
		s := s.DstGeo()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstMac = ")
	if err != nil {
		return err
	}
	{
		s := s.DstMac()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("headerLen = ")
	if err != nil {
		return err
	}
	{
		s := s.HeaderLen()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("inBytes = ")
	if err != nil {
		return err
	}
	{
		s := s.InBytes()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("inPkts = ")
	if err != nil {
		return err
	}
	{
		s := s.InPkts()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("inputPort = ")
	if err != nil {
		return err
	}
	{
		s := s.InputPort()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipSize = ")
	if err != nil {
		return err
	}
	{
		s := s.IpSize()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv4DstAddr = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4DstAddr()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv4SrcAddr = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4SrcAddr()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("l4DstPort = ")
	if err != nil {
		return err
	}
	{
		s := s.L4DstPort()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("l4SrcPort = ")
	if err != nil {
		return err
	}
	{
		s := s.L4SrcPort()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("outputPort = ")
	if err != nil {
		return err
	}
	{
		s := s.OutputPort()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("protocol = ")
	if err != nil {
		return err
	}
	{
		s := s.Protocol()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("sampledPacketSize = ")
	if err != nil {
		return err
	}
	{
		s := s.SampledPacketSize()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcAs = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcAs()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcGeo = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcGeo()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcMac = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcMac()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("tcpFlags = ")
	if err != nil {
		return err
	}
	{
		s := s.TcpFlags()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("tos = ")
	if err != nil {
		return err
	}
	{
		s := s.Tos()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("vlanIn = ")
	if err != nil {
		return err
	}
	{
		s := s.VlanIn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("vlanOut = ")
	if err != nil {
		return err
	}
	{
		s := s.VlanOut()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv4NextHop = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4NextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("mplsType = ")
	if err != nil {
		return err
	}
	{
		s := s.MplsType()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("outBytes = ")
	if err != nil {
		return err
	}
	{
		s := s.OutBytes()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("outPkts = ")
	if err != nil {
		return err
	}
	{
		s := s.OutPkts()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("tcpRetransmit = ")
	if err != nil {
		return err
	}
	{
		s := s.TcpRetransmit()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcFlowTags = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcFlowTags()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstFlowTags = ")
	if err != nil {
		return err
	}
	{
		s := s.DstFlowTags()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("sampleRate = ")
	if err != nil {
		return err
	}
	{
		s := s.SampleRate()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("deviceId = ")
	if err != nil {
		return err
	}
	{
		s := s.DeviceId()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("flowTags = ")
	if err != nil {
		return err
	}
	{
		s := s.FlowTags()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("timestamp = ")
	if err != nil {
		return err
	}
	{
		s := s.Timestamp()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstBgpAsPath = ")
	if err != nil {
		return err
	}
	{
		s := s.DstBgpAsPath()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstBgpCommunity = ")
	if err != nil {
		return err
	}
	{
		s := s.DstBgpCommunity()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcBgpAsPath = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcBgpAsPath()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcBgpCommunity = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcBgpCommunity()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcNextHopAs = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcNextHopAs()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstNextHopAs = ")
	if err != nil {
		return err
	}
	{
		s := s.DstNextHopAs()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcGeoRegion = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcGeoRegion()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstGeoRegion = ")
	if err != nil {
		return err
	}
	{
		s := s.DstGeoRegion()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcGeoCity = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcGeoCity()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstGeoCity = ")
	if err != nil {
		return err
	}
	{
		s := s.DstGeoCity()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("big = ")
	if err != nil {
		return err
	}
	{
		s := s.Big()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("sampleAdj = ")
	if err != nil {
		return err
	}
	{
		s := s.SampleAdj()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv4DstNextHop = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4DstNextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv4SrcNextHop = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv4SrcNextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcRoutePrefix = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcRoutePrefix()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstRoutePrefix = ")
	if err != nil {
		return err
	}
	{
		s := s.DstRoutePrefix()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcRouteLength = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcRouteLength()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstRouteLength = ")
	if err != nil {
		return err
	}
	{
		s := s.DstRouteLength()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcSecondAsn = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcSecondAsn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstSecondAsn = ")
	if err != nil {
		return err
	}
	{
		s := s.DstSecondAsn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcThirdAsn = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcThirdAsn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstThirdAsn = ")
	if err != nil {
		return err
	}
	{
		s := s.DstThirdAsn()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv6DstAddr = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6DstAddr()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv6SrcAddr = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6SrcAddr()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("srcEthMac = ")
	if err != nil {
		return err
	}
	{
		s := s.SrcEthMac()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("dstEthMac = ")
	if err != nil {
		return err
	}
	{
		s := s.DstEthMac()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("custom = ")
	if err != nil {
		return err
	}
	{
		s := s.Custom()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteCapLit(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv6SrcNextHop = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6SrcNextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv6DstNextHop = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6DstNextHop()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv6SrcRoutePrefix = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6SrcRoutePrefix()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ipv6DstRoutePrefix = ")
	if err != nil {
		return err
	}
	{
		s := s.Ipv6DstRoutePrefix()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s CHF) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type CHF_List C.PointerList

func NewCHFList(s *C.Segment, sz int) CHF_List { return CHF_List(s.NewCompositeList(224, 14, sz)) }
func (s CHF_List) Len() int                    { return C.PointerList(s).Len() }
func (s CHF_List) At(i int) CHF                { return CHF(C.PointerList(s).At(i).ToStruct()) }
func (s CHF_List) ToArray() []CHF {
	n := s.Len()
	a := make([]CHF, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s CHF_List) Set(i int, item CHF) { C.PointerList(s).Set(i, C.Object(item)) }

type PackedCHF C.Struct

func NewPackedCHF(s *C.Segment) PackedCHF      { return PackedCHF(s.NewStruct(0, 1)) }
func NewRootPackedCHF(s *C.Segment) PackedCHF  { return PackedCHF(s.NewRootStruct(0, 1)) }
func AutoNewPackedCHF(s *C.Segment) PackedCHF  { return PackedCHF(s.NewStructAR(0, 1)) }
func ReadRootPackedCHF(s *C.Segment) PackedCHF { return PackedCHF(s.Root(0).ToStruct()) }
func (s PackedCHF) Msgs() CHF_List             { return CHF_List(C.Struct(s).GetObject(0)) }
func (s PackedCHF) SetMsgs(v CHF_List)         { C.Struct(s).SetObject(0, C.Object(v)) }
func (s PackedCHF) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"msgs\":")
	if err != nil {
		return err
	}
	{
		s := s.Msgs()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteJSON(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s PackedCHF) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s PackedCHF) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("msgs = ")
	if err != nil {
		return err
	}
	{
		s := s.Msgs()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteCapLit(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s PackedCHF) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type PackedCHF_List C.PointerList

func NewPackedCHFList(s *C.Segment, sz int) PackedCHF_List {
	return PackedCHF_List(s.NewCompositeList(0, 1, sz))
}
func (s PackedCHF_List) Len() int           { return C.PointerList(s).Len() }
func (s PackedCHF_List) At(i int) PackedCHF { return PackedCHF(C.PointerList(s).At(i).ToStruct()) }
func (s PackedCHF_List) ToArray() []PackedCHF {
	n := s.Len()
	a := make([]PackedCHF, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s PackedCHF_List) Set(i int, item PackedCHF) { C.PointerList(s).Set(i, C.Object(item)) }
