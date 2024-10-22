/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: result.proto */

#ifndef PROTOBUF_C_result_2eproto__INCLUDED
#define PROTOBUF_C_result_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1000000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1000002 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _Fdw__Result Fdw__Result;


/* --- enums --- */


/* --- messages --- */

struct  _Fdw__Result
{
  ProtobufCMessage base;
  char *request_id;
  char *query;
  char *tn;
  uint32_t user_id;
  uint32_t server_id;
  char *remote_host;
  size_t n_aggs;
  char **aggs;
  size_t n_orderby;
  char **orderby;
  size_t n_groupby;
  char **groupby;
};
#define FDW__RESULT__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&fdw__result__descriptor) \
    , NULL, NULL, NULL, 0, 0, NULL, 0,NULL, 0,NULL, 0,NULL }


/* Fdw__Result methods */
void   fdw__result__init
                     (Fdw__Result         *message);
size_t fdw__result__get_packed_size
                     (const Fdw__Result   *message);
size_t fdw__result__pack
                     (const Fdw__Result   *message,
                      uint8_t             *out);
size_t fdw__result__pack_to_buffer
                     (const Fdw__Result   *message,
                      ProtobufCBuffer     *buffer);
Fdw__Result *
       fdw__result__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   fdw__result__free_unpacked
                     (Fdw__Result *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*Fdw__Result_Closure)
                 (const Fdw__Result *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor fdw__result__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_result_2eproto__INCLUDED */
