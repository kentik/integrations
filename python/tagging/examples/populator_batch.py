import sys
sys.path.append("..")

import batch as b
import client as c

#
# Example HyperScale populator batch.
#
# This shows a batch of populators to the custom dimension 'c_my_column', replacing
# everything that's already there. Normally, for this bulk API, you'd be creating and
# appending tags/populators while looping over an input file or database connection.
# The linear nature is just to show a few different examples.
#


# -----
# initialize a batch that will replace all populators
# -----
batch = b.Batch(True)


# -----
# add a few populators with unique values - just IP addresses
# -----

crit = b.Criteria("dst")
crit.add_ip_address("1.2.3.4")
batch.add_upsert("src_ip1", crit)

crit = b.Criteria("src")
crit.add_ip_address("2.3.4.5")
batch.add_upsert("src_ip2", crit)

crit = b.Criteria("dst")
crit.add_ip_address("3.4.5.6")
batch.add_upsert("dst_ip1", crit)


# -----
# add a few populators with the same value - IP addresses and ports
# -----

crit = b.Criteria("dst")
crit.add_ip_address("10.10.10.1")
crit.add_ip_address("10.10.10.2")
crit.add_port_range(33, 55)
crit.add_port(66)
batch.add_upsert("multi", crit)

crit = b.Criteria("dst")
crit.add_ip_address("11.11.11.1")
crit.add_ip_address("11.11.11.2")
crit.add_port_range(66, 77)
crit.add_port(88)
batch.add_upsert("multi", crit)

crit = b.Criteria("dst")
crit.add_ip_address("12.12.12.1")
crit.add_ip_address("12.12.12.2")
crit.add_port_range(88, 99)
crit.add_port(111)
batch.add_upsert("multi", crit)


# -----
# add a complicated populator that uses all fields in the criteria
# -----

crit = b.Criteria("either")

# ports
crit.add_port(4)
crit.add_port(9)
crit.add_port(1000)
crit.add_port(555)
crit.add_port(666)
crit.add_port(44)

# protocols
crit.add_protocol(1)
crit.add_protocol(2)
crit.add_protocol(4)
crit.add_protocol(5)

# asn
crit.add_asn_range(12345, 20000)
crit.add_asn_range(10, 100)
crit.add_asn(500)

# last hop asn names
crit.add_last_hop_asn_name("asn2")
crit.add_last_hop_asn_name("asn3")

# next hop asn
crit.add_next_hop_asn_range(100, 200)
crit.add_next_hop_asn(35628)
crit.add_next_hop_asn_range(60000, 70000)

# next hop asn names
crit.add_next_hop_asn_name("asn2")
crit.add_next_hop_asn_name("asn3")

# bgp as paths
crit.add_bgp_as_path("^3737 1212")
crit.add_bgp_as_path("_7801_")
crit.add_bgp_as_path("2906$")

# bgp communities
crit.add_bgp_community("2096:2212")
crit.add_bgp_community("2097:2213")

# TCP flags
crit.add_tcp_flag(2)
crit.add_tcp_flag(4)
crit.add_tcp_flag(64)

# IP addresses
crit.add_ip_address("100.100.100.0/24")
crit.add_ip_address("200.100.0.0/16")
crit.add_ip_address("123.123.123.123")

# MAC addresses
crit.add_mac_address("01:42:12:ae:92:bf")
crit.add_mac_address("03:42:1f:1e:22:bf")

# country codes
crit.add_country_code("US")
crit.add_country_code("EN")

# site names
crit.add_site_name("site1")
crit.add_site_name("site2")

# device types
crit.add_device_type("router")

# interface names
crit.add_interface_name("interface_1")
crit.add_interface_name("interface_2")

# device names
crit.add_device_name("device_1")
crit.add_device_name("device_2")

# next hop IP address
crit.add_next_hop_ip_address("200.1.2.3")
crit.add_next_hop_ip_address("30.1.2.0/24")
crit.add_next_hop_ip_address("1.0.0.0/8")

batch.add_upsert("complicated_tag", crit)


# -----
# Delete some populators.
# Note: This doesn't do anything since the batch is replacing all populators.
#       It's just here for demo.
# -----
batch.add_delete("old_tag_1")
batch.add_delete("old_tag_2")
batch.add_delete("old_tag_3")


# -----
# Showtime! Submit the batch as populators for the custom dimension 'c_my_column'
# -----
client = c.Client('me@email.com', '673244805e76d01d01e42d32c56d64c2')
client.submit_populator_batch('c_my_column', batch)
