window.BENCHMARK_DATA = {
  "lastUpdate": 1727859794130,
  "repoUrl": "https://github.com/ucan-wg/go-ucan",
  "entries": {
    "Go Benchmark": [
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "50ea43e3fa317d1e3b202c5e7ca176b35e19b924",
          "message": "Merge pull request #33 from ucan-wg/continuous-bench\n\ngha: add a continuous benchmark action",
          "timestamp": "2024-10-01T17:40:50+02:00",
          "tree_id": "1e61b975dacff32071db7329167226eda6aeaa00",
          "url": "https://github.com/ucan-wg/go-ucan/commit/50ea43e3fa317d1e3b202c5e7ca176b35e19b924"
        },
        "date": 1727797732051,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGlob",
            "value": 42.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.37,
            "unit": "ns/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15176,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15176,
            "unit": "ns/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 111582,
            "unit": "ns/op\t   15358 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 111582,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15358,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54064,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54064,
            "unit": "ns/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 99378,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 99378,
            "unit": "ns/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 62659,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 62659,
            "unit": "ns/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 111723,
            "unit": "ns/op\t   15309 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 111723,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15309,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55260,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55260,
            "unit": "ns/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 100876,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 100876,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58579,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58579,
            "unit": "ns/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "19824 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6b8fbcee0adc4a289da32668616a3f49bad526a1",
          "message": "Merge pull request #34 from ucan-wg/invocation-stub\n\ntoken: add invocation partial stub",
          "timestamp": "2024-10-02T11:02:32+02:00",
          "tree_id": "f11c633892fe52609f43e85a407013ded5a0ce7d",
          "url": "https://github.com/ucan-wg/go-ucan/commit/6b8fbcee0adc4a289da32668616a3f49bad526a1"
        },
        "date": 1727859793240,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGlob",
            "value": 41.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.8,
            "unit": "ns/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15132,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15132,
            "unit": "ns/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109556,
            "unit": "ns/op\t   15358 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109556,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15358,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53087,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53087,
            "unit": "ns/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97627,
            "unit": "ns/op\t   14414 B/op\t     305 allocs/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97627,
            "unit": "ns/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14414,
            "unit": "B/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58314,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58314,
            "unit": "ns/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108815,
            "unit": "ns/op\t   15310 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108815,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15310,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55681,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55681,
            "unit": "ns/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97985,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97985,
            "unit": "ns/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57119,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57119,
            "unit": "ns/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "20989 times\n4 procs"
          }
        ]
      }
    ]
  }
}