import json

# Batch collects tags or populators as values and criteria.
class Batch:
    def __init__(self, replace_all):
        self.replace_all = replace_all
        self.upserts = dict()
        self.deletes = set()

    # add a tag or populator to the batch by value and criteria
    def add_upsert(self, value, criteria):
        v = value.lower()
        criteria_array = self.upserts.get(v)
        if criteria_array == None:
            criteria_array = []
        criteria_array.append(criteria.to_dict())
        self.upserts[v] = criteria_array


    # delete a tag or populator by value - these are processed before upserts
    def add_delete(self, value):
        v = value.strip().lower()
        if len(v) == 0:
            raise ValueError("Invalid value for delete. Value is empty.")

        self.deletes.add(v)

    # return an array of batch parts to submit
    def parts(self):
        parts = []

        upserts = dict()
        deletes = []

        # 20,000 records per batch
        max_records = 20000

        # loop upserts first - fit the deletes in afterward
        record_count = 0
        for value in self.upserts:
            record_count += 1
            upserts[value] = self.upserts[value]
            if record_count >= max_records:
                parts.append(BatchPart(self.replace_all, upserts, deletes))
                upserts = dict()
                deletes = []
                record_count = 0

        for value in self.deletes:
            record_count += 1
            deletes.append({ 'value': value })
            if record_count > max_records:
                parts.append(BatchPart(self.replace_all, upserts, deletes))
                upserts = dict()
                deletes = []
                record_count = 0

        if len(upserts) + len(deletes) > 0:
            # finish the batch
            parts.append(BatchPart(self.replace_all, upserts, deletes))

        if len(parts) == 0:
            raise ValueError("Batch has no data.")

        # last part finishes the batch
        parts[-1].set_last_part()
        return parts


# BatchPart contains tags/populators to be sent as part of a (potentially) multi-part logical batch
class BatchPart:
    def __init__(self, replace_all, upserts, deletes):
        if replace_all not in [True, False]:
            raise ValueError("Invalid value for replace_all. Must be True or False.")

        self.replace_all = replace_all
        self.complete = False
        self.upserts = upserts
        self.deletes = deletes


    # marks this part as the last to be sent for a logical batch
    def set_last_part(self):
        self.complete = True


    # build JSON with the input guid
    def build_json(self, guid):
        upserts = []
        for value in self.upserts:
            upserts.append({ "value": value, "criteria": self.upserts[value] })
        return json.dumps({'replace_all': self.replace_all, 'guid': guid, 'complete': self.complete, 'upserts': upserts, 'deletes': self.deletes})


