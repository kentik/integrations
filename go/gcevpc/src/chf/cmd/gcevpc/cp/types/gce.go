package types

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	flowclient "chf/cmd/gcevpc/cp/client"

	"github.com/kentik/gohippo"
	"github.com/kentik/libkflow/api"
	"github.com/kentik/libkflow/flow"
)

/**
{
  "insertId": "8vn8gwf16y866",
  "jsonPayload": {
    "bytes_sent": "520429",
    "connection": {
      "dest_ip": "10.128.0.2",
      "dest_port": 43350,
      "protocol": 6,
      "src_ip": "104.91.207.177",
      "src_port": 443
    },
    "dest_instance": {
      "project_id": "kentik-continuous-delivery",
      "region": "us-central1",
      "vm_name": "avi-flow1",
      "zone": "us-central1-c"
    },
    "dest_vpc": {
      "project_id": "kentik-continuous-delivery",
      "subnetwork_name": "default",
      "vpc_name": "default"
    },
    "end_time": "2018-04-17T18:48:30.867049953Z",
    "packets_sent": "185",
    "reporter": "DEST",
    "rtt_msec": "13",
    "src_location": {},
    "start_time": "2018-04-17T18:48:30.776952635Z"
  },
  "logName": "projects/kentik-continuous-delivery/logs/compute.googleapis.com%2Fvpc_flows",
  "receiveTimestamp": "2018-04-17T18:48:36.288089Z",
  "resource": {
    "labels": {
      "location": "us-central1-c",
      "project_id": "kentik-continuous-delivery",
      "subnetwork_id": "7839590128170438108",
      "subnetwork_name": "default"
    },
    "type": "gce_subnetwork"
  },
  "timestamp": "2018-04-17T18:48:36.288089Z"
}
*/

const (
	SRC_PROJECT_ID = "c_gce_src_project_id"
	SRC_VM_NAME    = "c_gce_src_vm_name"
	SRC_ZONE       = "c_gce_src_zone"
	SRC_VPC_SNN    = "c_gce_src_vpc_snn"
	DST_PROJECT_ID = "c_gce_dst_project_id"
	DST_VM_NAME    = "c_gce_dst_vm_name"
	DST_ZONE       = "c_gce_dst_zone"
	DST_VPC_SNN    = "c_gce_dst_vpc_snn"
	REPORTER       = "c_gce_reporter"

	RECV_WINDOW = -1 * 5 * 60 * time.Second
)

var (
	GCEColumns = []string{
		SRC_PROJECT_ID,
		SRC_VM_NAME,
		SRC_ZONE,
		SRC_VPC_SNN,
		DST_PROJECT_ID,
		DST_VM_NAME,
		DST_ZONE,
		DST_VPC_SNN,
		REPORTER,
	}
)

type GCELogLine struct {
	InsertID  string    `json:"insertId"`
	Payload   *Payload  `json:"jsonPayload"`
	LogName   string    `json:"logName"`
	RecvTs    string    `json:"receiveTimestamp"`
	Resource  *Resource `json:"resource"`
	Timestamp string    `json:"timestamp"`
}

type Connection struct {
	DestIP   string `json:"dest_ip"`
	DestPort int    `json:"dest_port"`
	Protocol int    `json:"protocol"`
	SrcIP    string `json:"src_ip"`
	SrcPort  int    `json:"src_port"`
}

type Instance struct {
	ProjectID string `json:"project_id"`
	Region    string `json:"region"`
	VMName    string `json:"vm_name"`
	Zone      string `json:"zone"`
}

type VPC struct {
	ProjectID      string `json:"project_id"`
	SubnetworkName string `json:"subnetwork_name"`
	Name           string `json:"vpc_name"`
}

type Payload struct {
	Bytes        string      `json:"bytes_sent"`
	Connection   *Connection `json:"connection"`
	DestInstance *Instance   `json:"dest_instance"`
	SrcInstance  *Instance   `json:"src_instance"`
	DestVPC      *VPC        `json:"dest_vpc"`
	SrcVPC       *VPC        `json:"src_vpc"`
	EndTime      string      `json:"end_time"`
	Pkts         string      `json:"packets_sent"`
	Reporter     string      `json:"reporter"`
	RTT          string      `json:"rtt_msec"`
	SrcLocation  *Location   `json:"src_location"`
	DstLocation  *Location   `json:"dest_location"`
	StartTime    string      `json:"start_time"`
}

type Location struct {
	City      string `json:"city"`
	Continent string `json:"continent"`
	Country   string `json:"country"`
	Region    string `json:"region"`
}

type Resource struct {
	Labels *Labels `json:"labels"`
	Type   string  `json:"type"`
}

type Labels struct {
	Location       string `json:"location"`
	ProjectID      string `json:"project_id"`
	SubnetworkID   string `json:"subnetwork_id"`
	SubnetworkName string `json:"subnetwork_name"`
}

func (m *GCELogLine) GetTimestamp() time.Time {
	t, _ := time.Parse(time.RFC3339, m.Payload.EndTime)
	return t
}

