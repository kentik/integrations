package kt

import (
	"strconv"
)

type IntId uint64
type Cid IntId // company id

func AtoiCid(s string) (Cid, error) { id, err := strconv.Atoi(s); return Cid(id), err }