# Criteria defines a set of rules that must match for a tag or populator.
# A flow record is tagged with this value if it matches at least one value from each non-empty criteria.
class Criteria:
    def __init__(self, direction):
        self._json_dict = dict()

        v = direction.lower()
        if v not in ["src", "dst", "either"]:
            raise ValueError("Invalid value for direction. Valid: src, dst, either.")
        self._json_dict['direction'] = v


    def to_dict(self):
        return self._json_dict


    def _ensure_array(self, key):
        if key not in self._json_dict:
            self._json_dict[key] = []


    def add_port(self, port):
        if port < 0 or port > 65535:
            raise ValueError("Invalid port. Valid: 0-65535.")
        self._ensure_array('ports')
        self._json_dict['ports'].append(str(port))


    def add_port_range(self, start, end):
        if start < 0 or start > 65535:
            raise ValueError("Invalid start port. Valid: 0-65535.")
        if end < 0 or end > 65535:
            raise ValueError("Invalid end port. Valid: 0-65535.")
        self._ensure_array('ports')

        if start == end:
            self._json_dict['ports'].append(str(start))
            return

        self._json_dict['ports'].append("%d-%d" % (start, end))


    def add_protocol(self, protocol):
        if protocol < 0 or protocol > 255:
            raise ValueError("Invalid protocol. Valid: 0-255.")
        self._ensure_array('protocols')
        self._json_dict['protocols'].append(protocol)


    def add_asn(self, asn):
        _validate_asn(asn)
        self._ensure_array('asns')
        self._json_dict['asns'].append(str(asn))


    def add_asn_range(self, start, end):
        _validate_asn(start)
        _validate_asn(end)
        self._ensure_array('asns')

        if start == end:
            self._json_dict['asns'].append(str(start))
            return

        if start > end:
            raise ValueError("Invalid ASN range. Start must be before end.")

        self._json_dict['asns'].append("%d-%d" % (start, end))


    def add_last_hop_asn_name(self, last_hop_asn_name):
        v = last_hop_asn_name.strip()
        if len(v) == 0:
            raise ValueError("Invalid last_hop_asn_name. Value is empty.")

        self._ensure_array('last_hop_asn_names')
        self._json_dict['last_hop_asn_names'].append(v)


    def add_next_hop_asn_range(self, start, end):
        _validate_asn(start)
        _validate_asn(end)
        self._ensure_array('next_hop_asns')

        if start == end:
            self._json_dict['next_hop_asns'].append(str(start))
            return
        self._json_dict['next_hop_asns'].append("%d-%d" % (start, end))


    def add_next_hop_asn(self, next_hop_asn):
        _validate_asn(next_hop_asn)
        self._ensure_array('next_hop_asns')
        self._json_dict['next_hop_asns'].append(str(next_hop_asn))


    def add_next_hop_asn_name(self, next_hop_asn_name):
        v = next_hop_asn_name.strip()
        if len(v) == 0:
            raise ValueError("Invalid next_hop_asn_name. Value is empty.")

        self._ensure_array('next_hop_asn_names')
        self._json_dict['next_hop_asn_names'].append(v)


    def add_bgp_as_path(self, bgp_as_path):
        v = bgp_as_path.strip()
        if len(v) == 0:
            raise ValueError("Invalid bgp_as_path. Value is empty.")

        # TODO: validate
        self._ensure_array('bgp_as_paths')
        self._json_dict['bgp_as_paths'].append(v)


    def add_bgp_community(self, bgp_community):
        v = bgp_community.strip()
        if len(v) == 0:
            raise ValueError("Invalid bgp_community. Value is empty.")

        # TODO: validate
        self._ensure_array('bgp_communities')
        self._json_dict['bgp_communities'].append(v)


    # add a single TCP flag - will be OR'd into the existing bitmask
    def add_tcp_flag(self, tcp_flag):
        if tcp_flag not in [1, 2, 4, 8, 16, 32, 64, 128]:
            raise ValueError("Invalid TCP flag. Valid: [1, 2, 4, 8, 16,32, 64, 128]")

        if self._json_dict.get('tcp_flags') == None:
            self._json_dict['tcp_flags'] = 0
        self._json_dict['tcp_flags'] |= tcp_flag


    # set the complete tcp flag bitmask
    def set_tcp_flags(self, tcp_flags):
        if tcp_flags < 0 or tcp_flags > 255:
            raise ValueError("Invalid tcp_flags. Valid: 0-255.")

        self._json_dict['tcp_flags'] = tcp_flags


    def add_ip_address(self, ip_address):
        v = ip_address.strip()
        if len(v) == 0:
            raise ValueError("Invalid ip_address. Value is empty.")

        # TODO: validate?
        self._ensure_array('ip_addresses')
        self._json_dict['ip_addresses'].append(v)


    def add_mac_address(self, mac_address):
        v = mac_address.strip()
        if len(v) == 0:
            raise ValueError("Invalid mac_address. Value is empty.")

        # TODO: validate?
        self._ensure_array('mac_addresses')
        self._json_dict['mac_addresses'].append(v)


    def add_country_code(self, country_code):
        v = country_code.strip()
        if len(v) == 0:
            raise ValueError("Invalid country_code. Value is empty.")

        # TODO: validate?
        self._ensure_array('country_codes')
        self._json_dict['country_codes'].append(v)


    def add_site_name(self, site_name):
        v = site_name.strip()
        if len(v) == 0:
            raise ValueError("Invalid site_name. Value is empty.")

        self._ensure_array('site_names')
        self._json_dict['site_names'].append(v)


    def add_device_type(self, device_type):
        v = device_type.strip()
        if len(v) == 0:
            raise ValueError("Invalid device_type. Value is empty.")

        self._ensure_array('device_types')
        self._json_dict['device_types'].append(v)


    def add_interface_name(self, interface_name):
        v = interface_name.strip()
        if len(v) == 0:
            raise ValueError("Invalid interface_name. Value is empty.")

        self._ensure_array('interface_names')
        self._json_dict['interface_names'].append(v)


    def add_device_name(self, device_name):
        v = device_name.strip()
        if len(v) == 0:
            raise ValueError("Invalid device_name. Value is empty.")

        self._ensure_array('device_names')
        self._json_dict['device_names'].append(v)


    def add_next_hop_ip_address(self, next_hop_ip_address):
        v = next_hop_ip_address.strip()
        if v == 0:
            raise ValueError("Invalid next_hop_ip_address. Value is empty.")

        self._ensure_array('next_hop_ip_addresses')
        self._json_dict['next_hop_ip_addresses'].append(v)


def _validate_asn(asn):
    if asn < 0 or asn > 4294967295:
        raise ValueError("Invalid ASN. Valid: 0-4294967295")