func (m *GCELogLine) GetHost(deviceMap string) (host string, err error) {
	defer func() {
		if r := recover(); r != nil {
			json, errI := json.Marshal(m)
			if errI != nil {
				err = errI
			} else {
				err = fmt.Errorf("%v -> %s", r, string(json))
			}
		}
	}()

	switch deviceMap {
	case "subnet":
		if m.IsIn() {
			host = m.Payload.SrcVPC.SubnetworkName
		} else {
			host = m.Payload.DestVPC.SubnetworkName
		}
	case "vmname":
		return m.GetVMName()
	case "project":
		if m.IsIn() {
			host = strings.Join([]string{m.Payload.SrcInstance.ProjectID, m.Payload.SrcInstance.Region, m.Payload.SrcInstance.Zone}, "_")
		} else {
			host = strings.Join([]string{m.Payload.DestInstance.ProjectID, m.Payload.DestInstance.Region, m.Payload.DestInstance.Zone}, "_")
		}
	default:
		return "", fmt.Errorf("Invalid device map: %s", deviceMap)
	}

	// Hack to avoid breaking Kentik.
	if host == "default" {
		host = "gce_" + host
	}

	return host, nil
}

func (m *GCELogLine) GetVMName() (host string, err error) {
	defer func() {
		if r := recover(); r != nil {
			json, errI := json.Marshal(m)
			if errI != nil {
				err = errI
			} else {
				err = fmt.Errorf("%v -> %s", r, string(json))
			}
		}
	}()

	if m.IsIn() {
		host = m.Payload.SrcInstance.VMName
	} else {
		host = m.Payload.DestInstance.VMName
	}

	return host, nil
}

func (m *GCELogLine) GetInterface() (*api.InterfaceUpdate, error) {
	vm, err := m.GetVMName()
	if err != nil {
		return nil, err
	}

	if m.IsIn() {
		return &api.InterfaceUpdate{
			Alias:   vm,
			Address: m.Payload.Connection.SrcIP,
		}, nil
	} else {
		return &api.InterfaceUpdate{
			Alias:   vm,
			Address: m.Payload.Connection.DestIP,
		}, nil
	}
}

func (m *GCELogLine) GetDeviceConfig(plan int, site int, host string) *api.DeviceCreate {
	dev := &api.DeviceCreate{
		Name:        "",
		Type:        "host-nprobe-dns-www",
		Description: "",
		SampleRate:  1,
		BgpType:     "none",
		PlanID:      plan,
		SiteID:      site,
		IPs:         []net.IP{},
		CdnAttr:     "N",
	}

	if m.IsIn() {
		dev.Name = host
		dev.Description = fmt.Sprintf("GCE VM %s %s", m.Payload.SrcInstance.ProjectID, m.Payload.SrcVPC.Name)
		dev.IPs = append(dev.IPs, net.ParseIP(m.Payload.Connection.SrcIP))
	} else {
		dev.Name = host
		dev.Description = fmt.Sprintf("GCE VM %s %s", m.Payload.DestInstance.ProjectID, m.Payload.DestVPC.Name)
		dev.IPs = append(dev.IPs, net.ParseIP(m.Payload.Connection.DestIP))
	}

	return dev
}

func (m *GCELogLine) SetTags(upserts map[string][]hippo.Upsert) (map[string][]hippo.Upsert, int, error) {
	done := 0

	// Pre-populate this.
	fullUpserts := map[string][]hippo.Upsert{}
	for s, v := range upserts {
		fullUpserts[s] = v
	}

	for _, col := range GCEColumns {
		var req *hippo.Req

		if m.IsIn() {
			req = &hippo.Req{
				Replace:  false,
				Complete: true,
				Upserts: []hippo.Upsert{
					{
						Val: "",
						Rules: []hippo.Rule{
							{
								Dir:         "src",
								IPAddresses: []string{m.Payload.Connection.SrcIP},
							},
						},
					},
				},
			}
			switch col {
			case SRC_PROJECT_ID:
				req.Upserts[0].Val = m.Payload.SrcInstance.ProjectID
			case SRC_VM_NAME:
				req.Upserts[0].Val = m.Payload.SrcInstance.VMName
			case SRC_ZONE:
				req.Upserts[0].Val = m.Payload.SrcInstance.Zone
			case SRC_VPC_SNN:
				req.Upserts[0].Val = m.Payload.SrcVPC.SubnetworkName
			case REPORTER:
				req.Upserts[0].Val = m.Payload.Reporter
			}
		} else {
			req = &hippo.Req{
				Replace:  false,
				Complete: true,
				Upserts: []hippo.Upsert{
					{
						Val: "",
						Rules: []hippo.Rule{
							{
								Dir:         "dst",
								IPAddresses: []string{m.Payload.Connection.DestIP},
							},
						},
					},
				},
			}
			switch col {
			case DST_PROJECT_ID:
				req.Upserts[0].Val = m.Payload.DestInstance.ProjectID
			case DST_VM_NAME:
				req.Upserts[0].Val = m.Payload.DestInstance.VMName
			case DST_ZONE:
				req.Upserts[0].Val = m.Payload.DestInstance.Zone
			case DST_VPC_SNN:
				req.Upserts[0].Val = m.Payload.DestVPC.SubnetworkName
			case REPORTER:
				req.Upserts[0].Val = m.Payload.Reporter
			}
		}

		if req.Upserts[0].Val != "" {
			if old, ok := upserts[col]; ok {
				for _, oldCol := range old {
					if oldCol.Val != "" {
						if oldCol.Val == req.Upserts[0].Val {
							req.Upserts[0].Rules[0].IPAddresses = append(req.Upserts[0].Rules[0].IPAddresses, oldCol.Rules[0].IPAddresses...)
						} else {
							req.Upserts = append(req.Upserts, oldCol)
						}
					}
				}
			}

			newUps := []hippo.Upsert{}
			for _, u := range req.Upserts {
				if u.Val != "" {
					newUps = append(newUps, u)
				}
			}
			req.Upserts = newUps
			done++
			fullUpserts[col] = req.Upserts
		}
	}

	return fullUpserts, done, nil
}

