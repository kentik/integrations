// +build ignore

/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: result.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "result.pb-c.h"
void   fdw__result__init
                     (Fdw__Result         *message)
{
  static Fdw__Result init_value = FDW__RESULT__INIT;
  *message = init_value;
}
size_t fdw__result__get_packed_size
                     (const Fdw__Result *message)
{
  assert(message->base.descriptor == &fdw__result__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t fdw__result__pack
                     (const Fdw__Result *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &fdw__result__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t fdw__result__pack_to_buffer
                     (const Fdw__Result *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &fdw__result__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Fdw__Result *
       fdw__result__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Fdw__Result *)
     protobuf_c_message_unpack (&fdw__result__descriptor,
                                allocator, len, data);
}
void   fdw__result__free_unpacked
                     (Fdw__Result *message,
                      ProtobufCAllocator *allocator)
{
  assert(message->base.descriptor == &fdw__result__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor fdw__result__field_descriptors[9] =
{
  {
    "request_id",
    1,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Fdw__Result, request_id),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "query",
    2,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Fdw__Result, query),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "tn",
    3,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Fdw__Result, tn),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "user_id",
    4,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Fdw__Result, user_id),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "server_id",
    5,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_UINT32,
    0,   /* quantifier_offset */
    offsetof(Fdw__Result, server_id),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "remote_host",
    6,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Fdw__Result, remote_host),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "aggs",
    7,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_STRING,
    offsetof(Fdw__Result, n_aggs),
    offsetof(Fdw__Result, aggs),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "orderby",
    8,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_STRING,
    offsetof(Fdw__Result, n_orderby),
    offsetof(Fdw__Result, orderby),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "groupby",
    9,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_STRING,
    offsetof(Fdw__Result, n_groupby),
    offsetof(Fdw__Result, groupby),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned fdw__result__field_indices_by_name[] = {
  6,   /* field[6] = aggs */
  8,   /* field[8] = groupby */
  7,   /* field[7] = orderby */
  1,   /* field[1] = query */
  5,   /* field[5] = remote_host */
  0,   /* field[0] = request_id */
  4,   /* field[4] = server_id */
  2,   /* field[2] = tn */
  3,   /* field[3] = user_id */
};
static const ProtobufCIntRange fdw__result__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 9 }
};
const ProtobufCMessageDescriptor fdw__result__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "fdw.Result",
  "Result",
  "Fdw__Result",
  "fdw",
  sizeof(Fdw__Result),
  9,
  fdw__result__field_descriptors,
  fdw__result__field_indices_by_name,
  1,  fdw__result__number_ranges,
  (ProtobufCMessageInit) fdw__result__init,
  NULL,NULL,NULL    /* reserved[123] */
};
