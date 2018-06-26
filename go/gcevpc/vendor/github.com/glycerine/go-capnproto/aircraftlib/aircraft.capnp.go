package aircraftlib

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
	"math"
	"net"
)

type Zdate C.Struct

func NewZdate(s *C.Segment) Zdate      { return Zdate(s.NewStruct(8, 0)) }
func NewRootZdate(s *C.Segment) Zdate  { return Zdate(s.NewRootStruct(8, 0)) }
func AutoNewZdate(s *C.Segment) Zdate  { return Zdate(s.NewStructAR(8, 0)) }
func ReadRootZdate(s *C.Segment) Zdate { return Zdate(s.Root(0).ToStruct()) }
func (s Zdate) Year() int16            { return int16(C.Struct(s).Get16(0)) }
func (s Zdate) SetYear(v int16)        { C.Struct(s).Set16(0, uint16(v)) }
func (s Zdate) Month() uint8           { return C.Struct(s).Get8(2) }
func (s Zdate) SetMonth(v uint8)       { C.Struct(s).Set8(2, v) }
func (s Zdate) Day() uint8             { return C.Struct(s).Get8(3) }
func (s Zdate) SetDay(v uint8)         { C.Struct(s).Set8(3, v) }
func (s Zdate) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"year\":")
	if err != nil {
		return err
	}
	{
		s := s.Year()
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
	_, err = b.WriteString("\"month\":")
	if err != nil {
		return err
	}
	{
		s := s.Month()
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
	_, err = b.WriteString("\"day\":")
	if err != nil {
		return err
	}
	{
		s := s.Day()
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
func (s Zdate) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Zdate) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("year = ")
	if err != nil {
		return err
	}
	{
		s := s.Year()
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
	_, err = b.WriteString("month = ")
	if err != nil {
		return err
	}
	{
		s := s.Month()
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
	_, err = b.WriteString("day = ")
	if err != nil {
		return err
	}
	{
		s := s.Day()
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
func (s Zdate) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Zdate_List C.PointerList

func NewZdateList(s *C.Segment, sz int) Zdate_List { return Zdate_List(s.NewCompositeList(8, 0, sz)) }
func (s Zdate_List) Len() int                      { return C.PointerList(s).Len() }
func (s Zdate_List) At(i int) Zdate                { return Zdate(C.PointerList(s).At(i).ToStruct()) }
func (s Zdate_List) ToArray() []Zdate {
	n := s.Len()
	a := make([]Zdate, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Zdate_List) Set(i int, item Zdate) { C.PointerList(s).Set(i, C.Object(item)) }

type Zdata C.Struct

func NewZdata(s *C.Segment) Zdata      { return Zdata(s.NewStruct(0, 1)) }
func NewRootZdata(s *C.Segment) Zdata  { return Zdata(s.NewRootStruct(0, 1)) }
func AutoNewZdata(s *C.Segment) Zdata  { return Zdata(s.NewStructAR(0, 1)) }
func ReadRootZdata(s *C.Segment) Zdata { return Zdata(s.Root(0).ToStruct()) }
func (s Zdata) Data() []byte           { return C.Struct(s).GetObject(0).ToData() }
func (s Zdata) SetData(v []byte)       { C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s Zdata) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"data\":")
	if err != nil {
		return err
	}
	{
		s := s.Data()
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
func (s Zdata) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Zdata) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("data = ")
	if err != nil {
		return err
	}
	{
		s := s.Data()
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
func (s Zdata) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Zdata_List C.PointerList

func NewZdataList(s *C.Segment, sz int) Zdata_List { return Zdata_List(s.NewCompositeList(0, 1, sz)) }
func (s Zdata_List) Len() int                      { return C.PointerList(s).Len() }
func (s Zdata_List) At(i int) Zdata                { return Zdata(C.PointerList(s).At(i).ToStruct()) }
func (s Zdata_List) ToArray() []Zdata {
	n := s.Len()
	a := make([]Zdata, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Zdata_List) Set(i int, item Zdata) { C.PointerList(s).Set(i, C.Object(item)) }

type Airport uint16

const (
	AIRPORT_NONE Airport = 0
	AIRPORT_JFK  Airport = 1
	AIRPORT_LAX  Airport = 2
	AIRPORT_SFO  Airport = 3
	AIRPORT_LUV  Airport = 4
	AIRPORT_DFW  Airport = 5
	AIRPORT_TEST Airport = 6
)

func (c Airport) String() string {
	switch c {
	case AIRPORT_NONE:
		return "none"
	case AIRPORT_JFK:
		return "jfk"
	case AIRPORT_LAX:
		return "lax"
	case AIRPORT_SFO:
		return "sfo"
	case AIRPORT_LUV:
		return "luv"
	case AIRPORT_DFW:
		return "dfw"
	case AIRPORT_TEST:
		return "test"
	default:
		return ""
	}
}

func AirportFromString(c string) Airport {
	switch c {
	case "none":
		return AIRPORT_NONE
	case "jfk":
		return AIRPORT_JFK
	case "lax":
		return AIRPORT_LAX
	case "sfo":
		return AIRPORT_SFO
	case "luv":
		return AIRPORT_LUV
	case "dfw":
		return AIRPORT_DFW
	case "test":
		return AIRPORT_TEST
	default:
		return 0
	}
}

type Airport_List C.PointerList

func NewAirportList(s *C.Segment, sz int) Airport_List { return Airport_List(s.NewUInt16List(sz)) }
func (s Airport_List) Len() int                        { return C.UInt16List(s).Len() }
func (s Airport_List) At(i int) Airport                { return Airport(C.UInt16List(s).At(i)) }
func (s Airport_List) ToArray() []Airport {
	n := s.Len()
	a := make([]Airport, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Airport_List) Set(i int, item Airport) { C.UInt16List(s).Set(i, uint16(item)) }
func (s Airport) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Airport) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Airport) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	_, err = b.WriteString(s.String())
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Airport) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type PlaneBase C.Struct

func NewPlaneBase(s *C.Segment) PlaneBase      { return PlaneBase(s.NewStruct(32, 2)) }
func NewRootPlaneBase(s *C.Segment) PlaneBase  { return PlaneBase(s.NewRootStruct(32, 2)) }
func AutoNewPlaneBase(s *C.Segment) PlaneBase  { return PlaneBase(s.NewStructAR(32, 2)) }
func ReadRootPlaneBase(s *C.Segment) PlaneBase { return PlaneBase(s.Root(0).ToStruct()) }
func (s PlaneBase) Name() string               { return C.Struct(s).GetObject(0).ToText() }
func (s PlaneBase) NameBytes() []byte          { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s PlaneBase) SetName(v string)           { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s PlaneBase) Homes() Airport_List        { return Airport_List(C.Struct(s).GetObject(1)) }
func (s PlaneBase) SetHomes(v Airport_List)    { C.Struct(s).SetObject(1, C.Object(v)) }
func (s PlaneBase) Rating() int64              { return int64(C.Struct(s).Get64(0)) }
func (s PlaneBase) SetRating(v int64)          { C.Struct(s).Set64(0, uint64(v)) }
func (s PlaneBase) CanFly() bool               { return C.Struct(s).Get1(64) }
func (s PlaneBase) SetCanFly(v bool)           { C.Struct(s).Set1(64, v) }
func (s PlaneBase) Capacity() int64            { return int64(C.Struct(s).Get64(16)) }
func (s PlaneBase) SetCapacity(v int64)        { C.Struct(s).Set64(16, uint64(v)) }
func (s PlaneBase) MaxSpeed() float64          { return math.Float64frombits(C.Struct(s).Get64(24)) }
func (s PlaneBase) SetMaxSpeed(v float64)      { C.Struct(s).Set64(24, math.Float64bits(v)) }
func (s PlaneBase) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"name\":")
	if err != nil {
		return err
	}
	{
		s := s.Name()
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
	_, err = b.WriteString("\"homes\":")
	if err != nil {
		return err
	}
	{
		s := s.Homes()
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
	_, err = b.WriteString("\"rating\":")
	if err != nil {
		return err
	}
	{
		s := s.Rating()
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
	_, err = b.WriteString("\"canFly\":")
	if err != nil {
		return err
	}
	{
		s := s.CanFly()
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
	_, err = b.WriteString("\"capacity\":")
	if err != nil {
		return err
	}
	{
		s := s.Capacity()
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
	_, err = b.WriteString("\"maxSpeed\":")
	if err != nil {
		return err
	}
	{
		s := s.MaxSpeed()
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
func (s PlaneBase) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s PlaneBase) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("name = ")
	if err != nil {
		return err
	}
	{
		s := s.Name()
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
	_, err = b.WriteString("homes = ")
	if err != nil {
		return err
	}
	{
		s := s.Homes()
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
	_, err = b.WriteString("rating = ")
	if err != nil {
		return err
	}
	{
		s := s.Rating()
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
	_, err = b.WriteString("canFly = ")
	if err != nil {
		return err
	}
	{
		s := s.CanFly()
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
	_, err = b.WriteString("capacity = ")
	if err != nil {
		return err
	}
	{
		s := s.Capacity()
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
	_, err = b.WriteString("maxSpeed = ")
	if err != nil {
		return err
	}
	{
		s := s.MaxSpeed()
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
func (s PlaneBase) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type PlaneBase_List C.PointerList

func NewPlaneBaseList(s *C.Segment, sz int) PlaneBase_List {
	return PlaneBase_List(s.NewCompositeList(32, 2, sz))
}
func (s PlaneBase_List) Len() int           { return C.PointerList(s).Len() }
func (s PlaneBase_List) At(i int) PlaneBase { return PlaneBase(C.PointerList(s).At(i).ToStruct()) }
func (s PlaneBase_List) ToArray() []PlaneBase {
	n := s.Len()
	a := make([]PlaneBase, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s PlaneBase_List) Set(i int, item PlaneBase) { C.PointerList(s).Set(i, C.Object(item)) }

type B737 C.Struct

func NewB737(s *C.Segment) B737      { return B737(s.NewStruct(0, 1)) }
func NewRootB737(s *C.Segment) B737  { return B737(s.NewRootStruct(0, 1)) }
func AutoNewB737(s *C.Segment) B737  { return B737(s.NewStructAR(0, 1)) }
func ReadRootB737(s *C.Segment) B737 { return B737(s.Root(0).ToStruct()) }
func (s B737) Base() PlaneBase       { return PlaneBase(C.Struct(s).GetObject(0).ToStruct()) }
func (s B737) SetBase(v PlaneBase)   { C.Struct(s).SetObject(0, C.Object(v)) }
func (s B737) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"base\":")
	if err != nil {
		return err
	}
	{
		s := s.Base()
		err = s.WriteJSON(b)
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
func (s B737) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s B737) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("base = ")
	if err != nil {
		return err
	}
	{
		s := s.Base()
		err = s.WriteCapLit(b)
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
func (s B737) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type B737_List C.PointerList

func NewB737List(s *C.Segment, sz int) B737_List { return B737_List(s.NewCompositeList(0, 1, sz)) }
func (s B737_List) Len() int                     { return C.PointerList(s).Len() }
func (s B737_List) At(i int) B737                { return B737(C.PointerList(s).At(i).ToStruct()) }
func (s B737_List) ToArray() []B737 {
	n := s.Len()
	a := make([]B737, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s B737_List) Set(i int, item B737) { C.PointerList(s).Set(i, C.Object(item)) }

type A320 C.Struct

func NewA320(s *C.Segment) A320      { return A320(s.NewStruct(0, 1)) }
func NewRootA320(s *C.Segment) A320  { return A320(s.NewRootStruct(0, 1)) }
func AutoNewA320(s *C.Segment) A320  { return A320(s.NewStructAR(0, 1)) }
func ReadRootA320(s *C.Segment) A320 { return A320(s.Root(0).ToStruct()) }
func (s A320) Base() PlaneBase       { return PlaneBase(C.Struct(s).GetObject(0).ToStruct()) }
func (s A320) SetBase(v PlaneBase)   { C.Struct(s).SetObject(0, C.Object(v)) }
func (s A320) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"base\":")
	if err != nil {
		return err
	}
	{
		s := s.Base()
		err = s.WriteJSON(b)
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
func (s A320) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s A320) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("base = ")
	if err != nil {
		return err
	}
	{
		s := s.Base()
		err = s.WriteCapLit(b)
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
func (s A320) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type A320_List C.PointerList

func NewA320List(s *C.Segment, sz int) A320_List { return A320_List(s.NewCompositeList(0, 1, sz)) }
func (s A320_List) Len() int                     { return C.PointerList(s).Len() }
func (s A320_List) At(i int) A320                { return A320(C.PointerList(s).At(i).ToStruct()) }
func (s A320_List) ToArray() []A320 {
	n := s.Len()
	a := make([]A320, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s A320_List) Set(i int, item A320) { C.PointerList(s).Set(i, C.Object(item)) }

type F16 C.Struct

func NewF16(s *C.Segment) F16      { return F16(s.NewStruct(0, 1)) }
func NewRootF16(s *C.Segment) F16  { return F16(s.NewRootStruct(0, 1)) }
func AutoNewF16(s *C.Segment) F16  { return F16(s.NewStructAR(0, 1)) }
func ReadRootF16(s *C.Segment) F16 { return F16(s.Root(0).ToStruct()) }
func (s F16) Base() PlaneBase      { return PlaneBase(C.Struct(s).GetObject(0).ToStruct()) }
func (s F16) SetBase(v PlaneBase)  { C.Struct(s).SetObject(0, C.Object(v)) }
func (s F16) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"base\":")
	if err != nil {
		return err
	}
	{
		s := s.Base()
		err = s.WriteJSON(b)
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
func (s F16) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s F16) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("base = ")
	if err != nil {
		return err
	}
	{
		s := s.Base()
		err = s.WriteCapLit(b)
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
func (s F16) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type F16_List C.PointerList

func NewF16List(s *C.Segment, sz int) F16_List { return F16_List(s.NewCompositeList(0, 1, sz)) }
func (s F16_List) Len() int                    { return C.PointerList(s).Len() }
func (s F16_List) At(i int) F16                { return F16(C.PointerList(s).At(i).ToStruct()) }
func (s F16_List) ToArray() []F16 {
	n := s.Len()
	a := make([]F16, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s F16_List) Set(i int, item F16) { C.PointerList(s).Set(i, C.Object(item)) }

type Regression C.Struct

func NewRegression(s *C.Segment) Regression      { return Regression(s.NewStruct(24, 3)) }
func NewRootRegression(s *C.Segment) Regression  { return Regression(s.NewRootStruct(24, 3)) }
func AutoNewRegression(s *C.Segment) Regression  { return Regression(s.NewStructAR(24, 3)) }
func ReadRootRegression(s *C.Segment) Regression { return Regression(s.Root(0).ToStruct()) }
func (s Regression) Base() PlaneBase             { return PlaneBase(C.Struct(s).GetObject(0).ToStruct()) }
func (s Regression) SetBase(v PlaneBase)         { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Regression) B0() float64                 { return math.Float64frombits(C.Struct(s).Get64(0)) }
func (s Regression) SetB0(v float64)             { C.Struct(s).Set64(0, math.Float64bits(v)) }
func (s Regression) Beta() C.Float64List         { return C.Float64List(C.Struct(s).GetObject(1)) }
func (s Regression) SetBeta(v C.Float64List)     { C.Struct(s).SetObject(1, C.Object(v)) }
func (s Regression) Planes() Aircraft_List       { return Aircraft_List(C.Struct(s).GetObject(2)) }
func (s Regression) SetPlanes(v Aircraft_List)   { C.Struct(s).SetObject(2, C.Object(v)) }
func (s Regression) Ymu() float64                { return math.Float64frombits(C.Struct(s).Get64(8)) }
func (s Regression) SetYmu(v float64)            { C.Struct(s).Set64(8, math.Float64bits(v)) }
func (s Regression) Ysd() float64                { return math.Float64frombits(C.Struct(s).Get64(16)) }
func (s Regression) SetYsd(v float64)            { C.Struct(s).Set64(16, math.Float64bits(v)) }
func (s Regression) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"base\":")
	if err != nil {
		return err
	}
	{
		s := s.Base()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"b0\":")
	if err != nil {
		return err
	}
	{
		s := s.B0()
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
	_, err = b.WriteString("\"beta\":")
	if err != nil {
		return err
	}
	{
		s := s.Beta()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
	_, err = b.WriteString("\"planes\":")
	if err != nil {
		return err
	}
	{
		s := s.Planes()
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
	_, err = b.WriteString("\"ymu\":")
	if err != nil {
		return err
	}
	{
		s := s.Ymu()
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
	_, err = b.WriteString("\"ysd\":")
	if err != nil {
		return err
	}
	{
		s := s.Ysd()
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
func (s Regression) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Regression) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("base = ")
	if err != nil {
		return err
	}
	{
		s := s.Base()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("b0 = ")
	if err != nil {
		return err
	}
	{
		s := s.B0()
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
	_, err = b.WriteString("beta = ")
	if err != nil {
		return err
	}
	{
		s := s.Beta()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
	_, err = b.WriteString("planes = ")
	if err != nil {
		return err
	}
	{
		s := s.Planes()
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
	_, err = b.WriteString("ymu = ")
	if err != nil {
		return err
	}
	{
		s := s.Ymu()
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
	_, err = b.WriteString("ysd = ")
	if err != nil {
		return err
	}
	{
		s := s.Ysd()
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
func (s Regression) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Regression_List C.PointerList

func NewRegressionList(s *C.Segment, sz int) Regression_List {
	return Regression_List(s.NewCompositeList(24, 3, sz))
}
func (s Regression_List) Len() int            { return C.PointerList(s).Len() }
func (s Regression_List) At(i int) Regression { return Regression(C.PointerList(s).At(i).ToStruct()) }
func (s Regression_List) ToArray() []Regression {
	n := s.Len()
	a := make([]Regression, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Regression_List) Set(i int, item Regression) { C.PointerList(s).Set(i, C.Object(item)) }

type Aircraft C.Struct
type Aircraft_Which uint16

const (
	AIRCRAFT_VOID Aircraft_Which = 0
	AIRCRAFT_B737 Aircraft_Which = 1
	AIRCRAFT_A320 Aircraft_Which = 2
	AIRCRAFT_F16  Aircraft_Which = 3
)

func NewAircraft(s *C.Segment) Aircraft      { return Aircraft(s.NewStruct(8, 1)) }
func NewRootAircraft(s *C.Segment) Aircraft  { return Aircraft(s.NewRootStruct(8, 1)) }
func AutoNewAircraft(s *C.Segment) Aircraft  { return Aircraft(s.NewStructAR(8, 1)) }
func ReadRootAircraft(s *C.Segment) Aircraft { return Aircraft(s.Root(0).ToStruct()) }
func (s Aircraft) Which() Aircraft_Which     { return Aircraft_Which(C.Struct(s).Get16(0)) }
func (s Aircraft) SetVoid()                  { C.Struct(s).Set16(0, 0) }
func (s Aircraft) B737() B737                { return B737(C.Struct(s).GetObject(0).ToStruct()) }
func (s Aircraft) SetB737(v B737)            { C.Struct(s).Set16(0, 1); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Aircraft) A320() A320                { return A320(C.Struct(s).GetObject(0).ToStruct()) }
func (s Aircraft) SetA320(v A320)            { C.Struct(s).Set16(0, 2); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Aircraft) F16() F16                  { return F16(C.Struct(s).GetObject(0).ToStruct()) }
func (s Aircraft) SetF16(v F16)              { C.Struct(s).Set16(0, 3); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Aircraft) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	if s.Which() == AIRCRAFT_VOID {
		_, err = b.WriteString("\"void\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == AIRCRAFT_B737 {
		_, err = b.WriteString("\"b737\":")
		if err != nil {
			return err
		}
		{
			s := s.B737()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == AIRCRAFT_A320 {
		_, err = b.WriteString("\"a320\":")
		if err != nil {
			return err
		}
		{
			s := s.A320()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == AIRCRAFT_F16 {
		_, err = b.WriteString("\"f16\":")
		if err != nil {
			return err
		}
		{
			s := s.F16()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Aircraft) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Aircraft) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	if s.Which() == AIRCRAFT_VOID {
		_, err = b.WriteString("void = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == AIRCRAFT_B737 {
		_, err = b.WriteString("b737 = ")
		if err != nil {
			return err
		}
		{
			s := s.B737()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == AIRCRAFT_A320 {
		_, err = b.WriteString("a320 = ")
		if err != nil {
			return err
		}
		{
			s := s.A320()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == AIRCRAFT_F16 {
		_, err = b.WriteString("f16 = ")
		if err != nil {
			return err
		}
		{
			s := s.F16()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Aircraft) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Aircraft_List C.PointerList

func NewAircraftList(s *C.Segment, sz int) Aircraft_List {
	return Aircraft_List(s.NewCompositeList(8, 1, sz))
}
func (s Aircraft_List) Len() int          { return C.PointerList(s).Len() }
func (s Aircraft_List) At(i int) Aircraft { return Aircraft(C.PointerList(s).At(i).ToStruct()) }
func (s Aircraft_List) ToArray() []Aircraft {
	n := s.Len()
	a := make([]Aircraft, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Aircraft_List) Set(i int, item Aircraft) { C.PointerList(s).Set(i, C.Object(item)) }

type Z C.Struct
type Z_Which uint16

const (
	Z_VOID        Z_Which = 0
	Z_ZZ          Z_Which = 1
	Z_F64         Z_Which = 2
	Z_F32         Z_Which = 3
	Z_I64         Z_Which = 4
	Z_I32         Z_Which = 5
	Z_I16         Z_Which = 6
	Z_I8          Z_Which = 7
	Z_U64         Z_Which = 8
	Z_U32         Z_Which = 9
	Z_U16         Z_Which = 10
	Z_U8          Z_Which = 11
	Z_BOOL        Z_Which = 12
	Z_TEXT        Z_Which = 13
	Z_BLOB        Z_Which = 14
	Z_F64VEC      Z_Which = 15
	Z_F32VEC      Z_Which = 16
	Z_I64VEC      Z_Which = 17
	Z_I32VEC      Z_Which = 18
	Z_I16VEC      Z_Which = 19
	Z_I8VEC       Z_Which = 20
	Z_U64VEC      Z_Which = 21
	Z_U32VEC      Z_Which = 22
	Z_U16VEC      Z_Which = 23
	Z_U8VEC       Z_Which = 24
	Z_ZVEC        Z_Which = 25
	Z_ZVECVEC     Z_Which = 26
	Z_ZDATE       Z_Which = 27
	Z_ZDATA       Z_Which = 28
	Z_AIRCRAFTVEC Z_Which = 29
	Z_AIRCRAFT    Z_Which = 30
	Z_REGRESSION  Z_Which = 31
	Z_PLANEBASE   Z_Which = 32
	Z_AIRPORT     Z_Which = 33
	Z_B737        Z_Which = 34
	Z_A320        Z_Which = 35
	Z_F16         Z_Which = 36
	Z_ZDATEVEC    Z_Which = 37
	Z_ZDATAVEC    Z_Which = 38
	Z_BOOLVEC     Z_Which = 39
)

func NewZ(s *C.Segment) Z             { return Z(s.NewStruct(16, 1)) }
func NewRootZ(s *C.Segment) Z         { return Z(s.NewRootStruct(16, 1)) }
func AutoNewZ(s *C.Segment) Z         { return Z(s.NewStructAR(16, 1)) }
func ReadRootZ(s *C.Segment) Z        { return Z(s.Root(0).ToStruct()) }
func (s Z) Which() Z_Which            { return Z_Which(C.Struct(s).Get16(0)) }
func (s Z) SetVoid()                  { C.Struct(s).Set16(0, 0) }
func (s Z) Zz() Z                     { return Z(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetZz(v Z)                 { C.Struct(s).Set16(0, 1); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) F64() float64              { return math.Float64frombits(C.Struct(s).Get64(8)) }
func (s Z) SetF64(v float64)          { C.Struct(s).Set16(0, 2); C.Struct(s).Set64(8, math.Float64bits(v)) }
func (s Z) F32() float32              { return math.Float32frombits(C.Struct(s).Get32(8)) }
func (s Z) SetF32(v float32)          { C.Struct(s).Set16(0, 3); C.Struct(s).Set32(8, math.Float32bits(v)) }
func (s Z) I64() int64                { return int64(C.Struct(s).Get64(8)) }
func (s Z) SetI64(v int64)            { C.Struct(s).Set16(0, 4); C.Struct(s).Set64(8, uint64(v)) }
func (s Z) I32() int32                { return int32(C.Struct(s).Get32(8)) }
func (s Z) SetI32(v int32)            { C.Struct(s).Set16(0, 5); C.Struct(s).Set32(8, uint32(v)) }
func (s Z) I16() int16                { return int16(C.Struct(s).Get16(8)) }
func (s Z) SetI16(v int16)            { C.Struct(s).Set16(0, 6); C.Struct(s).Set16(8, uint16(v)) }
func (s Z) I8() int8                  { return int8(C.Struct(s).Get8(8)) }
func (s Z) SetI8(v int8)              { C.Struct(s).Set16(0, 7); C.Struct(s).Set8(8, uint8(v)) }
func (s Z) U64() uint64               { return C.Struct(s).Get64(8) }
func (s Z) SetU64(v uint64)           { C.Struct(s).Set16(0, 8); C.Struct(s).Set64(8, v) }
func (s Z) U32() uint32               { return C.Struct(s).Get32(8) }
func (s Z) SetU32(v uint32)           { C.Struct(s).Set16(0, 9); C.Struct(s).Set32(8, v) }
func (s Z) U16() uint16               { return C.Struct(s).Get16(8) }
func (s Z) SetU16(v uint16)           { C.Struct(s).Set16(0, 10); C.Struct(s).Set16(8, v) }
func (s Z) U8() uint8                 { return C.Struct(s).Get8(8) }
func (s Z) SetU8(v uint8)             { C.Struct(s).Set16(0, 11); C.Struct(s).Set8(8, v) }
func (s Z) Bool() bool                { return C.Struct(s).Get1(64) }
func (s Z) SetBool(v bool)            { C.Struct(s).Set16(0, 12); C.Struct(s).Set1(64, v) }
func (s Z) Text() string              { return C.Struct(s).GetObject(0).ToText() }
func (s Z) TextBytes() []byte         { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s Z) SetText(v string)          { C.Struct(s).Set16(0, 13); C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Z) Blob() []byte              { return C.Struct(s).GetObject(0).ToData() }
func (s Z) SetBlob(v []byte)          { C.Struct(s).Set16(0, 14); C.Struct(s).SetObject(0, s.Segment.NewData(v)) }
func (s Z) F64vec() C.Float64List     { return C.Float64List(C.Struct(s).GetObject(0)) }
func (s Z) SetF64vec(v C.Float64List) { C.Struct(s).Set16(0, 15); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) F32vec() C.Float32List     { return C.Float32List(C.Struct(s).GetObject(0)) }
func (s Z) SetF32vec(v C.Float32List) { C.Struct(s).Set16(0, 16); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) I64vec() C.Int64List       { return C.Int64List(C.Struct(s).GetObject(0)) }
func (s Z) SetI64vec(v C.Int64List)   { C.Struct(s).Set16(0, 17); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) I32vec() C.Int32List       { return C.Int32List(C.Struct(s).GetObject(0)) }
func (s Z) SetI32vec(v C.Int32List)   { C.Struct(s).Set16(0, 18); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) I16vec() C.Int16List       { return C.Int16List(C.Struct(s).GetObject(0)) }
func (s Z) SetI16vec(v C.Int16List)   { C.Struct(s).Set16(0, 19); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) I8vec() C.Int8List         { return C.Int8List(C.Struct(s).GetObject(0)) }
func (s Z) SetI8vec(v C.Int8List)     { C.Struct(s).Set16(0, 20); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) U64vec() C.UInt64List      { return C.UInt64List(C.Struct(s).GetObject(0)) }
func (s Z) SetU64vec(v C.UInt64List)  { C.Struct(s).Set16(0, 21); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) U32vec() C.UInt32List      { return C.UInt32List(C.Struct(s).GetObject(0)) }
func (s Z) SetU32vec(v C.UInt32List)  { C.Struct(s).Set16(0, 22); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) U16vec() C.UInt16List      { return C.UInt16List(C.Struct(s).GetObject(0)) }
func (s Z) SetU16vec(v C.UInt16List)  { C.Struct(s).Set16(0, 23); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) U8vec() C.UInt8List        { return C.UInt8List(C.Struct(s).GetObject(0)) }
func (s Z) SetU8vec(v C.UInt8List)    { C.Struct(s).Set16(0, 24); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Zvec() Z_List              { return Z_List(C.Struct(s).GetObject(0)) }
func (s Z) SetZvec(v Z_List)          { C.Struct(s).Set16(0, 25); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Zvecvec() C.PointerList    { return C.PointerList(C.Struct(s).GetObject(0)) }
func (s Z) SetZvecvec(v C.PointerList) {
	C.Struct(s).Set16(0, 26)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Z) Zdate() Zdate               { return Zdate(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetZdate(v Zdate)           { C.Struct(s).Set16(0, 27); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Zdata() Zdata               { return Zdata(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetZdata(v Zdata)           { C.Struct(s).Set16(0, 28); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Aircraftvec() Aircraft_List { return Aircraft_List(C.Struct(s).GetObject(0)) }
func (s Z) SetAircraftvec(v Aircraft_List) {
	C.Struct(s).Set16(0, 29)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Z) Aircraft() Aircraft     { return Aircraft(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetAircraft(v Aircraft) { C.Struct(s).Set16(0, 30); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Regression() Regression { return Regression(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetRegression(v Regression) {
	C.Struct(s).Set16(0, 31)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Z) Planebase() PlaneBase     { return PlaneBase(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetPlanebase(v PlaneBase) { C.Struct(s).Set16(0, 32); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Airport() Airport         { return Airport(C.Struct(s).Get16(8)) }
func (s Z) SetAirport(v Airport)     { C.Struct(s).Set16(0, 33); C.Struct(s).Set16(8, uint16(v)) }
func (s Z) B737() B737               { return B737(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetB737(v B737)           { C.Struct(s).Set16(0, 34); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) A320() A320               { return A320(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetA320(v A320)           { C.Struct(s).Set16(0, 35); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) F16() F16                 { return F16(C.Struct(s).GetObject(0).ToStruct()) }
func (s Z) SetF16(v F16)             { C.Struct(s).Set16(0, 36); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Zdatevec() Zdate_List     { return Zdate_List(C.Struct(s).GetObject(0)) }
func (s Z) SetZdatevec(v Zdate_List) { C.Struct(s).Set16(0, 37); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Zdatavec() Zdata_List     { return Zdata_List(C.Struct(s).GetObject(0)) }
func (s Z) SetZdatavec(v Zdata_List) { C.Struct(s).Set16(0, 38); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) Boolvec() C.BitList       { return C.BitList(C.Struct(s).GetObject(0)) }
func (s Z) SetBoolvec(v C.BitList)   { C.Struct(s).Set16(0, 39); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Z) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	if s.Which() == Z_VOID {
		_, err = b.WriteString("\"void\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == Z_ZZ {
		_, err = b.WriteString("\"zz\":")
		if err != nil {
			return err
		}
		{
			s := s.Zz()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_F64 {
		_, err = b.WriteString("\"f64\":")
		if err != nil {
			return err
		}
		{
			s := s.F64()
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
	if s.Which() == Z_F32 {
		_, err = b.WriteString("\"f32\":")
		if err != nil {
			return err
		}
		{
			s := s.F32()
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
	if s.Which() == Z_I64 {
		_, err = b.WriteString("\"i64\":")
		if err != nil {
			return err
		}
		{
			s := s.I64()
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
	if s.Which() == Z_I32 {
		_, err = b.WriteString("\"i32\":")
		if err != nil {
			return err
		}
		{
			s := s.I32()
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
	if s.Which() == Z_I16 {
		_, err = b.WriteString("\"i16\":")
		if err != nil {
			return err
		}
		{
			s := s.I16()
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
	if s.Which() == Z_I8 {
		_, err = b.WriteString("\"i8\":")
		if err != nil {
			return err
		}
		{
			s := s.I8()
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
	if s.Which() == Z_U64 {
		_, err = b.WriteString("\"u64\":")
		if err != nil {
			return err
		}
		{
			s := s.U64()
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
	if s.Which() == Z_U32 {
		_, err = b.WriteString("\"u32\":")
		if err != nil {
			return err
		}
		{
			s := s.U32()
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
	if s.Which() == Z_U16 {
		_, err = b.WriteString("\"u16\":")
		if err != nil {
			return err
		}
		{
			s := s.U16()
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
	if s.Which() == Z_U8 {
		_, err = b.WriteString("\"u8\":")
		if err != nil {
			return err
		}
		{
			s := s.U8()
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
	if s.Which() == Z_BOOL {
		_, err = b.WriteString("\"bool\":")
		if err != nil {
			return err
		}
		{
			s := s.Bool()
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
	if s.Which() == Z_TEXT {
		_, err = b.WriteString("\"text\":")
		if err != nil {
			return err
		}
		{
			s := s.Text()
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
	if s.Which() == Z_BLOB {
		_, err = b.WriteString("\"blob\":")
		if err != nil {
			return err
		}
		{
			s := s.Blob()
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
	if s.Which() == Z_F64VEC {
		_, err = b.WriteString("\"f64vec\":")
		if err != nil {
			return err
		}
		{
			s := s.F64vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_F32VEC {
		_, err = b.WriteString("\"f32vec\":")
		if err != nil {
			return err
		}
		{
			s := s.F32vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_I64VEC {
		_, err = b.WriteString("\"i64vec\":")
		if err != nil {
			return err
		}
		{
			s := s.I64vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_I32VEC {
		_, err = b.WriteString("\"i32vec\":")
		if err != nil {
			return err
		}
		{
			s := s.I32vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_I16VEC {
		_, err = b.WriteString("\"i16vec\":")
		if err != nil {
			return err
		}
		{
			s := s.I16vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_I8VEC {
		_, err = b.WriteString("\"i8vec\":")
		if err != nil {
			return err
		}
		{
			s := s.I8vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_U64VEC {
		_, err = b.WriteString("\"u64vec\":")
		if err != nil {
			return err
		}
		{
			s := s.U64vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_U32VEC {
		_, err = b.WriteString("\"u32vec\":")
		if err != nil {
			return err
		}
		{
			s := s.U32vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_U16VEC {
		_, err = b.WriteString("\"u16vec\":")
		if err != nil {
			return err
		}
		{
			s := s.U16vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_U8VEC {
		_, err = b.WriteString("\"u8vec\":")
		if err != nil {
			return err
		}
		{
			s := s.U8vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_ZVEC {
		_, err = b.WriteString("\"zvec\":")
		if err != nil {
			return err
		}
		{
			s := s.Zvec()
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
	}
	if s.Which() == Z_ZVECVEC {
		_, err = b.WriteString("\"zvecvec\":")
		if err != nil {
			return err
		}
		{
			s := s.Zvecvec()
			_ = s
			_, err = b.WriteString("\"untyped list\"")
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_ZDATE {
		_, err = b.WriteString("\"zdate\":")
		if err != nil {
			return err
		}
		{
			s := s.Zdate()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_ZDATA {
		_, err = b.WriteString("\"zdata\":")
		if err != nil {
			return err
		}
		{
			s := s.Zdata()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_AIRCRAFTVEC {
		_, err = b.WriteString("\"aircraftvec\":")
		if err != nil {
			return err
		}
		{
			s := s.Aircraftvec()
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
	}
	if s.Which() == Z_AIRCRAFT {
		_, err = b.WriteString("\"aircraft\":")
		if err != nil {
			return err
		}
		{
			s := s.Aircraft()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_REGRESSION {
		_, err = b.WriteString("\"regression\":")
		if err != nil {
			return err
		}
		{
			s := s.Regression()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_PLANEBASE {
		_, err = b.WriteString("\"planebase\":")
		if err != nil {
			return err
		}
		{
			s := s.Planebase()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_AIRPORT {
		_, err = b.WriteString("\"airport\":")
		if err != nil {
			return err
		}
		{
			s := s.Airport()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_B737 {
		_, err = b.WriteString("\"b737\":")
		if err != nil {
			return err
		}
		{
			s := s.B737()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_A320 {
		_, err = b.WriteString("\"a320\":")
		if err != nil {
			return err
		}
		{
			s := s.A320()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_F16 {
		_, err = b.WriteString("\"f16\":")
		if err != nil {
			return err
		}
		{
			s := s.F16()
			err = s.WriteJSON(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_ZDATEVEC {
		_, err = b.WriteString("\"zdatevec\":")
		if err != nil {
			return err
		}
		{
			s := s.Zdatevec()
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
	}
	if s.Which() == Z_ZDATAVEC {
		_, err = b.WriteString("\"zdatavec\":")
		if err != nil {
			return err
		}
		{
			s := s.Zdatavec()
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
	}
	if s.Which() == Z_BOOLVEC {
		_, err = b.WriteString("\"boolvec\":")
		if err != nil {
			return err
		}
		{
			s := s.Boolvec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Z) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Z) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	if s.Which() == Z_VOID {
		_, err = b.WriteString("void = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == Z_ZZ {
		_, err = b.WriteString("zz = ")
		if err != nil {
			return err
		}
		{
			s := s.Zz()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_F64 {
		_, err = b.WriteString("f64 = ")
		if err != nil {
			return err
		}
		{
			s := s.F64()
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
	if s.Which() == Z_F32 {
		_, err = b.WriteString("f32 = ")
		if err != nil {
			return err
		}
		{
			s := s.F32()
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
	if s.Which() == Z_I64 {
		_, err = b.WriteString("i64 = ")
		if err != nil {
			return err
		}
		{
			s := s.I64()
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
	if s.Which() == Z_I32 {
		_, err = b.WriteString("i32 = ")
		if err != nil {
			return err
		}
		{
			s := s.I32()
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
	if s.Which() == Z_I16 {
		_, err = b.WriteString("i16 = ")
		if err != nil {
			return err
		}
		{
			s := s.I16()
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
	if s.Which() == Z_I8 {
		_, err = b.WriteString("i8 = ")
		if err != nil {
			return err
		}
		{
			s := s.I8()
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
	if s.Which() == Z_U64 {
		_, err = b.WriteString("u64 = ")
		if err != nil {
			return err
		}
		{
			s := s.U64()
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
	if s.Which() == Z_U32 {
		_, err = b.WriteString("u32 = ")
		if err != nil {
			return err
		}
		{
			s := s.U32()
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
	if s.Which() == Z_U16 {
		_, err = b.WriteString("u16 = ")
		if err != nil {
			return err
		}
		{
			s := s.U16()
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
	if s.Which() == Z_U8 {
		_, err = b.WriteString("u8 = ")
		if err != nil {
			return err
		}
		{
			s := s.U8()
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
	if s.Which() == Z_BOOL {
		_, err = b.WriteString("bool = ")
		if err != nil {
			return err
		}
		{
			s := s.Bool()
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
	if s.Which() == Z_TEXT {
		_, err = b.WriteString("text = ")
		if err != nil {
			return err
		}
		{
			s := s.Text()
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
	if s.Which() == Z_BLOB {
		_, err = b.WriteString("blob = ")
		if err != nil {
			return err
		}
		{
			s := s.Blob()
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
	if s.Which() == Z_F64VEC {
		_, err = b.WriteString("f64vec = ")
		if err != nil {
			return err
		}
		{
			s := s.F64vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_F32VEC {
		_, err = b.WriteString("f32vec = ")
		if err != nil {
			return err
		}
		{
			s := s.F32vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_I64VEC {
		_, err = b.WriteString("i64vec = ")
		if err != nil {
			return err
		}
		{
			s := s.I64vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_I32VEC {
		_, err = b.WriteString("i32vec = ")
		if err != nil {
			return err
		}
		{
			s := s.I32vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_I16VEC {
		_, err = b.WriteString("i16vec = ")
		if err != nil {
			return err
		}
		{
			s := s.I16vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_I8VEC {
		_, err = b.WriteString("i8vec = ")
		if err != nil {
			return err
		}
		{
			s := s.I8vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_U64VEC {
		_, err = b.WriteString("u64vec = ")
		if err != nil {
			return err
		}
		{
			s := s.U64vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_U32VEC {
		_, err = b.WriteString("u32vec = ")
		if err != nil {
			return err
		}
		{
			s := s.U32vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_U16VEC {
		_, err = b.WriteString("u16vec = ")
		if err != nil {
			return err
		}
		{
			s := s.U16vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_U8VEC {
		_, err = b.WriteString("u8vec = ")
		if err != nil {
			return err
		}
		{
			s := s.U8vec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	if s.Which() == Z_ZVEC {
		_, err = b.WriteString("zvec = ")
		if err != nil {
			return err
		}
		{
			s := s.Zvec()
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
	}
	if s.Which() == Z_ZVECVEC {
		_, err = b.WriteString("zvecvec = ")
		if err != nil {
			return err
		}
		{
			s := s.Zvecvec()
			_ = s
			_, err = b.WriteString("\"untyped list\"")
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_ZDATE {
		_, err = b.WriteString("zdate = ")
		if err != nil {
			return err
		}
		{
			s := s.Zdate()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_ZDATA {
		_, err = b.WriteString("zdata = ")
		if err != nil {
			return err
		}
		{
			s := s.Zdata()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_AIRCRAFTVEC {
		_, err = b.WriteString("aircraftvec = ")
		if err != nil {
			return err
		}
		{
			s := s.Aircraftvec()
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
	}
	if s.Which() == Z_AIRCRAFT {
		_, err = b.WriteString("aircraft = ")
		if err != nil {
			return err
		}
		{
			s := s.Aircraft()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_REGRESSION {
		_, err = b.WriteString("regression = ")
		if err != nil {
			return err
		}
		{
			s := s.Regression()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_PLANEBASE {
		_, err = b.WriteString("planebase = ")
		if err != nil {
			return err
		}
		{
			s := s.Planebase()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_AIRPORT {
		_, err = b.WriteString("airport = ")
		if err != nil {
			return err
		}
		{
			s := s.Airport()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_B737 {
		_, err = b.WriteString("b737 = ")
		if err != nil {
			return err
		}
		{
			s := s.B737()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_A320 {
		_, err = b.WriteString("a320 = ")
		if err != nil {
			return err
		}
		{
			s := s.A320()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_F16 {
		_, err = b.WriteString("f16 = ")
		if err != nil {
			return err
		}
		{
			s := s.F16()
			err = s.WriteCapLit(b)
			if err != nil {
				return err
			}
		}
	}
	if s.Which() == Z_ZDATEVEC {
		_, err = b.WriteString("zdatevec = ")
		if err != nil {
			return err
		}
		{
			s := s.Zdatevec()
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
	}
	if s.Which() == Z_ZDATAVEC {
		_, err = b.WriteString("zdatavec = ")
		if err != nil {
			return err
		}
		{
			s := s.Zdatavec()
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
	}
	if s.Which() == Z_BOOLVEC {
		_, err = b.WriteString("boolvec = ")
		if err != nil {
			return err
		}
		{
			s := s.Boolvec()
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
					buf, err = json.Marshal(s)
					if err != nil {
						return err
					}
					_, err = b.Write(buf)
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
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Z) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Z_List C.PointerList

func NewZList(s *C.Segment, sz int) Z_List { return Z_List(s.NewCompositeList(16, 1, sz)) }
func (s Z_List) Len() int                  { return C.PointerList(s).Len() }
func (s Z_List) At(i int) Z                { return Z(C.PointerList(s).At(i).ToStruct()) }
func (s Z_List) ToArray() []Z {
	n := s.Len()
	a := make([]Z, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Z_List) Set(i int, item Z) { C.PointerList(s).Set(i, C.Object(item)) }

type Counter C.Struct

func NewCounter(s *C.Segment) Counter      { return Counter(s.NewStruct(8, 2)) }
func NewRootCounter(s *C.Segment) Counter  { return Counter(s.NewRootStruct(8, 2)) }
func AutoNewCounter(s *C.Segment) Counter  { return Counter(s.NewStructAR(8, 2)) }
func ReadRootCounter(s *C.Segment) Counter { return Counter(s.Root(0).ToStruct()) }
func (s Counter) Size() int64              { return int64(C.Struct(s).Get64(0)) }
func (s Counter) SetSize(v int64)          { C.Struct(s).Set64(0, uint64(v)) }
func (s Counter) Words() string            { return C.Struct(s).GetObject(0).ToText() }
func (s Counter) WordsBytes() []byte       { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s Counter) SetWords(v string)        { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Counter) Wordlist() C.TextList     { return C.TextList(C.Struct(s).GetObject(1)) }
func (s Counter) SetWordlist(v C.TextList) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s Counter) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"size\":")
	if err != nil {
		return err
	}
	{
		s := s.Size()
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
	_, err = b.WriteString("\"words\":")
	if err != nil {
		return err
	}
	{
		s := s.Words()
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
	_, err = b.WriteString("\"wordlist\":")
	if err != nil {
		return err
	}
	{
		s := s.Wordlist()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
func (s Counter) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Counter) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("size = ")
	if err != nil {
		return err
	}
	{
		s := s.Size()
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
	_, err = b.WriteString("words = ")
	if err != nil {
		return err
	}
	{
		s := s.Words()
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
	_, err = b.WriteString("wordlist = ")
	if err != nil {
		return err
	}
	{
		s := s.Wordlist()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
func (s Counter) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Counter_List C.PointerList

func NewCounterList(s *C.Segment, sz int) Counter_List {
	return Counter_List(s.NewCompositeList(8, 2, sz))
}
func (s Counter_List) Len() int         { return C.PointerList(s).Len() }
func (s Counter_List) At(i int) Counter { return Counter(C.PointerList(s).At(i).ToStruct()) }
func (s Counter_List) ToArray() []Counter {
	n := s.Len()
	a := make([]Counter, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Counter_List) Set(i int, item Counter) { C.PointerList(s).Set(i, C.Object(item)) }

type Bag C.Struct

func NewBag(s *C.Segment) Bag      { return Bag(s.NewStruct(0, 1)) }
func NewRootBag(s *C.Segment) Bag  { return Bag(s.NewRootStruct(0, 1)) }
func AutoNewBag(s *C.Segment) Bag  { return Bag(s.NewStructAR(0, 1)) }
func ReadRootBag(s *C.Segment) Bag { return Bag(s.Root(0).ToStruct()) }
func (s Bag) Counter() Counter     { return Counter(C.Struct(s).GetObject(0).ToStruct()) }
func (s Bag) SetCounter(v Counter) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Bag) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"counter\":")
	if err != nil {
		return err
	}
	{
		s := s.Counter()
		err = s.WriteJSON(b)
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
func (s Bag) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Bag) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("counter = ")
	if err != nil {
		return err
	}
	{
		s := s.Counter()
		err = s.WriteCapLit(b)
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
func (s Bag) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Bag_List C.PointerList

func NewBagList(s *C.Segment, sz int) Bag_List { return Bag_List(s.NewCompositeList(0, 1, sz)) }
func (s Bag_List) Len() int                    { return C.PointerList(s).Len() }
func (s Bag_List) At(i int) Bag                { return Bag(C.PointerList(s).At(i).ToStruct()) }
func (s Bag_List) ToArray() []Bag {
	n := s.Len()
	a := make([]Bag, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Bag_List) Set(i int, item Bag) { C.PointerList(s).Set(i, C.Object(item)) }

type Zserver C.Struct

func NewZserver(s *C.Segment) Zserver        { return Zserver(s.NewStruct(0, 1)) }
func NewRootZserver(s *C.Segment) Zserver    { return Zserver(s.NewRootStruct(0, 1)) }
func AutoNewZserver(s *C.Segment) Zserver    { return Zserver(s.NewStructAR(0, 1)) }
func ReadRootZserver(s *C.Segment) Zserver   { return Zserver(s.Root(0).ToStruct()) }
func (s Zserver) Waitingjobs() Zjob_List     { return Zjob_List(C.Struct(s).GetObject(0)) }
func (s Zserver) SetWaitingjobs(v Zjob_List) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Zserver) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"waitingjobs\":")
	if err != nil {
		return err
	}
	{
		s := s.Waitingjobs()
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
func (s Zserver) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Zserver) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("waitingjobs = ")
	if err != nil {
		return err
	}
	{
		s := s.Waitingjobs()
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
func (s Zserver) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Zserver_List C.PointerList

func NewZserverList(s *C.Segment, sz int) Zserver_List {
	return Zserver_List(s.NewCompositeList(0, 1, sz))
}
func (s Zserver_List) Len() int         { return C.PointerList(s).Len() }
func (s Zserver_List) At(i int) Zserver { return Zserver(C.PointerList(s).At(i).ToStruct()) }
func (s Zserver_List) ToArray() []Zserver {
	n := s.Len()
	a := make([]Zserver, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Zserver_List) Set(i int, item Zserver) { C.PointerList(s).Set(i, C.Object(item)) }

type Zjob C.Struct

func NewZjob(s *C.Segment) Zjob      { return Zjob(s.NewStruct(0, 2)) }
func NewRootZjob(s *C.Segment) Zjob  { return Zjob(s.NewRootStruct(0, 2)) }
func AutoNewZjob(s *C.Segment) Zjob  { return Zjob(s.NewStructAR(0, 2)) }
func ReadRootZjob(s *C.Segment) Zjob { return Zjob(s.Root(0).ToStruct()) }
func (s Zjob) Cmd() string           { return C.Struct(s).GetObject(0).ToText() }
func (s Zjob) CmdBytes() []byte      { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s Zjob) SetCmd(v string)       { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Zjob) Args() C.TextList      { return C.TextList(C.Struct(s).GetObject(1)) }
func (s Zjob) SetArgs(v C.TextList)  { C.Struct(s).SetObject(1, C.Object(v)) }
func (s Zjob) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"cmd\":")
	if err != nil {
		return err
	}
	{
		s := s.Cmd()
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
	_, err = b.WriteString("\"args\":")
	if err != nil {
		return err
	}
	{
		s := s.Args()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
func (s Zjob) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Zjob) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("cmd = ")
	if err != nil {
		return err
	}
	{
		s := s.Cmd()
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
	_, err = b.WriteString("args = ")
	if err != nil {
		return err
	}
	{
		s := s.Args()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
func (s Zjob) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Zjob_List C.PointerList

func NewZjobList(s *C.Segment, sz int) Zjob_List { return Zjob_List(s.NewCompositeList(0, 2, sz)) }
func (s Zjob_List) Len() int                     { return C.PointerList(s).Len() }
func (s Zjob_List) At(i int) Zjob                { return Zjob(C.PointerList(s).At(i).ToStruct()) }
func (s Zjob_List) ToArray() []Zjob {
	n := s.Len()
	a := make([]Zjob, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Zjob_List) Set(i int, item Zjob) { C.PointerList(s).Set(i, C.Object(item)) }

type VerEmpty C.Struct

func NewVerEmpty(s *C.Segment) VerEmpty      { return VerEmpty(s.NewStruct(0, 0)) }
func NewRootVerEmpty(s *C.Segment) VerEmpty  { return VerEmpty(s.NewRootStruct(0, 0)) }
func AutoNewVerEmpty(s *C.Segment) VerEmpty  { return VerEmpty(s.NewStructAR(0, 0)) }
func ReadRootVerEmpty(s *C.Segment) VerEmpty { return VerEmpty(s.Root(0).ToStruct()) }
func (s VerEmpty) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s VerEmpty) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VerEmpty) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s VerEmpty) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VerEmpty_List C.PointerList

func NewVerEmptyList(s *C.Segment, sz int) VerEmpty_List {
	return VerEmpty_List(s.NewCompositeList(0, 0, sz))
}
func (s VerEmpty_List) Len() int          { return C.PointerList(s).Len() }
func (s VerEmpty_List) At(i int) VerEmpty { return VerEmpty(C.PointerList(s).At(i).ToStruct()) }
func (s VerEmpty_List) ToArray() []VerEmpty {
	n := s.Len()
	a := make([]VerEmpty, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VerEmpty_List) Set(i int, item VerEmpty) { C.PointerList(s).Set(i, C.Object(item)) }

type VerOneData C.Struct

func NewVerOneData(s *C.Segment) VerOneData      { return VerOneData(s.NewStruct(8, 0)) }
func NewRootVerOneData(s *C.Segment) VerOneData  { return VerOneData(s.NewRootStruct(8, 0)) }
func AutoNewVerOneData(s *C.Segment) VerOneData  { return VerOneData(s.NewStructAR(8, 0)) }
func ReadRootVerOneData(s *C.Segment) VerOneData { return VerOneData(s.Root(0).ToStruct()) }
func (s VerOneData) Val() int16                  { return int16(C.Struct(s).Get16(0)) }
func (s VerOneData) SetVal(v int16)              { C.Struct(s).Set16(0, uint16(v)) }
func (s VerOneData) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"val\":")
	if err != nil {
		return err
	}
	{
		s := s.Val()
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
func (s VerOneData) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VerOneData) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("val = ")
	if err != nil {
		return err
	}
	{
		s := s.Val()
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
func (s VerOneData) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VerOneData_List C.PointerList

func NewVerOneDataList(s *C.Segment, sz int) VerOneData_List {
	return VerOneData_List(s.NewCompositeList(8, 0, sz))
}
func (s VerOneData_List) Len() int            { return C.PointerList(s).Len() }
func (s VerOneData_List) At(i int) VerOneData { return VerOneData(C.PointerList(s).At(i).ToStruct()) }
func (s VerOneData_List) ToArray() []VerOneData {
	n := s.Len()
	a := make([]VerOneData, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VerOneData_List) Set(i int, item VerOneData) { C.PointerList(s).Set(i, C.Object(item)) }

type VerTwoData C.Struct

func NewVerTwoData(s *C.Segment) VerTwoData      { return VerTwoData(s.NewStruct(16, 0)) }
func NewRootVerTwoData(s *C.Segment) VerTwoData  { return VerTwoData(s.NewRootStruct(16, 0)) }
func AutoNewVerTwoData(s *C.Segment) VerTwoData  { return VerTwoData(s.NewStructAR(16, 0)) }
func ReadRootVerTwoData(s *C.Segment) VerTwoData { return VerTwoData(s.Root(0).ToStruct()) }
func (s VerTwoData) Val() int16                  { return int16(C.Struct(s).Get16(0)) }
func (s VerTwoData) SetVal(v int16)              { C.Struct(s).Set16(0, uint16(v)) }
func (s VerTwoData) Duo() int64                  { return int64(C.Struct(s).Get64(8)) }
func (s VerTwoData) SetDuo(v int64)              { C.Struct(s).Set64(8, uint64(v)) }
func (s VerTwoData) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"val\":")
	if err != nil {
		return err
	}
	{
		s := s.Val()
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
	_, err = b.WriteString("\"duo\":")
	if err != nil {
		return err
	}
	{
		s := s.Duo()
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
func (s VerTwoData) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VerTwoData) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("val = ")
	if err != nil {
		return err
	}
	{
		s := s.Val()
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
	_, err = b.WriteString("duo = ")
	if err != nil {
		return err
	}
	{
		s := s.Duo()
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
func (s VerTwoData) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VerTwoData_List C.PointerList

func NewVerTwoDataList(s *C.Segment, sz int) VerTwoData_List {
	return VerTwoData_List(s.NewCompositeList(16, 0, sz))
}
func (s VerTwoData_List) Len() int            { return C.PointerList(s).Len() }
func (s VerTwoData_List) At(i int) VerTwoData { return VerTwoData(C.PointerList(s).At(i).ToStruct()) }
func (s VerTwoData_List) ToArray() []VerTwoData {
	n := s.Len()
	a := make([]VerTwoData, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VerTwoData_List) Set(i int, item VerTwoData) { C.PointerList(s).Set(i, C.Object(item)) }

type VerOnePtr C.Struct

func NewVerOnePtr(s *C.Segment) VerOnePtr      { return VerOnePtr(s.NewStruct(0, 1)) }
func NewRootVerOnePtr(s *C.Segment) VerOnePtr  { return VerOnePtr(s.NewRootStruct(0, 1)) }
func AutoNewVerOnePtr(s *C.Segment) VerOnePtr  { return VerOnePtr(s.NewStructAR(0, 1)) }
func ReadRootVerOnePtr(s *C.Segment) VerOnePtr { return VerOnePtr(s.Root(0).ToStruct()) }
func (s VerOnePtr) Ptr() VerOneData            { return VerOneData(C.Struct(s).GetObject(0).ToStruct()) }
func (s VerOnePtr) SetPtr(v VerOneData)        { C.Struct(s).SetObject(0, C.Object(v)) }
func (s VerOnePtr) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ptr\":")
	if err != nil {
		return err
	}
	{
		s := s.Ptr()
		err = s.WriteJSON(b)
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
func (s VerOnePtr) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VerOnePtr) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("ptr = ")
	if err != nil {
		return err
	}
	{
		s := s.Ptr()
		err = s.WriteCapLit(b)
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
func (s VerOnePtr) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VerOnePtr_List C.PointerList

func NewVerOnePtrList(s *C.Segment, sz int) VerOnePtr_List {
	return VerOnePtr_List(s.NewCompositeList(0, 1, sz))
}
func (s VerOnePtr_List) Len() int           { return C.PointerList(s).Len() }
func (s VerOnePtr_List) At(i int) VerOnePtr { return VerOnePtr(C.PointerList(s).At(i).ToStruct()) }
func (s VerOnePtr_List) ToArray() []VerOnePtr {
	n := s.Len()
	a := make([]VerOnePtr, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VerOnePtr_List) Set(i int, item VerOnePtr) { C.PointerList(s).Set(i, C.Object(item)) }

type VerTwoPtr C.Struct

func NewVerTwoPtr(s *C.Segment) VerTwoPtr      { return VerTwoPtr(s.NewStruct(0, 2)) }
func NewRootVerTwoPtr(s *C.Segment) VerTwoPtr  { return VerTwoPtr(s.NewRootStruct(0, 2)) }
func AutoNewVerTwoPtr(s *C.Segment) VerTwoPtr  { return VerTwoPtr(s.NewStructAR(0, 2)) }
func ReadRootVerTwoPtr(s *C.Segment) VerTwoPtr { return VerTwoPtr(s.Root(0).ToStruct()) }
func (s VerTwoPtr) Ptr1() VerOneData           { return VerOneData(C.Struct(s).GetObject(0).ToStruct()) }
func (s VerTwoPtr) SetPtr1(v VerOneData)       { C.Struct(s).SetObject(0, C.Object(v)) }
func (s VerTwoPtr) Ptr2() VerOneData           { return VerOneData(C.Struct(s).GetObject(1).ToStruct()) }
func (s VerTwoPtr) SetPtr2(v VerOneData)       { C.Struct(s).SetObject(1, C.Object(v)) }
func (s VerTwoPtr) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ptr1\":")
	if err != nil {
		return err
	}
	{
		s := s.Ptr1()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ptr2\":")
	if err != nil {
		return err
	}
	{
		s := s.Ptr2()
		err = s.WriteJSON(b)
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
func (s VerTwoPtr) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VerTwoPtr) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("ptr1 = ")
	if err != nil {
		return err
	}
	{
		s := s.Ptr1()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ptr2 = ")
	if err != nil {
		return err
	}
	{
		s := s.Ptr2()
		err = s.WriteCapLit(b)
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
func (s VerTwoPtr) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VerTwoPtr_List C.PointerList

func NewVerTwoPtrList(s *C.Segment, sz int) VerTwoPtr_List {
	return VerTwoPtr_List(s.NewCompositeList(0, 2, sz))
}
func (s VerTwoPtr_List) Len() int           { return C.PointerList(s).Len() }
func (s VerTwoPtr_List) At(i int) VerTwoPtr { return VerTwoPtr(C.PointerList(s).At(i).ToStruct()) }
func (s VerTwoPtr_List) ToArray() []VerTwoPtr {
	n := s.Len()
	a := make([]VerTwoPtr, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VerTwoPtr_List) Set(i int, item VerTwoPtr) { C.PointerList(s).Set(i, C.Object(item)) }

type VerTwoDataTwoPtr C.Struct

func NewVerTwoDataTwoPtr(s *C.Segment) VerTwoDataTwoPtr { return VerTwoDataTwoPtr(s.NewStruct(16, 2)) }
func NewRootVerTwoDataTwoPtr(s *C.Segment) VerTwoDataTwoPtr {
	return VerTwoDataTwoPtr(s.NewRootStruct(16, 2))
}
func AutoNewVerTwoDataTwoPtr(s *C.Segment) VerTwoDataTwoPtr {
	return VerTwoDataTwoPtr(s.NewStructAR(16, 2))
}
func ReadRootVerTwoDataTwoPtr(s *C.Segment) VerTwoDataTwoPtr {
	return VerTwoDataTwoPtr(s.Root(0).ToStruct())
}
func (s VerTwoDataTwoPtr) Val() int16           { return int16(C.Struct(s).Get16(0)) }
func (s VerTwoDataTwoPtr) SetVal(v int16)       { C.Struct(s).Set16(0, uint16(v)) }
func (s VerTwoDataTwoPtr) Duo() int64           { return int64(C.Struct(s).Get64(8)) }
func (s VerTwoDataTwoPtr) SetDuo(v int64)       { C.Struct(s).Set64(8, uint64(v)) }
func (s VerTwoDataTwoPtr) Ptr1() VerOneData     { return VerOneData(C.Struct(s).GetObject(0).ToStruct()) }
func (s VerTwoDataTwoPtr) SetPtr1(v VerOneData) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s VerTwoDataTwoPtr) Ptr2() VerOneData     { return VerOneData(C.Struct(s).GetObject(1).ToStruct()) }
func (s VerTwoDataTwoPtr) SetPtr2(v VerOneData) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s VerTwoDataTwoPtr) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"val\":")
	if err != nil {
		return err
	}
	{
		s := s.Val()
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
	_, err = b.WriteString("\"duo\":")
	if err != nil {
		return err
	}
	{
		s := s.Duo()
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
	_, err = b.WriteString("\"ptr1\":")
	if err != nil {
		return err
	}
	{
		s := s.Ptr1()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ptr2\":")
	if err != nil {
		return err
	}
	{
		s := s.Ptr2()
		err = s.WriteJSON(b)
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
func (s VerTwoDataTwoPtr) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VerTwoDataTwoPtr) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("val = ")
	if err != nil {
		return err
	}
	{
		s := s.Val()
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
	_, err = b.WriteString("duo = ")
	if err != nil {
		return err
	}
	{
		s := s.Duo()
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
	_, err = b.WriteString("ptr1 = ")
	if err != nil {
		return err
	}
	{
		s := s.Ptr1()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ptr2 = ")
	if err != nil {
		return err
	}
	{
		s := s.Ptr2()
		err = s.WriteCapLit(b)
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
func (s VerTwoDataTwoPtr) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VerTwoDataTwoPtr_List C.PointerList

func NewVerTwoDataTwoPtrList(s *C.Segment, sz int) VerTwoDataTwoPtr_List {
	return VerTwoDataTwoPtr_List(s.NewCompositeList(16, 2, sz))
}
func (s VerTwoDataTwoPtr_List) Len() int { return C.PointerList(s).Len() }
func (s VerTwoDataTwoPtr_List) At(i int) VerTwoDataTwoPtr {
	return VerTwoDataTwoPtr(C.PointerList(s).At(i).ToStruct())
}
func (s VerTwoDataTwoPtr_List) ToArray() []VerTwoDataTwoPtr {
	n := s.Len()
	a := make([]VerTwoDataTwoPtr, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VerTwoDataTwoPtr_List) Set(i int, item VerTwoDataTwoPtr) {
	C.PointerList(s).Set(i, C.Object(item))
}

type HoldsVerEmptyList C.Struct

func NewHoldsVerEmptyList(s *C.Segment) HoldsVerEmptyList { return HoldsVerEmptyList(s.NewStruct(0, 1)) }
func NewRootHoldsVerEmptyList(s *C.Segment) HoldsVerEmptyList {
	return HoldsVerEmptyList(s.NewRootStruct(0, 1))
}
func AutoNewHoldsVerEmptyList(s *C.Segment) HoldsVerEmptyList {
	return HoldsVerEmptyList(s.NewStructAR(0, 1))
}
func ReadRootHoldsVerEmptyList(s *C.Segment) HoldsVerEmptyList {
	return HoldsVerEmptyList(s.Root(0).ToStruct())
}
func (s HoldsVerEmptyList) Mylist() VerEmpty_List     { return VerEmpty_List(C.Struct(s).GetObject(0)) }
func (s HoldsVerEmptyList) SetMylist(v VerEmpty_List) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s HoldsVerEmptyList) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mylist\":")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerEmptyList) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HoldsVerEmptyList) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mylist = ")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerEmptyList) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HoldsVerEmptyList_List C.PointerList

func NewHoldsVerEmptyListList(s *C.Segment, sz int) HoldsVerEmptyList_List {
	return HoldsVerEmptyList_List(s.NewCompositeList(0, 1, sz))
}
func (s HoldsVerEmptyList_List) Len() int { return C.PointerList(s).Len() }
func (s HoldsVerEmptyList_List) At(i int) HoldsVerEmptyList {
	return HoldsVerEmptyList(C.PointerList(s).At(i).ToStruct())
}
func (s HoldsVerEmptyList_List) ToArray() []HoldsVerEmptyList {
	n := s.Len()
	a := make([]HoldsVerEmptyList, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HoldsVerEmptyList_List) Set(i int, item HoldsVerEmptyList) {
	C.PointerList(s).Set(i, C.Object(item))
}

type HoldsVerOneDataList C.Struct

func NewHoldsVerOneDataList(s *C.Segment) HoldsVerOneDataList {
	return HoldsVerOneDataList(s.NewStruct(0, 1))
}
func NewRootHoldsVerOneDataList(s *C.Segment) HoldsVerOneDataList {
	return HoldsVerOneDataList(s.NewRootStruct(0, 1))
}
func AutoNewHoldsVerOneDataList(s *C.Segment) HoldsVerOneDataList {
	return HoldsVerOneDataList(s.NewStructAR(0, 1))
}
func ReadRootHoldsVerOneDataList(s *C.Segment) HoldsVerOneDataList {
	return HoldsVerOneDataList(s.Root(0).ToStruct())
}
func (s HoldsVerOneDataList) Mylist() VerOneData_List {
	return VerOneData_List(C.Struct(s).GetObject(0))
}
func (s HoldsVerOneDataList) SetMylist(v VerOneData_List) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s HoldsVerOneDataList) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mylist\":")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerOneDataList) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HoldsVerOneDataList) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mylist = ")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerOneDataList) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HoldsVerOneDataList_List C.PointerList

func NewHoldsVerOneDataListList(s *C.Segment, sz int) HoldsVerOneDataList_List {
	return HoldsVerOneDataList_List(s.NewCompositeList(0, 1, sz))
}
func (s HoldsVerOneDataList_List) Len() int { return C.PointerList(s).Len() }
func (s HoldsVerOneDataList_List) At(i int) HoldsVerOneDataList {
	return HoldsVerOneDataList(C.PointerList(s).At(i).ToStruct())
}
func (s HoldsVerOneDataList_List) ToArray() []HoldsVerOneDataList {
	n := s.Len()
	a := make([]HoldsVerOneDataList, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HoldsVerOneDataList_List) Set(i int, item HoldsVerOneDataList) {
	C.PointerList(s).Set(i, C.Object(item))
}

type HoldsVerTwoDataList C.Struct

func NewHoldsVerTwoDataList(s *C.Segment) HoldsVerTwoDataList {
	return HoldsVerTwoDataList(s.NewStruct(0, 1))
}
func NewRootHoldsVerTwoDataList(s *C.Segment) HoldsVerTwoDataList {
	return HoldsVerTwoDataList(s.NewRootStruct(0, 1))
}
func AutoNewHoldsVerTwoDataList(s *C.Segment) HoldsVerTwoDataList {
	return HoldsVerTwoDataList(s.NewStructAR(0, 1))
}
func ReadRootHoldsVerTwoDataList(s *C.Segment) HoldsVerTwoDataList {
	return HoldsVerTwoDataList(s.Root(0).ToStruct())
}
func (s HoldsVerTwoDataList) Mylist() VerTwoData_List {
	return VerTwoData_List(C.Struct(s).GetObject(0))
}
func (s HoldsVerTwoDataList) SetMylist(v VerTwoData_List) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s HoldsVerTwoDataList) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mylist\":")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerTwoDataList) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HoldsVerTwoDataList) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mylist = ")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerTwoDataList) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HoldsVerTwoDataList_List C.PointerList

func NewHoldsVerTwoDataListList(s *C.Segment, sz int) HoldsVerTwoDataList_List {
	return HoldsVerTwoDataList_List(s.NewCompositeList(0, 1, sz))
}
func (s HoldsVerTwoDataList_List) Len() int { return C.PointerList(s).Len() }
func (s HoldsVerTwoDataList_List) At(i int) HoldsVerTwoDataList {
	return HoldsVerTwoDataList(C.PointerList(s).At(i).ToStruct())
}
func (s HoldsVerTwoDataList_List) ToArray() []HoldsVerTwoDataList {
	n := s.Len()
	a := make([]HoldsVerTwoDataList, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HoldsVerTwoDataList_List) Set(i int, item HoldsVerTwoDataList) {
	C.PointerList(s).Set(i, C.Object(item))
}

type HoldsVerOnePtrList C.Struct

func NewHoldsVerOnePtrList(s *C.Segment) HoldsVerOnePtrList {
	return HoldsVerOnePtrList(s.NewStruct(0, 1))
}
func NewRootHoldsVerOnePtrList(s *C.Segment) HoldsVerOnePtrList {
	return HoldsVerOnePtrList(s.NewRootStruct(0, 1))
}
func AutoNewHoldsVerOnePtrList(s *C.Segment) HoldsVerOnePtrList {
	return HoldsVerOnePtrList(s.NewStructAR(0, 1))
}
func ReadRootHoldsVerOnePtrList(s *C.Segment) HoldsVerOnePtrList {
	return HoldsVerOnePtrList(s.Root(0).ToStruct())
}
func (s HoldsVerOnePtrList) Mylist() VerOnePtr_List     { return VerOnePtr_List(C.Struct(s).GetObject(0)) }
func (s HoldsVerOnePtrList) SetMylist(v VerOnePtr_List) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s HoldsVerOnePtrList) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mylist\":")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerOnePtrList) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HoldsVerOnePtrList) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mylist = ")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerOnePtrList) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HoldsVerOnePtrList_List C.PointerList

func NewHoldsVerOnePtrListList(s *C.Segment, sz int) HoldsVerOnePtrList_List {
	return HoldsVerOnePtrList_List(s.NewCompositeList(0, 1, sz))
}
func (s HoldsVerOnePtrList_List) Len() int { return C.PointerList(s).Len() }
func (s HoldsVerOnePtrList_List) At(i int) HoldsVerOnePtrList {
	return HoldsVerOnePtrList(C.PointerList(s).At(i).ToStruct())
}
func (s HoldsVerOnePtrList_List) ToArray() []HoldsVerOnePtrList {
	n := s.Len()
	a := make([]HoldsVerOnePtrList, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HoldsVerOnePtrList_List) Set(i int, item HoldsVerOnePtrList) {
	C.PointerList(s).Set(i, C.Object(item))
}

type HoldsVerTwoPtrList C.Struct

func NewHoldsVerTwoPtrList(s *C.Segment) HoldsVerTwoPtrList {
	return HoldsVerTwoPtrList(s.NewStruct(0, 1))
}
func NewRootHoldsVerTwoPtrList(s *C.Segment) HoldsVerTwoPtrList {
	return HoldsVerTwoPtrList(s.NewRootStruct(0, 1))
}
func AutoNewHoldsVerTwoPtrList(s *C.Segment) HoldsVerTwoPtrList {
	return HoldsVerTwoPtrList(s.NewStructAR(0, 1))
}
func ReadRootHoldsVerTwoPtrList(s *C.Segment) HoldsVerTwoPtrList {
	return HoldsVerTwoPtrList(s.Root(0).ToStruct())
}
func (s HoldsVerTwoPtrList) Mylist() VerTwoPtr_List     { return VerTwoPtr_List(C.Struct(s).GetObject(0)) }
func (s HoldsVerTwoPtrList) SetMylist(v VerTwoPtr_List) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s HoldsVerTwoPtrList) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mylist\":")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerTwoPtrList) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HoldsVerTwoPtrList) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mylist = ")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerTwoPtrList) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HoldsVerTwoPtrList_List C.PointerList

func NewHoldsVerTwoPtrListList(s *C.Segment, sz int) HoldsVerTwoPtrList_List {
	return HoldsVerTwoPtrList_List(s.NewCompositeList(0, 1, sz))
}
func (s HoldsVerTwoPtrList_List) Len() int { return C.PointerList(s).Len() }
func (s HoldsVerTwoPtrList_List) At(i int) HoldsVerTwoPtrList {
	return HoldsVerTwoPtrList(C.PointerList(s).At(i).ToStruct())
}
func (s HoldsVerTwoPtrList_List) ToArray() []HoldsVerTwoPtrList {
	n := s.Len()
	a := make([]HoldsVerTwoPtrList, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HoldsVerTwoPtrList_List) Set(i int, item HoldsVerTwoPtrList) {
	C.PointerList(s).Set(i, C.Object(item))
}

type HoldsVerTwoTwoList C.Struct

func NewHoldsVerTwoTwoList(s *C.Segment) HoldsVerTwoTwoList {
	return HoldsVerTwoTwoList(s.NewStruct(0, 1))
}
func NewRootHoldsVerTwoTwoList(s *C.Segment) HoldsVerTwoTwoList {
	return HoldsVerTwoTwoList(s.NewRootStruct(0, 1))
}
func AutoNewHoldsVerTwoTwoList(s *C.Segment) HoldsVerTwoTwoList {
	return HoldsVerTwoTwoList(s.NewStructAR(0, 1))
}
func ReadRootHoldsVerTwoTwoList(s *C.Segment) HoldsVerTwoTwoList {
	return HoldsVerTwoTwoList(s.Root(0).ToStruct())
}
func (s HoldsVerTwoTwoList) Mylist() VerTwoDataTwoPtr_List {
	return VerTwoDataTwoPtr_List(C.Struct(s).GetObject(0))
}
func (s HoldsVerTwoTwoList) SetMylist(v VerTwoDataTwoPtr_List) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s HoldsVerTwoTwoList) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mylist\":")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerTwoTwoList) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HoldsVerTwoTwoList) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mylist = ")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerTwoTwoList) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HoldsVerTwoTwoList_List C.PointerList

func NewHoldsVerTwoTwoListList(s *C.Segment, sz int) HoldsVerTwoTwoList_List {
	return HoldsVerTwoTwoList_List(s.NewCompositeList(0, 1, sz))
}
func (s HoldsVerTwoTwoList_List) Len() int { return C.PointerList(s).Len() }
func (s HoldsVerTwoTwoList_List) At(i int) HoldsVerTwoTwoList {
	return HoldsVerTwoTwoList(C.PointerList(s).At(i).ToStruct())
}
func (s HoldsVerTwoTwoList_List) ToArray() []HoldsVerTwoTwoList {
	n := s.Len()
	a := make([]HoldsVerTwoTwoList, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HoldsVerTwoTwoList_List) Set(i int, item HoldsVerTwoTwoList) {
	C.PointerList(s).Set(i, C.Object(item))
}

type HoldsVerTwoTwoPlus C.Struct

func NewHoldsVerTwoTwoPlus(s *C.Segment) HoldsVerTwoTwoPlus {
	return HoldsVerTwoTwoPlus(s.NewStruct(0, 1))
}
func NewRootHoldsVerTwoTwoPlus(s *C.Segment) HoldsVerTwoTwoPlus {
	return HoldsVerTwoTwoPlus(s.NewRootStruct(0, 1))
}
func AutoNewHoldsVerTwoTwoPlus(s *C.Segment) HoldsVerTwoTwoPlus {
	return HoldsVerTwoTwoPlus(s.NewStructAR(0, 1))
}
func ReadRootHoldsVerTwoTwoPlus(s *C.Segment) HoldsVerTwoTwoPlus {
	return HoldsVerTwoTwoPlus(s.Root(0).ToStruct())
}
func (s HoldsVerTwoTwoPlus) Mylist() VerTwoTwoPlus_List {
	return VerTwoTwoPlus_List(C.Struct(s).GetObject(0))
}
func (s HoldsVerTwoTwoPlus) SetMylist(v VerTwoTwoPlus_List) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s HoldsVerTwoTwoPlus) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mylist\":")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerTwoTwoPlus) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HoldsVerTwoTwoPlus) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mylist = ")
	if err != nil {
		return err
	}
	{
		s := s.Mylist()
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
func (s HoldsVerTwoTwoPlus) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HoldsVerTwoTwoPlus_List C.PointerList

func NewHoldsVerTwoTwoPlusList(s *C.Segment, sz int) HoldsVerTwoTwoPlus_List {
	return HoldsVerTwoTwoPlus_List(s.NewCompositeList(0, 1, sz))
}
func (s HoldsVerTwoTwoPlus_List) Len() int { return C.PointerList(s).Len() }
func (s HoldsVerTwoTwoPlus_List) At(i int) HoldsVerTwoTwoPlus {
	return HoldsVerTwoTwoPlus(C.PointerList(s).At(i).ToStruct())
}
func (s HoldsVerTwoTwoPlus_List) ToArray() []HoldsVerTwoTwoPlus {
	n := s.Len()
	a := make([]HoldsVerTwoTwoPlus, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HoldsVerTwoTwoPlus_List) Set(i int, item HoldsVerTwoTwoPlus) {
	C.PointerList(s).Set(i, C.Object(item))
}

type VerTwoTwoPlus C.Struct

func NewVerTwoTwoPlus(s *C.Segment) VerTwoTwoPlus      { return VerTwoTwoPlus(s.NewStruct(24, 3)) }
func NewRootVerTwoTwoPlus(s *C.Segment) VerTwoTwoPlus  { return VerTwoTwoPlus(s.NewRootStruct(24, 3)) }
func AutoNewVerTwoTwoPlus(s *C.Segment) VerTwoTwoPlus  { return VerTwoTwoPlus(s.NewStructAR(24, 3)) }
func ReadRootVerTwoTwoPlus(s *C.Segment) VerTwoTwoPlus { return VerTwoTwoPlus(s.Root(0).ToStruct()) }
func (s VerTwoTwoPlus) Val() int16                     { return int16(C.Struct(s).Get16(0)) }
func (s VerTwoTwoPlus) SetVal(v int16)                 { C.Struct(s).Set16(0, uint16(v)) }
func (s VerTwoTwoPlus) Duo() int64                     { return int64(C.Struct(s).Get64(8)) }
func (s VerTwoTwoPlus) SetDuo(v int64)                 { C.Struct(s).Set64(8, uint64(v)) }
func (s VerTwoTwoPlus) Ptr1() VerTwoDataTwoPtr {
	return VerTwoDataTwoPtr(C.Struct(s).GetObject(0).ToStruct())
}
func (s VerTwoTwoPlus) SetPtr1(v VerTwoDataTwoPtr) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s VerTwoTwoPlus) Ptr2() VerTwoDataTwoPtr {
	return VerTwoDataTwoPtr(C.Struct(s).GetObject(1).ToStruct())
}
func (s VerTwoTwoPlus) SetPtr2(v VerTwoDataTwoPtr) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s VerTwoTwoPlus) Tre() int64                 { return int64(C.Struct(s).Get64(16)) }
func (s VerTwoTwoPlus) SetTre(v int64)             { C.Struct(s).Set64(16, uint64(v)) }
func (s VerTwoTwoPlus) Lst3() C.Int64List          { return C.Int64List(C.Struct(s).GetObject(2)) }
func (s VerTwoTwoPlus) SetLst3(v C.Int64List)      { C.Struct(s).SetObject(2, C.Object(v)) }
func (s VerTwoTwoPlus) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"val\":")
	if err != nil {
		return err
	}
	{
		s := s.Val()
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
	_, err = b.WriteString("\"duo\":")
	if err != nil {
		return err
	}
	{
		s := s.Duo()
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
	_, err = b.WriteString("\"ptr1\":")
	if err != nil {
		return err
	}
	{
		s := s.Ptr1()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ptr2\":")
	if err != nil {
		return err
	}
	{
		s := s.Ptr2()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"tre\":")
	if err != nil {
		return err
	}
	{
		s := s.Tre()
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
	_, err = b.WriteString("\"lst3\":")
	if err != nil {
		return err
	}
	{
		s := s.Lst3()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
func (s VerTwoTwoPlus) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VerTwoTwoPlus) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("val = ")
	if err != nil {
		return err
	}
	{
		s := s.Val()
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
	_, err = b.WriteString("duo = ")
	if err != nil {
		return err
	}
	{
		s := s.Duo()
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
	_, err = b.WriteString("ptr1 = ")
	if err != nil {
		return err
	}
	{
		s := s.Ptr1()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("ptr2 = ")
	if err != nil {
		return err
	}
	{
		s := s.Ptr2()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("tre = ")
	if err != nil {
		return err
	}
	{
		s := s.Tre()
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
	_, err = b.WriteString("lst3 = ")
	if err != nil {
		return err
	}
	{
		s := s.Lst3()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
func (s VerTwoTwoPlus) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VerTwoTwoPlus_List C.PointerList

func NewVerTwoTwoPlusList(s *C.Segment, sz int) VerTwoTwoPlus_List {
	return VerTwoTwoPlus_List(s.NewCompositeList(24, 3, sz))
}
func (s VerTwoTwoPlus_List) Len() int { return C.PointerList(s).Len() }
func (s VerTwoTwoPlus_List) At(i int) VerTwoTwoPlus {
	return VerTwoTwoPlus(C.PointerList(s).At(i).ToStruct())
}
func (s VerTwoTwoPlus_List) ToArray() []VerTwoTwoPlus {
	n := s.Len()
	a := make([]VerTwoTwoPlus, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VerTwoTwoPlus_List) Set(i int, item VerTwoTwoPlus) { C.PointerList(s).Set(i, C.Object(item)) }

type HoldsText C.Struct

func NewHoldsText(s *C.Segment) HoldsText      { return HoldsText(s.NewStruct(0, 3)) }
func NewRootHoldsText(s *C.Segment) HoldsText  { return HoldsText(s.NewRootStruct(0, 3)) }
func AutoNewHoldsText(s *C.Segment) HoldsText  { return HoldsText(s.NewStructAR(0, 3)) }
func ReadRootHoldsText(s *C.Segment) HoldsText { return HoldsText(s.Root(0).ToStruct()) }
func (s HoldsText) Txt() string                { return C.Struct(s).GetObject(0).ToText() }
func (s HoldsText) TxtBytes() []byte           { return C.Struct(s).GetObject(0).ToDataTrimLastByte() }
func (s HoldsText) SetTxt(v string)            { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s HoldsText) Lst() C.TextList            { return C.TextList(C.Struct(s).GetObject(1)) }
func (s HoldsText) SetLst(v C.TextList)        { C.Struct(s).SetObject(1, C.Object(v)) }
func (s HoldsText) Lstlst() C.PointerList      { return C.PointerList(C.Struct(s).GetObject(2)) }
func (s HoldsText) SetLstlst(v C.PointerList)  { C.Struct(s).SetObject(2, C.Object(v)) }
func (s HoldsText) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"txt\":")
	if err != nil {
		return err
	}
	{
		s := s.Txt()
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
	_, err = b.WriteString("\"lst\":")
	if err != nil {
		return err
	}
	{
		s := s.Lst()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
	_, err = b.WriteString("\"lstlst\":")
	if err != nil {
		return err
	}
	{
		s := s.Lstlst()
		_ = s
		_, err = b.WriteString("\"untyped list\"")
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
func (s HoldsText) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s HoldsText) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("txt = ")
	if err != nil {
		return err
	}
	{
		s := s.Txt()
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
	_, err = b.WriteString("lst = ")
	if err != nil {
		return err
	}
	{
		s := s.Lst()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
	_, err = b.WriteString("lstlst = ")
	if err != nil {
		return err
	}
	{
		s := s.Lstlst()
		_ = s
		_, err = b.WriteString("\"untyped list\"")
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
func (s HoldsText) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type HoldsText_List C.PointerList

func NewHoldsTextList(s *C.Segment, sz int) HoldsText_List {
	return HoldsText_List(s.NewCompositeList(0, 3, sz))
}
func (s HoldsText_List) Len() int           { return C.PointerList(s).Len() }
func (s HoldsText_List) At(i int) HoldsText { return HoldsText(C.PointerList(s).At(i).ToStruct()) }
func (s HoldsText_List) ToArray() []HoldsText {
	n := s.Len()
	a := make([]HoldsText, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s HoldsText_List) Set(i int, item HoldsText) { C.PointerList(s).Set(i, C.Object(item)) }

type WrapEmpty C.Struct

func NewWrapEmpty(s *C.Segment) WrapEmpty      { return WrapEmpty(s.NewStruct(0, 1)) }
func NewRootWrapEmpty(s *C.Segment) WrapEmpty  { return WrapEmpty(s.NewRootStruct(0, 1)) }
func AutoNewWrapEmpty(s *C.Segment) WrapEmpty  { return WrapEmpty(s.NewStructAR(0, 1)) }
func ReadRootWrapEmpty(s *C.Segment) WrapEmpty { return WrapEmpty(s.Root(0).ToStruct()) }
func (s WrapEmpty) MightNotBeReallyEmpty() VerEmpty {
	return VerEmpty(C.Struct(s).GetObject(0).ToStruct())
}
func (s WrapEmpty) SetMightNotBeReallyEmpty(v VerEmpty) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s WrapEmpty) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mightNotBeReallyEmpty\":")
	if err != nil {
		return err
	}
	{
		s := s.MightNotBeReallyEmpty()
		err = s.WriteJSON(b)
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
func (s WrapEmpty) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s WrapEmpty) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mightNotBeReallyEmpty = ")
	if err != nil {
		return err
	}
	{
		s := s.MightNotBeReallyEmpty()
		err = s.WriteCapLit(b)
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
func (s WrapEmpty) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type WrapEmpty_List C.PointerList

func NewWrapEmptyList(s *C.Segment, sz int) WrapEmpty_List {
	return WrapEmpty_List(s.NewCompositeList(0, 1, sz))
}
func (s WrapEmpty_List) Len() int           { return C.PointerList(s).Len() }
func (s WrapEmpty_List) At(i int) WrapEmpty { return WrapEmpty(C.PointerList(s).At(i).ToStruct()) }
func (s WrapEmpty_List) ToArray() []WrapEmpty {
	n := s.Len()
	a := make([]WrapEmpty, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s WrapEmpty_List) Set(i int, item WrapEmpty) { C.PointerList(s).Set(i, C.Object(item)) }

type Wrap2x2 C.Struct

func NewWrap2x2(s *C.Segment) Wrap2x2      { return Wrap2x2(s.NewStruct(0, 1)) }
func NewRootWrap2x2(s *C.Segment) Wrap2x2  { return Wrap2x2(s.NewRootStruct(0, 1)) }
func AutoNewWrap2x2(s *C.Segment) Wrap2x2  { return Wrap2x2(s.NewStructAR(0, 1)) }
func ReadRootWrap2x2(s *C.Segment) Wrap2x2 { return Wrap2x2(s.Root(0).ToStruct()) }
func (s Wrap2x2) MightNotBeReallyEmpty() VerTwoDataTwoPtr {
	return VerTwoDataTwoPtr(C.Struct(s).GetObject(0).ToStruct())
}
func (s Wrap2x2) SetMightNotBeReallyEmpty(v VerTwoDataTwoPtr) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Wrap2x2) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mightNotBeReallyEmpty\":")
	if err != nil {
		return err
	}
	{
		s := s.MightNotBeReallyEmpty()
		err = s.WriteJSON(b)
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
func (s Wrap2x2) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Wrap2x2) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mightNotBeReallyEmpty = ")
	if err != nil {
		return err
	}
	{
		s := s.MightNotBeReallyEmpty()
		err = s.WriteCapLit(b)
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
func (s Wrap2x2) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Wrap2x2_List C.PointerList

func NewWrap2x2List(s *C.Segment, sz int) Wrap2x2_List {
	return Wrap2x2_List(s.NewCompositeList(0, 1, sz))
}
func (s Wrap2x2_List) Len() int         { return C.PointerList(s).Len() }
func (s Wrap2x2_List) At(i int) Wrap2x2 { return Wrap2x2(C.PointerList(s).At(i).ToStruct()) }
func (s Wrap2x2_List) ToArray() []Wrap2x2 {
	n := s.Len()
	a := make([]Wrap2x2, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Wrap2x2_List) Set(i int, item Wrap2x2) { C.PointerList(s).Set(i, C.Object(item)) }

type Wrap2x2plus C.Struct

func NewWrap2x2plus(s *C.Segment) Wrap2x2plus      { return Wrap2x2plus(s.NewStruct(0, 1)) }
func NewRootWrap2x2plus(s *C.Segment) Wrap2x2plus  { return Wrap2x2plus(s.NewRootStruct(0, 1)) }
func AutoNewWrap2x2plus(s *C.Segment) Wrap2x2plus  { return Wrap2x2plus(s.NewStructAR(0, 1)) }
func ReadRootWrap2x2plus(s *C.Segment) Wrap2x2plus { return Wrap2x2plus(s.Root(0).ToStruct()) }
func (s Wrap2x2plus) MightNotBeReallyEmpty() VerTwoTwoPlus {
	return VerTwoTwoPlus(C.Struct(s).GetObject(0).ToStruct())
}
func (s Wrap2x2plus) SetMightNotBeReallyEmpty(v VerTwoTwoPlus) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Wrap2x2plus) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"mightNotBeReallyEmpty\":")
	if err != nil {
		return err
	}
	{
		s := s.MightNotBeReallyEmpty()
		err = s.WriteJSON(b)
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
func (s Wrap2x2plus) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Wrap2x2plus) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("mightNotBeReallyEmpty = ")
	if err != nil {
		return err
	}
	{
		s := s.MightNotBeReallyEmpty()
		err = s.WriteCapLit(b)
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
func (s Wrap2x2plus) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Wrap2x2plus_List C.PointerList

func NewWrap2x2plusList(s *C.Segment, sz int) Wrap2x2plus_List {
	return Wrap2x2plus_List(s.NewCompositeList(0, 1, sz))
}
func (s Wrap2x2plus_List) Len() int             { return C.PointerList(s).Len() }
func (s Wrap2x2plus_List) At(i int) Wrap2x2plus { return Wrap2x2plus(C.PointerList(s).At(i).ToStruct()) }
func (s Wrap2x2plus_List) ToArray() []Wrap2x2plus {
	n := s.Len()
	a := make([]Wrap2x2plus, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Wrap2x2plus_List) Set(i int, item Wrap2x2plus) { C.PointerList(s).Set(i, C.Object(item)) }

type Endpoint C.Struct

func NewEndpoint(s *C.Segment) Endpoint      { return Endpoint(s.NewStruct(8, 2)) }
func NewRootEndpoint(s *C.Segment) Endpoint  { return Endpoint(s.NewRootStruct(8, 2)) }
func AutoNewEndpoint(s *C.Segment) Endpoint  { return Endpoint(s.NewStructAR(8, 2)) }
func ReadRootEndpoint(s *C.Segment) Endpoint { return Endpoint(s.Root(0).ToStruct()) }
func (s Endpoint) Ip() net.IP                { return net.IP(C.Struct(s).GetObject(0).ToData()) }
func (s Endpoint) SetIp(v net.IP)            { C.Struct(s).SetObject(0, s.Segment.NewData([]byte(v))) }
func (s Endpoint) Port() int16               { return int16(C.Struct(s).Get16(0)) }
func (s Endpoint) SetPort(v int16)           { C.Struct(s).Set16(0, uint16(v)) }
func (s Endpoint) Hostname() string          { return C.Struct(s).GetObject(1).ToText() }
func (s Endpoint) HostnameBytes() []byte     { return C.Struct(s).GetObject(1).ToDataTrimLastByte() }
func (s Endpoint) SetHostname(v string)      { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s Endpoint) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"ip\":")
	if err != nil {
		return err
	}
	{
		s := s.Ip()
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
	_, err = b.WriteString("\"port\":")
	if err != nil {
		return err
	}
	{
		s := s.Port()
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
	_, err = b.WriteString("\"hostname\":")
	if err != nil {
		return err
	}
	{
		s := s.Hostname()
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
func (s Endpoint) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Endpoint) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("ip = ")
	if err != nil {
		return err
	}
	{
		s := s.Ip()
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
	_, err = b.WriteString("port = ")
	if err != nil {
		return err
	}
	{
		s := s.Port()
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
	_, err = b.WriteString("hostname = ")
	if err != nil {
		return err
	}
	{
		s := s.Hostname()
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
func (s Endpoint) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Endpoint_List C.PointerList

func NewEndpointList(s *C.Segment, sz int) Endpoint_List {
	return Endpoint_List(s.NewCompositeList(8, 2, sz))
}
func (s Endpoint_List) Len() int          { return C.PointerList(s).Len() }
func (s Endpoint_List) At(i int) Endpoint { return Endpoint(C.PointerList(s).At(i).ToStruct()) }
func (s Endpoint_List) ToArray() []Endpoint {
	n := s.Len()
	a := make([]Endpoint, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Endpoint_List) Set(i int, item Endpoint) { C.PointerList(s).Set(i, C.Object(item)) }

type VoidUnion C.Struct
type VoidUnion_Which uint16

const (
	VOIDUNION_A VoidUnion_Which = 0
	VOIDUNION_B VoidUnion_Which = 1
)

func NewVoidUnion(s *C.Segment) VoidUnion      { return VoidUnion(s.NewStruct(8, 0)) }
func NewRootVoidUnion(s *C.Segment) VoidUnion  { return VoidUnion(s.NewRootStruct(8, 0)) }
func AutoNewVoidUnion(s *C.Segment) VoidUnion  { return VoidUnion(s.NewStructAR(8, 0)) }
func ReadRootVoidUnion(s *C.Segment) VoidUnion { return VoidUnion(s.Root(0).ToStruct()) }
func (s VoidUnion) Which() VoidUnion_Which     { return VoidUnion_Which(C.Struct(s).Get16(0)) }
func (s VoidUnion) SetA()                      { C.Struct(s).Set16(0, 0) }
func (s VoidUnion) SetB()                      { C.Struct(s).Set16(0, 1) }
func (s VoidUnion) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	if s.Which() == VOIDUNION_A {
		_, err = b.WriteString("\"a\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == VOIDUNION_B {
		_, err = b.WriteString("\"b\":")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
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
func (s VoidUnion) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s VoidUnion) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	if s.Which() == VOIDUNION_A {
		_, err = b.WriteString("a = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
		if err != nil {
			return err
		}
	}
	if s.Which() == VOIDUNION_B {
		_, err = b.WriteString("b = ")
		if err != nil {
			return err
		}
		_ = s
		_, err = b.WriteString("null")
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
func (s VoidUnion) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type VoidUnion_List C.PointerList

func NewVoidUnionList(s *C.Segment, sz int) VoidUnion_List {
	return VoidUnion_List(s.NewCompositeList(8, 0, sz))
}
func (s VoidUnion_List) Len() int           { return C.PointerList(s).Len() }
func (s VoidUnion_List) At(i int) VoidUnion { return VoidUnion(C.PointerList(s).At(i).ToStruct()) }
func (s VoidUnion_List) ToArray() []VoidUnion {
	n := s.Len()
	a := make([]VoidUnion, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s VoidUnion_List) Set(i int, item VoidUnion) { C.PointerList(s).Set(i, C.Object(item)) }

type Nester1Capn C.Struct

func NewNester1Capn(s *C.Segment) Nester1Capn      { return Nester1Capn(s.NewStruct(0, 1)) }
func NewRootNester1Capn(s *C.Segment) Nester1Capn  { return Nester1Capn(s.NewRootStruct(0, 1)) }
func AutoNewNester1Capn(s *C.Segment) Nester1Capn  { return Nester1Capn(s.NewStructAR(0, 1)) }
func ReadRootNester1Capn(s *C.Segment) Nester1Capn { return Nester1Capn(s.Root(0).ToStruct()) }
func (s Nester1Capn) Strs() C.TextList             { return C.TextList(C.Struct(s).GetObject(0)) }
func (s Nester1Capn) SetStrs(v C.TextList)         { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Nester1Capn) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"strs\":")
	if err != nil {
		return err
	}
	{
		s := s.Strs()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
func (s Nester1Capn) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Nester1Capn) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("strs = ")
	if err != nil {
		return err
	}
	{
		s := s.Strs()
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
				buf, err = json.Marshal(s)
				if err != nil {
					return err
				}
				_, err = b.Write(buf)
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
func (s Nester1Capn) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Nester1Capn_List C.PointerList

func NewNester1CapnList(s *C.Segment, sz int) Nester1Capn_List {
	return Nester1Capn_List(s.NewCompositeList(0, 1, sz))
}
func (s Nester1Capn_List) Len() int             { return C.PointerList(s).Len() }
func (s Nester1Capn_List) At(i int) Nester1Capn { return Nester1Capn(C.PointerList(s).At(i).ToStruct()) }
func (s Nester1Capn_List) ToArray() []Nester1Capn {
	n := s.Len()
	a := make([]Nester1Capn, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Nester1Capn_List) Set(i int, item Nester1Capn) { C.PointerList(s).Set(i, C.Object(item)) }

type RWTestCapn C.Struct

func NewRWTestCapn(s *C.Segment) RWTestCapn        { return RWTestCapn(s.NewStruct(0, 1)) }
func NewRootRWTestCapn(s *C.Segment) RWTestCapn    { return RWTestCapn(s.NewRootStruct(0, 1)) }
func AutoNewRWTestCapn(s *C.Segment) RWTestCapn    { return RWTestCapn(s.NewStructAR(0, 1)) }
func ReadRootRWTestCapn(s *C.Segment) RWTestCapn   { return RWTestCapn(s.Root(0).ToStruct()) }
func (s RWTestCapn) NestMatrix() C.PointerList     { return C.PointerList(C.Struct(s).GetObject(0)) }
func (s RWTestCapn) SetNestMatrix(v C.PointerList) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s RWTestCapn) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"nestMatrix\":")
	if err != nil {
		return err
	}
	{
		s := s.NestMatrix()
		_ = s
		_, err = b.WriteString("\"untyped list\"")
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
func (s RWTestCapn) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s RWTestCapn) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("nestMatrix = ")
	if err != nil {
		return err
	}
	{
		s := s.NestMatrix()
		_ = s
		_, err = b.WriteString("\"untyped list\"")
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
func (s RWTestCapn) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type RWTestCapn_List C.PointerList

func NewRWTestCapnList(s *C.Segment, sz int) RWTestCapn_List {
	return RWTestCapn_List(s.NewCompositeList(0, 1, sz))
}
func (s RWTestCapn_List) Len() int            { return C.PointerList(s).Len() }
func (s RWTestCapn_List) At(i int) RWTestCapn { return RWTestCapn(C.PointerList(s).At(i).ToStruct()) }
func (s RWTestCapn_List) ToArray() []RWTestCapn {
	n := s.Len()
	a := make([]RWTestCapn, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s RWTestCapn_List) Set(i int, item RWTestCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type ListStructCapn C.Struct

func NewListStructCapn(s *C.Segment) ListStructCapn      { return ListStructCapn(s.NewStruct(0, 1)) }
func NewRootListStructCapn(s *C.Segment) ListStructCapn  { return ListStructCapn(s.NewRootStruct(0, 1)) }
func AutoNewListStructCapn(s *C.Segment) ListStructCapn  { return ListStructCapn(s.NewStructAR(0, 1)) }
func ReadRootListStructCapn(s *C.Segment) ListStructCapn { return ListStructCapn(s.Root(0).ToStruct()) }
func (s ListStructCapn) Vec() Nester1Capn_List           { return Nester1Capn_List(C.Struct(s).GetObject(0)) }
func (s ListStructCapn) SetVec(v Nester1Capn_List)       { C.Struct(s).SetObject(0, C.Object(v)) }
func (s ListStructCapn) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"vec\":")
	if err != nil {
		return err
	}
	{
		s := s.Vec()
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
func (s ListStructCapn) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ListStructCapn) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("vec = ")
	if err != nil {
		return err
	}
	{
		s := s.Vec()
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
func (s ListStructCapn) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type ListStructCapn_List C.PointerList

func NewListStructCapnList(s *C.Segment, sz int) ListStructCapn_List {
	return ListStructCapn_List(s.NewCompositeList(0, 1, sz))
}
func (s ListStructCapn_List) Len() int { return C.PointerList(s).Len() }
func (s ListStructCapn_List) At(i int) ListStructCapn {
	return ListStructCapn(C.PointerList(s).At(i).ToStruct())
}
func (s ListStructCapn_List) ToArray() []ListStructCapn {
	n := s.Len()
	a := make([]ListStructCapn, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ListStructCapn_List) Set(i int, item ListStructCapn) { C.PointerList(s).Set(i, C.Object(item)) }

type StackingRoot C.Struct

func NewStackingRoot(s *C.Segment) StackingRoot      { return StackingRoot(s.NewStruct(0, 2)) }
func NewRootStackingRoot(s *C.Segment) StackingRoot  { return StackingRoot(s.NewRootStruct(0, 2)) }
func AutoNewStackingRoot(s *C.Segment) StackingRoot  { return StackingRoot(s.NewStructAR(0, 2)) }
func ReadRootStackingRoot(s *C.Segment) StackingRoot { return StackingRoot(s.Root(0).ToStruct()) }
func (s StackingRoot) A() StackingA                  { return StackingA(C.Struct(s).GetObject(1).ToStruct()) }
func (s StackingRoot) SetA(v StackingA)              { C.Struct(s).SetObject(1, C.Object(v)) }
func (s StackingRoot) AWithDefault() StackingA {
	return StackingA(C.Struct(s).GetObject(0).ToStructDefault(x_832bcc6686a26d56, 0))
}
func (s StackingRoot) SetAWithDefault(v StackingA) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s StackingRoot) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"a\":")
	if err != nil {
		return err
	}
	{
		s := s.A()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"aWithDefault\":")
	if err != nil {
		return err
	}
	{
		s := s.AWithDefault()
		err = s.WriteJSON(b)
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
func (s StackingRoot) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s StackingRoot) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("a = ")
	if err != nil {
		return err
	}
	{
		s := s.A()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("aWithDefault = ")
	if err != nil {
		return err
	}
	{
		s := s.AWithDefault()
		err = s.WriteCapLit(b)
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
func (s StackingRoot) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type StackingRoot_List C.PointerList

func NewStackingRootList(s *C.Segment, sz int) StackingRoot_List {
	return StackingRoot_List(s.NewCompositeList(0, 2, sz))
}
func (s StackingRoot_List) Len() int { return C.PointerList(s).Len() }
func (s StackingRoot_List) At(i int) StackingRoot {
	return StackingRoot(C.PointerList(s).At(i).ToStruct())
}
func (s StackingRoot_List) ToArray() []StackingRoot {
	n := s.Len()
	a := make([]StackingRoot, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s StackingRoot_List) Set(i int, item StackingRoot) { C.PointerList(s).Set(i, C.Object(item)) }

type StackingA C.Struct

func NewStackingA(s *C.Segment) StackingA      { return StackingA(s.NewStruct(8, 1)) }
func NewRootStackingA(s *C.Segment) StackingA  { return StackingA(s.NewRootStruct(8, 1)) }
func AutoNewStackingA(s *C.Segment) StackingA  { return StackingA(s.NewStructAR(8, 1)) }
func ReadRootStackingA(s *C.Segment) StackingA { return StackingA(s.Root(0).ToStruct()) }
func (s StackingA) Num() int32                 { return int32(C.Struct(s).Get32(0)) }
func (s StackingA) SetNum(v int32)             { C.Struct(s).Set32(0, uint32(v)) }
func (s StackingA) B() StackingB               { return StackingB(C.Struct(s).GetObject(0).ToStruct()) }
func (s StackingA) SetB(v StackingB)           { C.Struct(s).SetObject(0, C.Object(v)) }
func (s StackingA) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"num\":")
	if err != nil {
		return err
	}
	{
		s := s.Num()
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
	_, err = b.WriteString("\"b\":")
	if err != nil {
		return err
	}
	{
		s := s.B()
		err = s.WriteJSON(b)
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
func (s StackingA) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s StackingA) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("num = ")
	if err != nil {
		return err
	}
	{
		s := s.Num()
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
	_, err = b.WriteString("b = ")
	if err != nil {
		return err
	}
	{
		s := s.B()
		err = s.WriteCapLit(b)
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
func (s StackingA) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type StackingA_List C.PointerList

func NewStackingAList(s *C.Segment, sz int) StackingA_List {
	return StackingA_List(s.NewCompositeList(8, 1, sz))
}
func (s StackingA_List) Len() int           { return C.PointerList(s).Len() }
func (s StackingA_List) At(i int) StackingA { return StackingA(C.PointerList(s).At(i).ToStruct()) }
func (s StackingA_List) ToArray() []StackingA {
	n := s.Len()
	a := make([]StackingA, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s StackingA_List) Set(i int, item StackingA) { C.PointerList(s).Set(i, C.Object(item)) }

type StackingB C.Struct

func NewStackingB(s *C.Segment) StackingB      { return StackingB(s.NewStruct(8, 0)) }
func NewRootStackingB(s *C.Segment) StackingB  { return StackingB(s.NewRootStruct(8, 0)) }
func AutoNewStackingB(s *C.Segment) StackingB  { return StackingB(s.NewStructAR(8, 0)) }
func ReadRootStackingB(s *C.Segment) StackingB { return StackingB(s.Root(0).ToStruct()) }
func (s StackingB) Num() int32                 { return int32(C.Struct(s).Get32(0)) }
func (s StackingB) SetNum(v int32)             { C.Struct(s).Set32(0, uint32(v)) }
func (s StackingB) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"num\":")
	if err != nil {
		return err
	}
	{
		s := s.Num()
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
func (s StackingB) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s StackingB) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("num = ")
	if err != nil {
		return err
	}
	{
		s := s.Num()
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
func (s StackingB) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type StackingB_List C.PointerList

func NewStackingBList(s *C.Segment, sz int) StackingB_List {
	return StackingB_List(s.NewCompositeList(8, 0, sz))
}
func (s StackingB_List) Len() int           { return C.PointerList(s).Len() }
func (s StackingB_List) At(i int) StackingB { return StackingB(C.PointerList(s).At(i).ToStruct()) }
func (s StackingB_List) ToArray() []StackingB {
	n := s.Len()
	a := make([]StackingB, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s StackingB_List) Set(i int, item StackingB) { C.PointerList(s).Set(i, C.Object(item)) }

var x_832bcc6686a26d56 = C.NewBuffer([]byte{
	0, 0, 0, 0, 1, 0, 1, 0,
	42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
})
