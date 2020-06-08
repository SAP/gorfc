## Usage

ABAP data types supported by RFC API are mapped to GO data types:

| ABAP type | C typedef   | Length (bytes) | Description                                    | GO type                               |
| --------- | ----------- | -------------- | ---------------------------------------------- | ------------------------------------- |
| c         | RFC_CHAR    | 1-65535        | Characters, padded trailing blanks             | string                                |
| n         | RFC_NUM     | 1-65535        | Digits, fixed size, padded leading '0'         | string                                |
| x         | RFC_BYTE    | 1-65535        | Binary data                                    | []byte                                |
| p         | RFC_BCD     | 1-16           | BCD numbers (Binary Coded Decimals)            | float64 or string, from abap: string  |
| i         | RFC_INT     | 4              | Integer                                        | int32                                 |
| b         | RFC_INT1    | 1              | 1-byte integer, not directly supported by ABAP | uint8                                 |
| s         | RFC_INT2    | 2              | 2-byte integer, not directly supported by ABAP | int16                                 |
| 8         | RFC_INT8    | 8              | 8-byte integer                                 | int64                                 |
| p         | RFC_UTCLONG | 8              | Timestamp with high precision - 8 bytes        | string                                |
| f         | RFC_FLOAT   | 8              | Floating point numbers                         | float64 or string, from abap: float64 |
| d         | RFC_DATE    | 8 or 16        | Date ("YYYYMMDD")                              | Time or string, from abap: Time       |
| t         | RFC_TIME    | 6 or 12        | Time ("HHMMSS")                                | Time or string, from abap: Time       |
| a         | RFC_DECF16  | 8              | Decimal floating point 8 bytes (IEEE 754r)     | float64 or string, from abap: string  |
| e         | RFC_DECF34  | 16             | Decimal floating point 16 bytes (IEEE 754r)    | float64 or string, from abap: string  |
| g         | RFC_CHAR\*  |                | Variable-length, zero terminated string        | string                                |
| y         | RFC_BYTE\*  |                | Variable-length raw string, length in bytes    | []byte                                |