func (m *GCELogLine) IsValid() bool {
	if m.Payload != nil {
		t := m.GetTimestamp()
		return t.After(time.Now().Add(RECV_WINDOW))
	}

	return false
}

func (m *GCELogLine) IsIn() bool {
	return m.Payload.SrcInstance != nil && m.Payload.SrcInstance.VMName != ""
}

func (m *GCELogLine) IsInternal() bool {
	return (m.Payload.SrcInstance != nil && m.Payload.SrcInstance.VMName != "") && (m.Payload.DestInstance != nil && m.Payload.DestInstance.VMName != "")
}

func (m *GCELogLine) ToJson() []byte {
	json, _ := json.Marshal(m)
	return json
}

func (m *GCELogLine) ToFlow(customs map[string]uint32, client *flowclient.FlowClient, dropIntraDest, dropIntraSrc bool) (in *flow.Flow, err error) {
	defer func() {
		if r := recover(); r != nil {
			json, errI := json.Marshal(m)
			if errI != nil {
				err = errI
			} else {
				err = fmt.Errorf("%v -> %s", r, string(json))
			}
		}
	}()

	srcVM := ""
	if m.Payload.SrcInstance != nil {
		srcVM = m.Payload.SrcInstance.VMName
	}

	dstVM := ""
	if m.Payload.DestInstance != nil {
		dstVM = m.Payload.DestInstance.VMName
	}

	if m.IsIn() {
		if dropIntraSrc && m.IsInternal() {
			return nil, nil
		}

		in = &flow.Flow{
			TimestampNano: time.Now().Unix(),
			InBytes:       getUInt64(&m.Payload.Bytes),
			InPkts:        getUInt64(&m.Payload.Pkts),
			OutBytes:      0,
			OutPkts:       0,
			InputPort:     client.GetInterfaceID(srcVM),
			OutputPort:    client.GetInterfaceID(dstVM),
			L4DstPort:     uint32(m.Payload.Connection.DestPort),
			L4SrcPort:     uint32(m.Payload.Connection.SrcPort),
			Protocol:      uint32(m.Payload.Connection.Protocol),
			SampleRate:    1,
			SampleAdj:     true,
			Customs: []flow.Custom{
				flow.Custom{
					ID:   customs[CLIENT_NW_LATENCY_MS],
					Type: flow.U32,
					U32:  getUInt32(&m.Payload.RTT),
				},
			},
		}
	} else {
		if dropIntraDest && m.IsInternal() {
			return nil, nil
		}

		in = &flow.Flow{
			TimestampNano: time.Now().Unix(),
			OutBytes:      getUInt64(&m.Payload.Bytes),
			OutPkts:       getUInt64(&m.Payload.Pkts),
			InBytes:       0,
			InPkts:        0,
			InputPort:     client.GetInterfaceID(dstVM),
			OutputPort:    client.GetInterfaceID(srcVM),
			L4SrcPort:     uint32(m.Payload.Connection.DestPort),
			L4DstPort:     uint32(m.Payload.Connection.SrcPort),
			Protocol:      uint32(m.Payload.Connection.Protocol),
			SampleRate:    1,
			SampleAdj:     true,
			Customs: []flow.Custom{
				flow.Custom{
					ID:   customs[CLIENT_NW_LATENCY_MS],
					Type: flow.U32,
					U32:  getUInt32(&m.Payload.RTT),
				},
			},
		}
	}

	v4Src, v6Src := PackIP(&m.Payload.Connection.SrcIP)
	v4Dst, v6Dst := PackIP(&m.Payload.Connection.DestIP)

	if m.IsIn() {
		if v6Src != nil {
			in.Ipv6SrcAddr = v6Src
		} else {
			in.Ipv4SrcAddr = v4Src
		}

		if v6Dst != nil {
			in.Ipv6DstAddr = v6Dst
		} else {
			in.Ipv4DstAddr = v4Dst
		}
	} else {
		if v6Src != nil {
			in.Ipv6DstAddr = v6Src
		} else {
			in.Ipv4DstAddr = v4Src
		}

		if v6Dst != nil {
			in.Ipv6SrcAddr = v6Dst
		} else {
			in.Ipv4SrcAddr = v4Dst
		}
	}

	return in, err
}
