# gRPC

- Convenience
  - Protocol buffers: like a contract between the APIs. Like a Schema registry
- Performance
  - protocol buffer are sent as a serialized and sent as a binary
  - JSON is not space efficient as it is not flat and compressed.
  - Uses HTTP 2.0 
- Helpful video for gRPC and building API: https://www.youtube.com/watch?v=1MPWPq2N768

## RPC requirement
1. Function needs to be a method, belongs to the type
2. Function needs to be exported, using the capital letter on first word
3. Function needs to have two arguments. Both has to be type
4. Return type has to be error type
5. Function needs to have second argument with pointer (*)
