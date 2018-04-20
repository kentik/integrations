// Contains constants and maps related to interface classification types, i.e.
// network boundary types and connectivity types.
// ic = interface classification
// nb = network boundary
// ct = connectivity type
package ic

import (
	"fmt"
)

func NameFromCTInt(ic int) string {
	if v, ok := CONNECTIVITY_TYPE_INT_TO_NAME[ic]; ok {
		return v
	}
	return NONE_NAME
}

func IntFromCTName(ic string) int {
	if v, ok := CONNECTIVITY_TYPE_NAME_TO_INT[ic]; ok {
		return v
	}
	return NONE_INT
}

func NameFromNBInt(ic int) string {
	if v, ok := NETWORK_BOUNDARY_INT_TO_NAME[ic]; ok {
		return v
	}
	return NONE_NAME
}

func IntFromNBName(ic string) int {
	if v, ok := NETWORK_BOUNDARY_NAME_TO_INT[ic]; ok {
		return v
	}
	return NONE_INT
}

func TrafficProfileNumbersFromName(prof string) (uint32, uint32) {
	var (
		ov uint32
		dv uint32
	)

	switch prof {
	case NETWORK_SRC_INTERNAL_NAME:
		ov = NETWORK_SRC_INTERNAL
		dv = NETWORK_SRC_INTERNAL
	case NETWORK_SRC_THROUGH_NAME:
		ov = NETWORK_SRC_EXTERNAL
		dv = NETWORK_SRC_EXTERNAL
	case NETWORK_SRC_TERMINATED_NAME, NETWORK_SRC_TERMINATED_NAME_TRUNCATED:
		ov = NETWORK_SRC_EXTERNAL
		dv = NETWORK_SRC_INTERNAL
	case NETWORK_SRC_ORIGINATED_NAME, NETWORK_SRC_ORIGINATED_NAME_TRUNCATED:
		ov = NETWORK_SRC_INTERNAL
		dv = NETWORK_SRC_EXTERNAL
	}

	return ov, dv
}

func TrafficNameFromNumbers(ov uint32, dv uint32) string {
	if ov == NETWORK_SRC_INTERNAL && dv == NETWORK_SRC_EXTERNAL {
		return NETWORK_SRC_ORIGINATED_NAME
	} else if ov == NETWORK_SRC_EXTERNAL && dv == NETWORK_SRC_INTERNAL {
		return NETWORK_SRC_TERMINATED_NAME
	} else if ov == NETWORK_SRC_EXTERNAL && dv == NETWORK_SRC_EXTERNAL {
		return NETWORK_SRC_THROUGH_NAME
	} else if ov == NETWORK_SRC_INTERNAL && dv == NETWORK_SRC_INTERNAL {
		return NETWORK_SRC_INTERNAL_NAME
	}
	return fmt.Sprintf("%d -> %d", ov, dv)
}

const (
	NB_EXTERNAL = 10
	NB_INTERNAL = 20

	NB_EXTERNAL_NAME = "external"
	NB_INTERNAL_NAME = "internal"

	CT_FREE_PNI  = 5
	CT_CUSTOMER  = 15
	CT_HOST      = 25
	CT_BACKBONE  = 35
	CT_PAID_PNI  = 45
	CT_OTHER     = 55
	CT_TRANSIT   = 65
	CT_IX        = 75
	CT_RESERVED  = 85
	CT_AVAILABLE = 95
	CT_DC_IC     = 105
	CT_AGG_IC    = 115

	CT_FREE_PNI_NAME  = "free_pni"
	CT_CUSTOMER_NAME  = "customer"
	CT_HOST_NAME      = "host"
	CT_BACKBONE_NAME  = "backbone"
	CT_PAID_PNI_NAME  = "paid_pni"
	CT_OTHER_NAME     = "other"
	CT_TRANSIT_NAME   = "transit"
	CT_IX_NAME        = "ix"
	CT_RESERVED_NAME  = "reserved"
	CT_AVAILABLE_NAME = "available"
	CT_DC_IC_NAME     = "datacenter_interconnect"
	CT_AGG_IC_NAME    = "aggregation_interconnect"

	NETWORK_CL_INTERNAL_NAME              = "inside"
	NETWORK_CL_EXTERNAL_NAME              = "outside"
	NETWORK_SRC_INTERNAL                  = uint32(10)
	NETWORK_SRC_EXTERNAL                  = uint32(20)
	NETWORK_SRC_INTERNAL_NAME             = "internal"
	NETWORK_SRC_THROUGH_NAME              = "through"
	NETWORK_SRC_TERMINATED_NAME           = "from outside, terminated inside"
	NETWORK_SRC_ORIGINATED_NAME           = "originated inside, to outside"
	NETWORK_SRC_TERMINATED_NAME_TRUNCATED = "from outside"
	NETWORK_SRC_ORIGINATED_NAME_TRUNCATED = "originated inside"
	NETWORK_DIR_IN_NAME                   = "in"
	NETWORK_DIR_OUT_NAME                  = "out"
	NETWORK_DIR_NOT_HOST_NAME             = "not_a_host"
	NETWORK_DIR_ERR_NAME                  = "error"
	NETWORK_DIR_IN                        = uint32(40)
	NETWORK_DIR_OUT                       = uint32(50)
	NETWORK_DIR_NOT_HOST                  = uint32(60)
	NETWORK_DIR_ERR                       = uint32(70)

	NONE_INT = 0

	NONE_NAME = "none"
)

var (
	NETWORK_BOUNDARY_INT_TO_NAME = map[int]string{
		NB_EXTERNAL: NB_EXTERNAL_NAME,
		NB_INTERNAL: NB_INTERNAL_NAME,
	}

	NETWORK_CLASS_INT_TO_NAME = map[uint32]string{
		NETWORK_SRC_EXTERNAL: NETWORK_CL_EXTERNAL_NAME,
		NETWORK_SRC_INTERNAL: NETWORK_CL_INTERNAL_NAME,
		NETWORK_DIR_IN:       NETWORK_DIR_IN_NAME,
		NETWORK_DIR_OUT:      NETWORK_DIR_OUT_NAME,
		NETWORK_DIR_NOT_HOST: NETWORK_DIR_NOT_HOST_NAME,
		NETWORK_DIR_ERR:      NETWORK_DIR_ERR_NAME,
	}

	CONNECTIVITY_TYPE_INT_TO_NAME = map[int]string{
		CT_FREE_PNI:  CT_FREE_PNI_NAME,
		CT_CUSTOMER:  CT_CUSTOMER_NAME,
		CT_HOST:      CT_HOST_NAME,
		CT_BACKBONE:  CT_BACKBONE_NAME,
		CT_PAID_PNI:  CT_PAID_PNI_NAME,
		CT_OTHER:     CT_OTHER_NAME,
		CT_TRANSIT:   CT_TRANSIT_NAME,
		CT_IX:        CT_IX_NAME,
		CT_RESERVED:  CT_RESERVED_NAME,
		CT_AVAILABLE: CT_AVAILABLE_NAME,
		CT_DC_IC:     CT_DC_IC_NAME,
		CT_AGG_IC:    CT_AGG_IC_NAME,
	}

	NETWORK_BOUNDARY_NAME_TO_INT = func() map[string]int {
		m := make(map[string]int)
		for k, v := range NETWORK_BOUNDARY_INT_TO_NAME {
			m[v] = k
		}
		return m
	}()

	CONNECTIVITY_TYPE_NAME_TO_INT = func() map[string]int {
		m := make(map[string]int)
		for k, v := range CONNECTIVITY_TYPE_INT_TO_NAME {
			m[v] = k
		}
		return m
	}()

	NETWORK_CLASS_NAME_TO_INT = func() map[string]uint32 {
		m := make(map[string]uint32)
		for k, v := range NETWORK_CLASS_INT_TO_NAME {
			m[v] = k
		}
		return m
	}()
)
